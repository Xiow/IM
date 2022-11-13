package router

import (
	"IM/middlerwares"
	"IM/service"
	"github.com/gin-gonic/gin"
)

func Router()*gin.Engine  {
	r:=gin.Default()

	r.POST("/login",service.Login)
	//发送验证码
	r.POST("/send/code",service.SendCode)
	r.GET("/websocket/message",middlerwares.AuthCheck1(),service.WebsocketMessage)
	auth:=r.Group("/u",middlerwares.AuthCheck())
	//获取用户详情
	auth.GET("/user/detail",service.UserDetails)
	auth.GET("/websocket/message",service.WebsocketMessage)
	auth.GET("/test",service.Test)
	return r
}
