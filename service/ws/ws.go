package ws

import (
	"gchat-gzh/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var aiChan = make(chan int, 5)
var msgChan = make(chan []string, 20)
var msgChanMap = make(map[string][]string)
var connMap sync.Map

var wsupgreder = websocket.Upgrader{

	HandshakeTimeout: 5 * time.Second,
	//取消跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MsgWsHandler(w http.ResponseWriter, r *http.Request, id string) {

	var conn *websocket.Conn
	conn, err := wsupgreder.Upgrade(w, r, nil)
	if err != nil {
		logger.Log.Errorf("Upgrade ws failed,err:%v", err)
		return
	}

	connMap.Store(id, conn)
	logger.Log.Infof("ws connect success,start receiving messages... ")

	//go func() {
	//	for {
	//		time.Sleep(time.Second * 3)
	//		msgChanMap[id] <- "ping"
	//	}
	//}()
	//发送消息均为异步发送
	go func() {
		for {
			msg := <-msgChan
			iDconn, ok := connMap.Load(msg[0])
			if ok {
				cc, oks := iDconn.(*websocket.Conn)
				if oks {
					err = cc.WriteMessage(websocket.TextMessage, []byte(msg[1]))
					if err != nil {
						logger.Log.Error("write msg error", err)
						return
					}
				}
			}
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
		MsgDeal(data, id)
		<-aiChan
	}
}
