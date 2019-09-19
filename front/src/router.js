import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Registration from './components/Registration'
import Login from './components/Login'
import ArtistsList from "./components/ArtistsList"

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    }, {
      path: '/registration',
      name: 'registration',
      component: Registration
    }, {
      path: '/login',
      name: 'login',
      component: Login
    }, {
      path: '/artists',
      name: 'artists',
      component: ArtistsList
    }
  ]
})
