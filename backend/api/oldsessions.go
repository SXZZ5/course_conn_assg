package api

import (
	"assignment/db"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ReturnOldSessions(conn *gin.Context) {
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
	fmt.Println("id:", id, ", sesssion:",session)
	res, err := storage.OldSessionCookies(id, session)
	if err != nil {
		fmt.Println(err.Error())
		conn.JSON(400, gin.H{
			"message": "could not get old cookies for this user",
		})
		return
	}

	conn.JSON(200, gin.H{
		"message": "Fetched old cookies successfully",
		"cookies": res,
	})
}
