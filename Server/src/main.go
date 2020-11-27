package main

import (
	"kronos/src/ogcservice"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ReqData struct {
	Op      int         `json:"op"`
	LayerID int         `json:"layer_id"`
	Feat    interface{} `json:"feats"`
}

func main() {
	ogcservice.Connect()

	lm := ogcservice.NewManager()

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
		layerIDParam := c.Query("id")
		layerID, _ := strconv.Atoi(layerIDParam)
		c.JSON(200, lm.GetLayerContent(layerID))
	})

	r.POST("/post", func(c *gin.Context) {
		var req ReqData

		if c.BindJSON(&req) == nil {
			lm.OperOnLayer(req.Op, req.LayerID, req.Feat.(map[string]interface{}))
			c.JSON(200, gin.H{
				"status": "SUCCESS",
			})
		} else {
			c.JSON(400, gin.H{"status": "FAILED"})
		}

	})

	r.Run(":8088")

}
