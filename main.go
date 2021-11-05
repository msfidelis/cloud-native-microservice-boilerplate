package main

import (
	"github.com/gin-contrib/logger"
	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/liveness"
	"github.com/msfidelis/change-me/controllers/readiness"
	"github.com/msfidelis/change-me/controllers/version"
	"github.com/msfidelis/change-me/pkg/memory_cache"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/gin-gonic/gin"
	"github.com/Depado/ginprom"

	_ "github.com/msfidelis/change-me/docs"

	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

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

	//Middlewares
	router.Use(p.Instrument())
	router.Use(gin.Recovery())
	router.Use(chaos.Load())
	router.Use(logger.SetLogger(
		logger.WithSkipPath([]string{"/skip"}),
		logger.WithUTC(true),
		logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
			return zerolog.New(out).With().
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Dur("latency", latency).
				Int("status", c.Writer.Status()).
				Str("user_agent", c.Request.UserAgent()).
				Logger()
		}),
	))


	// Healthcheck Router
	router.GET("/healthcheck", healthcheck.Ok)

	// Version Router
	router.GET("/version", version.Get)

	// Liveness and Readiness
	router.GET("/liveness", liveness.Ok)
	router.GET("/readiness", readiness.Ok)

	router.Run()
}
