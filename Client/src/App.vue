<template>
  <v-app>
    <NavHeader /> 
 <!-- <v-footer padless>
    <v-col
      class="text-center"
      cols="12"
    >
      Powered by Vue.js & Gin
    </v-col>
  </v-footer> -->
   
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
import NavHeader from '@/components/navheader'
// import SideBar from '@/components/sidebar'

export default {
  name: 'App',
  components: {
    NavHeader,
    // SideBar
  },
  data () {
    return {
    }
  },
  created() {
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
