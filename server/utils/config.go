package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ServerConfig struct {
  CORS              string    `json:"cors"`
  Address           string    `json:"address"`
  Password          string    `json:"password"`
  Domain            string    `json:"domain"`
  Page404           string    `json:"404page"`
  Rootpath          string    `json:"rootpath"`
  DatabaseDir       string    `json:"database_dir"`
  ServerHeader      string    `json:"header"`
  HttpOnlyCookie    bool      `json:"cookie_httponly"`
  SecureCookie      bool      `json:"cookie_secure"`
}

var Config ServerConfig = ServerConfig{}
var Version string = "1.2"

// reads and loads config from "config.json"
func ReadConfig(){
  Config.Address        = "localhost:8080"
  Config.Password       = "changeme"
  Config.Domain         = "localhost"
  Config.Page404        = "404.html"
  Config.Rootpath       = "/"
  Config.DatabaseDir    = "db"
  Config.ServerHeader   = "nginx"
  Config.HttpOnlyCookie = false
  Config.SecureCookie   = true 
  Config.CORS           = "" 

  raw, err := ioutil.ReadFile("config.json")
  if err != nil {
    log.Println("Configuration file not found, using the default config") 
    return
  }

  err = json.Unmarshal(raw, &Config)
  if err != nil {
    log.Fatal("Error reading configuration file: "+err.Error()) 
  }
}
