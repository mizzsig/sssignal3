<template>
  <section class="container">
    <div v-for="(star, index) in stars" v-bind:key="index" :style="star"></div>
    <img class="main-img" src="/index/sss3_top.png" />
    <img class="logo" src="/index/sss3_top_logo.png" />
  </section>
</template>

<script>
export default {
  head() {
    return {
      title: "Starch Syrup Signal3",
      meta: [
        { hid: "description", name: "description", content: "ウェブサイトです" }
      ]
    };
  },
  data() {
    return {
      stars: [],
      starsValue: [],
      updateStars: null
    };
  },
  beforeMount() {
    this.$store.commit("character/setShow", false);

    this.updateStars = setInterval(() => {
      const starNum = 300;
      const starDiff = 6;
      const starCalc = starNum / starDiff / 2;

      // 古いの消す
      if (this.stars.length > starNum) {
        this.stars.splice(0, starDiff);
        this.starsValue.splice(0, starDiff);
      }
      // 新しいやつ追加
      for (let i = 0; i < starDiff; i++) {
        this.stars.push({
          width: "2px",
          height: "2px",
          background: "#A0A0A0",
          position: "fixed",
          top: `${Math.random() * window.innerHeight}px`,
          left: `${Math.random() * window.innerWidth}px`
        });
        this.starsValue.push(0);
      }
      // あるやつ更新
      this.starsValue.forEach((value, index) => {
        this.starsValue[index]++;
        this.stars[index].opacity =
          (starCalc - Math.abs(this.starsValue[index] - starCalc)) / starCalc;
      });
    }, 40);
  },
  beforeDestroy() {
    clearInterval(this.updateStars);
  },
  mounted() {},
  methods: {}
};
</script>

<style lang="scss" scoped>
.container {
  text-align: left;
  .logo {
    position: fixed;
    transform: translate(-50%, 0%);
    left: 50%;
    width: 70%;
    max-width: 700px;
    padding-top: 10px;
  }

  .main-img {
    position: fixed;
    object-fit: cover;
    width: 100%;
    height: 100%;
    min-height: 600px;
  }
}
</style>
