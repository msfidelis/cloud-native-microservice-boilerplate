package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/change-me/pkg/configuration"
)

type Response struct {
	Application string `json:"application" binding:"required"`
	Version     string `json:"version" binding:"required"`
}

func Get(c *gin.Context) {
	configs := configuration.Load()

	response := Response{
		Version:     configs.Version,
		Application: configs.Application,
	}

	c.JSON(http.StatusOK, response)
}
