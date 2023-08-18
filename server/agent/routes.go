package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
)

func Setup(group gin.RouterGroup){
  // load agents from the db
  bridge.LoadAgents()

  // setup the auth middleware
  group.Use(AuthMiddleware)

  // setup the routes
  group.POST("/register", AuthRegister) 
  group.POST("/result", JobRes) 
  group.GET("/job", GetJob) 
}
