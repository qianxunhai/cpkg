package template

var (
	Route = `package routers

import (
	"{{.Name}}/controllers/exampleController"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()

	apiRoute := r.Group("/api")
	apiRoute.GET("example", exampleController.Example)
	return r
}`
)
