package api

import (
	"fmt"
	"net/http"

	"math/rand"
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
	conn.JSON(http.StatusOK, gin.H{
		"message": "your hashed password is: " + string(hashedpwd),
	})
}
