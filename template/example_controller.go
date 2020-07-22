package template

var (
	ExampleController = `package exampleController

import (
	"{{.Name}}/pkg/app"
	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {
	appG := app.Gin{C: c}
	//agent := c.GetHeader("User-Agent")
	//// raw取参数
	//rawdata, err := c.GetRawData()
    //if err != nil {
	//	appG.ResponnseFailure(10001, "参数错误")
	//	return
	//}
	appG.ResoponseSucess(map[string]interface{}{
		"example key": "exapmle value",
	})
}`
)
