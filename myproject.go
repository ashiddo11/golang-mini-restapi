package main

import (
	"net/http"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main(){
	name := "api"
	db, err = sql.Open("mysql", "myuser:mypass@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+name)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE "+name)
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}


	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
