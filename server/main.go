/*
 *
 *    Viper | Command & Control Server
 *    ############################################
 *    HTTP/HTTPS C2 server for viper toolkit,
 *    see README.md for build instructions
 *
 *    This project is licensed under GNU Public License Version 2,
 *    please see LICENSE.txt
 *
 *    Written by ngn
 *
 */

package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/agent"
	"github.com/ngn13/viper/server/operator"
	"github.com/ngn13/viper/server/utils"
)

func main(){
  gin.SetMode(gin.ReleaseMode)

  // read config, seed random, init database
  utils.RandInit()
  utils.ReadConfig()
  utils.InitDatabase()

  // creating the default router
  app := gin.Default()
  
  // setup groups
  app.Use(func(c *gin.Context){
    c.Header("Server", utils.Config.ServerHeader)
  })
  topGroup    := app.Group(utils.Config.Rootpath)
  agentGroup  := topGroup.Group("/agent")
  opGroup     := topGroup.Group("/operator")
  agent.Setup(*agentGroup)
  operator.Setup(*opGroup)
  app.NoRoute(func (c *gin.Context){
    raw, err := ioutil.ReadFile(utils.Config.Page404)
    if err != nil {
      c.Data(http.StatusNotFound, "text/html", []byte("<h1>404 Page Not Found</h1>"))
      return
    }
    c.Data(http.StatusNotFound, "text/html", raw)
  }) 

  // starting the application on port 8080
  log.Printf("Starting the C2 server at address: %s", utils.Config.Address)
  app.Run(utils.Config.Address)
}
