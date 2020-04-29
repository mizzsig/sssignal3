<template>
  <div class="container" @mousemove="moveCircle" @mousedown="toggleCircle">
    <div v-show="circle.isShow" :style="circleStyle"></div>
    <img id="main-img" :style="imageStyle" src="~/assets/top/scene1/kari.png" />
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  data() {
    return {
      circle: {
        defaultRadius: 300,
        radius: 0,
        isShow: true
      },
      mainImage: {
        width: 0,
        height: 0
      },
      circleStyle: {
        filter: "blur(30px)",
        position: "fixed",
        top: "0px",
        left: "0px",
        width: "",
        height: "",
        borderRadius: "",
        background:
          "radial-gradient(rgba(255, 255, 255, 0.2), rgba(255,255,255, 0.0))"
      },
      imageStyle: {
        position: "absolute",
        top: "0px",
        left: "0px",
        right: "0px",
        margin: "auto",
        maxWidth: "",
        maxHeight: ""
      }
    };
  },
  mounted() {
    if (this.$store.state.character.scene === 1) {
      // 初回のリサイズを実施
      window.addEventListener("resize", this.onResize);
      this.mainImage.width = document.getElementById("main-img").scrollWidth;
      this.mainImage.height = document.getElementById("main-img").scrollHeight;
      this.onResize();

      // 最初は光を左上に出す
      this.circleStyle.top = `${-this.circle.radius + 50}px`;
      this.circleStyle.left = `${-this.circle.radius}px`;
    }
  },
  beforeDestroy() {
    if (this.$store.state.character.scene === 1) {
      window.removeEventListener("resize", this.onResize);
    }
  },
  methods: {
    toggleCircle() {
      this.circle.isShow = !this.circle.isShow;
    },
    moveCircle(event) {
      this.circleStyle.top = `${event.clientY - this.circle.radius}px`;
      this.circleStyle.left = `${event.clientX - this.circle.radius}px`;
    },
    onResize() {
      // 画像がウィンドウからはみ出さないようにする
      this.imageStyle.maxWidth = `${window.innerWidth}px`;
      this.imageStyle.maxHeight = `${window.innerHeight - 50}px`;

      // 円の半径を更新
      this.circle.radius = this.circle.defaultRadius;
      if (this.circle.radius * 3 > window.innerWidth) {
        this.circle.radius = window.innerWidth / 3;
      }
      if (this.circle.radius * 3 > window.innerHeight) {
        this.circle.radius = window.innerHeight / 3;
      }
      // 半径を元に円のサイズ更新
      this.circleStyle.width = `${this.circle.radius * 2}px`;
      this.circleStyle.height = `${this.circle.radius * 2}px`;
      this.circleStyle.borderRadius = `${this.circle.radius}px`;
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  height: 100%;
  position: relative;
}
</style>
