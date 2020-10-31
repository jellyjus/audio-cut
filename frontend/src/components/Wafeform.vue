<template>
    <div>
        <div>
            <canvas id="canvas1"></canvas>
        </div>
        <div>
            <canvas id="canvas2"></canvas>
        </div>
    </div>
</template>

<script>
    import WaveformData from 'waveform-data';
    export default {
        name: "Wafeform",
        async created() {
            // this.dat();
            this.json()
        },
        methods: {
            async dat() {
                const raw_data = await (await fetch('http://localhost:8080/audio.dat')).arrayBuffer()
                const waveform = WaveformData.create(raw_data);
                console.log(waveform)

                const scaleY = (amplitude, height) => {
                    const range = 256;
                    const offset = 128;

                    return height - ((amplitude + offset) * height) / range;
                }

                const canvas = document.getElementById('canvas1');
                const ctx = canvas.getContext('2d');
                console.log(canvas.height)
                ctx.beginPath();

                const channel = waveform.channel(0);
                console.log(channel)

// Loop forwards, drawing the upper half of the waveform
                for (let x = 0; x < waveform.length; x++) {
                    const val = channel.max_sample(x);

                    ctx.lineTo(x + 0.5, scaleY(val, canvas.height) + 0.5);
                }

// Loop backwards, drawing the lower half of the waveform
                for (let x = waveform.length - 1; x >= 0; x--) {
                    const val = channel.min_sample(x);

                    ctx.lineTo(x + 0.5, scaleY(val, canvas.height) + 0.5);
                }

                ctx.closePath();
                ctx.stroke();
                ctx.fill();
            },
            async json() {
                const raw_data = await (await fetch('http://localhost:8080/audio.json')).json()
                const waveform = WaveformData.create(raw_data);

                this.draw(raw_data.data, raw_data.data.length)
            },
            draw(json, width) {

                var c = document.getElementById("canvas2");
                var ctx = c.getContext("2d");
                c.height = 120;
                c.width = width;
                json.forEach(function(item, i, arr) {

                    // ctx.fillStyle = "white";
                    // ctx.fillRect(i * 3, c.height, 2, item - c.height)
                    ctx.fillStyle = "black";
                    // draw a small

                    let h2;
                    if (item <= json[i+1]) {
                        h2 = json[i+ 1];
                    } else {
                        h2 = item;
                    }

                    ctx.fillRect(i, c.height, 1, item* 0.3);
                });
            }
        }
    }
</script>

<style scoped>
    #canvas1 {
        width: 800px;
        height: 250px;
    }
</style>