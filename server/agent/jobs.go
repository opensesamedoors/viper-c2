package agent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
	"github.com/ngn13/viper/server/utils"
)

type JobResult struct {
  Done      bool     `json:"done"`
  Result    string   `json:"result"`
}

// saves a job result for an agent 
func JobRes(c *gin.Context){
  indx := bridge.FindJob(c.Query("uid"))
  if indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("Job not found"))
    return
  }

  res := JobResult{}
  res.Done = true

  if c.BindJSON(&res) != nil {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Bad JSON"))
    return
  }
  
  if res.Done {
    bridge.Jobs[indx].Done = true
  }

  indx = bridge.FindAgent(c.Query("uid"))
  bridge.Agents[indx].JobRes = res.Result
  bridge.SaveAgents()

  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}

// returns the first job for the agent
func GetJob(c *gin.Context){
  indx := bridge.FindJob(c.Query("uid"))
  if indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("No avaliable job for the agent"))
    return
  }

  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": bridge.Jobs[indx],
  })
} 
