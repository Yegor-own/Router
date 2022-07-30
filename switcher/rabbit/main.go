package rabbit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListenRabbitMQ() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":2100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
