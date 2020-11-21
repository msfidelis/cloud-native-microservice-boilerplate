package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int `json:"status" binding:"required"`
}

func Ok(c *gin.Context) {
	var response Response
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
