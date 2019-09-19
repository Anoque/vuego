<template>
  <div id="home-page">
    <h2>Home page</h2>
    <article id="home-page-blocks">
      <section id="home-page-links">
        <router-link to="/artists">Artists list</router-link>
      </section>
    </article>
  </div>
</template>

<script>
export default {
  name: 'home'
}
</script>
