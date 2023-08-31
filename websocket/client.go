package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
	"z0nix007/service"

	"github.com/gorilla/websocket"
)

func main() {
	connectWebSocket()
	select {}
	// Wait forever
}

func connectWebSocket() {
	service.Session = service.Login()
	service.Token = service.GetToken(service.Session)

	go func ()  {
		for {
			service.Session = service.Login()
			service.Token = service.GetToken(service.Session)
			time.Sleep(5 * time.Minute)
		}
	}()

	u := url.URL{Scheme: "wss", Host: "msg-api-th.fleet.vip", Path: "/connection/websocket"}
	fmt.Printf("Connecting to %s\n", u.String())

	for {
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Println("Error connecting to WebSocket:", err)
			connectWebSocket()
			continue
		}
		defer c.Close()

		// Send the first message
		// สร้างข้อความที่ต้องการส่ง
		connectData := map[string]interface{}{
			"connect": map[string]string{
				"token": service.Token,
				"name":  "js",
			},
			"id": 1,
		}

		subscribeData := map[string]interface{}{
			"subscribe": map[string]string{
				"channel": "fms_grab_order",
			},
			"id": 2,
		}

		// แปลงข้อมูล JSON object เป็น JSON string
		connectMessage, err := json.Marshal(connectData)
		if err != nil {
			log.Println("Error marshaling 'connect' message:", err)
			return
		}

		subscribeMessage, err := json.Marshal(subscribeData)
		if err != nil {
			log.Println("Error marshaling 'subscribe' message:", err)
			return
		}

		// ส่งข้อความที่คุณสร้างไปยังเซิร์ฟเวอร์ WebSocket
		err = c.WriteMessage(websocket.TextMessage, connectMessage)
		if err != nil {
			log.Println("Error sending 'connect' message:", err)
			return
		}

		err = c.WriteMessage(websocket.TextMessage, subscribeMessage)
		if err != nil {
			log.Println("Error sending 'subscribe' message:", err)
			return
		}

		readMessages(c)
		// If the connection is closed, attempt to reconnect
		log.Println("WebSocket connection closed. Reconnecting...")
	}

}

func readMessages(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			connectWebSocket()
			return
		}
		log.Println(string(message))
		if string(message) == "{}" {
			connectMessage := []byte(`{}`)
			err = c.WriteMessage(websocket.TextMessage, connectMessage)
			if err != nil {
				log.Println("Error sending 'connect' message:", err)
				return
			}
		} else {
			service.GetJob()
		}
		// Check if the received message is "{}" and replace it with "{]"
	}
}
