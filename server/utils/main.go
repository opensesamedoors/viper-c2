package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// just a list for chars, used in MakeRandom
var Chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// seeds random number generator with current nanosecond
func RandInit(){
  rand.Seed(time.Now().UnixNano())
}

// generates a random string
func MakeRandom(l int) string{
  b := make([]rune, l)
  for i := range b {
    b[i] = Chars[rand.Intn(len(Chars))]
  }
  return string(b)
}

// create a gin.H JSON with the given error
func ErrorJSON(msg string) gin.H {
  return gin.H {
    "error": msg,
  }
}

// get current time in the HH:MM DD/MM/YYYY format
func GetTime() string {
  now := time.Now()
  return fmt.Sprintf("%02d:%02d:%02d %d/%02d/%02d", 
    now.Hour(), now.Minute(), now.Second(), 
    now.Day(), now.Month(), now.Year(),
  )
}
