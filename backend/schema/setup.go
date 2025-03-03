package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	sqlcomm := `
	CREATE TABLE IF NOT EXISTS users (
		id int not null auto_increment primary key,
    	username varchar(250) not null unique,
    	pwdhash varchar(250) not null,
    	salt varchar(250) not null,
    	session varchar(250)
	);
	`

	cfg := mysql.Config{
		User:   "root",
		Passwd: "jdleegrb",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "regsys",
	}
	dsnstring := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsnstring)
	if err != nil {
		panic(err)
	}

	if pingerr := db.Ping(); pingerr != nil {
		panic(pingerr)
	}

	_, err = db.Exec(sqlcomm)
	if err != nil {
		log.Fatalln("could not create the table")
	}
}
