<template>
  <div>
    <div class="column-body" v-html="column.Body"></div>
  </div>
</template>

<script>
export default {
  async asyncData({ params }) {
    return fetch(
      process.env.SSSIGNAL_API_DOMAIN + "/column/" + params.url,
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
        const column = resultJson;

        // 記事の内容なかった時用の
        if (column.Url === "") {
          column.Body = "記事ないよー(泣)";
        }

        column.slicedDate = column.Date.slice(0, 10);
        return { column: column };
      });
  },
  head() {
    return {
      title: `${this.column.Title} - Starch Syrup Signal3`,
      meta: [
        { hid: "description", name: "description", content: this.column.CharacterComment },
        { hid: "og:description", name: "og:description", content: this.column.CharacterComment },
        { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
        { hid: "og:title", name: "og:title", content: `${this.column.Title} - Starch Syrup Signal3` },
        { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
        { hid: "og:image", name: "og:image", content: this.column.ImageUrl }
      ]
    };
  },
  data() {
    return {
    };
  },
  mounted() {
    console.log(this.column);
    if (!this.column.ShowCharacter) {
        this.$store.commit("character/setShow", false);
    } else {
        this.$store.commit("character/setShow", true);
        this.$store.commit(
            "character/setComment",
            this.column.CharacterComment
        );
    }
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/column.scss";
</style>
