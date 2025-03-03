package db

import (
	"errors"
	"strconv"
)

func (z *Storage) CheckSession(idstr, session string) error {
	id, _ := strconv.Atoi(idstr)
	sqlcomm := "SELECT (id, session) from users where id=?"
	row := z.Dbptr.QueryRow(sqlcomm, id)

	var (
		fetched_id      int
		fetched_session string
	)
	if err := row.Scan(&fetched_id, &fetched_session); err != nil {
		return err
	}
	fetched_idstr := strconv.Itoa(fetched_id)
	if idstr == fetched_idstr && fetched_session == session {
		return nil
	} else {
		return errors.New("session not valid")
	}
}
