package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/call", func(ctx *gin.Context) {
		ctx.String(200, "Micro3")
	})
	return r
}

func main() {
	r := setupRouter()
	//tracer := opentracing.GlobalTracer()
	//r.Use(ginhttp.Middleware(tracer))
	r.Run(":8082")
}
