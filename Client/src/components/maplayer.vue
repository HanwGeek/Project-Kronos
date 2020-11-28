<template>
  <div>
    <div id="map"></div>
    <v-layout class="tool-panel">
      <div>
        <v-fab-transition>
          <v-btn v-show="add" @click="addPoint" color="green" fab dark small>
            <v-icon>mdi-vector-point</v-icon>
          </v-btn>
        </v-fab-transition>
        <v-fab-transition>
          <v-btn v-show="add" @click="addLine" color="blue" fab dark small>
            <v-icon>mdi-vector-polyline</v-icon>
          </v-btn>
        </v-fab-transition>
        <v-fab-transition>
          <v-btn v-show="add" @click="addPolygon" color="orange" fab dark small>
            <v-icon>mdi-vector-polygon</v-icon>
          </v-btn>
        </v-fab-transition>
        <v-fab-transition>
          <v-btn
            v-show="enableEdit"
            @click="addFeature"
            color="pink"
            fab
            dark
            small
          >
            <v-icon v-if="add">
              mdi-content-save
            </v-icon>
            <v-icon v-else>
              mdi-plus
            </v-icon>
          </v-btn>
        </v-fab-transition>
      </div>
    </v-layout>
  </div>
</template>

<script>
import Map from 'ol/Map'
//import GeoJSON from 'ol/format/GeoJSON';
import OpenLayersView from 'ol/View'
import GeoJSON from 'ol/format/GeoJSON'
import { bbox as bboxStrategy } from 'ol/loadingstrategy'
import { Stroke, Style } from 'ol/style'
import { Tile as TileLayer, Vector as VectorLayer } from 'ol/layer'
import { OSM, Vector as VectorSource } from 'ol/source'
import { Draw, Modify, Snap, Select } from 'ol/interaction'
import shp from 'shpjs'
// import { platformModifierKeyOnly } from 'ol/events/condition'

//import qs from 'qs';
//import FileSaver from 'file-saver';
//import  writeFeaturesObject from 'ol/format/GeoJSON';
//import Projection from 'ol/proj/Projection';

