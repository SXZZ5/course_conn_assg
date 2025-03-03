package api

import (
	"assignment/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Logout(conn *gin.Context) {
	tmp, boolerr := conn.Get("storage")
	if boolerr == false {
		log.Println("looks like database dependency injection failed")
	}
	storage, ok := tmp.(*db.Storage)
	if !ok {
		log.Println("Failed to convert storage to *db.Storage")
		conn.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	id, err := conn.Cookie("id")
	if err != nil {
		fmt.Println("id cookie", err.Error())
		conn.JSON(400, gin.H{
			"message": "cookie not valid",
		})
		return
	}
	session, err := conn.Cookie("session")
	if err != nil {
		fmt.Println("session cookie", err.Error())
		conn.JSON(400, gin.H{
			"message": "cookie not valid",
		})
		return
	}

	err = storage.CheckSession(id, session)
	if err != nil {
		fmt.Println(err.Error())
		conn.JSON(400, gin.H{
			"message": "cookies not valid",
		})
		return
	}

	_, clearSessionErr := storage.ClearSession(id)
	if clearSessionErr != nil {
		fmt.Println(clearSessionErr.Error())
		conn.JSON(400, gin.H{
			"message": "Could Not Log out",
		})
		return
	}

	conn.SetCookie("session", "null", 7200, "/", "", false, true)
	conn.JSON(200, gin.H{
		"message": "Logged Out successfully",
	})
	return
}
