package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, GetPing())
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

/*
func GetPing() map[string]string {
	result := map[string]string{"message": "pong1"}
	return result
}
*/