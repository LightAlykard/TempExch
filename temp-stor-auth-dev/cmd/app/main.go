package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/TempExch/temp-stor-auth-dev/internal/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	go application.Start(ctx)
	<-ctx.Done()
	application.Stop(ctx)
}
