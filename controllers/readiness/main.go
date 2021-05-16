package readiness

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/change-me/pkg/memory_cache"
)

type Response struct {
	Status string `json:"status" binding:"required"`
}

// Ok godoc
// @Summary Return 200 status Ok in readiness
// @Tags readiness
// @Produce json
// @Success 200 {object} Response
// @Router /readiness [get]
func Ok(c *gin.Context) {
	m := memory_cache.GetInstance()
	var response Response

	_, found := m.Get("readiness.ok")
	if found {
		response.Status = "NotReady"
		c.JSON(http.StatusServiceUnavailable, response)
	} else {
		response.Status = "Ready"
		c.JSON(http.StatusOK, response)
	}
}
