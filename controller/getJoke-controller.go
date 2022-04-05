package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parseRandomJSON-test/service"
)

type GetJokeController interface {
	GetJoke(c *gin.Context)
}

//type getJokeController struct {
//	getJokeService service.JokeService
//}

func GetJoke(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ParseJokeNameStruct())
}
