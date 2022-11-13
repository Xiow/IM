package test

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var ws=make(map[*websocket.Conn]struct{})
var upgrader = websocket.Upgrader{} // use default options
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ws[c]= struct{}{}
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		for conn:=range ws{
			err = conn.WriteMessage(mt,message )
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
func TestWebsocket(t *testing.T){
	//flag.Parse()
	//log.SetFlags(0)
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func TestGinWebsocket(t *testing.T){
	r:=gin.Default()
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer,ctx.Request)
	})
	r.Run(":8080")
}