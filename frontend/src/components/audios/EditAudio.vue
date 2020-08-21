<template>
    <div class="card">
        {{selectedAudio}}
        <div id="waveform"></div>
    </div>
</template>

<script>
    import {mapState, mapMutations} from 'vuex';
    import WaveSurfer from 'wavesurfer.js'
    import WaveSurferRegions from 'wavesurfer.js/dist/plugin/wavesurfer.regions'
    export default {
        name: "EditAudio",
        data() {
            return {
                wavesurfer: null,
                selectedAudio: null
            }
        },
        computed: {
            // ...mapState(['selectedAudio'])
        },
        created() {
            this.selectedAudio = {
                title: 'qwe',
                artist: 'sss',
                url: "https://cs1-74v4.vkuseraudio.net/p3/e2b68c3380f1c7.mp3?extra=Q4RpUHXzRX3sW7HgaVlrS802ke44HqWROJL4pgE_gJGagigriLmkXirsmCMI2713T63zrvdp7CWgb6sEU-hL7oAz8MHEauQUEYciBbraxgaDICqcMGB2e45spv-vN3wka5fkTz8II4beWeonwnSMAXM&long_chunk=1"
            }
        },
        mounted() {
            let region = null;
            this.wavesurfer = WaveSurfer.create({
                container: '#waveform',
                waveColor: '#dcd7d7',
                progressColor: '#42E2EEF1',
                normalize: true,
                responsive: true,
                plugins: [
                    WaveSurferRegions.create({})
                ]

            });
            this.wavesurfer.load(this.selectedAudio.url);

            this.wavesurfer.on('seek', val => {
                // region.play();
            })
            this.wavesurfer.on('ready', val => {
                region = this.wavesurfer.addRegion({
                    loop: true,
                    end: this.wavesurfer.getDuration(),
                    color: "rgba(0,0,0,0.06)"
                })
                // region.play();
            })

            this.wavesurfer.on('region-update-end', region => {
                // region.play()
            })
        }
    }
</script>

<style scoped>

</style>

<style>
    wave {
        overflow: visible !important;
    }

    .wavesurfer-region {
        z-index: 4 !important;
        top: -5% !important;
        height: 110% !important;
    }

    .wavesurfer-handle {
        max-width: 10px !important;
        width: 10px !important;
        background-color: #920890 !important;
        /*background-color: #00e0ff !important;*/
    }

    .wavesurfer-handle-start {
        border-top-left-radius: 5px;
        border-bottom-left-radius: 5px;
    }

    .wavesurfer-handle-end {
        border-top-right-radius: 5px;
        border-bottom-right-radius: 5px;
    }

    .wavesurfer-handle-end:before {
        content: '•\A•\A•';
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
        font-size: 12px;
        margin-top: 0;
        white-space: pre;
        line-height: 6px;
    }

    .wavesurfer-handle-end:hover .wavesurfer-handle-end:before {
        color: white !important;
    }
</style>