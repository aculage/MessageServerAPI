//This module is responsible for config struct and its behaviour
package mservapi

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

type Config struct {
	BindAddress string `toml:"bind_address"`
}
func NewConfig() *Config{
	return &Config{
		BindAddress: ":9000",
	}
}
func GetConfig() *Config{

	_,err := os.Stat("configs/mservapi.toml") //Try open the cfg file
	//If file read-only or nonexistent:
	if err != nil{
		fmt.Println("Server config file missing or unreachable.\nUsing default config.")

		config := NewConfig() //new config created
		cfg, _ := toml.Marshal(config) //config marshaled into toml format
		_ = os.Mkdir("configs",0755)

		//config written into a new file
		cfgfile, err := os.Create("configs/mservapi.toml")
		if err != nil{
			defer fmt.Println(err)
			panic("Something went wrong with config file, current config will not be saved.")
		}

		cfgfile.Write(cfg) //actual byte[] input into the file
		return config //return ptr to object created
	}
	//If file exists and accessible
		config := Config{}
		cfgfromfile, _ := ioutil.ReadFile("configs/mservapi.toml")
		toml.Unmarshal(cfgfromfile, &config)
		return &config

}