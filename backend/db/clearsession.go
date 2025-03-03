package db

import (
	"github.com/google/uuid"
	"strconv"
)

func (z *Storage) ClearSession(idstr string) (string, error) {
	id, _ := strconv.Atoi(idstr)

	session_id := uuid.NewString()

	sqlcomm := "UPDATE users SET session = ? where id = ? "
	if _, err := z.Dbptr.Query(sqlcomm, session_id, id); err != nil {
		return "", err
	}

	//don't actually return the random overwritten session id to the caller.
	return "null", nil
}
