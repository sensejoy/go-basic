package main

import (
	"fmt"
	"go-basic/controller"
	"go-basic/lib"
	"os"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	server = gin.New()
	controller.InitRoute(server)
}

func main() {
	if err := server.Run(fmt.Sprintf(":%d", lib.Conf["server"]["port"].(int))); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
