package main

import (
	"context"
	"gitlab.com/g6834/team17/task-service/internal/application"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	go application.Start(ctx)
	<-ctx.Done()
	application.Stop()
}
