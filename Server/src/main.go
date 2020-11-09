package main

import "kronos/src/ogcservice"

func main() {
	ogcservice.Connect()
	// r := gin.Default()

	// // Allow CORS
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// r.Use(cors.New(config))

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

	// r.Run()

}
