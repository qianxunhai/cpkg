package template

var ApiResponse =`package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `+"`"+ `json:"error_code"`+"`"+`
	Msg  string      `+"`"+ `json:"error_message"`+"`"+`
	Data interface{} `+"`"+ `json:"data"`+"`"+`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, errorMsg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  errorMsg,
		Data: data,
	})
	return
}

func (g *Gin) ResoponseSucess(data interface{}){
	g.Response(200,0,"",data)
	return
}

func (g *Gin) ResponnseFailure(error_code int, error_msg string){
	g.Response(200,error_code,error_msg,make(map[string]interface{}))
	return
}

`
