<template>
  <div>
    <div class="gallery-menu">
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/images"
      >
        <span>images</span>
      </nuxt-link>
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/movies"
      >
        <span class="active">movies</span>
      </nuxt-link>
    </div>
    ぎゃらりー！ <br />動画のまとめページです <br />上のやつが新しい
    <div class="container">
      <a
        v-for="movie in movies"
        v-bind:key="movie.MovieUrl"
        target="_blank"
        :href="movie.TweetUrl"
      >
        <div class="content">
          <video
            class="content-video"
            @mouseenter="videoPlay($event, movie.Comment)"
            @mouseleave="videoPause"
            loop
            muted
            :src="movie.MovieUrl"
          ></video>
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
      title: "Movies - Starch Syrup Signal3",
      meta: [
        { hid: "description", name: "description", content: "動画一覧" },
        { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
        { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
        { hid: "og:description", property: "og:description", content: "動画一覧" },
        { hid: "og:title", property: "og:title", content: "Movies - Starch Syrup Signal3" },
        { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/temporary_twitter_card.png" }
      ]
    };
  },
  data() {
    return {
      movies: null
    };
  },
  mounted() {
    this.$store.commit("character/setShow", true);
    // キャラクターコメントの初期化
    this.commentInit();

    // APIから動画一覧読み込み
    fetch(process.env.SSSIGNAL_API_DOMAIN + "/gallery/movies", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(moviesJson => {
        this.movies = moviesJson;
      });
  },
  methods: {
    commentInit() {
      this.$store.commit("character/setComment", "");
    },
    isActive(path) {
      return path === this.$route.path;
    },
    videoPlay(e, comment) {
      e.target.play();

      this.$store.commit("character/setComment", comment);
    },
    videoPause(e) {
      e.target.pause();

      this.commentInit();
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/gallery";
</style>
