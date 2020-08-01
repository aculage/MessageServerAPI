//This is program start point
package main

import (
	mservapi "MessageServer/cmd/internal/app/apiserver"
	"log"
)


func main() {
	//server start point
	serverInstance := mservapi.New(mservapi.GetConfig())
	if err := serverInstance.Start(); err != nil {
		log.Fatal(err)
	}
}
