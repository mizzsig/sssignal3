<template>
  <div>
    <div class="gallery-menu">
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/images"
      >
        <span class="active">images</span>
      </nuxt-link>
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/movies"
      >
        <span>movies</span>
      </nuxt-link>
    </div>
    ぎゃらりー！ <br />画像のまとめページです <br />上のやつが新しい
    <div class="container">
      <a
        v-for="image in images"
        v-bind:key="image.ImageUrl"
        target="_blank"
        :href="image.TweetUrl"
      >
        <div class="content">
          <img
            class="content-img"
            :src="image.ImageUrl"
            @mouseenter="$store.commit('character/setComment', image.Comment)"
            @mouseleave="commentInit"
          />
        </div>
      </a>
    </div>
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  head() {
    return {
      title: "Images - Starch Syrup Signal3",
      meta: [
        { hid: "description", name: "description", content: "画像一覧" },
        { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
        { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
        { hid: "og:description", property: "og:description", content: "画像一覧" },
        { hid: "og:title", property: "og:title", content: "Images - Starch Syrup Signal3" },
        { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/temporary_twitter_card.png" }
      ]
    };
  },
  data() {
    return {
      images: null
    };
  },
  mounted() {
    this.$store.commit("character/setShow", true);
    // キャラクターコメントの初期化
    this.commentInit();

    // APIから画像一覧読み込み
    fetch(process.env.SSSIGNAL_API_DOMAIN + "/gallery/images", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(imagesJson => {
        this.images = imagesJson;
      });
  },
  methods: {
    commentInit() {
      this.$store.commit("character/setComment", "");
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/gallery";
</style>
