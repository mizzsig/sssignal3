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
      this.$store.commit("character/setComment", "画像で振り返る水飴信号！");
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/gallery";
</style>
