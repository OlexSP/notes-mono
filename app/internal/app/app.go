package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/OlexSP/notes-mono/internal/config"
	"github.com/OlexSP/notes-mono/pkg/logging"
	"github.com/OlexSP/notes-mono/pkg/metric"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net"
	"net/http"
	"os"
)

type App struct {
	cfg *config.Config

	logger     *slog.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(config *config.Config, logger *slog.Logger) (App, error) {
	logger.Info("router initialization")
	router := httprouter.New()

	logger.Info("swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	logger.Info("heartbeat metrics initialization")
	metricsHandler := metric.Handler{}
	metricsHandler.Register(router)

	return App{
		cfg:    config,
		logger: logger,
		router: router,
	}, nil
}

func (a *App) Run() error {
	return a.startHTTP()
}

func (a *App) startHTTP() error {

	a.logger.With(
		"IP", a.cfg.HTTP.IP,
		"Port", a.cfg.HTTP.Port,
	).Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		a.logger.Error("failed to listen", logging.Err(err))
		os.Exit(1)
	}

	a.logger.With(
		"IP", a.cfg.HTTP.IP,
		"Port", a.cfg.HTTP.Port,
		"AllowedMethods", a.cfg.HTTP.CORS.AllowedMethods,
		"AllowedOrigins", a.cfg.HTTP.CORS.AllowedOrigins,
		"AllowCredentials", a.cfg.HTTP.CORS.AllowCredentials,
		"AllowedHeaders", a.cfg.HTTP.CORS.AllowedHeaders,
		"OptionsPassthrough", a.cfg.HTTP.CORS.OptionsPassthrough,
		"ExposedHeaders", a.cfg.HTTP.CORS.ExposedHeaders,
		"Debug", a.cfg.HTTP.CORS.Debug,
	).Info("CORS initializing")

	c := cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Error("failed to start server", logging.Err(err))
			os.Exit(1)
		}
	}

	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Error("failed to shutdown server", logging.Err(err))
	}

	return err
}
