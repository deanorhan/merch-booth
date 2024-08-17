package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := Open()
	if err != nil {
		log.Panic(err.Error())
	}
	defer Close()

	r := getRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
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
