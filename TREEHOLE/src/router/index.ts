import { createRouter, createWebHistory } from 'vue-router'
import homepage from '../components/homepage/homepage.vue';
import login from '../components/login/login.vue';
import idcard from '../components/idcard/idcard.vue';

const routes = [
  {
    path: '/home',
    name: 'homepage',
    component: homepage
  },
  {
    path: '/login',
    name: 'login',
    component: login
  },
  // {
  //   path:'/idcard',
  //   name: 'idcard',
  //   component:idcard
  // },
  {
    path: '/',
    redirect: '/home'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router