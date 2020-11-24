package ogcservice

import (
	"encoding/json"
	"fmt"
	"log"

	"kronos/src/geom"

	"github.com/go-pg/pg/v10" // version error, wtf!!!!!!
)

const (
	hostname = "162.105.17.227:5432"
	user     = "han"
	password = "wh123"
	dbname   = "features"
)

var db *pg.DB

type LayerInfo struct {
	//tableName struct{} `pg:"layers"`
	LayerID   int    `pg:"layer_id,pk,use_zero"`
	LayerName string `pg:"layer_name"`
	Type      int    `pg:"type"`
	Count     int    `pg:"count"`
	//Attributes string   `pg:"attributes"`
	//Wkt        string   `pg:"wkt"`
}

type FeatInfo struct {
	//tableName  struct{} `pg:"features"`
	LayerID    int    `pg:"layer_id,pk,use_zero"`
	FeatID     int    `pg:"feat_id,pk,use_zero"`
	Type       int    `pg:"type"`
	Attributes string `pg:"attributes"`
	Wkt        string `pg:"wkt"`
}

/*
// Structures for parsing json.
type AttrJSON struct {
	Attributes []map[string]interface{} `json:"attributes"`
}

type WktJSON struct {
	Wkt []string `json:"wkt"`
}*/

// AttrJSON is to parse a feature's attributes(json format).
type AttrJSON struct {
	Attributes map[string]interface{} `json:"attributes"`
}

// Connect is to connect postgresql database
func Connect() {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     hostname,
			User:     user,
			Password: password,
			Database: dbname,
		})
	}

	fmt.Printf("[Database] Connect to %v successfully\n", hostname)

	// get a layer.
	// fmt.Printf("test for query data.\n\n")
	// layers := GetLayers()
	// fmt.Printf("test1: layer count = %d\n", len(layers))
	//fmt.Printf("test2: layers = %v\n", layers)

	// // add a layer.
	// layer := geom.NewLayer(1, "test_layer2", 1) // layer pointer.
	// //p1 := geom.Point{geom.Coord{1, 2}}
	// //p2 := geom.Point{geom.Coord{3, 4}}
	// p1 := geom.Point{}
	// p1.SetPos(1, 2)
	// p2 := geom.Point{}
	// p2.SetPos(3, 4)
	// attr1 := map[string]interface{}{"fid": 0, "name": "point1", "class": 1}
	// attr2 := map[string]interface{}{"fid": 1, "name": "point2", "class": 2}
	// feat1 := geom.NewFeature(p1, attr1) // feature pointer.
	// feat2 := geom.NewFeature(p2, attr2) // feature pointer.
	// fmt.Printf("the feat1 = %v\n", *feat1)
	// fmt.Printf("the feat2 = %v\n", *feat2)
	// layer.AddFeature(*feat1)
	// layer.AddFeature(*feat2)
	// AddLayer(*layer)
	// fmt.Printf("the new layer = %v\n", *layer)

	// fmt.Printf("\n\ntest for add data.\n\n")
	// layers = GetLayers()
	// fmt.Printf("test3: layer count = %d\n", len(layers))
	// fmt.Printf("test4: layers = %v\n", layers[1:])

	// // update a layer.
	// layer.SetName("test_layer3")
	// f := layer.GetFeature(0)
	// (&f).SetAttribute("name", "point1.1")
	// //(&(layer.GetFeature(0))).SetAttribute("name", "point1.1")
	// //layer.GetFeature(0).SetAttribute("name", "point1.1")
	// UpdateLayer(*layer)

	// fmt.Printf("test5: layer count = %d\n", len(layers))
	// fmt.Printf("test6: layers = %v\n", layers[1:])

	// // delete a layer.
	// DelLayer(layer.GetId())

	// fmt.Printf("\n\ntest for delete data.\n\n")
	// fmt.Printf("layer.getid = %d\n", layer.GetId())
	// layers = GetLayers()
	// fmt.Printf("test7: layer count = %d\n", len(layers))

}

