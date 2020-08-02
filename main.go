package main

import (
	"github.com/msfidelis/change-me/controllers/healthcheck"
	"github.com/msfidelis/change-me/controllers/version"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Healthcheck
	router.GET("/healthcheck", healthcheck.Ok)

	// Version
	router.GET("/version", version.Get)

	router.Run()
}
