package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//var wsConnections []*websocket.Conn

func WebChat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// panic(err)
		log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//wsConnections = append(wsConnections, conn)
	for {
		// Read message from client
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			// panic(err)
			log.Printf("%s, error while reading message\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			break
		}

		// Echo message back to client
		severMsg := "From Server : " + string(p)
		err = conn.WriteMessage(messageType, []byte(severMsg))
		go responseToWs(conn, messageType)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while writing message\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			break
		}
	}
}

func responseToWs(wsCon *websocket.Conn, mtype int) {
	time.Sleep(10 * time.Second)
	for i := 0; i < 2; i++ {
		severMsg := "From Server : " + strconv.Itoa(i)
		time.Sleep(10 * time.Second)
		wsCon.WriteMessage(mtype, []byte(severMsg))
	}
}
