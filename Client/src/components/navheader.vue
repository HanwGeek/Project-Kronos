<template>
  <div id='header'>
    <v-navigation-drawer
      v-model="drawer"
      :width="show ? 256 : 56"
      permanent
      expand-on-hover
      app>

    <v-list-item class="px-2">
      <v-list-item-avatar>
        <v-img :src="avatar"></v-img>
      </v-list-item-avatar>
      <v-list-item-title>Layers</v-list-item-title>
    </v-list-item>
    
    <v-divider></v-divider>

    <v-list v-show="show">
      <v-list-item-group v-model="selected">
        <v-list-item
          v-for="(layer, i) in layers"
          input-value="false"
          :key="i">
        <v-list-item-action>
          <v-btn 
            @click="changeVisible(i)"
            icon>
            <v-icon
              color="blue darken-2"
              v-if="layer.show">
              mdi-eye
            </v-icon>
            <v-icon
              v-else>
              mdi-eye-off
            </v-icon>
          </v-btn>
        </v-list-item-action>
        <v-list-item-title v-text="layer.name"></v-list-item-title>
        <v-list-item-action>
          <v-btn
            @click="changeEdit(i)" 
            icon>
             <v-icon
              color="blue darken-2"
              v-if="layer.edit">
              mdi-pencil
            </v-icon>
            <v-icon
              v-else>
              mdi-pencil
            </v-icon>
          </v-btn>
        </v-list-item-action>
        </v-list-item>
      </v-list-item-group>
    </v-list>

    </v-navigation-drawer>

    <v-app-bar 
      short
      app>
      <v-toolbar-title>KRONOS</v-toolbar-title>

        <!-- <a id='avatar' href='https://hanwgeek.github.io/'>
          <img :src='avatar_wang'>
        </a>
        <a id='avatar_2' href='#'>
          <img :src='ava_chang'>
        </a>
        <a id='avatar_3' href='#'>
          <img :src='ava_yin'>
        </a>
        <a id='avatar_4' href='#'>
          <img :src='ava_guo'>
        </a> -->
    </v-app-bar>


  </div>
</template>


<script>
// import SideBar from "@/components/sidebar"

export default {
  name: 'NavHeader',
  components: {
    // SideBar
  },
  data () {
    return {
      drawer: null,
      show: false,
      selected: null,
      layers: [],
      title: 'KRONOS',
      avatar: require("../assets/greek.jpeg"),
      avatar_wang: require("../assets/wang.jpeg"),
      ava_guo: require("../assets/guo.jpg"),
      ava_yin: require("../assets/yin.jpg"),
      ava_chang: require("../assets/chang.jpg"),
    }
  },
  created() {
    this.show = false;
    this.$bus.$on("show-drawer", () => {
      this.show = true;
    });
    this.$bus.$on("layer-names", (layerNames) => {
      this.layers = layerNames.map(o => {return {
        "name": o,
        "show": false,
        "edit": false,
      }});
    })
  },
  methods: {
    changeVisible(idx) {
      this.layers[idx].show = !this.layers[idx].show;
      this.$bus.$emit("change-visible", idx);
    },
    changeEdit(idx) {
      this.layers[idx].edit = !this.layers[idx].edit;
      this.$bus.$emit("change-edit", idx);
    }
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
  right: 15px;
  height: 35px;
  border-radius: 50%;
}
#avatar_2 img{
  position: fixed;
  right: 65px;
  height: 35px;
  border-radius: 50%;
}
#avatar_3 img{
  position: fixed;
  right: 115px;
  height: 35px;
  border-radius: 50%;
}
#avatar_4 img{
  position: fixed;
  right: 165px;
  height: 35px;
  border-radius: 50%;
}
</style>