package main

import (
	"fmt"
	"kronos/src/ogcservice"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ReqData struct {
	Op      int         `json:"op" binding:"required"`
	LayerID int         `json:"layer_id" binding:"required"`
	Feat    interface{} `json:feats`
}

func main() {
	ogcservice.Connect()
	r := gin.Default()

	// Allow CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// file, err := ioutil.ReadFile("./data/crop.json")
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	// data, err := geojson.UnmarshalFeatureCollection(file)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }
	// r.GET("/data", func(c *gin.Context) {
	// 	c.JSON(200, data)
	// })

	r.GET("/getlayerinfo", func(c *gin.Context) {
		info := ogcservice.GetLayerInfo()
		c.JSON(200, info)
	})

	r.GET("/getlayer", func(c *gin.Context) {
		name := c.Query("name")
		fmt.Printf("%v\n", name)
	})

	r.POST("/post", func(c *gin.Context) {
		var req ReqData

		if c.BindJSON(&req) == nil {

			c.JSON(200, gin.H{
				"status": "SUCCESS",
				"op":     req.Op,
				"layer":  req.LayerID,
				"feat":   req.Feat,
			})
		} else {
			c.JSON(400, gin.H{"status": "FAILED"})
		}

	})

	r.Run()

}
