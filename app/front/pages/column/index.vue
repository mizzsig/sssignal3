<template>
  <div>
    こらむ！
    <br />記事とか載せたりするページです
    <br />工事中
    <div>
      <div class="column" v-for="column in columns" v-bind:key="column.Date">
        <nuxt-link
          :to="'/column/' + column.Url"
          @mouseenter.native="$store.commit('character/setComment', column.CharacterComment)"
          @mouseleave.native="commentInit"
        >
          <div>
            <div>{{ column.Title }}</div>
            <div>{{ column.Date }}</div>
          </div>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      columns: null
    };
  },
  mounted() {
    this.commentInit();

    // APIから記事一覧読み込み
    fetch(process.env.SSSIGNAL_API_DOMAIN + "/column/list", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(columnsJson => {
        this.columns = columnsJson;
      });
  },
  methods: {
    commentInit() {
      console.log("init");
      this.$store.commit("character/setComment", "");
    }
  }
};
</script>

<style lang="scss" scoped>
.column {
  display: inline-block;
  width: 300px;
  background: rgb(87, 87, 87);
  margin: 10px;
  transition: all 0.2s;

  &:hover {
    background: rgb(98, 119, 156);
  }

  a {
    text-decoration: none;
  }
}
</style>