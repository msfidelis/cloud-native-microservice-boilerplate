package healthcheck

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int `json:"status" binding:"required"`
}

func Ok(c *gin.Context) {

	fmt.Println("kkkkkkk")

	var response Response
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}
