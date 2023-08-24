# Viper | Go API Library 
Go API libary can be used to interact with a viper web server,
**this library only supports agent features**

## Setup 
Add this library to your project in order to use it:
```
go get github.com/ngn13/viper/lib/go 
```

## Examples
### Example Agent
```go
package main 

import (
	"fmt"
    viper "github.com/ngn13/viper/lib/go"
)

func main(){
  agent := viper.Agent{
    Token: "123",
    Uid: "123",
    Server: "http://localhost:8080",
  }

  err := viper.Register(&agent, "A cool agent", "And a cool message")
  if err != nil {
    return
  }

  job, err := viper.Jobs(&agent)
  if err != nil {
    return
  }

  fmt.Printf("Job action %s", job.Action)
  fmt.Printf("Job data %s", job.Data)
  fmt.Printf("Job target %s", job.Target)

  err = viper.Result(&agent, "I completed the job!", true)
  if err != nil {
    return
  }

  fmt.Println("Completed!")
}
```
