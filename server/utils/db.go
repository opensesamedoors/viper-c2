package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func InitDatabase() {
  st, err := os.Stat(Config.DatabaseDir)

  if os.IsNotExist(err) {
    err := os.Mkdir(Config.DatabaseDir, 0777)
    if err != nil {
      log.Fatal("Cannot create database directory: "+err.Error())
      return
    }
    return
  }

  if err != nil {
    log.Fatal("Error checking database: "+err.Error())
    return
  }

  if !st.IsDir(){
    log.Fatal("Database is not a directory")
  }
}

func LoadFromDatabase(file string, obj any) bool{
  raw, err := ioutil.ReadFile(path.Join(Config.DatabaseDir, file))

  if os.IsNotExist(err) {
    return false
  }

  if err != nil {
    log.Fatal("Can't read database: "+err.Error())
    return false
  }

  err = json.Unmarshal(raw, obj)
  if err != nil {
    log.Fatal("JSON unmarshal error: "+err.Error())
    return false
  }

  return true
}

func SaveToDatabase(file string, obj any) bool{
  raw, err := json.Marshal(obj)
  if err != nil {
    log.Fatal("JSON marshal error: "+err.Error())
    return false
  }

  err = ioutil.WriteFile(path.Join(Config.DatabaseDir, file), raw, 0666)
  if err != nil {
    log.Fatal("Can't write database: "+err.Error())
    return false
  }

  return true
}
