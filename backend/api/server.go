package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	svreng *gin.Engine
}

func (svr *Server) Init() {
	svr.svreng = gin.Default()
	svr.svreng.POST("/signup", Signup)
}

func (svr *Server) Run() {
	svr.svreng.Run(":8080")
}
