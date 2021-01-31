package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping/:id", getPing)
	// route.Run(":1997")
	route.Run(":3333")
}
func getPing(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}
