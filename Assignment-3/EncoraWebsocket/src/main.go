package main

import (
	api "EncoraWebsocket/src/api"
	config "EncoraWebsocket/src/config"
	"fmt"

)

func main() {
	fmt.Println("web socket application started")
	config.ReadEnvironmentVariable()
	api.ApiServer()
	

}
