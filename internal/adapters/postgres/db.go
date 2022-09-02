package postgres

import (
	"errors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gitlab.com/g6834/team17/task-service/internal/config"
)

type Database struct {
	db     *sqlx.DB
	logger *zerolog.Logger
}

func New(cfg *config.Config, l *zerolog.Logger) (*Database, error) {
	db, err := sqlx.Open("pgx", connectionString(cfg))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("failed ping the database")
	}

	//db.SetMaxOpenConns(cfg.MaxOpenConns)
	//db.SetMaxIdleConns(cfg.MaxIdleConns)
	//db.SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	return &Database{
		db:     db,
		logger: l,
	}, nil
}

func (d *Database) DB() *sqlx.DB {
	return d.db
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		d.logger.Error().Err(err)
	}
}
