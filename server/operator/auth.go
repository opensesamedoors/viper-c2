package operator

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/utils"
)

type LoginData struct {
  Username  string  `json:"username"`
  Password  string  `json:"password"`
} 

type Cookie struct {
  Token     string 
  Lastcon   string
  Username  string
}

// list of operator cookies and the length for a cookie
var cookies []Cookie = []Cookie{}
var COOKIE_LEN int = 20

// generates an operator cookie and appends it to 
// the cookies list then returns it
func AddCookie(username string) Cookie{
  cookie := Cookie{}
  cookie.Token = utils.MakeRandom(COOKIE_LEN) 
  cookie.Username = username
  cookies = append(cookies, cookie)
  return cookie
}

// checks if a given cookie is valid or not 
// if so then it returns the cookie index
func FindCookie(token string) int {
  for i := range cookies {
    if token == cookies[i].Token { 
      return i 
    }
  }

  return -1 
}

// removes a given cookie from the cookies list
func RemoveCookie(token string) error {
  indx := FindCookie(token) 

  if indx == -1 {
    return errors.New("Cookie not found")
  }

  cookies = append(cookies[:indx], cookies[indx+1:]...)
  return nil
}

// checks if operator is authenticated, if so calls the 
// handler, if not returns 404 page
func AuthMiddleware(c *gin.Context){
  if strings.Contains(c.FullPath(), "login") {
    c.Next()
    return
  }

  token, err := c.Cookie("token")

  if err != nil {
    c.JSON(http.StatusUnauthorized, utils.ErrorJSON("You are not authenticated"))
    c.Abort()
    return 
  }

  indx := FindCookie(token)
  if indx == -1 {
    c.JSON(http.StatusUnauthorized, utils.ErrorJSON("You are not authenticated"))
    c.Abort()
    return
  }

  cookies[indx].Lastcon = utils.GetTime()
  c.Next()
}

// lets an operator to login
func AuthLogin(c *gin.Context){
  body := LoginData{}
  c.BindJSON(&body)

  if body.Username == "" || body.Password == "" {
    c.JSON(http.StatusBadRequest, utils.ErrorJSON("Bad JSON"))
    return 
  }

  for i := range body.Username {
    if !strings.Contains(string(utils.Chars), string(body.Username[i])){
      c.JSON(http.StatusBadRequest, utils.ErrorJSON("Bad username"))
      return
    }
  }

  if body.Password != utils.Config.Password {
    c.JSON(http.StatusForbidden, utils.ErrorJSON("Invalid password"))
    return
  } 

  for i := range cookies {
    if cookies[i].Username == body.Username {
      c.JSON(http.StatusConflict, utils.ErrorJSON("An operator with the same name exists"))
      return
    }
  }

  cookie := AddCookie(body.Username)
  cookie.Lastcon = utils.GetTime()
  c.SetCookie(
    "token", cookie.Token, 3600, "/", utils.Config.Domain, utils.Config.SecureCookie, utils.Config.HttpOnlyCookie,
  )
  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}

// lets an operator to logout
func AuthLogout(c *gin.Context){
  token, _ := c.Cookie("token")
  RemoveCookie(token)

  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}

// sends a list of operators that are logged in 
func AuthList(c *gin.Context) {
  operators := []string{}
  for i := range cookies {
    operators = append(operators, cookies[i].Username+","+cookies[i].Lastcon)
  }

  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": operators,
  })
  return
}
