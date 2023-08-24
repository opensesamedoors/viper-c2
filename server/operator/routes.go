package operator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
	"github.com/ngn13/viper/server/utils"
)

func Version(c *gin.Context){
  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": utils.Version, 
  })
}

func Setup(router gin.RouterGroup){
  // load tokens from the db
  bridge.LoadTokens()

  // authentication middleware and the version route
  router.Use(AuthMiddleware)
  router.GET("/version",    Version)

  // auth routes
  router.POST("/login",     AuthLogin) 
  router.GET("/logout",     AuthLogout) 
  router.GET("/operators",  AuthList) 
  
  // agent routes
  router.POST("/job",       AddJob) 
  router.GET("/results",    JobResults) 
  router.GET("/jobs",       ListJobs) 
  router.POST("/tokens",    AddToken)
  router.GET("/tokens",     GetTokens)
  router.DELETE("/tokens",  DelToken)
  router.GET("/agents",     GetAgents) 
} 
