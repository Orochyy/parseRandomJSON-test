package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parseRandomJSON-test/service"
)

type GetJokeController interface {
	GetJoke(c *gin.Context)
}

func GetJoke(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ParseJokeNameStruct())
}
