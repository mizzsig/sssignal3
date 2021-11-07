<template>
  <div>
    こらむ！
    <br />記事とか載せたりするページです
    <div>
      <div class="column" v-for="column in columns" v-bind:key="column.Date">
        <nuxt-link
          :to="'/column/' + column.Url + '/'"
          @mouseenter.native="
            $store.commit('character/setComment', column.CharacterComment)
          "
          @mouseleave.native="commentInit"
          class="column-inner"
        >
          <div>
            <div v-if="column.ImageUrl !== ''">
              <img class="thumbnail" :src="column.ImageUrl" />
            </div>
            <div v-if="column.ImageUrl === ''" style="padding: 5px 5px 0px 5px;"></div>
            <div class="date">{{ column.Date.slice(0, 10) }}</div>
            <div class="title">{{ column.Title }}</div>
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
      meta: [
        { hid: "description", name: "description", content: "記事一覧" },
        { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
        { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
        { hid: "og:description", property: "og:description", content: "記事一覧" },
        { hid: "og:title", property: "og:title", content: "Columns - Starch Syrup Signal3" },
        { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/temporary_twitter_card.png" }
      ]
    };
  },
  asyncData() {
    // APIから記事一覧読み込み
    return fetch(process.env.SSSIGNAL_API_DOMAIN + "/column/list", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(columnsJson => {
        const columns = columnsJson;
        return { columns: columns };
      });
  },
  mounted() {
    this.commentInit();
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
  max-width: 360px;
  background: rgb(87, 87, 87);
  box-sizing: border-box;
  margin: 10px;
  padding: 0px 0px 5px 0px;
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

  .thumbnail {
    width: 100%;
    height: 240px;
    object-fit: cover;
    border-radius: 2px;
  }
}

.date {
  text-align: left;
  font-size: 12px;
  padding-left: 10px;
}

.title {
  text-align: left;
  font-size: 20px;
  padding-left: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
