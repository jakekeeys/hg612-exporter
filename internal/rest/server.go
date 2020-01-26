package rest

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Server struct {
	httpSrv *http.Server
}

func New(bind string) *Server {
	var s Server

	r := http.NewServeMux()
	r.Handle("/prometheus", promhttp.Handler())

	server := http.Server{
		Addr:    bind,
		Handler: handlers.LoggingHandler(os.Stdout, r),
	}

	s.httpSrv = &server

	return &s
}

func (s Server) Start() {
	go func() {
		err := s.httpSrv.ListenAndServe()
		if err != nil {
			logrus.Panic(err)
		}
	}()
}

func (s Server) Stop() {
	err := s.httpSrv.Shutdown(context.Background())
	if err != nil {
		logrus.Warn(err)
	}
}
