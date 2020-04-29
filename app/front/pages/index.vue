<template>
  <section class="container">
    <scene1 v-show="scene === 1"></scene1>
    <scene2 v-show="scene === 2"></scene2>
    <scene3 v-show="scene === 3"></scene3>
  </section>
</template>

<script>
import { mapMutations } from "vuex";
import Scene1 from "../components/index/scene1.vue";
import Scene2 from "../components/index/scene2.vue";
import Scene3 from "../components/index/scene3.vue";

export default {
  components: {
    Scene1,
    Scene2,
    Scene3
  },
  data() {
    return {
      scene: 1,
      components: 3 // 本番はここシーンの数にする
    };
  },
  head() {
    return {
      title: "sssignal3 Teaser Website"
    };
  },
  beforeMount() {
    this.scene = Math.floor(Math.random() * this.components) + 1; // 本番は1にする
    // パラメータで指定されている時
    if (this.$route.query.scene !== undefined) {
      this.scene = parseInt(this.$route.query.scene);
    }

    this.$store.commit("character/setScene", this.scene);
  },
  methods: {}
};
</script>

<style lang="scss" scoped></style>