/*
* Below are 4 interfaces to query data from database.
* 1. GetFeatureById
* 2. GetFeatures
* 3. GetLayerById
* 4. GetLayer/*
		// get a layer.
		fmt.Printf("test for query data.\n\n")
		layers := GetLayers()
		fmt.Printf("test1: layer count = %d\n", len(layers))
		//fmt.Printf("test2: layers = %v\n", layers)

		// add a layer.
		layer := geom.NewLayer(1, "test_layer2", 1) // layer pointer.
		//p1 := geom.Point{geom.Coord{1, 2}}
		//p2 := geom.Point{geom.Coord{3, 4}}
		p1 := geom.Point{}
		p1.SetPos(1, 2)
		p2 := geom.Point{}
		p2.SetPos(3, 4)
		attr1 := map[string]interface{}{"fid": 0, "name": "point1", "class": 1}
		attr2 := map[string]interface{}{"fid": 1, "name": "point2", "class": 2}
		feat1 := geom.NewFeature(p1, attr1) // feature pointer.
		feat2 := geom.NewFeature(p2, attr2) // feature pointer.
		fmt.Printf("the feat1 = %v\n", *feat1)
		fmt.Printf("the feat2 = %v\n", *feat2)
		layer.AddFeature(*feat1)
		layer.AddFeature(*feat2)
		AddLayer(*layer)
		fmt.Printf("the new layer = %v\n", *layer)

		fmt.Printf("\n\ntest for add data.\n\n")
		layers = GetLayers()
		fmt.Printf("test3: layer count = %d\n", len(layers))
		fmt.Printf("test4: layers = %v\n", layers[1:])

		// update a layer.
		layer.SetName("test_layer3")
		f := layer.GetFeature(0)
		(&f).SetAttribute("name", "point1.1")
		//(&(layer.GetFeature(0))).SetAttribute("name", "point1.1")
		//layer.GetFeature(0).SetAttribute("name", "point1.1")
		UpdateLayer(*layer)
s
*/

// GetFeatureById return a feature with a given layerId & featId (as the format we defined in "geom" module).
func GetFeatureById(layerId int, featId int) geom.Feature {
	var feature geom.Feature
	featinfo := GetFeatInfoById(layerId, featId)

	// Parse attrbutes.
	attrData := []byte(featinfo.Attributes)
	var attrJson AttrJSON
	err := json.Unmarshal(attrData, &attrJson)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	attr := attrJson.Attributes
	wkt := featinfo.Wkt

	geometry := geom.NewGeomByWKT(wkt)
	feature = *geom.NewFeature(geometry, attr)

	return feature
}

func parseFeat(featinfo FeatInfo) geom.Feature {
	// Parse attrbutes.
	attrData := []byte(featinfo.Attributes)
	var attrJson interface{}
	err := json.Unmarshal(attrData, &attrJson)
	if err != nil {
		log.Fatal(err)
	}

	// attr := attrJson.Attributes
	wkt := featinfo.Wkt

	geometry := geom.NewGeomByWKT(wkt)
	feature := *geom.NewFeature(geometry, attrJson.(map[string]interface{}))
	feature.SetID(featinfo.FeatID)
	return feature
}

// GetFeatures return all features with a given layerId (as the format we defined in "geom" module).
func GetFeatures(layerID int) []geom.Feature {
	var features []geom.Feature
	featinfos := GetFeatInfo(layerID)

	for _, featinfo := range featinfos {
		// feature := GetFeatureById(layerID, featinfo.FeatID)
		features = append(features, parseFeat(featinfo))
	}

	return features
}

// GetLayerById return a layer with a given layerId (as the format we defined in "geom" module).
func GetLayerById(layerId int) geom.Layer {
	//var layer geom.Layer
	layerinfo := GetLayerInfoById(layerId)

	layer := geom.NewLayer(layerinfo.LayerID, layerinfo.LayerName, layerinfo.Type)
	feats := GetFeatures(layerinfo.LayerID)
	for _, feat := range feats {
		layer.AddFeature(feat)
	}

	return *layer
}

// GetLayers return all layers (as the format we defined in the "geom" module).
func GetLayers() []geom.Layer {
	var layers []geom.Layer
	layerinfos := GetLayerInfo()
	// Read all layers.
	for _, layerinfo := range layerinfos {
		layer := GetLayerById(layerinfo.LayerID)
		layers = append(layers, layer)
	}

	return layers
}

/*
* Below are 2 interfaces to insert data into database.
* 1. AddFeature
* 2. AddLayer
 */

