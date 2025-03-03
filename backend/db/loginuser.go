package db

import "fmt"

func (z *Storage) LoginUser(username, password string) (int, string, error) {
	if err := z.MatchPwd(username, password); err != nil {
		fmt.Println(err.Error())
		return -1, "", err
	}

	sqlcomm := "SELECT id FROM users where username = ?"
	row := z.Dbptr.QueryRow(sqlcomm, username)
	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Println(err.Error())
		return -1, "", err
	}
	session_id, err := z.SetSession(username)
	if err != nil {
		fmt.Println(err.Error());
		return -1, "", err
	} else {
		return id, session_id, nil
	}
}
