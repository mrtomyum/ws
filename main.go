package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.Run(":9999")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	go func() {
			for {
				t, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}
				msg, err = json.Marshal(gin.H{"message":"Hello"})
				if err != nil {
					log.Print("Eror Marshal gin.H")
				}
				conn.WriteMessage(t, msg)
			}
	}()
}
