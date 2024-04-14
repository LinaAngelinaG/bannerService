package server

import (
	"bannerService/config"
	bannerChi "bannerService/internal/banner/httpChi"
	"context"
	"gitlab.com/piorun102/lg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var stop context.CancelFunc

type Server struct {
	httpServer *http.Server
	cfg        *config.Config
	ctx        context.Context
}

func New(cfg *config.Config) *Server {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	stop = cancel
	return &Server{
		cfg:        cfg,
		ctx:        ctx,
		httpServer: &http.Server{Addr: "0.0.0.0:" + "9090"},
	}
}

func StopContext() context.CancelFunc {
	return stop
}

func (s *Server) Run() {
	defer bannerChi.HandlePanic()
	var err error
	if err = s.MapHandlers(); err != nil {
		lg.Panicf("Cannot map router. Error: {%s}", err)
	}
	go func() {
		lg.Infof("starting server on %v", s.httpServer.Addr)
		if err = s.httpServer.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			lg.Panicf("ListenAndServe: {%s}", err)
		}
	}()
	s.listen()
}
