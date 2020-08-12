//This module is responsible for config structs and their behaviour
package mservapi

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BindAddress string `toml:"bind_address"`
}

type DBConfig struct{
	DatabaseURL string `toml:"database-url"`
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
		log.Println("Server config file missing or unreachable.\nUsing default config.")

		config := NewConfig() //new config created
		cfg, _ := toml.Marshal(config) //config marshaled into toml format
		_ = os.Mkdir("configs",0755)

		//config written into a new file
		cfgfile, err := os.Create("configs/mservapi.toml")
		if err != nil{
			defer log.Println(err)
			panic("Something went wrong with config file, current config will not be saved.")
		}

		cfgfile.Write(cfg) //actual byte[] input into the file
		return config //return ptr to object created
	}
	//If file exists and accessible
		config := NewConfig()
		cfgfromfile, _ := ioutil.ReadFile("configs/mservapi.toml")
		toml.Unmarshal(cfgfromfile, &config)
		return config

}

func NewDBConfig() *DBConfig{
	return &DBConfig{
		DatabaseURL: "user=client password=client host=db port=5432 dbname=mservapi_deb",
	}
}
func GetDBConfig() *DBConfig{

	_,err := os.Stat("configs/db.toml") //Try open the cfg file
	//If file read-only or nonexistent:
	if err != nil{
		log.Println("Database config file missing or unreachable.\nUsing default config.")

		config := NewDBConfig() //new config created
		cfg, _ := toml.Marshal(config) //config marshaled into toml format
		_ = os.Mkdir("configs",0755)

		//config written into a new file
		cfgfile, err := os.Create("configs/db.toml")
		if err != nil{
			defer log.Println(err)
			panic("Something went wrong with database config file, current database config will not be saved.")
		}

		cfgfile.Write(cfg) //actual byte[] input into the file
		return config //return ptr to object created
	}
	//If file exists and accessible
	config := NewDBConfig()
	cfgfromfile, _ := ioutil.ReadFile("configs/db.toml")
	toml.Unmarshal(cfgfromfile, config)
	return config

}