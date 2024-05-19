package main

import (
	"context"
	"github.com/sergey4qb/rate/internal/cmd/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.Create()
	err := application.Start(ctx)
	if err != nil {
		log.Println(err)
	}

	<-stop
	log.Println("Shutting down server...")
}
