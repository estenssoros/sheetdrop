import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Orgs from '../views/Orgs.vue'
import Resource from '../views/Resource.vue'
import Schema from '../views/Schema.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/orgs/:id',
    name: 'orgs',
    component: Orgs
  },
  {
    path: '/resource/:id',
    name: 'resource',
    component: Resource
  },
  {
    path: '/schema/:id',
    name: 'schema',
    component: Schema
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
