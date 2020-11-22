<template>
  <div class="auth" :style="styles">
    <v-card min-width="450px" outlined elevation="2">
      <v-card-title> 请登录 </v-card-title>
      <v-form>
        <v-container>
          <v-row align="center">
            <v-col align-self="center" cols="10" offset="1">
              <v-text-field
                label="Username"
                v-model="auth.username"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row align="center">
            <v-col cols="10" offset="1">
              <v-text-field
                label="Password"
                v-model="auth.password"
                :append-icon="showEye ? 'mdi-eye' : 'mdi-eye-off'"
                @click:append="showEye = !showEye"
                :type="showEye ? 'text' : 'password'"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-card-actions>
            <v-btn color="primary" @click="login" block>
              Login
            </v-btn>
          </v-card-actions>
        </v-container>
      </v-form>
    </v-card>
  </div>
</template>

<script>
import picture from '@/assets/background.jpg'

export default {
  name: 'auth',
  backImg: `url(${picture})`,
  data() {
    return {
      clientHeight: '',
      showEye: false,
      auth: {
        username: '',
        password: '',
      },
    }
  },
  computed: {
    styles: function() {
      return {
        backgroundImage: `url(${picture})`,
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'cover',
        height: this.clientHeight + 'px',
      }
    },
  },
  mounted() {
    this.clientHeight = `${document.documentElement.clientHeight}` - 64
  },
  methods: {
    login() {
      this.$bus.$emit('show-drawer')
      this.$router.push({
        name: 'map',
      })
    },
  },
}
</script>

<style scoped>
.auth {
  width: 100%;
  /* height: 100%; */
  display: flex;
  justify-content: center;
  align-items: center;
}

.el-card {
  width: 480px;
}
</style>
