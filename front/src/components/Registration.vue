<template>
  <AuthForm v-bind:is_auth="false"/>
</template>

<script>
import AuthForm from './AuthForm'

export default {
  name: 'Registration',
  components: {
    AuthForm
  }
}
</script>

<style scoped>

</style>
