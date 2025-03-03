package db

import (
	"golang.org/x/crypto/bcrypt"
)


func (z *Storage) MatchPwd(username, password string) error {
	salt, err := z.getSalt(username)
	if err != nil {
		return err
	}
	hashedpwd, err := z.getHashedpwd(username)
	if err != nil {
		return err
	}
	fullpwd := password + salt
	err = bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(fullpwd))
	return err
}

func (z *Storage) getSalt(username string) (string, error) {
	sqlcomm := "SELECT (salt) from users where username = ? "
	row := z.Dbptr.QueryRow(sqlcomm, username)
	var salt string
	if err := row.Scan(&salt); err != nil {
		return "", err
	} else {
		return salt, nil
	}
}

func (z *Storage) getHashedpwd(username string) (string, error) {
	sqlcomm := "SELECT (pwdhash) from users where username = ?"
	row := z.Dbptr.QueryRow(sqlcomm, username)
	var pwdhash string
	if err := row.Scan(&pwdhash); err != nil {
		return "", err
	} else {
		return pwdhash, nil
	}
}
