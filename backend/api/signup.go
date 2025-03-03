package api

import (
	"assignment/db"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(conn *gin.Context) {
	username := conn.PostForm("username")
	pwd := conn.PostForm("password")
	salt := rand.Int()
	pwd += strconv.Itoa(salt)
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 4)
	if err != nil {
		fmt.Println("err in hashing", err.Error())
	} else {
		fmt.Println("hashed password is:", hashedpwd)
	}
	fmt.Print("username", username)
	fmt.Print("pwd:", pwd)

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
	id, session, err := storage.AddUser(username, string(hashedpwd), strconv.Itoa(salt))
	if err != nil {
		panic(err.Error())
	}

	conn.SetCookie("id", strconv.Itoa(id), 7200, "/", "", false, true)
	conn.SetCookie("session", session, 7200, "/","",false,true)

	conn.JSON(http.StatusOK, gin.H{
		"message": "your hashed password is: " + string(hashedpwd),
	})
}
