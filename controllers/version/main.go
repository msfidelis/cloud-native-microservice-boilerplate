package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/change-me/pkg/configuration"
)

func Get(c *gin.Context) {
	configs := configuration.Load()
	c.JSON(http.StatusOK, gin.H{
		"version": configs.Version,
	})
}
