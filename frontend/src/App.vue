<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: "App",
    async created() {
        // const res = await this.$api.login();
        // console.log(res)
    },
    methods: {

    }
  }
</script>

<style>
  #app {
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 100%;
    padding: 20px;
    box-sizing: border-box;
  }
</style>
