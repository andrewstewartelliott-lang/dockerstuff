package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type InfoResponse struct {
	GoVersion string `json:"go_version"`
}

func main() {
	router := gin.Default()

	router.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, InfoResponse{
			GoVersion: runtime.Version(),
		})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
