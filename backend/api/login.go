package api

import (
	"assignment/db"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(conn *gin.Context) {
	username := conn.PostForm("username")
	rawpwd := conn.PostForm("password")

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
	}

	id, session_id, err := storage.LoginUser(username, rawpwd)
	if err != nil {
		log.Println("Login Failed")
		conn.JSON(400, gin.H{
			"message": "Login Failed",
		})
		return
	}

	conn.SetCookie("id", strconv.Itoa(id), 7200, "/", "", false, true)
	conn.SetCookie("session", session_id, 7200, "/", "", false, true)

	conn.JSON(200, gin.H{
		"message": "Login Succesful",
	})
}
