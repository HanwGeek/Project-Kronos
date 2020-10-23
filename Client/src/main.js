import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import x2js from 'x2js'
import router from './router'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import ElementUI from 'element-ui'

Vue.config.productionTip = false;
Vue.use(Vuetify);
Vue.use(ElementUI);

Vue.prototype.$http = axios;
Vue.prototype.$bus = new Vue();
Vue.prototype.$x2js = new x2js();

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
