package operator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngn13/viper/server/bridge"
	"github.com/ngn13/viper/server/utils"
)

// deletes a token, making this prevents any agent registering
// with this token
func DelToken(c *gin.Context){
  token := c.Query("token")
  indx := -1

  for i := range bridge.Tokens {
    if bridge.Tokens[i] == token {
      indx = i
      break
    }
  }

  if indx == -1 {
    c.JSON(http.StatusNotFound, utils.ErrorJSON("Token not found"))
    return
  }

  bridge.Tokens = append(bridge.Tokens[:indx], bridge.Tokens[indx+1:]...)
  bridge.SaveTokens()
  c.JSON(http.StatusOK, utils.ErrorJSON(""))
}

// creates a new access token that can be used in agents
func AddToken(c *gin.Context){
  token := utils.MakeRandom(24)
  bridge.Tokens = append(bridge.Tokens, token)

  tokensjson := bridge.TokensJson{}
  tokensjson.List = bridge.Tokens

  bridge.SaveTokens()
  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": token,
  })
}

// returns a list of tokens
func GetTokens(c *gin.Context){ 
  c.JSON(http.StatusOK, gin.H {
    "error": "",
    "result": bridge.Tokens,
  })
}
