package analytics

import (
	"context"
	"github.com/rs/zerolog"
	"gitlab.com/g6834/team17/task-service/internal/config"
	"gitlab.com/g6834/team17/task-service/internal/constants"
	"gitlab.com/g6834/team17/task-service/internal/domain/models"
	pb "gitlab.com/g6834/team17/task-service/pkg/analytics_messaging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"strconv"
)

type GrpcClient struct {
	pb.AnalyticsMsgServiceClient
	conn   *grpc.ClientConn
	logger *zerolog.Logger
	cfg    *config.Analytics
}

func NewGrpcClient(cfg *config.Config, l *zerolog.Logger) (*GrpcClient, error) {
	conn, err := grpc.Dial(net.JoinHostPort(cfg.Analytics.Host, strconv.Itoa(cfg.Analytics.Port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	analyticsClient := pb.NewAnalyticsMsgServiceClient(conn)

	return &GrpcClient{
		conn:                      conn,
		AnalyticsMsgServiceClient: analyticsClient,
		logger:                    l,
		cfg:                       &cfg.Analytics,
	}, nil
}

func (c *GrpcClient) SendEvent(ctx context.Context, task *models.Task, event models.EventType) error {
	user, ok := ctx.Value(constants.CTX_USER).(*models.User)
	if !ok {
		return models.ErrCastUser
	}

	message := &pb.EventMessage{
		TaskUuid: task.ID.String(),
	}

	switch event {
	case models.Created:
		message.EventType = pb.EventType_CREATED
		message.UserUuid = task.CreatorID.String()
		message.Timestamp = timestamppb.New(task.CreatedDate)

	case models.ApprovedBy:
		message.EventType = pb.EventType_APPROVED_BY
		message.UserUuid = user.ID.String()
		message.Timestamp = timestamppb.New(task.UpdatedDate)

	case models.RejectedBy:
		message.EventType = pb.EventType_REJECTED_BY
		message.UserUuid = user.ID.String()
		message.Timestamp = timestamppb.New(task.UpdatedDate)

	case models.Signed:
		message.EventType = pb.EventType_SIGNED
		message.UserUuid = task.CreatorID.String()
		message.Timestamp = timestamppb.New(task.UpdatedDate)

	case models.Sent:
		message.EventType = pb.EventType_SENT
		message.UserUuid = task.CreatorID.String()
		message.Timestamp = timestamppb.Now()

	default:
		message.EventType = pb.EventType_UNKNOWN
		message.UserUuid = task.CreatorID.String()
		message.Timestamp = timestamppb.Now()
	}

	_, err := c.AnalyticsMsgServiceClient.SendMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
