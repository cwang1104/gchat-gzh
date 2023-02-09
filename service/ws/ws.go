package ws

import (
	"gchat-gzh/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var aiChan = make(chan int, 1)
var msgChan = make(chan string, 10)

var wsupgreder = websocket.Upgrader{

	HandshakeTimeout: 5 * time.Second,
	//取消跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MsgWsHandler(w http.ResponseWriter, r *http.Request) {

	var conn *websocket.Conn
	conn, err := wsupgreder.Upgrade(w, r, nil)
	if err != nil {
		logger.Log.Errorf("Upgrade ws failed,err:%v", err)
		return
	}

	logger.Log.Infof("ws connect success,start receiving messages... ")

	//发送消息均为异步发送
	go func() {
		for {
			_ = conn.WriteMessage(websocket.TextMessage, []byte(<-msgChan))
		}
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			logger.Log.Errorf("read msg failed,err:%v", err)
			return
		}
		aiChan <- 1
		//消息处理
		MsgDeal(data)
		<-aiChan
	}
}
