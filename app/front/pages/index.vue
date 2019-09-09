<template>
  <section class="container">
    <div class="button" @click="prevScene">←</div>
    {{ scene }}
    <div class="button" @click="nextScene">→</div>
    {{ render }}
  </section>
</template>

<script>
export default {
  data() {
    return {
      scene: 1,
      components: [],
      render: "<div>test</div>"
    }
  },
  head() {
    return {
      title: "sssignal3 Teaser Website"
    };
  },
  methods: {
    prevScene() {
      this.scene--;
      console.log(this.scene);
    },
    nextScene() {
      this.scene++;
      console.log(this.scene);
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  display: inline-flex;
  max-height: 100%;
}
img {
  max-width: 100%;
  max-height: 100%;
}

.button {
    width: 100px;
    height: 20px;
    cursor: pointer;
    border: 1px solid #777777;
    border-radius: 4px;
    background-color: #555555;
}
</style>
