package service

import (
	"IM/define"
	"IM/models"

	//"IM/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	//"github.com/gogo/protobuf/test/castvalue"
	"github.com/gorilla/websocket"
)

var upgrader=websocket.Upgrader{}
var wc=make(map[string]*websocket.Conn)
func WebsocketMessage(c *gin.Context)  {
	log.Println("come IN")
	conn,err:=upgrader.Upgrade(c.Writer,c.Request,nil)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"系统异常"+err.Error(),
		})
	}
	defer conn.Close()
	//uc:=c.MustGet("user_claims").(*helper.UserClaims)
	randName:=c.MustGet("userName").(string)
	log.Println(randName)
	wc[randName]=conn
	for{
		ms:=new(define.MessageStruct)
		err:=conn.ReadJSON(ms)
		if err!=nil{
			log.Printf("Read Error:%v\n",err)
			return
		}
		//TODO:判断用户是否属于消息体的房间
		_,err=models.GetUserRoomByUserIdentityRoomIdntity(randName)
		if err!=nil{
			log.Printf("没有找到useridentity相关信息:%v\n",err)
			return
		}
		//TODO:保存消息
		//TODO:获取特定房间的在线用户
	//	userRooms,err:=models.GetUserRoomByRoomidentity(randName)
	//	if err!=nil{
	//		log.Printf("[DB ERROR]:%v\n",err)
	//		return
	//	}
	//	//遍历所有的user_room
	//	for i,v:=range userRooms{
	//	log.Println(i,v)
	//	}
	//	for _,cc:=range wc{
	//		err:=cc.WriteMessage(websocket.TextMessage,[]byte(ms.Message))
	//		if err!=nil{
	//			log.Printf("writer Send Error:%v\n",err)
	//			return
	//		}
	//	}
	//}
		userRooms,err:=models.GetUserRoomByRoomidentity(randName)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			return
		}
		for _, room := range userRooms {
			if cc, ok := wc[room.UserIdentity]; ok {
				err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					log.Printf("Write Message Error:%v\n", err)
					return
				}
			}
		}
	}
}

