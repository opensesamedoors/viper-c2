package bridge

import "github.com/ngn13/viper/server/utils"

// ############ AGENTS ############
var Agents []Agent = []Agent{}
type Agent struct {
	Uid      string `json:"uid"`
	Desc     string `json:"desc"`
	Token    string `json:"token"`
	Message  string `json:"message"`
  JobRes   string `json:"jobres"`
	Lastcon  string `json:"lastcon"`
}

type AgentsJson struct {
  List     []Agent `json:"list"`
}

// saves all the agents to the database
func SaveAgents() {
  agentsjson := AgentsJson{}
  agentsjson.List = Agents
  utils.SaveToDatabase("agents.json", agentsjson)
}

// loads agents from the database
func LoadAgents() {
  agentsjson := AgentsJson{}
  if !utils.LoadFromDatabase("agents.json", &agentsjson) {
    agentsjson.List = []Agent{}
  }
  Agents = agentsjson.List
}

// finds an agent by its uid
func FindAgent(uid string) int {
  for i := range Agents {
    if Agents[i].Uid == uid {
      return i
    }
  }

  return -1
}

// ############ TOKENS ############ 
var Tokens []string = []string{}
type TokensJson struct{
  List  []string  `json:"list"`
}

func FindToken(tok string) int{
  for i := range Tokens {
    if tok == Tokens[i] {
      return i
    }
  }

  return -1
}

func SaveTokens() {
  tokenjson := TokensJson{}
  tokenjson.List = Tokens
  utils.SaveToDatabase("tokens.json", tokenjson)
}

func LoadTokens() {
  tokensjson := TokensJson{}
  if !utils.LoadFromDatabase("tokens.json", &tokensjson){
    tokensjson.List = []string{}
  }
  Tokens = tokensjson.List
}

// ############ JOBS ############ 
var Jobs []JobData = []JobData{}
type JobData struct{
  Action  string  `json:"action"`
  Data    string  `json:"data"`
  Target  string  `json:"target"`
  Done    bool    `json:"done"`
}

// finds jobs by target id
func FindJob(target string) int{
  for i := range Jobs {
    if target == Jobs[i].Target && !Jobs[i].Done {
      return i
    }
  }

  return -1
}
