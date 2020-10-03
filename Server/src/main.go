package main

import (
	"log"
	"net/http"
	"ogcservice"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/wfs", ogcservice.WFSHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	// shape, err := shp.Open("./data/crop.shp")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer shape.Close()
	// // fields from the attribute table (DBF)
	// fields := shape.Fields()

	// // loop through all features in the shapefile
	// for shape.Next() {
	// 	n, p := shape.Shape()

	// 	// print feature
	// 	fmt.Fprintf(w, "%v %v", reflect.TypeOf(p).Elem(), p.BBox())

	// 	// print attributes
	// 	for k, f := range fields {
	// 		val := shape.ReadAttribute(n, k)
	// 		fmt.Fprintf(w, "\t%v: %v\n", f, val)
	// 	}
	// 	fmt.Fprintf(w, "\n")
	// }
}
