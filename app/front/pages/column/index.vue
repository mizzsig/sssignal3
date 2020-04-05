<template>
  <div>
    こらむ！
    <br />記事とか載せたりするページです
    <br />工事中
  </div>
</template>

<script>
export default {
  mounted() {
    this.commentInit();
  },
  methods: {
    commentInit() {
      this.$store.commit("character/setComment", "");
    }
  }
};
</script>

<style lang="scss" scoped>
</style>