package middlerwares

import (
	"IM/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AuthCheck() gin.HandlerFunc{

	return func(c *gin.Context) {
		token:=c.GetHeader("token")
		connection:=c.GetHeader("connection")
		log.Println(connection)
		userClaims,err:=helper.AnalyseToken(token)
		if err!=nil{
			c.Abort()
			c.JSON(http.StatusOK,gin.H{
				"code":-1,
				"msg":"用户认证不通过",
			})
			return
		}
		c.Set("user_claims",userClaims)
		c.Set("userName",helper.RandName())
		c.Next()
	}
}


func AuthCheck1() gin.HandlerFunc{

	return func(c *gin.Context) {
		//token:=c.GetHeader("token")
		//connection:=c.GetHeader("connection")
		//log.Println(connection)
		//userClaims,err:=helper.AnalyseToken(token)
		//if err!=nil{
		//	c.Abort()
		//	c.JSON(http.StatusOK,gin.H{
		//		"code":-1,
		//		"msg":"用户认证不通过",
		//	})
		//	return
		//}
		//c.Set("user_claims",userClaims)
		ID:=helper.GenerateRandInt(0,10)
		c.Set("userName","Xiaowei:"+strconv.Itoa(ID))
		c.Next()
	}
}