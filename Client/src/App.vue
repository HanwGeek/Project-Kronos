<template>
  <div id="app">
    <Header />
    <!-- <MapLayer /> -->
        <!-- <router-view/> -->
      <el-container>
        <el-aside width="20vw" v-show="show">
         <el-collapse-transition> 
          <SideBar />
         </el-collapse-transition>
        </el-aside>
        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    <Footer />
  </div>
</template>

<script>
import Header from '@/components/header'
import Footer from '@/components/footer'
import SideBar from '@/components/sidebar'

export default {
  name: 'App',
  components: {
    Header,
    Footer,
    SideBar
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

.el-container {
  padding: 0px;
  margin: 0px;
}

.el-main {
  padding: 0px;
  margin: 0px;
}
</style>
