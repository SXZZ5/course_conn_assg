package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Storage struct {
	Dbptr *sql.DB
}

func (z *Storage) Connect() {
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
	z.Dbptr = db
	if pingerr := z.Dbptr.Ping(); pingerr != nil {
		panic(pingerr)
	}
	log.Printf("db connection successful")
}

type entry struct {
	Name string
}

func (z *Storage) PrintAll() {
	sqlcomm := `
	SELECT username FROM users LIMIT 20
	`
	rows, err := z.Dbptr.Query(sqlcomm)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		var entryy entry
		if err := rows.Scan(&entryy.Name); err != nil {
			log.Println("row scan error:", err.Error())
		} else {
			log.Println(entryy)
		}
	}
}
