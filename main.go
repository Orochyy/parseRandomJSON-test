package main

import (
	"github.com/gin-gonic/gin"
	"parseRandomJSON-test/controller"
)

var getJokeController = controller.GetJoke

func main() {
	router := gin.Default()

	router.GET("/joke", getJokeController)

	router.Run("localhost:8080")
}
