<template>
  <div>
    url
    <div>{{ column.Date }}</div>
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
      });
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
</style>