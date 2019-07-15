<template>
  <AuthForm v-bind:is_auth="true"/>
</template>

<script>
import AuthForm from './AuthForm'

export default {
  name: 'Login',
  components: {
    AuthForm
  }
}
</script>

<style scoped>

</style>
