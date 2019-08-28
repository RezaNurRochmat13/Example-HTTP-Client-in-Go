package main

import (
	"httpClientGo/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	routerCoonfig := gin.Default()

	versioningAPIs := routerCoonfig.Group("/public/api/v1/")
	{
		versioningAPIs.GET("user", controller.GetResponse)
	}

	routerCoonfig.Run(":8080")
}
