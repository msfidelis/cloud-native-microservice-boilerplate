package main

import (
	"os"

	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/version"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

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

	router := gin.New()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	subLog := zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Logger()

	router.Use(logger.SetLogger(logger.Config{
		Logger: &subLog,
		UTC:    true,
	}))

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)

	// Version
	router.GET("/version", version.Get)

	router.Run()
}
