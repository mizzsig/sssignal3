<template>
  <div>
    こらむ！
    <br />記事とか載せたりするページです <br />工事中
    <div>
      <div class="column" v-for="column in columns" v-bind:key="column.Date">
        <nuxt-link
          :to="'/column/' + column.Url"
          @mouseenter.native="
            $store.commit('character/setComment', column.CharacterComment)
          "
          @mouseleave.native="commentInit"
          class="column-inner"
        >
          <div>
            <div>{{ column.Title }}</div>
            <div>{{ column.Date.slice(0, 10) }}</div>
          </div>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "Columns - Starch Syrup Signal3",
      meta: [{ hid: "description", name: "description", content: "記事一覧" }]
    };
  },
  data() {
    return {
      columns: []
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
      this.$store.commit("character/setShow", true);
      this.$store.commit("character/setComment", "");
    }
  }
};
</script>

<style lang="scss" scoped>
.column {
  display: inline-block;
  width: calc(100% - 20px);
  max-width: 400px;
  background: rgb(87, 87, 87);
  box-sizing: border-box;
  margin: 10px;
  padding: 5px;
  transition: all 0.5s;
  border-radius: 2px;
  position: relative;

  &:hover {
    background: rgb(68, 88, 124);
  }

  // 左右
  &::before,
  &::after {
    content: "";
    position: absolute;
    background: rgb(124, 144, 180);
    width: 2px;
    height: 0px;
    border-radius: 2px;
    transition: all 0.25s;
    transition-delay: 0.25s;
  }

  &::before {
    right: 0px;
    bottom: 0px;
  }

  &::after {
    top: 0px;
    left: 0px;
  }

  &:hover::before,
  &:hover::after {
    height: 100%;
    transition-delay: 0s;
  }

  // 上下
  .column-inner::before,
  .column-inner::after {
    content: "";
    position: absolute;
    background: rgb(124, 144, 180);
    width: 0px;
    height: 2px;
    border-radius: 2px;
    transition: all 0.25s;
  }

  .column-inner::before {
    right: 0px;
    top: 0px;
  }

  .column-inner::after {
    bottom: 0px;
    left: 0px;
  }

  &:hover .column-inner::before,
  &:hover .column-inner::after {
    width: 100%;
    transition-delay: 0.25s;
  }

  .column-inner {
    text-decoration: none;
  }
}
</style>
