package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := getRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getRouter() *gin.Engine {
	r := gin.Default()

	merchGroup := r.Group("/api/v1/merch")
	{
		merchGroup.GET("/", GetMerch)
		merchGroup.POST("/", PostMerch)

		merchGroup.GET("/:id", GetMerchItem)
		merchGroup.PUT("/:id", PutMerchItem)
	}

	return r
}
