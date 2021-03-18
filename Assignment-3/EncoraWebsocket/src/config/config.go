package config

import (
	"fmt"
	"os"
)

type Apicredentials struct {
	Ip   string
	Port string
}

var Apidetails Apicredentials

var Api_config string

func ReadEnvironmentVariable() {

	_, ApiIpStatus := os.LookupEnv("API_IP")
	if !ApiIpStatus {
		fmt.Println("api ip  Not found in Environment variable file")
		os.Exit(3)
	} else {
		Apidetails.Ip = os.Getenv("API_IP")
		fmt.Println("api ip :", Apidetails.Ip)
	}
	_, ApiPortStatus := os.LookupEnv("API_PORT")
	if !ApiPortStatus {
		fmt.Println("api port  Not found in Environment variable file")
		os.Exit(3)
	} else {
		Apidetails.Port = os.Getenv("API_PORT")
		fmt.Println("api port :", Apidetails.Port)
	}
	Api_config = Apidetails.Ip + ":" + Apidetails.Port
	fmt.Println("Api_config --->", Api_config)
}
