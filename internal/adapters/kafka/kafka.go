package kafka

import (
	"github.com/segmentio/kafka-go"
	"gitlab.com/g6834/team17/task-service/internal/config"
	"gitlab.com/g6834/team17/task-service/internal/ports"
)

type Producer struct {
	*kafka.Writer
	taskService ports.Task
}

func NewProducer(cfg *config.Config) *Producer {
	producer := &Producer{}
	producer.Writer = &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Broker),
		Balancer: &kafka.LeastBytes{},
	}
	return producer
}
