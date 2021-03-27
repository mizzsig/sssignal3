<template>
    <div class="wrapper">
        <video autoplay class="video-bg" :srcObject.prop="video.srcObject"></video>
        <div class="video-circle">
            <video autoplay :style="{ borderRadius: `${video.width / 2}px` }" :srcObject.prop="video.srcObject" :width="video.width" :height="video.height"></video>
            <div :style="{ position: 'absolute', top: '-5%', left: '50%', transform: 'translate(-50%, -50%)' }">aaaaa</div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return ({
            video: {
                srcObject: '',
                width: 300,
                height: 300
            }
        });
    },
    mounted() {
        // webカメラに接続
        navigator.mediaDevices.getUserMedia({
            audio: false,
            video: {
                width: { ideal: this.video.width },
                height: { ideal: this.video.height }
            }
        }).then(stream => {
            this.video.srcObject = stream;
        });
        this.$store.commit("character/setShow", false);
    },
    beforeDestroy() {
        const tracks = this.video.srcObject.getTracks();
        tracks.forEach(function(track) {
            track.stop();
        });
        this.video.srcObject = null;
    },
    methods: {}
}
</script>

<style lang="scss" scoped>
.wrapper {
    position: relative;
    overflow: hidden;
    width: 100%;
    height: calc(100vh - 50px);
}

.video-bg {
    position: absolute;
    left: 50%;
    top: 50%;
    min-width: 100%;
    min-height: 100%;
    transform: translate(-50%, -50%);
    filter: blur(7px) grayscale(80%);
    opacity: 0.4;
}

.video-circle {
    position: absolute;
    width: 300px;
    height: 300px;
    transform: translate(-50%, -50%);
    margin-top: 30%;
    margin-left: 30%;
}
</style>