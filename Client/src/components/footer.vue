<template>
  <div id='footer'>
    <div id='title'>{{title}}</div>
  </div>
</template>


<script>
export default {
  name: 'Footer',
  data () {
    return {
      title: "Powered By Vue.js & Gin",
    }
  },
  methods: {
  }
}
</script>

<style scoped>
#footer {
  position: fixed;
  height: 8px;
  width: 100%;
  padding: 10px 0px;
  line-height: 8px;
  left: 0;
  bottom: 0;
  z-index: 100;
  background-color: #fff;
  box-shadow: 0 0 1px rgba(0,0,0,0.25);
}
#title {
  font-size: 12px;
  vertical-align: center;
}
</style>