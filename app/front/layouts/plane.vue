<template>
  <div>
    <nuxt />
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  computed: {
  },
  methods: {
  }
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/defaultLayout";
</style>

<style lang="scss">
html {
  font-family: "Avenir", "Helvetica Neue", "Helvetica", "Arial", "Hiragino Sans",
    "ヒラギノ角ゴシック", YuGothic, "Yu Gothic", "メイリオ", Meiryo,
    "ＭＳ Ｐゴシック", "MS PGothic", sans-serif;
  font-size: 16px;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}
// html,
// body {
//   background-color: #000;
//   height: 100%;
// }
html,
body,
div {
  text-align: center;
  color: #bbbbbb;
  margin: 0px;
  padding: 0px;
}

a {
  color: rgb(54, 149, 222);

  &:hover {
    color: rgb(18, 111, 182);
  }
}
</style>
