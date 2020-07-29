//This is the server module itself
package mservapi

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"net/http"
	"time"
)

//ApiServer
type APIServer struct{
	config *Config
}

type User struct{
	id rune
	username string
	created_at time.Time
}

type Chat struct{
	id rune
	name string
	users []User
	created_at time.Time
}

type Message struct{
	id int64
	chat *rune
	author *rune
	text string
	created_at time.Time
}

//New ...
func New(config *Config) *APIServer{
	f,_ := toml.Marshal(config)
	fmt.Println("Starting server.\nCurrent config:\n" + string(f))
	return &APIServer{
		config: config,
	}
}

func (server *APIServer) Start() error{


	http.HandleFunc("/users/add", userAdd)
	http.HandleFunc("/chats/add", chatAdd)
	http.HandleFunc("/messages/add", messageAdd)
	http.HandleFunc("/chats/get", chatGet)
	http.HandleFunc("/messages/get", messageGet)
	log.Fatal(http.ListenAndServe(server.config.BindAddress,nil))
	return nil
}

 //TODO rewrite functions into post responsive
func userAdd(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "userAdd")
}
func chatAdd(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "chatAdd")
}
func messageAdd(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "messageAdd")
}
func chatGet(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "chatGet")
}
func messageGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "messageGet")
}


