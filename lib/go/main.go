package main 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
	"bytes"
)

type Agent struct {
  Server string 
  Uid    string
  Token  string
}

type RegisterReqData struct {
  Uid   string  `json:"uid"`
  Msg   string  `json:"message"`
  Desc  string  `json:"desc"`
  Token string  `json:"token"`
}

type JobData struct{
  Action  string  `json:"action"`
  Data    string  `json:"data"`
  Target  string  `json:"target"`
  Done    bool    `json:"done"`
}

type ResData struct {
  Error  string  `json:"error"`
  Result string  `json:"result"`
}

type ResultReqData struct {
  Result string `json:"result"`
  Done   bool   `json:"done"`
}

func GetAuthPath(agent *Agent, pth string) string {
  return agent.Server+"/agent/"+pth+"?uid="+agent.Uid+"&token="+agent.Token
}

func Register(agent *Agent, desc string, msg string) error {
  var reqdata RegisterReqData = RegisterReqData{
    Uid: agent.Uid,
    Msg: msg,
    Desc: desc,
    Token: agent.Token,
  }
  var resdata ResData

  postData, err := json.Marshal(reqdata)
  if err != nil {
    return err
  }

  res, err := http.Post(agent.Server+"/agent/register", "application/json", bytes.NewBuffer(postData))
  if err != nil {
    return err
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return err
  }
  
  json.Unmarshal(body, &resdata)
  if resdata.Error != "" {
    return errors.New(resdata.Error)
  }

  return nil
}

func Result(agent *Agent, result string, done bool) error {
  var reqdata ResultReqData = ResultReqData{
    Result: result,
    Done: done,
  }
  var resdata ResData

  postData, err := json.Marshal(reqdata)
  if err != nil {
    return err
  }

  res, err := http.Post(GetAuthPath(agent, "result"), "application/json", bytes.NewBuffer(postData))
  if err != nil {
    return err
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return err
  }
  
  json.Unmarshal(body, &resdata)
  if resdata.Error != "" {
    return errors.New(resdata.Error)
  }

  return nil
}

func Jobs(agent *Agent) (JobData, error) {
  var resdata ResData
  var job JobData

  res, err := http.Get(GetAuthPath(agent, "job"))
  if err != nil {
    return job, err
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return job, err
  }

  json.Unmarshal(body, &resdata)
  if resdata.Error != "" {
    return job, errors.New(resdata.Error)
  }

  json.Unmarshal([]byte(resdata.Result), &job)
  if job.Action == "" {
    return job, errors.New("Bad job")
  }

  return job, nil
} 
