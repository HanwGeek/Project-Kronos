package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/jonas-p/go-shp"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}

func handler(w http.ResponseWriter, r *http.Request) {
	shape, err := shp.Open("./data/crop.shp")
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()
	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		// print feature
		fmt.Fprintf(w, "%v %v", reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Fprintf(w, "\t%v: %v\n", f, val)
		}
		fmt.Fprintf(w, "\n")
	}
}
