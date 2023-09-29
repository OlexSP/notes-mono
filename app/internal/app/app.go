package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/OlexSP/notes-mono/internal/config"
	"github.com/OlexSP/notes-mono/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
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

	a.logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
	if err != nil {
		a.logger.Info("failed to listen")
		os.Exit(1)
	}

	a.logger.Info("CORS initializing")

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch,
		},
		AllowedOrigins:     []string{"http://localhost:10000", "https://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Authorization", "Content-Type", "X-CSRF-Token", "Location", "Charset", "Access-Control-Allow-Origin", "Origin", "Accept", "Content-Length", "Accept-Encoding"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Content-Disposition", "Authorization"},
		// enable debug mode for testing in prod
		Debug: false,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
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
