package main

import (
	"IM/router"
)

func main(){
	//fmt.Println(helper.GetMd5("123456"))
	r:=router.Router()
	r.Run(":8080")
}