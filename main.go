package main

import (
	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/version"

	chaos "github.com/msfidelis/gin-chaos-monkey"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(chaos.Load())

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)

	// Version
	router.GET("/version", version.Get)

	router.Run()
}
