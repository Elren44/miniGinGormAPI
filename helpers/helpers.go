package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func WriteJSON(c *gin.Context, statusCode int, data interface{}) {
	var message Message
	log.Println("status code : ", statusCode)

	message.StatusCode = statusCode
	message.Data = data
	c.JSON(200, message)
}
