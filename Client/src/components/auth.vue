<template>
<div class="auth" :style="styles">
    <el-card>
      <div slot="header">
        请登录
      </div>
      <el-form :model="auth">
        <el-form-item>
          <el-input v-model="auth.username" placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="auth.password" placeholder="密码" show-password></el-input>
        </el-form-item>
         <el-form-item>
          <el-button type="primary" @click="login">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
</div>
</template>

<script>
import picture from '@/assets/background.jpg'

export default {
  name: 'auth',
  backImg: `url(${picture})`,
  data () {
    return {
      clientHeight: '',
      auth:{
        username: '',
        password: ''
      }
    }
  },
  computed: {
    styles: function() {
      return {
        backgroundImage: `url(${picture})`,
        backgroundRepeat: "no-repeat",
        backgroundSize: "cover",
        height: this.clientHeight + 'px'
      }
    }
  },
  mounted() {
    this.clientHeight = `${document.documentElement.clientHeight}` - 85;
  },
  methods: {
    login() {
      this.$bus.$emit("showMenuButton");
      this.$router.push({
        name: 'map'
      });
    },

  }
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