// AddFeature is to add a feature into the database with a given layerId.
func AddFeature(layerId int, feat geom.Feature) {
	// insert a feature into the table "feat_infos".
	var featinfo FeatInfo
	attr := AttrJSON{
		Attributes: feat.GetAttributes(),
	}
	attrStr, err := json.Marshal(attr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	featinfo = FeatInfo{
		LayerID: layerId,
		FeatID:  feat.GetID(),
		Type:    feat.GetGeomType(),
		//Attributes: string(feat.GetAttributes()),
		Attributes: string(attrStr), //json.Marshal(attr)
		Wkt:        feat.GetGeometry().ExportWKT(),
	}

	_, err = db.Model(&featinfo).Insert()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// update the "count" of the given layer(count++) in the table "layer_infos".
	layerinfo := GetLayerInfoById(layerId)
	layerinfo.Count += 1

	_, err = db.Model(&layerinfo).WherePK().Update() // wherepk
	//_, err = db.Model(&layerinfo).Set("count = ?count").Where("layer_id = ?layer_id").Update() // wherepk
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// AddLayer is to add a layer into the database.
func AddLayer(layer geom.Layer) {
	// insert a layer into the table "layer_infos".
	var layerinfo LayerInfo
	layerinfo = LayerInfo{
		LayerID:   layer.GetId(),
		LayerName: layer.GetName(),
		Type:      layer.GetGeomType(),
		Count:     0,
		//Count: 		layer.FeatureCount(),
	}

	_, err := db.Model(&layerinfo).Insert()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	//fmt.Printf("feature count: %d\n", layer.FeatureCount())

	// insert all the layer's features into the table "feat_infos".
	//for i := 0; i < layerinfo.Count; i++ {
	// woc, wojuranzaizhelibeikengle, wofule.
	for i := 0; i < layer.FeatureCount(); i++ {
		feat := layer.GetFeature(i)
		AddFeature(layerinfo.LayerID, feat)
	}

}

/*
* Below are 2 interfaces to delete data from database.
* 1. DelFeature
* 2. DelLayer
 */

// DelFeature is to delete a feature from database with a given layerId and featId.
func DelFeature(layerId int, featId int) {
	// delete a feature from the table "feat_infos".
	featinfo := new(FeatInfo)
	_, err := db.Model(featinfo).Where("layer_id = ?", layerId).Where("feat_id = ?", featId).Delete()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// update the "count" of the given layer(count--) in the table "layer_infos".
	layerinfo := GetLayerInfoById(layerId)
	layerinfo.Count -= 1

	_, err = db.Model(&layerinfo).WherePK().Update() // wherepk.
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// DelLayer is to delete a layer from database with a given layerId.
func DelLayer(layerId int) {
	// delete all features of the layer from the table "feat_infos".
	featinfos := GetFeatInfo(layerId)
	for _, featinfo := range featinfos {
		DelFeature(layerId, featinfo.FeatID)
	}

	// "delete a layer" should be behind the "delete all features", because there is a "layer.count--" operation in the function "DelFeature".
	// delete a layer from the table "layer_infos".
	layerinfo := new(LayerInfo)
	_, err := db.Model(layerinfo).Where("layer_id = ?", layerId).Delete()
	if err != nil {
		fmt.Printf("here is dellayer\n")
		fmt.Printf("error: %v\n", err)
	}

}

/*
* Below are 2 interfaces to update data in database.
* 1. UpdateFeature
* 2. UpdateLayer
 */

// UpdataFeature is to update a feature in database with a given layerId.
func UpdateFeature(layerId int, feat geom.Feature) {
	// update a feature in the table "feat_infos".
	var featinfo FeatInfo
	attr := AttrJSON{
		Attributes: feat.GetAttributes(),
	}
	attrStr, err := json.Marshal(attr)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	featinfo = FeatInfo{
		LayerID:    layerId,
		FeatID:     feat.GetID(),
		Type:       feat.GetGeomType(),
		Attributes: string(attrStr), //string(feat.GetAttributes()),
		Wkt:        feat.GetGeometry().ExportWKT(),
	}

	//err := db.Model(&featinfo).Where("layer_id = ?", layerId).Where("feat_id = ?", featinfo.FeatID).Update()
	_, err = db.Model(&featinfo).WherePK().Update() // wherepk.
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// UpdateLayer is to update a layer in database.
func UpdateLayer(layer geom.Layer) {
	// update a layer in the table "layer_infos".
	var layerinfo LayerInfo
	layerinfo = LayerInfo{
		LayerID:   layer.GetId(),
		LayerName: layer.GetName(),
		Type:      layer.GetGeomType(),
		Count:     layer.FeatureCount(),
	}

	_, err := db.Model(&layerinfo).WherePK().Update() // wherepk.
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// update all features of the layer in the table "feat_infos".
	// 1. delete
	featinfos := GetFeatInfo(layerinfo.LayerID)
	for _, featinfo := range featinfos {
		DelFeature(layerinfo.LayerID, featinfo.FeatID)
	}
	// 2. insert
	for i := 0; i < layerinfo.Count; i++ {
		feat := layer.GetFeature(i)
		AddFeature(layerinfo.LayerID, feat)
	}

}

/*
* The functions below are private functions, and they should not be called publicly.
* They are implemented to implement the above public functions,
* and the returned results are all structures defined in this file.
 */

// GetLayerInfo return all layers' metadata.
func GetLayerInfo() []LayerInfo {
	var model []LayerInfo
	err := db.Model(&model).Select()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil
	}
	return model
}

// GetLayerInfoById return layer's metadata with a given layer id.
func GetLayerInfoById(layerId int) LayerInfo {
	var model LayerInfo
	err := db.Model(&model).Where("layer_id = ?", layerId).Select()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		//return nil
	}
	return model
}

// GetFeatInfo return all features from the same layer.
func GetFeatInfo(layerId int) []FeatInfo {
	var model []FeatInfo
	err := db.Model(&model).Order("feat_id ASC").Where("layer_id = ?", layerId).Select()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil
	}
	return model
}

// GetFeatInfo return feature with a given layerId and featId.
func GetFeatInfoById(layerId int, featId int) FeatInfo {
	var model FeatInfo
	err := db.Model(&model).Where("layer_id = ?", layerId).Where("feat_id = ?", featId).Select()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		//return nil
	}
	return model
}
