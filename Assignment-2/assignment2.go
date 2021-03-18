package main

import (
	"log"
	"time"
)


func main() {

	for v := 0; v < 10; v++ {
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}

	time.Sleep(time.Second * 10)
}

func callDouble(v int) int {
	
	//Adjust code to call double only up to 5 times concurrently
	if v < 5 {
	return double(v)
	}
	
	return 0
	
}

func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}