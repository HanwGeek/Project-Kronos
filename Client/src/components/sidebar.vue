<template>
  <div id="side-bar" v-show="show">
    <el-drawer 
      :visible.sync="show"
      :direction="direction"
      :modal="false"
      :wrapper-closable="false">

    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'SideBar',
  data () {
    return {
      show: false,
      direction: 'ltr',
      layers: ["crop", "1"]
    }
  },
  created() {
    this.$bus.$on("showAside", (show) => {
      this.show = show;
    })
  },
  methods: {

  }
}
</script>

<style scoped>
#side-bar {
  width: 100px;
}

.el-drawer {
  height: 100px;
}

el-table {
  font-size: 15px;
}

.el-table-column {
  padding: 0px;
}
</style>