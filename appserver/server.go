package appserver

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/zehiraa/golang_test_project/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	Logger *log.Logger
	Router *chi.Mux
}

func (s *Server) Start() {
	server := &http.Server{Addr: ":" + config.Port(), Handler: s.Router}
	go gracefulShutdown(server, s.Logger)
	s.Logger.Info("server started on port: %s", config.Port())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.Logger.Error("cannot start server.", err)
		panic("cannot start server")
	}
}

func gracefulShutdown(srv *http.Server, logger *log.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown Error: ", err)
	}

	logger.Info("Server exiting")
}

func (s Server) HealthCheck() {
	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, 200)
		render.JSON(w, r, map[string]interface{}{"status": "ok"})
		return
	})
}
