package main

import (
	"github.com/gin-contrib/logger"
	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/gin-gonic/gin"

	_ "github.com/msfidelis/change-me/docs"

	"os"
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

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().Logger()

	router := gin.New()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	router.Use(logger.SetLogger(logger.Config{
		Logger:   &subLog,
		UTC:      true,
		SkipPath: []string{"/skip"},
	}))

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)

	// Version
	router.GET("/version", version.Get)

	router.Run()
}
