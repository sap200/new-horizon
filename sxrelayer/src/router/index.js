import { createRouter, createWebHistory } from 'vue-router'
import AssetBridgeView from '../views/AssetBridgeView.vue'
import LMView from '../views/LMView.vue'

const routes = [
  {
    path: '/',
    name: 'assetBridgeView',
    component: AssetBridgeView
  },
  {
    path: '/lmview/:myAddress',
    name: "LMView",
    component: LMView,
  },
  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  // }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
