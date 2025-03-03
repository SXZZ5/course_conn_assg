package db

import (
	"github.com/google/uuid"
)

func (z *Storage) SetSession(username string) (string, error) {

	session_id := uuid.NewString()

	sqlcomm := "UPDATE users SET session = ? where username = ? "
	if _, err := z.Dbptr.Query(sqlcomm, session_id, username); err != nil {
		return "", err
	}

	cookietable := "cookie_"+username
	sqlcomm = "INSERT INTO " + cookietable + " (sessions) value (?)"
	if _, err := z.Dbptr.Exec(sqlcomm, session_id); err != nil {
		return "", err
	} 

	return session_id, nil
}
