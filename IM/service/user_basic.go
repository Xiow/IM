package service

import (
	"IM/helper"
	"IM/models"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func Login(c *gin.Context){
	account:=c.PostForm("account")
	password:=c.PostForm("password")
	if account==""||password==""{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"用户名或者密码为空",
		})
		return
	}
	//ub,err:=models.GetUserBasicBYAccountPassword(account,password)

	ub,err:=models.GetUserBasicBYAccountPassword(account,helper.GetMd5(password))

	log.Println(ub)

	log.Println(err)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"用户名或密码错误",
		})
		return
	}
	//生成token
	token,err:=helper.GenerateToken(ub.Identity,ub.Email)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"系统错误"+err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"登录成功",
		"data":gin.H{
			"token":token,
		},

	})
}
func UserDetails(c *gin.Context){
	u,_:=c.Get("user_claims")
	uc:=u.(*helper.UserClaims)
	log.Println(uc)
	userbasic,err:=models.GetUserBasicByIdentity(uc.Identity)
	if err!=nil{
		log.Printf("[DB error]:%v\n",err)
	c.JSON(http.StatusOK,gin.H{
		"code":-1,
		"msg":"数据查询异常",
	})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":userbasic,
	})
}
func SendCode(c *gin.Context){
	email:=c.PostForm("email")
	if email==""{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"邮箱不能为空",
		})
		return
	}
	cnt,err:=models.GetUserBasicCountByEmail(email)
	if err!=nil{
		log.Printf("[DB ERROR]:%v\n",err)
		return
	}
	if cnt>0{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"当前邮箱已被注册",
		})
		return
	}
	err=helper.MailSendCode(email,"一起浪迹天涯!")
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"系统错误,请稍后再试",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"验证码发送成功",
	})
}
func Test(c *gin.Context){
	uc:=c.MustGet("user_claims").(*helper.UserClaims)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":uc,
	})
}