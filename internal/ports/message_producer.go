package ports

type MessageProducer[T any] interface {
	Send(key string, data *T) error
}
