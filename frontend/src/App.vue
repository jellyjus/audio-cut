<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: "App",
    async created() {
        const res = await this.$api.login();
        console.log(res)
    },
    methods: {

    }
  }
</script>

<style></style>
