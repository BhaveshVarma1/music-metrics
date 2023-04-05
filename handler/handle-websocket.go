package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var connections []*websocket.Conn

func HandleWebsocket(c echo.Context) error {

	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	addConnection(conn)

	// Handle WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("WebSocket error: %v", err)
			}
			removeConnection(conn)
			break
		}

		// Handle the received message
		fmt.Println(string(message))
		fmt.Println("Connections: ", len(connections))

		// Send response (for testing)
		/*err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Wassap homes. This bitch sent from main.go")))
		if err != nil {
			fmt.Println("Error encountered while sending message.")
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("WebSocket error: %v", err)
			}
			removeConnection(conn)
			break
		}*/
	}

	return nil
}

func addConnection(conn *websocket.Conn) {
	connections = append(connections, conn)
	notifyClients(len(connections))
}

func removeConnection(conn *websocket.Conn) {
	for i, c := range connections {
		if c == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
	notifyClients(len(connections))
}

func notifyClients(count int) {
	for _, conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Total connections: %d", count)))
		if err != nil {
			fmt.Println("Error encountered while sending message.")
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("WebSocket error: %v", err)
			}
			removeConnection(conn)
			break
		}
	}
}
