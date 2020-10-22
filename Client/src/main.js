import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import x2js from 'x2js'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false;
Vue.use(ElementUI);

Vue.prototype.$http = axios;
Vue.prototype.$bus = new Vue();
Vue.prototype.$x2js = new x2js();

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
