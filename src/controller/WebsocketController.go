package controller

import (
	"fmt"
	"log"

	// "io"
	// "net"
	"net/http"
	"os"
	// "strconv"
	"time"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"github.com/gorilla/websocket"
	// "golang.org/x/net/websocket"

	// "github.com/cloud-barista/cb-webtool/src/service"
	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"

	modelsocket "github.com/cloud-barista/cb-webtool/src/model/websocket"
)

var SpiderURL = os.Getenv("SPIDER_IP_PORT")
var TumbleBugURL = os.Getenv("TUMBLE_IP_PORT")
var DragonFlyURL = os.Getenv("DRAGONFLY_IP_PORT")
var LadyBugURL = os.Getenv("LADYBUG_IP_PORT")

var retryInterval = os.Getenv("KEEP_ALIVE_INTERVAL")
var checkInterval = 5

// type FrameworkAlive struct {
// 	FrameworkName string `json:"frameworkName"`
// 	IpPort        string `json:"ipPort"`
// 	IsAlive       bool   `json:"isAlive"`
// 	Message       string
// }

// websocket에서 호출이 있을 때마다 go routin - channel을 이용하여 결과를 set 하고
// 변경된 정보가 있으면 return.
// func GetStoredWebSocket(c echo.Context, socketName string) (map[string]FrameworkAlive, bool) {
// 	store := echosession.FromContext(c)
// 	result, ok := store.Get(socketName)

// 	// store에 저장된 게 있으면 push
// 	// 저장된 게 없으면 skip

// 	// map 에 object 를 넣는다.

// 	return result, ok
// }

//

// func HelloNetWebSocket(c echo.Context) error {
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer ws.Close()
// 		for {
// 			// Write
// 			err := websocket.Message.Send(ws, "Hello, Client!")
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}

// 			// Read
// 			msg := ""
// 			err = websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			fmt.Printf("%s\n", msg)
// 		}
// 	}).ServeHTTP(c.Response(), c.Request())
// 	return nil
// }

// Websocket 호출 Test form
func HelloForm(c echo.Context) error {
	fmt.Println("============== HelloForm ===============")

	// ws, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	// conn, wserr := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// if wserr != nil {
	// 	log.Println(wserr)
	// }
	// defer conn.Close()
	//go Echo(conn)

	log.Println("aaa1")
	//go HelloGorillaWebSocket(c)

	return echotemplate.Render(c, http.StatusOK,
		"WebsocketTest", // 파일명
		map[string]interface{}{})
}

// Gorilla WebSocket 호출 Test
var (
	upgrader = websocket.Upgrader{}
)

func HelloGorillaWebSocket(c echo.Context) error {
	log.Println("aaa9")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	log.Println("aaa911")
	for {
		//// Write

		// store에 저장된 keepAlive, mcisStatus, vmState of mcis
		// sendData := map[string]interface{}{
		// 	"event": "res",
		// 	"data":  nil,
		//  }
		// sendData["data"] = objmap["data"]
		// refineSendData, err := json.Marshal(sendData)
		// err = c.WriteMessage(mt, refineSendData)

		// err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Gorilla Client!"))
		// if err != nil {
		// 	c.Logger().Error(err)
		// }

		// //// Read
		// _, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	c.Logger().Error(err)
		// }
		// fmt.Printf("%s\n", msg)

		// TODO : 코딩 준비
		//------- Echo Session에서 변경값 조회하여 있는경우 write
		// if getEchossion.Request 에서 처리완료된 것들이 있으면 모두 꺼낸다.  "SOCKET_DATA"   : McisController.go 에서 set. McisHandler.go 에서 set
		//   => [
		// 			{taskId:MCIS_CREATE, taskKey:mcis01, status:request, time:2020-08-25 11:22:33.45}
		//			,{taskId:MCIS_CREATE, taskKey:mcis03, status:complete, time:2020-08-25 13:14:15.}
		//		]
		// 꺼낸 항목들을 json 객체에 넣는다.
		// write Message 한다.
		// echoSession에 있는 값을 지운다 또는 비운다.(어차피 다시 쓰일 애들이므로)
		//------- 1초에 1번 수행.
		log.Println("aaa912")
		store := echosession.FromContext(c)
		socketDataStore, ok := store.Get("socketdata")
		if ok {
			log.Println("aaa913")
			socketDataMap := socketDataStore.(map[string]modelsocket.WebSocketMessage)
			for key, val := range socketDataMap {
				log.Println("getsocketdata"+key+" ", val)
				websocketMessage := socketDataMap[key]
				//fmt.Println(key, val)
				// const messageType = 1 // TextMessage :1, BinaryMessage : 2, CloseMessage : 8, PingMessage = 0, PongMessage = 10
				if val.Status == "complete" {
					//sendErr := ws.WriteMessage(messageType, val)
					sendErr := ws.WriteJSON(websocketMessage)
					if sendErr == nil {
						delete(socketDataMap, key)
					}
				}
				log.Println("aaa914")
			}
		}
		log.Println("aaa915")
		// // socket의 key 생성 : ns + 구분 + id
		// taskKey := nameSpaceID + "||" + "mcis" + "||" + mcisInfo.Name // TODO : 공통 function으로 뺄 것.
		// websocketMessage := map[string][]model.WebSocketMessage{}
		// websocketMessage.TaskId = "mcis"
		// websocketMessage.TaskKey = taskKey
		// websocketMessage.Status = "request"
		// websocketMessage.ProcessTime = time.Now()
		// socketDataMap.put(taskKey, websocketMessage)

		time.Sleep(time.Second * 1)
	}
}

// Listener 에서 감지된 Data 변경을 UI 로 push
// func GorillaWebSocketPush(c echo.Context) error {
// 	err := c.WriteMessage(ws.TextMessage, {message})
// }

type pushMessage struct {
	pushpush string
}

func Echo(conn *websocket.Conn) {
	for {
		//m := pushMessage{pushpush: "echo"}

		// err := conn.ReadJSON(&m)
		// if err != nil {
		//     fmt.Println("Error reading json.", err)
		// }

		// fmt.Printf("Got message: %#v\n", m)

		// if err = conn.WriteJSON(m); err != nil {
		// 	fmt.Println(err)
		// }

		// m := msg{}
		// err := conn.ReadJSON(&m)

		err := conn.WriteMessage(websocket.TextMessage, []byte("Echo push"))
		if err != nil {
			// c.Logger().Error(err)
			// fmt.Printf(err)
			log.Println(err)
		}
	}
}
