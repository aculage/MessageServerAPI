//This module is responsible for database connection and management
package mservapi

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
type Storage struct{
	Config *DBConfig
	Db *sql.DB
}

func NewStorage(dbconfig *DBConfig) *Storage {

	return &Storage{
		Config: dbconfig,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.Config.DatabaseURL)
	if err != nil{
		log.Print("Database configuration failure")
		log.Fatal(err)
		return err
	}
	if err := db.Ping(); err != nil {
		log.Print("Database connection failure")
		log.Fatal(err)
		return err
	}
	storage.Db  = db
	log.Print("Database opened successfully")
	return nil
}

func (storage *Storage) Close() error{
	return nil
}