<template>
    <div class="card">
        {{selectedAudio}}
        <div id="waveform">
            <v-stage :config="configKonva">
                <v-layer>
                    <v-line v-for="(line, index) in lines" :key="index" :config="line"></v-line>
                </v-layer>
            </v-stage>
        </div>
    </div>
</template>

<script>
    import {mapState, mapMutations} from 'vuex';
    import WaveformData from 'waveform-data';
    export default {
        name: "EditAudio",
        data() {
            return {
                selectedAudio: null,
                configKonva: {
                    width: 800,
                    height: 300
                },
                lines: []
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
        async mounted() {
            const canvas = document.getElementById('waveform-canvas')
            const waveform = await this.getAudioData();
            console.log(waveform)

            const scaleY = (amplitude, height) => {
                const range = 256;
                const offset = 128;

                return height - ((amplitude + offset) * height) / range;
            }

            const channel = waveform.channel(0);
            const line1 = {
                points: [],
                stroke: 'red',
                strokeWidth: 2,
                lineCap: 'round',
                lineJoin: 'round',
                tension: 1,
            }
            const line2 = {
                points: [],
                stroke: 'red',
                strokeWidth: 2,
                lineCap: 'round',
                lineJoin: 'round',
                tension: 1,
            }
            for (let x = 0; x < waveform.length; x++) {
                line1.points.push(x);
                line1.points.push(scaleY(channel.max_sample(x), this.configKonva.height));
                line2.points.push(x);
                line2.points.push(scaleY(channel.min_sample(x), this.configKonva.height));
            }
            line1.points = this.interpolateArray(line1.points, this.configKonva.width);
            line2.points = this.interpolateArray(line2.points, this.configKonva.width);
            this.lines.push(line1)
            this.lines.push(line2)
        },
        methods: {
            async getAudioData() {
                const audioContext = new AudioContext();

                const buffer = await (await fetch(this.selectedAudio.url)).arrayBuffer();
                return new Promise((res, rej) => {
                    WaveformData.createFromAudio({
                        audio_context: audioContext,
                        array_buffer: buffer,
                        scale: 1024
                    }, (err, waveform) => {
                        if (err)
                            rej(err);
                        else
                            res(waveform);
                    })
                })
            },
            interpolateArray(data, fitCount) {
                var linearInterpolate = function (before, after, atPoint) {
                    return before + (after - before) * atPoint;
                };

                var newData = [];
                var springFactor = Number((data.length - 1) / (fitCount - 1));
                newData[0] = data[0]; // for new allocation
                for ( var i = 1; i < fitCount - 1; i++) {
                    var tmp = i * springFactor;
                    var before = Number(Math.floor(tmp)).toFixed();
                    var after = Number(Math.ceil(tmp)).toFixed();
                    var atPoint = tmp - before;
                    newData[i] = linearInterpolate(data[before], data[after], atPoint);
                }
                newData[fitCount - 1] = data[data.length - 1]; // for new allocation
                return newData;
            }
        }
    }
</script>

<style scoped>

</style>

<style>
    #waveform, #waveform-canvas {
        width: 100%;
    }
</style>