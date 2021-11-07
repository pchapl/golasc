package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pOng")
	})

	return r
}
func Run(port string) {
	r := setupRouter()

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
