<template>
  <div>
    <div v-html="column.Body"></div>
    <div>{{ column.slicedDate }}</div>
    <div>{{ column.Title }}</div>
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
.column-container {
  margin: 30px;

  .column-title {
    display: inline-block;
    font-size: 30px;
    border-bottom: 2px #556b9e solid;
    margin-bottom: 8px;
    position: relative;

    .column-title-date {
      display: inline-block;
      font-size: 12px;
      position: absolute;
      left: -10px;
      top: -12px;
    }
  }
}
</style>
