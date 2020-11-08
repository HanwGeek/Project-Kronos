package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	// ogcservice.Connect()
	r := gin.Default()

	file, err := ioutil.ReadFile("./data/crop.json")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, file)

	})

	r.Run()

}
