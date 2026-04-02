package main

import (
	"github.com/gin-gonic/gin"

	"server-slug/go-gin"
)

func
 main() {
	r := gin.Default()
	r.GET("/links", gogin.GetAll)
	r.GET("/links/:slug", gogin.GetAllSlug)
	r.POST("/links", gogin.Post)
	r.PATCH("/links/:id", gogin.Patch)
	r.DELETE("/links", gogin.Delete)
	r.GET("/r/:slug", gogin.Redirect)
	r.Run(":8080")
}


