package server

import (
	"context"
	"gitlab.com/piorun102/lg"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (s *Server) listen() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	<-ctx.Done()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	closeChan := make(chan struct{})
	go func(chan<- struct{}) {
		if err := s.httpServer.Shutdown(timeoutCtx); err != nil {
			lg.Error(err)
		}
		closeChan <- struct{}{}
	}(closeChan)

	select {
	case <-timeoutCtx.Done():
		lg.Warn("shutdown due to timeout")

	case <-closeChan:
		lg.Debug("Shutdown success")
	}
}
