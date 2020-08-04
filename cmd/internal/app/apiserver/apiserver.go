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
	Users []uuid.UUID
	Created_at time.Time
}

type Message struct{
	Id uuid.UUID
	Chat uuid.UUID
	Author uuid.UUID
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
		log.Print(err)
	}
	//now checking if request body is actually a valid username input
	log.Printf("userAdd: %s \n ",string(body))
	var u User
	//unmarshalling json body into a temp struct
	err = json.Unmarshal(body, &u)
	//if request body empty or invalid
	if (err != nil || u.Username == "") {
		w.WriteHeader(422)
		w.Write([]byte("Data unprocessable -> check your input\n"))
		log.Print(err)
		return
	}
	u.Created_at = time.Now()
	u.Id = uuid.New()
	//TODO check for existence of the same user in database

	ujson, _ := json.MarshalIndent(u,""," ") //make json output look readable

	//TODO push user into database
	log.Print("User added:\n", string(ujson)) //logging user addition
	w.Write([]byte(u.Id.String()))
	return


}

func chatAdd(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chatAdd: %s \n ",string(body))

	var c Chat
	err = json.Unmarshal(body, &c)
	//TODO make validator
	if err != nil{
		log.Print(err)
	}
	//TODO check for users to exist
	c.Created_at = time.Now()
	c.Id = uuid.New()
	cjson, _ := json.MarshalIndent(c,""," ") //make json output look readable

	//TODO push chat into database
	log.Print("Chat created:\n", string(cjson)) //logging chat creation


}
func messageAdd(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("messageAdd: %s \n ",string(body))

	var m Message
	err = json.Unmarshal(body, &m)
	//TODO make validator -- hint: if anything before text is invalid, text will be empty
	if err != nil{
		log.Print(err)
	}
	//TODO check if chat exists
	//TODO check if user is in the chat
	m.Created_at = time.Now()
	m.Id = uuid.New()
	mjson, _ := json.MarshalIndent(m,""," ") //make json output look readable

	//TODO push message into a chat
	log.Print("Message sent:\n", string(mjson)) //logging message

}

func chatGet(w http.ResponseWriter, r *http.Request){
	
}

func messageGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "messageGet")
}


