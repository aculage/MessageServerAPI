//This is the server module itself
package mservapi

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//ApiServer
type APIServer struct{
	config *Config
	router *mux.Router
	storage *Storage
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
func New(config *Config, dbconfig *DBConfig) *APIServer{
	f,_ := toml.Marshal(config)
	db,_ := toml.Marshal(dbconfig)
	log.Println("Starting server.\nCurrent config:\n" + string(f) + "\nCurrent database config:\n" + string(db))
	return &APIServer{
		config:  config,
		router:  mux.NewRouter(),
		storage: NewStorage(dbconfig),
	}
}

func (server *APIServer) Start() error{
	server.configureRouter()
	server.storage.Open()
	log.Fatal(http.ListenAndServe(server.config.BindAddress,server.router))
	return nil
}
func (server *APIServer) configureRouter(){
	server.router.HandleFunc("/users/add", server.userAdd)
	server.router.HandleFunc("/chats/add", server.chatAdd)
	server.router.HandleFunc("/messages/add", server.messageAdd)
	server.router.HandleFunc("/chats/get", server.chatGet)
	server.router.HandleFunc("/messages/get", server.messageGet)
}

 //TODO rewrite functions into post responsive
func (server *APIServer) userAdd(w http.ResponseWriter, r *http.Request){
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
	//solved: unique constraint added :: check for existence of the same user in database
	//solved: db.exec :: push user into database
	_,err = server.storage.Db.Exec("INSERT INTO users VALUES($1,$2,$3)", u.Id, u.Username, u.Created_at)
	if err != nil{
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	ujson, _ := json.MarshalIndent(u,""," ") //make json output look readable
	log.Print("User added:\n", string(ujson)) //logging user addition
	w.Write([]byte(u.Id.String()))
	return


}

func (server *APIServer) chatAdd(w http.ResponseWriter, r *http.Request){
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
	//solved::check for users to exist
	userExists := true
	for _,u:= range c.Users{
		res,_ :=server.storage.Db.Query("SELECT EXISTS(SELECT id FROM users WHERE id = $1)", u)
		user_found:=res.Next()
		if !user_found {
			//TODO return http code
			userExists = false
			log.Printf("User Id: $1 does not exist\n",u)
			w.Write([]byte("User Id: "+ u.String() + " does not exist\n"))
		}
	}
	if !userExists{
		w.WriteHeader(422)
		return
	}

	c.Created_at = time.Now()
	c.Id = uuid.New()
	cjson, _ := json.MarshalIndent(c,""," ") //make json output look readable
	//TODO check for chat existence
	//solved:: push chat into database
	_,err = server.storage.Db.Exec("INSERT INTO chats VALUES($1,$2,$3,$4)",c.Id,c.Name,pq.Array(c.Users),c.Created_at)
	if err != nil{
		 log.Println(err)
		 w.WriteHeader(500)
		 w.Write([]byte(err.Error()))
	}
	log.Print("Chat created:\n", string(cjson))//logging chat creation
	w.Write([]byte(c.Id.String()))

}
func (server *APIServer) messageAdd(w http.ResponseWriter, r *http.Request){
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

		res,_ :=server.storage.Db.Query("SELECT EXISTS(SELECT id FROM users WHERE id = $1)", m.Chat)

		if !res.Next() {
			//TODO return http code
			log.Printf("Chat Id: $1 does not exist\n",m.Chat)
			w.Write([]byte("Chat Id: "+ m.Chat.String() + " does not exist\n"))
		}

	//TODO check if user is in the chat
		res,_=server.storage.Db.Query("SELECT users FROM chats WHERE chats.id = $1", m.Chat)
	m.Created_at = time.Now()
	m.Id = uuid.New()
	mjson, _ := json.MarshalIndent(m,""," ") //make json output look readable

	//TODO push message into a chat
	log.Print("Message sent:\n", string(mjson)) //logging message

}

func (server *APIServer) chatGet(w http.ResponseWriter, r *http.Request){
	return
}

func (server *APIServer) messageGet(w http.ResponseWriter, r *http.Request) {
	return
}


