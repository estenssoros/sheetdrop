import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Callback from '../views/Callback.vue'
import Profile from '../views/Profile.vue'
import APIs from '../views/APIs.vue'
import API from '../views/API.vue'
import Orgs from '../views/Orgs.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/callback',
    name: 'callback',
    component: Callback
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile
  },
  {
    path: '/orgs',
    name: 'orgs',
    component: Orgs
  },
  {
    path: '/apis',
    name: 'apis',
    component: APIs
  },
  {
    path: '/api/:id',
    name: 'api',
    component: API
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.name == 'callback') { // check if "to"-route is "callback" and allow access
    next()
  } else if (to.name === 'home'){
    next()
  } else if (router.app.$auth.isAuthenticated()) { // if authenticated allow access
    next()
  } else { // trigger auth0 login
    router.app.$auth.login()
  }
})


export default router
