<template>
  <div>
    url
    <div>{{ column.slicedDate }}</div>
    <div>{{ column.Title }}</div>
    <div>{{ column.CharacterComment }}</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      column: []
    };
  },
  mounted() {
    fetch(
      process.env.SSSIGNAL_API_DOMAIN + "/column/" + this.$route.params.url,
      {
        method: "GET",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        }
      }
    )
      .then(response => response.json())
      .then(resultJson => {
        this.column = resultJson;
        this.$store.commit(
          "character/setComment",
          this.column.CharacterComment
        );
        this.column.slicedDate = this.column.Date.slice(0, 10);
      });
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
</style>