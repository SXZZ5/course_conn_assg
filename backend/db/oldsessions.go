package db

import (
	"errors"
	"fmt"
	"strconv"
)

func (z *Storage) OldSessionCookies(idstr, session string) ([]string, error) {
	fmt.Println("inside storage.OldSessionCookies()")
	fmt.Println("idstr:", idstr, " session:", session)
	id, _ := strconv.Atoi(idstr)
	sqlcomm := "SELECT username, session from users where id=?"
	row := z.Dbptr.QueryRow(sqlcomm, id)
	var fsession string
	var username string
	if err := row.Scan(&username, &fsession); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("retrieved username:", username, "sesssion:", fsession)
	if fsession != session {
		fmt.Println("sessions did not match")
		return nil, errors.New("no such session")
	}
	fmt.Println(fsession, session)

	cookietable := "cookie_" + username
	fmt.Println(cookietable)
	sqlcomm = "SELECT (sessions) from " + cookietable
	rows, err := z.Dbptr.Query(sqlcomm)
	if err != nil {
		fmt.Println("user's cookie table not found")
		return nil, err
	}

	res := []string{}
	for rows.Next() {
		var onesession string
		if err := rows.Scan(&onesession); err != nil {
			continue
		}
		res = append(res, onesession)
	}
	fmt.Println("res:", res)
	return res, nil
}
