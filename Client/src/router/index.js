import Vue from 'vue'
import VueRouter from 'vue-router'
import Auth from '@/components/auth'
import MapLayer from '@/components/maplayer'

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'auth',
    component: Auth
  },{
    path: '/map',
    name: 'map',
    component: MapLayer
  }
]

const router = new VueRouter({
  routes
});

export default router