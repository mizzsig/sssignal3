<template>
  <div>
    <div class="column-body" v-html="column.Body"></div>
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

        // 記事の内容なかった時用の
        if (this.column.Url === "") {
          this.column.Body = "記事ないよー(泣)";
        }

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
@import "@/assets/scss/column.scss";
</style>
