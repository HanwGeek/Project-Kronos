<template>
  <div>
    <div id="map"></div>
    <!-- <div>测试用：{{ type }}</div>
    <label>Shape type &nbsp;</label>
    <select v-model="type">
        <option value="Point">Point</option>
        <option value="LineString">LineString</option>
        <option value="Polygon">Polygon</option>
        <option value="Circle">Circle</option>
        <option value="Square">Square</option>
        <option value="Box">Box</option>
        <option value="None">None</option>
    </select>
    <div>a{{ type }}</div>  -->
  </div>

</template>

<script>
import Map from 'ol/Map'
//import GeoJSON from 'ol/format/GeoJSON';
import OpenLayersView from 'ol/View'
import GeoJSON from 'ol/format/GeoJSON';
import {bbox as bboxStrategy} from 'ol/loadingstrategy';
import {Stroke, Style} from 'ol/style';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer'
import {OSM, Vector as VectorSource} from 'ol/source'
import {Draw, Modify,Snap,Select,DragBox} from 'ol/interaction';

export default {
  name: 'MapLayer',
  data () {
    return {
      aaa: 0,
      type: '',
      map: null,
      Layers: null,
      draw: null,
      snap: null,
      sourceChosen: null,
      layerChosen: null,
      selectedFeatures: null,
      wfsSource: null,
      wfsLayer: null,
      layers: [],
      layerNames: ["crop", "landuse", "Beijing"],
      layerMap: null,
      visible: [],
      layerCnt: 3,
    }
  },
  created () {
    this.draw = new Draw();
    this.snap = new Snap();
    this.$bus.$on("changeVisble", (idx) => {
      if (this.layers[idx] == null) {
        // this.layerMap[idx] =  this.map.layers.length;
        this.loadLayer();
        this.layers[idx] = this.wfsLayer;
        this.map.addLayer(this.wfsLayer);
      } else {
        this.visible[idx] = !this.visible[idx];
        this.layers[idx].setVisible(this.visible[idx]);
      }
    })
  },
  mounted () {
    this.init();
    this.loadLayer();
  },
  methods: {
    init () {
      // Init layers metadata

      this.layers = Array.apply(null, Array(this.layerCnt)).map(function () {return null});
      this.visible = Array.apply(null, Array(this.layerCnt)).map(function () {return false;});
      this.layerMap = Array.apply(null, Array(this.layerCnt)).map(function (x, i) { return i});
      this.$bus.$emit("layerNames", this.layerNames);

      this.osmLayer = new TileLayer({
        source: new OSM()
      });

      this.map = new Map({
        controls: [],
        layers: [this.osmLayer],
        target: 'map',
        view: new OpenLayersView({
          projection: 'CRS:84',
          center: [116.3, 40],
          zoom: 12
        }),
      });
    },
    loadLayer() {
      this.wfsSource = new VectorSource({
        format: new GeoJSON(),
        url: require("../assets/crop.json"),
        strategy: bboxStrategy
      });

      this.wfsLayer = new VectorLayer({
        source: this.wfsSource,
        style: new Style({
          stroke: new Stroke({
            color: 'rgba(0, 0, 255, 1.0)',
            width: 2
          })
        })
      });
    },
    edit() {

      this.sourceChosen = new VectorSource();
      this.layerChosen = new VectorLayer({
        source:this.sourceChosen
      });
      //创建一个Modify控件，指定source参数来指定可以对哪些地图源进行图形编辑，
      //Map对象中加入Modify控件后，就可以使用鼠标对已绘制的图形进行编辑。除了可以用鼠标拖拽图形节点外，
      //也可以使用鼠标拖拽直线，这将会拖拽出新的节点。如果想删除某个节点，只需要按住键盘的Alt键，然后鼠标点击该节点即可
      var modify = new Modify({
        source: this.sourceChosen
      });

      // 将Modify控件加入到Map对象中
      this.map.addInteraction(modify);

      //点选、框选功能实现
      //实现鼠标点击选择
      var select = new Select();
      this.map.addInteraction(select);
      this.selectedFeatures = select.getFeatures();

      //鼠标框选
      var dragBox = new DragBox();

      this.map.addInteraction(dragBox);

      dragBox.on('boxend', () => {
      // 视图没有进行旋转变化时，框选范围可以视为和实际范围一致，因此矢量要素和框选相交时可以视为被选中
      var rotation = this.map.getView().getRotation();
      var oblique = rotation % (Math.PI / 2) !== 0;
      var candidateFeatures = oblique ? [] : this.selectedFeatures;
      var extent = dragBox.getGeometry().getExtent();
      this.sourceChosen.forEachFeatureIntersectingExtent(extent, function (feature) {
        candidateFeatures.push(feature);
      });

      // 如果视图存在旋转变化，需要先把方框和要素同时旋转
      if (oblique) {
          var anchor = [0, 0];
          var geometry = dragBox.getGeometry().clone();
          geometry.rotate(-rotation, anchor);
          var extent$1 = geometry.getExtent();
          candidateFeatures.forEach(function (feature) {
          var geometry = feature.getGeometry().clone();
          geometry.rotate(-rotation, anchor);
          if (geometry.intersectsExtent(extent$1)) {
              this.selectedFeatures.push(feature);
          }
          });
      }
      });

      //点击、绘制新的方框时清空选择列表
      dragBox.on('boxstart', () => {
        this.selectedFeatures.clear();
      });
    },
    //地图增加绘制与拖动控件
    AddInteraction(){
      if(this.type != "None"){
        var typeChosen = this.type;
        var geometryFunction;
        switch(this.type){
          // 生成规则的四边形的图形函数
          case "Square":
            typeChosen = 'Circle';
            geometryFunction = Draw.createRegularPolygon(4);
            break;
          // 生成盒状图形函数
          case "Box":
            typeChosen =  'Circle';
            geometryFunction = Draw.createBox();
            break;
          default:break;
        }

        //初始化Draw绘图控件
        this.draw = new Draw({
          source:this.sourceChosen,
          type: typeChosen,
          geometryFunction: geometryFunction
        });

        //将Draw控件加入地图
        this.map.addInteraction(this.draw);

        //初始化Snap控件
        this.snap = new Snap({
          source: this.sourceChosen
        })
        this.map.addInteraction(this.snap)

      }
    }
  },

  watch:{
    //监听要绘制的矢量要素类别是否发生了变化，发生变化时将绘制控件添加到地图上去
    type:function(val){
      this.type = val;
      this.map.removeInteraction(this.draw);
      this.map.removeInteraction(this.snap);
      this.AddInteraction();
    }

  }
  
}
</script>

<style scoped>
#map {
  width: 100%;
  height: 700px;
  left: 0;
  z-index: 5;
}
#ol-dragbox {
  background-color: rgba(255,255,255,0.4);
  border-color: rgba(100,150,0,1);
}
</style>

