<template>
  <div id="app">
    <Header />
    <MapLayer />
    <Footer />
  </div>
</template>

<script>
import Header from '@/components/header'
import Footer from '@/components/footer'
import MapLayer from '@/components/maplayer'

export default {
  name: 'App',
  components: {
    Header,
    Footer,
    MapLayer
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
  margin-top: 60px;
}
</style>
