package main

import (
	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/liveness"
	"github.com/msfidelis/change-me/controllers/readiness"
	"github.com/msfidelis/change-me/controllers/version"
	"github.com/msfidelis/change-me/middlewares"

	// "github.com/msfidelis/change-me/controllers/system"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	loggerInternal "github.com/msfidelis/change-me/pkg/log"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// "github.com/patrickmn/go-cache"
	"github.com/msfidelis/change-me/pkg/memory_cache"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/msfidelis/change-me/docs"
)

// @title github.com/msfidelis/change-me
// @version 1.0
// @description Cloud Native Toolset Running in Containers.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email matheus@nanoshots.com.br

// @license.name MIT
// @license.url https://github.com/mfidelis/github.com/msfidelis/change-me/blob/master/LICENSE

// @BasePath /
func main() {

	router := gin.New()

	// Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// Memory Cache Singleton
	c := memory_cache.GetInstance()

	// Readiness Probe Mock Config
	probe_time_raw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probe_time_raw == "" {
		probe_time_raw = "5"
	}
	probe_time, err := strconv.ParseUint(probe_time_raw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}
	c.Set("readiness.ok", "false", time.Duration(probe_time)*time.Second)

	// Prometheus Exporter Config
	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	// Logger
	logInternal := loggerInternal.Instance()

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	router.Use(chaos.Load())
	router.Use(middlewares.JsonLoggerMiddleware())
	// router.Use(logger.SetLogger(
	// 	logger.WithSkipPath([]string{"/skip"}),
	// 	logger.WithUTC(true),
	// ))

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)

	// Liveness
	router.GET("/liveness", liveness.Ok)

	// Readinesscurl
	router.GET("/readiness", readiness.Ok)

	// Version
	router.GET("/version", version.Get)

	// Graceful Shutdown Config
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logInternal.
				Error().
				Str("Error", err.Error()).
				Msg("Failed to listen")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logInternal.
		Warn().
		Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logInternal.
			Error().
			Str("Error", err.Error()).
			Msg("Server forced to shutdown: ")
	}

	fmt.Println("Server exiting")

}
