<template>
    <el-row class="container">
        <el-col :span="12">
            <Audios class="card"></Audios>
        </el-col>
    </el-row>
</template>

<script>
    import Audios from "./audios/Audios";
    export default {
        name: "Home",
        components: {Audios},
        created() {

        },
        methods: {

        }
    }
</script>

<style scoped>
    .container {
        height: 100%;
        display: flex;
        flex-direction: column;
    }
</style>