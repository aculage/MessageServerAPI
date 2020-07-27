package main

import (
	mservapi "MessageServer/cmd/internal/app/apiserver"
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
)

var confpath string

func init() {
	flag.StringVar(&confpath, "config-path", "configs/mservapi.toml", "path to cfg file")

}

func main() {
	flag.Parse()
	config := mservapi.NewConfig()

	cfg, _ := toml.Marshal(config)
	fmt.Print(string(cfg))
	fmt.Fprintf()
	fmt.Print(err)

	//server start point
	serverInstance := mservapi.New(config)
	if err := serverInstance.Start(); err != nil {
		log.Fatal(err)
	}
}
