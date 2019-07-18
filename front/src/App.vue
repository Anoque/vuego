<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-items class="hidden-sm-and-down">
        <v-btn>
          <router-link to="/">
            Home
          </router-link>
        </v-btn>
        <v-btn flat v-if="!logged">
          <router-link to="/login">
            Log in
          </router-link>
        </v-btn>
        <v-btn flat v-if="!logged">
          <router-link to="/registration">
            Sign up
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
/*function sendRequest(url, data) {
  axios.post(
    apiUrl + url, data, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    }
  ).then((response) => {
    console.log("Response: ", response)

    if (typeof response.error !== 'undefined') {
      console.error(response.error)

      return null
    } else if (typeof response.result !== 'undefined') {
      return response.result
    }
  }).catch((error) => {
    console.error(error)

    return null
  })
}*/

export default {
  name: 'App',
  data () {
    return {
      logged: false
    }
  }
}
</script>
