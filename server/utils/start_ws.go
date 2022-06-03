package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"github.com/zcubbs/pulse/models"
	"log"
)

type WsClient struct{} // Add more data to this type if needed

var WsClients = make(map[*websocket.Conn]WsClient) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var WsRegister = make(chan *websocket.Conn)
var WsBroadcast = make(chan string)
var WsUnregister = make(chan *websocket.Conn)

func RunHub() {
	for {
		select {
		case connection := <-WsRegister:
			WsClients[connection] = WsClient{}

		case message := <-WsBroadcast:
			unmarshalledMsg := &models.WSMessage{}
			json.Unmarshal([]byte(message), unmarshalledMsg)
			WriteMessage(unmarshalledMsg)

		case connection := <-WsUnregister:
			// Remove the WsClient from the hub
			delete(WsClients, connection)
		}
	}
}

func WriteMessage(message *models.WSMessage) {
	mMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send the message to all WsClients
	for connection := range WsClients {
		if err := connection.WriteMessage(websocket.TextMessage, mMessage); err != nil {
			log.Println("write error:", err)

			connection.WriteMessage(websocket.CloseMessage, []byte{})
			connection.Close()
			delete(WsClients, connection)
		}
	}
}
