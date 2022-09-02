package application

import (
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"gitlab.com/g6834/team17/task-service/internal/adapters/grpc/analytics"
	"gitlab.com/g6834/team17/task-service/internal/adapters/grpc/auth"
	"gitlab.com/g6834/team17/task-service/internal/adapters/kafka"
	"gitlab.com/g6834/team17/task-service/internal/sender"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	dbMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"gitlab.com/g6834/team17/task-service/internal/adapters/http"
	"gitlab.com/g6834/team17/task-service/internal/adapters/postgres"
	"gitlab.com/g6834/team17/task-service/internal/adapters/presenter"
	"gitlab.com/g6834/team17/task-service/internal/config"
	"gitlab.com/g6834/team17/task-service/internal/domain/task"
	"gitlab.com/g6834/team17/task-service/internal/utils"
	"golang.org/x/sync/errgroup"
)

var (
	srv    *http.Server
	logger *zerolog.Logger
)

func Start(ctx context.Context) {

	/* CONFIG init */
	if err := config.ReadConfigYML("config.yaml"); err != nil {
		log.Fatal("cannot read config file", err)
	}

	cfg := config.New()

	/* LOGGER init */
	logger = utils.NewLogger(cfg)

	/* DATABASE init */
	db, err := postgres.New(cfg, logger)
	if err != nil {
		logger.Error().Err(err).Msg("cannot initialize database")
	}
	defer db.Close()

	if cfg.Database.Migrations != "" {
		if err := runMigrations(db, cfg); err != nil {
			logger.Error().Err(err).Msg("cannot up migrate")
		}
	}

	/* SERVICES init */
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:   "https://40376d00d7b8408f8bc64950ce173be9@sentry.k8s.golang-mts-teta.ru/49",
		Debug: true,
	}); err != nil {
		logger.Error().Err(err).Msg("cannot init sentry")
	}
	defer sentry.Flush(2 * time.Second)

	/* Jaeger init */
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost(cfg.Jaeger.Host), jaeger.WithAgentPort(cfg.Jaeger.Port)))
	if err != nil {
		logger.Error().Err(err).Msg("cannot init jaeger collector")
		sentry.CaptureException(err)
	}

	/* Tracer provider init */
	traceProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(cfg.Jaeger.Service))))

	otel.SetTracerProvider(traceProvider)

	authS, err := auth.New(cfg, logger)
	if err != nil {
		logger.Error().Err(err).Msg("cannot initialize auth service")
		sentry.CaptureException(err)
	}
	defer authS.Close()

	clientGrpc, err := analytics.NewGrpcClient(cfg, logger)
	if err != nil {
		logger.Error().Err(err).Msg("cannot initialize analytics service")
		sentry.CaptureException(err)
	}

	clientKafka := kafka.NewProducer(cfg) //kafka client to analytics service
	clientKafka.Topic = cfg.Kafka.Topics.Tasks

	senderGrpc := sender.NewAnalyticsSender(clientGrpc) //wrapper on grpc and kafka client
	senderGrpcWithKafka := sender.NewAnalyticsSender(senderGrpc)

	taskS := task.New(db, authS, senderGrpcWithKafka)
	presenterS := presenter.New(logger)

	clientMessageService := kafka.NewProducer(cfg) //kafka client to mail service
	clientMessageService.Topic = cfg.Kafka.Topics.Messages

	senderToMailWithKafka := sender.NewMessageSender(clientMessageService)

	/* SERVERS Init */
	srv, err = http.New(cfg, logger, authS, taskS, presenterS)
	if err != nil {
		logger.Error().Err(err).Msg("cannot initialize server")
		sentry.CaptureException(err)
	}

	var g errgroup.Group
	g.Go(func() error {
		return srv.Start()
	})

	g.Go(func() error {
		for {
			select {
			case <-time.Tick(time.Duration(cfg.MessageHandler.RatePeriodMicroseconds)):
				if err := senderToMailWithKafka.WriteAll(ctx); err != nil {
					logger.Error().Err(err)
					continue
				}
				logger.Info().Msg("successfully sent msg")
			case <-ctx.Done():
				return nil
			}
		}
	})

	logger.Info().Msg("app is started")
	err = g.Wait()
	if err != nil {
		logger.Fatal().Err(err).Msg("http server start failed")
		sentry.CaptureException(err)
	}
}

func Stop() {
	logger.Warn().Msg("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()

	err := srv.Stop(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Error while stopping")
		sentry.CaptureException(err)
	}

	logger.Warn().Msg("app has stopped")
}

func runMigrations(pg *postgres.Database, cfg *config.Config) error {
	// Migrations block
	driver, err := dbMigrate.WithInstance(pg.DB().DB, &dbMigrate.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+cfg.Database.Migrations, cfg.Database.Name, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
