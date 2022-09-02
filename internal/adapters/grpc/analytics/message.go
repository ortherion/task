package analytics

//
//import (
//	"context"
//	"errors"
//	"gitlab.com/g6834/team17/task-service/internal/domain/models"
//)
//
//var ErrNoSuccess = errors.New("error while sending message to grpc server")
//
//type MessageProducerGrpc[T any] struct {
//	eventType string
//	analytics *AnalyticsClient
//	toTask    func(data *T) (*models.Task, error)
//}
//
//func NewMessageProducerGrpc[T any](analytics *AnalyticsClient, toTask func(data *T) (*models.Task, error)) *MessageProducerGrpc[T] {
//	return &MessageProducerGrpc[T]{
//		analytics: analytics,
//		toTask:    toTask,
//	}
//}
//
//func (mp *MessageProducerGrpc[T]) Send(key string, data *T) error {
//	task, err := mp.toTask(data)
//	if err != nil {
//		return err
//	}
//	ok, err := mp.analytics.SendMessage(context.Background(), task, mp.eventType)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		return ErrNoSuccess
//	}
//	return nil
//}
