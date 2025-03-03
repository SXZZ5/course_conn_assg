package api

import (
	"github.com/gin-gonic/gin"
	"assignment/db"
	"github.com/gin-contrib/cors"
)

type Server struct {
	svreng *gin.Engine
}

func (svr *Server) RoutesInit(storage *db.Storage) {
	svr.svreng = gin.Default()
	svr.svreng.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))
	svr.svreng.Use(InjectDB(storage))
	svr.svreng.POST("/signup", Signup)
	svr.svreng.POST("/login", Login)
	svr.svreng.GET("/oldcookies", ReturnOldSessions)
	svr.svreng.GET("/logout", Logout)
}


func (svr *Server) Run() {
	svr.svreng.Run(":8080")
}

func InjectDB(storage *db.Storage) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("storage", storage)
        c.Next()
    }
}
