<template>
    <div class="container">
        <div class="audio-title">
            <img v-if="played" @click="triggerPlay(false)" class="icon" src="../../assets/icons/pause.svg"/>
            <img v-else @click="triggerPlay(true)" class="icon" src="../../assets/icons/play.svg"/>
            {{audio.artist}} - {{audio.title}}
        </div>
        <el-collapse-transition>
            <div class="audio-info" v-show="selected">
                <el-button type="text">Download</el-button>
                <el-button type="text" @click="editAudio">Edit</el-button>
            </div>
        </el-collapse-transition>
    </div>
</template>

<script>
    export default {
        name: "Audio",
        props: ['audio', 'selected'],
        data() {
            return {
                player: null,
                played: false
            }
        },
        destroyed() {
            this.triggerPlay(false);
            this.player = null;
        },
        methods: {
            triggerPlay(play) {
                if (!this.player)
                    return;

                if (play) {
                    this.player.play();
                    this.played = true;
                } else {
                    this.player.pause()
                    this.played = false;
                }
            },
            editAudio() {
                this.$router.push('/edit')
            }
        },
        watch: {
            selected(val) {
                if (val) {
                    if (!this.player)
                        this.player = new Audio(this.audio.url);
                        this.triggerPlay(true)
                } else {
                    this.triggerPlay(false);
                    this.player = null;
                }
            }
        }
    }
</script>

<style scoped>
    .audio-title {
        display: flex;
        align-items: center;
    }

    .icon {
        width: 25px;
        height: 25px;
        margin-right: 5px;
    }

    .audio-info {
        margin-left: 30px;
    }
</style>