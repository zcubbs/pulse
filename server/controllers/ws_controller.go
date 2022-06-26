package controllers

import (
	"github.com/gofiber/websocket/v2"
	"github.com/zcubbs/pulse/server/utils"
	"log"
)

func HandleWS(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		utils.WsUnregister <- c
		c.Close()
	}()

	// Register the client
	utils.WsRegister <- c

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}

		if messageType == websocket.TextMessage {
			// Broadcast the received message
			utils.WsBroadcast <- string(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
