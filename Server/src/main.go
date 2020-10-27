package main

import (
	"kronos/src/ogcservice"
)

func main() {
	ogcservice.Connect()
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})

	// })

	// r.Run()

}
