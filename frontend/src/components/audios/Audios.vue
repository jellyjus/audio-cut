<template>
    <div>
        <ul class="list">
            <li class="list-item" :class="{selected: selectedAudio.id === audio.id}" v-for="(audio, index) in audios" @click="setSelectedAudio(audio)">
                <Audio :audio="audio" :selected="selectedAudio.id === audio.id"></Audio>
            </li>
        </ul>
    </div>
</template>

<script>
    import {mapState, mapMutations} from 'vuex';
    import Audio from "./Audio";
    export default {
        name: "Audios",
        components: {Audio},
        data() {
            return {
                audios: [],
            }
        },
        async created() {
            this.audios = await this.$api.getAudios();
            console.log(this.audios)
        },
        computed: {
            ...mapState(['selectedAudio'])
        },
        methods: {
            ...mapMutations(['setSelectedAudio'])
        }
    }
</script>

<style scoped>

</style>