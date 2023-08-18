package agent

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
	"github.com/ngn13/viper/server/utils"
)

// checks if a agent and its token exists 
func AuthMiddleware(c *gin.Context) {
  if strings.Contains(c.FullPath(), "register") {
    c.Next()
    return
  }

  uid := c.Query("uid")
  tok := c.Query("token")

  if uid == "" || tok == "" {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Required parameters provided"))
    c.Abort()
    return
  }

  indx := bridge.FindAgent(uid)
  if indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("Agent not found"))
    c.Abort()
    return
  }

  if bridge.Agents[indx].Token != tok {
    c.JSON(http.StatusUnauthorized, utils.ErrorJSON("Invalid token"))
    c.Abort()
    return
  }

  bridge.Agents[indx].Lastcon = utils.GetTime()
  c.Next()
}

// lets an agent to register
func AuthRegister(c *gin.Context) {
  agent := bridge.Agent{}
  c.BindJSON(&agent)

  if agent.Uid == "" || agent.Desc == "" || agent.Token == "" {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Bad JSON"))
    return
  }

  indx := bridge.FindToken(agent.Token)
  if indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("Invalid token"))
    return 
  }

  agent.JobRes = ""
  agent.Lastcon = utils.GetTime() 
  indx = bridge.FindAgent(agent.Uid)
  
  if indx == -1 {
    bridge.Agents = append(bridge.Agents, agent)
  }else {
    bridge.Agents[indx] = agent
  }

  bridge.SaveAgents()  
  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}
