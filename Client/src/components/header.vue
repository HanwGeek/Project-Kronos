<template>
  <div id='header'>
    <div id='title'>{{title}}</div>
    <div id='nav'>
      <a id='avatar' href='https://hanwgeek.github.io/'>
        <img :src='avatar'>
      </a>
      <a id='avatar_2' href='#'>
        <img :src='ava_chang'>
      </a>
      <a id='avatar_3' href='#'>
        <img :src='ava_yin'>
      </a>
      <a id='avatar_4' href='#'>
        <img :src='ava_guo'>
      </a>
    </div>
  </div>
</template>


<script>
export default {
  name: 'Header',
  data () {
    return {
      title: 'KRONOS',
      avatar: require("../assets/wang.jpeg"),
      ava_guo: require("../assets/guo.jpg"),
      ava_yin: require("../assets/yin.jpg"),
      ava_chang: require("../assets/chang.jpg"),
      isShowAside: false,
      btnType: 'info'
    }
  },
  methods: {
    updateMenuState() {
      this.isShowAside = !this.isShowAside;
      this.$bus.$emit("showAside", this.isShowAside);
      this.btnType = this.btnType == "info" ? "primary" : "info";
    },
  }
}
</script>

<style scoped>
#header {
  position: fixed;
  height: 40px;
  width: 100%;
  padding: 10px 50px;
  line-height: 40px;
  top: 0;
  left: 0;
  z-index: 100;
  background-color: #fff;
  box-shadow: 0 0 1px rgba(0,0,0,0.25);
}
#title {
  font-size: 24px;
  float: left;
  vertical-align: center;
}
.el-button {
  position: fixed;
  left: 10px;
  top: 15px;
  padding: 10px;
  /* height: 40px; */
}
#avatar img{
  position: fixed;
  right: 15px;
  height: 40px;
  border-radius: 50%;
}
#avatar_2 img{
  position: fixed;
  right: 65px;
  height: 40px;
  border-radius: 50%;
}
#avatar_3 img{
  position: fixed;
  right: 115px;
  height: 40px;
  border-radius: 50%;
}
#avatar_4 img{
  position: fixed;
  right: 165px;
  height: 40px;
  border-radius: 50%;
}
</style>