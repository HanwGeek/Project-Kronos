<template>
  <div id="side-bar" v-show="show">
    <v-navigation-drawer
      v-model="drawer"
      :width="show ? 256 : 56"
      permanent
      expand-on-hover
      app
    >
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: 'SideBar',
  data() {
    return {
      show: false,
      drawer: null,
      layers: ['crop', '1'],
    }
  },
  created() {
    this.$bus.$on('showAside', (show) => {
      this.show = show
    })
  },
  methods: {},
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
