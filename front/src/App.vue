<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-items class="hidden-sm-and-down">
        <router-link
          to="/"
          class="main-menu-link"
          v-ripple>
          Home
        </router-link>

        <router-link
          v-if="!logged"
          to="/login"
          class="main-menu-link"
          v-ripple>
          Log in
        </router-link>

        <router-link
          v-if="!logged"
          to="/registration"
          class="main-menu-link"
          v-ripple>
          Sign up
        </router-link>

        <span
          v-if="logged"
          class="main-menu-link"
          v-ripple
          v-on:click="logout()">
          Logout
        </span>
      </v-toolbar-items>
    </v-toolbar>

    <v-content>
      <v-container>
        <router-view/>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
  import Methods from './methods'
  import {STORE} from "./store"

  export default {
    name: 'App',
    methods: {
      logout: () => {
        STORE.commit(Methods.LOGOUT)
      }
    },
    computed: {
      logged: () => {
        return STORE.getters[Methods.GET_AUTH_STATUS]
      }
    }
  }
</script>
