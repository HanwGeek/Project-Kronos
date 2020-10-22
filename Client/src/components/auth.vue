<template>
<div class="auth" :style="{height:clientHeight + 'px'}">
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
export default {
  name: 'auth',
  data () {
    return {
      clientHeight: '',
      auth:{
        username: '',
        password: ''
      }
    }
  },
  mounted() {
    this.clientHeight = `${document.documentElement.clientHeight}` - 120;
  },
  methods: {
    login() {
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