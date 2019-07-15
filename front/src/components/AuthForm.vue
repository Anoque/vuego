<template>
  <div class="form-group" id="auth">
    <v-form v-model="valid">
      <h3>Form</h3>
      <v-text-field
        v-model="username"
        :rules="nameRules"
        :counter="64"
        label="User name"
        clearable
        required
      ></v-text-field>
      <v-text-field
        v-model="password"
        :rules="nameRules"
        :counter="64"
        label="Password"
        :type="'password'"
        clearable
        required
      ></v-text-field>
      <v-btn color="success" v-on:click="(is_auth) ? authUser() : createUser()">
        {{ (is_auth) ? 'Auth' : 'Create user' }}
      </v-btn>
    </v-form>
  </div>
</template>

<script>
  export default {
    name: "AuthForm",
    props: ['is_auth'],
    data: () => ({
      valid: false,
      username: '',
      password: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => v.length <= 64 || 'Name must be less than 10 characters'
      ]
    }),
    methods: {
      createUser () {
        this.$store.dispatch('createUser', {
          username: this.username,
          password: this.password
        })
      },
      authUser() {
        this.$store.dispatch('authUser', {
          username: this.username,
          password: this.password
        })
      }
    }
  }
</script>

<style scoped>

</style>
