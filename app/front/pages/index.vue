<template>
  <section class="container">
    <scene1 v-if="scene === 1"></scene1>
    <scene2 v-if="scene === 2"></scene2>
  </section>
</template>

<script>
import Vue from 'vue';
import Scene1 from '../components/index/scene1.vue';
import Scene2 from '../components/index/scene2.vue';

export default {
    components: {
        Scene1,
        Scene2,
    },
  data() {
    return {
      scene: 1,
      components: 1 // 本番はここシーンの数にする
    }
  },
  head() {
    return {
      title: "sssignal3 Teaser Website"
    };
  },
  beforeMount() {
      this.scene = Math.floor(Math.random() * this.components) + 2; // 本番は1にする
  },
  methods: {
  }
};
</script>

<style lang="scss" scoped>
// .container {
//   display: inline-flex;
//   max-height: 100%;
// }
// img {
//   max-width: 100%;
//   max-height: 100%;
// }

// .button {
//     width: 100px;
//     height: 20px;
//     cursor: pointer;
//     border: 1px solid #777777;
//     border-radius: 4px;
//     background-color: #555555;
// }
</style>
