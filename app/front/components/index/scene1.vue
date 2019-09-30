<template>
    <div class="container" @mousemove="moveCircle">
        <div :style="circleStyle"></div>
        <img :style="imageStyle" src="~/assets/top/scene1/kari.png">
    </div>
</template>

<script>
export default {
    data() {
        return {
            circleStyle: {
                position: "fixed",
                top: "0px",
                left: "0px",
                width: "400px",
                height: "400px",
                borderRadius: "200px",
                background: "radial-gradient(rgba(255, 255, 255, 0.4), rgba(255,255,255, 0.0))"
            },
            imageStyle: {
                position: "absolute",
                top: "0px",
                left: "0px"
            }
        }
    },
    methods: {
        moveCircle(event) {
            this.circleStyle.top = `${event.clientY - 200}px`;
            this.circleStyle.left = `${event.clientX - 200}px`;
        },
        imageResize() {
            if (window.width > image.width) {

            }
            if (window.height > image.height) {

            }
        }
    }
}
</script>

<style lang="scss" scoped>
.container {
    width: 100%;
    height: 100%;
    position: relative;
}
</style>