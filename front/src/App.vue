<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn flat>
          <router-link to="/">
            Home
          </router-link>
        </v-btn>
        <v-btn flat v-if="!logged">
          <router-link to="/registration">
            Registration
          </router-link>
        </v-btn>
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

export default {
  name: 'App',
  data () {
    return {
      logged: false
    }
  }
}
</script>
