<template>
  <el-collapse-transition>
  <div id="side-bar" v-show="show">

  </div>
  </el-collapse-transition>
</template>

<script>
export default {
  name: 'SideBar',
  data () {
    return {
      show: false
    }
  },
  created() {
  },

}
</script>

<style scoped>
#side-bar {
  width: 100px;
}
</style>