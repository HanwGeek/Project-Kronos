<template>
  <div id="app">
    <Header />
    
    <!-- <MapLayer /> -->
    <!-- <SideBar /> -->
    <router-view/>
 <v-footer padless>
    <v-col
      class="text-center"
      cols="12"
    >
      Powered by Vue.js & Gin
    </v-col>
  </v-footer>
  </div>
</template>

<script>
import Header from '@/components/header'
// import SideBar from '@/components/sidebar'

export default {
  name: 'App',
  components: {
    Header,
    // SideBar
  },
  data () {
    return {
      show: false
    }
  },
  created() {
    this.$bus.$on("showAside", () => {
      this.show = !this.show;
    })
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
  padding: 0px;
  margin: 0px;
  height: 100%;
}

.v-footer {
  font-size: 4px;
  /* padding: 0;
  margin: 0; */
}
</style>
