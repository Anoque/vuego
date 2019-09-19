<template>
  <div id="artists-list">
    <h2>Artists</h2>
    <div class="list">
      <span v-for="item in getArtists">
        {{ item.Name }}
      </span>
    </div>
  </div>
</template>

<script>
  import Methods from '../methods'
  import { STORE } from '../store'

  export default {
    name: "ArtistsList",
    methods: {
      getData() {
        STORE.dispatch(Methods.LOAD_ARTISTS)
      }
    },
    computed: {
      getArtists: () => {
        return STORE.getters.getArtists
      }
    },
    mounted() {
      this.getData()
    }
  }
</script>

<style scoped>

</style>
