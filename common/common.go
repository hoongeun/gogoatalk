package common

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Message struct {
	Id        string
	Text      string
	Userid    string
	CreatedAt int64
	UpdatedAt int64
}

const timeFormat = "03:04:05 PM"

func ClientLogf(ts time.Time, format string, args ...interface{}) {
	log.Printf("[%s] <<Client>>: "+format, append([]interface{}{ts.Format(timeFormat)}, args...)...)
}

func ServerLogf(ts time.Time, format string, args ...interface{}) {
	log.Printf("[%s] <<Server>>: "+format, append([]interface{}{ts.Format(timeFormat)}, args...)...)
}

func MessageLog(ts time.Time, name, msg string) {
	log.Printf("[%s] %s: %s", ts.Format(timeFormat), name, msg)
}

func SignalContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		ServerLogf(time.Now(), "listening for shutdown signal")
		<-sigs
		ServerLogf(time.Now(), "shutdown signal received")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	return ctx
}
