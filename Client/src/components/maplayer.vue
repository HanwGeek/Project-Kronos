<template>
  <div>
    <div id="map"></div>
  </div>
</template>

<script>
import Map from 'ol/Map'
import OpenLayersView from 'ol/View'
import {Tile as TileLayer} from 'ol/layer'
import {OSM} from 'ol/source'

export default {
  name: 'MapLayer',
  components: {

  },

  data () {
    return {
      map: null,
    }
  },
  created () {

  },
  mounted () {
    this.init();
  },
  methods: {
    init () {
      this.osmLayer = new TileLayer({
        source: new OSM()
      })
      this.map = new Map({
        layers: [this.osmLayer],
        target: 'map',
        view: new OpenLayersView({
          projection: 'CRS:84',
          center: [114, 30],
          zoom: 12
        })
      });
    },
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
</style>