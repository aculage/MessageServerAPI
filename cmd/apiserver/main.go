//This is program start point
package main

import (
	mservapi "MessageServer/cmd/internal/app/apiserver"
	//"flag"


	"log"

)

/*var confpath string

func init() {
	flag.StringVar(&confpath, "config-path", "configs/mservapi.toml", "path to cfg file")

}*/

func main() {
	//flag.Parse()

	//server start point
	serverInstance := mservapi.New(mservapi.GetConfig())
	if err := serverInstance.Start(); err != nil {
		log.Fatal(err)
	}
}
