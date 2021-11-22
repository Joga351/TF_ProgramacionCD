import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/Dashboard.vue';
import Buscar from '../views/Buscar.vue';

Vue.use(VueRouter)

const routes = [
  
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },

  {
    path: '/buscar/:id',
    name: 'Buscar',
    component: Buscar
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
