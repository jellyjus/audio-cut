<template>
  <div id="app">
    <Wafeform/>
<!--    <router-view/>-->
  </div>
</template>

<script>

  import Wafeform from "./components/Wafeform";
  export default {
    name: "App",
    components: {Wafeform},
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
