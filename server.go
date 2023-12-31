package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// graceful shutdown
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 別goroutineのエラーを受け取るためにerrgroupを利用
	eg, ctx := errgroup.WithContext(ctx)
	// 別goroutineでサーバーを起動
	eg.Go(
		func() error {
			// ErrServerClosedは正常なので除外
			if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
				log.Printf("failed to close %+v", err)
				return err
			}
			return nil
		},
	)

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
