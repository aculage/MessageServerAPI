//This is the server module itself
package mservapi

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/google/uuid"
)

//ApiServer
type APIServer struct{
	config *Config
}

type User struct{
	Username string
	Id uuid.UUID
	Created_at time.Time
}

type Chat struct{
	Id uuid.UUID
	Name string
	Users []User
	Created_at time.Time
}

type Message struct{
	Id int64
	Chat *rune
	Author *uuid.UUID
	Text string
	Created_at time.Time
}

//New server
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("userAdd: %s \n ",string(body))
	var u User
	err = json.Unmarshal(body, &u)
	u.Created_at = time.Now()
	u.Id = uuid.New()
	//todo check for existence of the same user in database
	if err != nil{
		log.Fatal(err)
	}
	ujson, _ := json.MarshalIndent(u,""," ")
	log.Print("User added:\n", string(ujson))
	w.Write([]byte(u.Id.String()))



}
func chatAdd(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
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