export default {
  name: 'MapLayer',
  data() {
    return {
      map: null,
      layers: [],
      mapLayers: [],
      styles: [],
      layerCnt: 0,
      metadata: null,
      tmpLayer: null,
      add: false,
      enableEdit: false,
      curIdx: 0,
      clientHeight: 0,
      type: '',
      draw: null,
      snap: null,
      sourceChosen: null,
      layerChosen: null,
      select: null,
      selectedFeatures: null,
      modify: null,
      modifyFeatures: [],
      wfsLayer: null,
      feature_to_save: null,
      f: null,
      data_Projection: null,
      op: null,
      surl: 'http://localhost:8088/',
    }
  },
  created() {
    this.$bus.$on('change-visible', (idx) => {
      if (this.mapLayers[idx] == null) {
        this.loadLayer(idx)
        this.mapLayers[idx] = this.wfsLayer
        this.map.addLayer(this.wfsLayer)
        this.styles.push('#000000')
        this.layers[idx].visible = true
      } else {
        this.layers[idx].visible = !this.layers[idx].visible
        this.mapLayers[idx].setVisible(this.layers[idx].visible)
      }
    })

    this.$bus.$on('change-edit', (idx) => {
      this.curIdx = idx
      if (this.layers[idx].edit == false) {
        this.layers[idx].edit = true
        this.enableEdit = true
        this.modifyFeatures = []
        this.edit()
      } else {
        this.layers[idx].edit = false
        this.enableEdit = false
        this.map.removeInteraction(this.modify)
        this.save()
      }
    })

    this.$bus.$on('show-osm', (show) => {
      this.osmLayer.setVisible(show)
    })

    this.$bus.$on('add-layer-data', (data) => {
      shp(data)
        .then((x) => {
          var newLayer = new VectorLayer({
            source: new VectorSource({
              features: new GeoJSON({
                featureProjection: 'EPSG:4326',
              }).readFeatures(x),
            }),
            style: new Style({
              stroke: new Stroke({
                color: '#000000',
                width: 2,
              }),
            }),
          })
          this.map.addLayer(newLayer)
        })
        .catch((e) => console.log('error', e))
    })

    this.$bus.$on('get-color', (idx) => {
      this.$bus.$emit('send-color', this.styles[idx])
    })

    this.$bus.$on('set-color', (param) => {
      this.mapLayers[param[0]].setStyle(
        new Style({
          stroke: new Stroke({
            color: param[1],
            width: 2,
          }),
        })
      )
    })
  },
  mounted() {
    this.clientHeight = `${document.documentElement.clientHeight}` - 64
    this.init()
  },
  methods: {
    init() {
      // Init layers metadata
      this.$http.get(this.surl + 'getlayerinfo').then((response) => {
        this.metadata = response.data
        this.layerCnt = this.metadata.length
        this.layers = this.metadata.map((o) => {
          return {
            id: o.LayerID,
            name: o.LayerName,
            cnt: o.Count,
            type: o.Type,
            show: false,
            edit: false,
          }
        })
        this.$bus.$emit(
          'layer-names',
          this.layers.map((o) => {
            return o.name
          })
        )
        this.mapLayers = Array.apply(null, Array(this.layerCnt)).map(
          function() {
            return null
          }
        )
      })

      this.osmLayer = new TileLayer({
        source: new OSM(),
      })

      this.osmLayer.setVisible(false)
      this.map = new Map({
        controls: [],
        layers: [this.osmLayer],
        target: 'map',
        view: new OpenLayersView({
          projection: 'CRS:84',
          center: [116.3, 40],
          zoom: 12,
        }),
      })
    },
    loadLayer(idx) {
      var wfsSource = new VectorSource({
        format: new GeoJSON(),
        url: this.surl + 'getlayer?id=' + this.layers[idx].id,
        strategy: bboxStrategy,
      })

      this.wfsLayer = new VectorLayer({
        source: wfsSource,
        style: new Style({
          stroke: new Stroke({
            color: '#000000',
            width: 2,
          }),
        }),
      })
    },
    edit() {
      this.enableSelect()
      this.enableModify()
    },
    save() {
      const format = new GeoJSON({ featureProjection: 'EPSG:4326' })

      this.$http({
        url: this.surl + 'post',
        method: 'post',
        //发送格式为json
        data: {
          op: 2,
          layer_id: this.curIdx,
          feats: format.writeFeaturesObject(this.modifyFeatures),
        },

        //headers: {'Content-Type':'application/x-www-form-urlencoded'}
        headers: { 'Content-Type': 'application/json; charset=UTF-8' },
      })
    },
    enableSelect() {
      //允许通过点选或者框选进行选择矢量要素
      //点选、框选功能实现
      //实现鼠标点击选择
      this.select = new Select()
      this.select.on('select', (e) => {
        this.modifyFeatures.push(e.selected[0])
      })
      this.map.addInteraction(this.select)
      this.selectedFeatures = this.select.getFeatures()

      //鼠标框选(ctrl+拖动)
      // this.dragBox = new DragBox({
      //   condition: platformModifierKeyOnly,
      // })
      // this.map.addInteraction(this.dragBox)

      // this.dragBox.on('boxend', () => {
      //   // 视图没有进行旋转变化时，框选范围可以视为和实际范围一致，因此矢量要素和框选相交时可以视为被选中
      //   var rotation = this.map.getView().getRotation()
      //   var oblique = rotation % (Math.PI / 2) !== 0
      //   var candidateFeatures = oblique ? [] : this.selectedFeatures
      //   var extent = this.dragBox.getGeometry().getExtent()

      //   this.mapLayers[this.curIdx]
      //     .getSource()
      //     .forEachFeatureIntersectingExtent(extent, function(feature) {
      //       candidateFeatures.push(feature)
      //     })

      //   // 如果视图存在旋转变化，需要先把方框和要素同时旋转
      //   if (oblique) {
      //     var anchor = [0, 0]
      //     var geometry = this.dragBox.getGeometry().clone()
      //     geometry.rotate(-rotation, anchor)
      //     var extent$1 = geometry.getExtent()
      //     candidateFeatures.forEach(function(feature) {
      //       var geometry = feature.getGeometry().clone()
      //       geometry.rotate(-rotation, anchor)
      //       if (geometry.intersectsExtent(extent$1)) {
      //         this.selectedFeatures.push(feature)
      //       }
      //     })
      //   }
      // })

      //点击、绘制新的方框时清空选择列表
      // this.dragBox.on('boxstart', () => {
      //   this.selectedFeatures.clear()
      // })
    },
    enableModify() {
      //创建一个Modify控件，指定source参数来指定可以对哪些地图源进行图形编辑，这里仅允许对选中的对象进行编辑
      //Map对象中加入Modify控件后，就可以使用鼠标对已绘制的图形进行编辑。除了可以用鼠标拖拽图形节点外，
      //也可以使用鼠标拖拽直线，这将会拖拽出新的节点。如果想删除某个节点，只需要按住键盘的Alt键，然后鼠标点击该节点即可
      this.modify = new Modify({
        //source: this.mapLayers[this.curIdx].getSource()
        features: this.select.getFeatures(),
      })

      // 将Modify控件加入到Map对象中
      this.map.addInteraction(this.modify)
    },
    addFeature() {
      if (this.add == true) {
        this.add = false
        this.enableModify()
        const format = new GeoJSON({ featureProjection: 'EPSG:4326' })

        this.$http({
          url: this.surl + 'post',
          method: 'post',
          //发送格式为json
          data: {
            op: 1,
            layer_id: this.curIdx,
            feats: format.writeFeaturesObject(
              this.tmpLayer.getSource().getFeatures()
            ),
          },
          headers: { 'Content-Type': 'application/json; charset=UTF-8' },
        })
      } else {
        this.add = true
        this.tmpLayer = new VectorLayer({
          source: new VectorSource(),
          style: new Style({
            stroke: new Stroke({
              color: '#000000',
              width: 2,
            }),
          }),
        })
        this.map.addLayer(this.tmpLayer)
        this.map.removeInteraction(this.draw)
      }
    },
    addPoint() {
      this.map.removeInteraction(this.draw)
      this.draw = new Draw({
        // source: this.mapLayers[this.curIdx].getSource(),
        source: this.tmpLayer.getSource(),
        type: 'Point',
      })
      this.map.addInteraction(this.draw)
    },
    addLine() {
      this.map.removeInteraction(this.draw)
      this.draw = new Draw({
        // source: this.mapLayers[this.curIdx].getSource(),
        source: this.tmpLayer.getSource(),
        type: 'LineString',
      })
      this.map.addInteraction(this.draw)
    },
    addPolygon() {
      this.map.removeInteraction(this.draw)
      this.draw = new Draw({
        // source: this.mapLayers[this.curIdx].getSource(),
        source: this.tmpLayer.getSource(),
        type: 'Polygon',
      })
      this.map.addInteraction(this.draw)
    },
    postData(_op, _layer_id) {
      const format = new GeoJSON({ featureProjection: 'EPSG:4326' })

      this.$http({
        url: this.surl + 'post',
        method: 'post',
        //发送格式为json
        data: {
          op: _op,
          Layer_id: _layer_id,
          geojson: JSON.stringify(
            format.writeFeaturesObject(this.tmpLayer.getSource().getFeatures())
          ),
        },
        headers: { 'Content-Type': 'application/json; charset=UTF-8' },
      })
    },

    //地图增加绘制与拖动控件
    AddInteraction() {
      if (this.type != 'None') {
        var typeChosen = this.type
        var geometryFunction
        switch (this.type) {
          // 生成规则的四边形的图形函数
          case 'Square':
            typeChosen = 'Circle'
            geometryFunction = Draw.createRegularPolygon(4)
            break
          // 生成盒状图形函数
          case 'Box':
            typeChosen = 'Circle'
            geometryFunction = Draw.createBox()
            break
          default:
            break
        }

        //初始化Draw绘图控件
        this.draw = new Draw({
          source: this.sourceChosen,
          type: typeChosen,
          geometryFunction: geometryFunction,
        })

        //将Draw控件加入地图
        this.map.addInteraction(this.draw)

        //初始化Snap控件
        this.snap = new Snap({
          source: this.sourceChosen,
        })
        this.map.addInteraction(this.snap)
      }
    },
  },
}
</script>

<style scoped>
#map {
  width: 100%;
  height: 800px;
  left: 0;
  z-index: 5;
}
#ol-dragbox {
  background-color: rgba(255, 255, 255, 0.4);
  border-color: rgba(100, 150, 0, 1);
}

.tool-panel {
  position: fixed;
  bottom: 0;
  right: 0;
}
</style>
