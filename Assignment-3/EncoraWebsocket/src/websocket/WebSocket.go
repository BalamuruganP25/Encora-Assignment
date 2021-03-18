package websocket

import (
	//INC "EncoraWebsocket/src/messages"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
)

/*Read , write the websocket message */

func Read_Websocket_Message(conn *websocket.Conn) ([]byte, int) {

	messageType, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)

	}

	return message, messageType

}

func Write_websocket_message(conn *websocket.Conn, msg_type int, response_msg interface{}) {
	response, _ := json.Marshal(response_msg)
	if err := conn.WriteMessage(msg_type, response); err != nil {
		fmt.Println("write error -->", err)
		return
	}
	fmt.Println("from message sender")

}
