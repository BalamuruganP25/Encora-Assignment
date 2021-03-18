package endpoints

import (
	socket "EncoraWebsocket/src/websocket"
	utils "EncoraWebsocket/src/utils"
	//"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	
)

// set buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1000000,
	WriteBufferSize: 1000000,
	CheckOrigin:     func(r *http.Request) bool { return true },
}



func GetXmlFileDetails(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	_, types := socket.Read_Websocket_Message(ws)

	 xmldata := utils.ParsingXml()

	socket.Write_websocket_message(ws, types,xmldata)
}
