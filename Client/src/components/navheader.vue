<template>
  <div id="header">
    <v-navigation-drawer
      v-model="drawer"
      :width="show ? 256 : 56"
      permanent
      expand-on-hover
      app
    >
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
            :key="i"
          >
            <v-list-item-action>
              <v-btn @click="changeVisible(i)" icon>
                <v-icon color="blue darken-2" v-if="layer.show">
                  mdi-eye
                </v-icon>
                <v-icon v-else>
                  mdi-eye-off
                </v-icon>
              </v-btn>
            </v-list-item-action>

            <v-list-item-title v-text="layer.name"></v-list-item-title>

            <v-list-item-action>
              <v-dialog width="300" v-model="settingDlg">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn v-bind="attrs" @click="editSetting(i)" v-on="on" icon>
                    <v-icon>
                      mdi-cog-outline
                    </v-icon>
                  </v-btn>
                </template>

                <v-card>
                  <v-color-picker v-model="color" show-swatches>
                  </v-color-picker>

                  <v-divider></v-divider>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary" text @click="acceptSetting">
                      Accept
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-list-item-action>
            <v-list-item-action>
              <v-btn @click="changeEdit(i)" icon>
                <v-icon color="blue darken-2" v-if="layer.edit">
                  mdi-pencil
                </v-icon>
                <v-icon v-else>
                  mdi-pencil
                </v-icon>
              </v-btn>
            </v-list-item-action>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar short app>
      <v-toolbar-title>KRONOS</v-toolbar-title>

      <v-spacer></v-spacer>
      <v-btn @click="upload" icon>
        <v-icon>
          mdi-cloud-upload
        </v-icon>
      </v-btn>
      <input type="file" id="uploader" ref="uploader" @change="changeFile" />
      <v-btn @click="showOsm" icon>
        <v-icon color="blue darken-2" v-if="osmShow">
          mdi-map
        </v-icon>
        <v-icon v-else>
          mdi-map
        </v-icon></v-btn
      >
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
  data() {
    return {
      drawer: null,
      show: false,
      curIdx: null,
      settingDlg: false,
      color: null,
      ccolor: null,
      setting: {
        color: null,
      },
      osmShow: false,
      selected: null,
      layers: [],
      title: 'KRONOS',
      avatar: require('../assets/greek.jpeg'),
    }
  },
  created() {
    this.show = false
    this.$bus.$on('show-drawer', () => {
      this.show = true
    })
    this.$bus.$on('layer-names', (layerNames) => {
      this.layers = layerNames.map((o) => {
        return {
          name: o,
          show: false,
          edit: false,
        }
      })
    })
    this.$bus.$on('send-color', (c) => {
      this.color = c
    })
  },
  methods: {
    changeVisible(idx) {
      this.layers[idx].show = !this.layers[idx].show
      this.$bus.$emit('change-visible', idx)
    },
    changeEdit(idx) {
      this.layers[idx].edit = !this.layers[idx].edit
      this.$bus.$emit('change-edit', idx)
    },
    showOsm() {
      this.osmShow = !this.osmShow
      this.$bus.$emit('show-osm', this.osmShow)
    },
    editSetting(idx) {
      this.curIdx = idx
      this.$bus.$emit('get-color', idx)
    },
    acceptSetting() {
      this.settingDlg = false
      this.$bus.$emit('set-color', [this.curIdx, this.color])
    },
    upload() {
      this.$refs.uploader.click()
    },
    changeFile(e) {
      const fileToArrayBuffer = (file) =>
        new Promise((resolve, reject) => {
          var reader = new FileReader()
          reader.onload = () => {
            resolve(reader.result)
          }
          reader.onerror = reject
          reader.readAsArrayBuffer(file)
        })

      fileToArrayBuffer(e.target.files[0]).then((data) => {
        this.$bus.$emit('add-layer-data', data)
      })
    },
  },
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
  z-index: 10;
  background-color: #fff;
  box-shadow: 0 0 1px rgba(0, 0, 0, 0.25);
}
#title {
  font-size: 24px;
  float: left;
  vertical-align: center;
}

#uploader {
  height: 0;
  width: 0;
  visibility: hidden;
}

.float {
  position: relative;
  left: 10px;
  z-index: 100;
}
</style>
