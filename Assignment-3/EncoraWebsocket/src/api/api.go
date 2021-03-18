package api

import (
	endpoints "EncoraWebsocket/src/api/v1/endpoints"
	config "EncoraWebsocket/src/config"
	"fmt"
	"net/http"
)

/*server start here*/
func ApiServer() {
	fmt.Println("Api server start")
	//One to one chat
	http.HandleFunc("/", endpoints.GetXmlFileDetails)

	http.ListenAndServe(config.Api_config, nil)

}
