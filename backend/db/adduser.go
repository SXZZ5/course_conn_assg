package db

import (
	"fmt"
)

func (z *Storage) AddUser(username, pwdhash, salt string) (int, string, error) {

	sqlcomm := "INSERT INTO users (username, pwdhash, salt) VALUES (?, ?, ?)"
	if _, err := z.Dbptr.Exec(sqlcomm, username, pwdhash, salt); err != nil {
		return -1, "", err
	}

	if err := z.createCookieTableForUser(username); err != nil {
		return -1, "", err
	}
	fmt.Println("User Registration successful")

	sqlcomm = "SELECT id from users where username = ?"
	row := z.Dbptr.QueryRow(sqlcomm, username)
	var id int
	if err := row.Scan(&id); err != nil {
		return -1, "", err
	}

	session_id, err := z.SetSession(username)
	return id, session_id, err
}

func (z *Storage) createCookieTableForUser(username string) error {
	tablename := "cookie_" + username
	sqlcomm := "CREATE TABLE IF NOT EXISTS " + tablename + " (sno int not null auto_increment primary key, sessions varchar(250) not null);"
	if _, err := z.Dbptr.Exec(sqlcomm); err != nil {
		return err
	} else {
		return nil
	}
}
