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
            @mouseenter="videoPlay"
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
  data() {
    return {
      movies: null
    };
  },
  mounted() {
    // APIから動画一覧読み込み
    fetch("http://api.sssignal.com/gallery/movies", {
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
    isActive(path) {
      return path === this.$route.path;
    },
    videoPlay(e) {
      e.target.play();
    },
    videoPause(e) {
      e.target.pause();
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-wrap: wrap;

  a {
    display: contents;

    .content {
      width: 300px;
      height: 200px;
      margin: 3px;
      overflow: hidden;

      .content-video {
        width: 100%;
        height: 100%;
        object-fit: cover;
        filter: contrast(70%);
        transition: all 0.3s;

        &:hover {
          filter: none;
        }
      }
    }
  }
}

.gallery-menu {
  margin-bottom: 10px;
}

.menu {
  text-decoration: none;
  color: #cccccc;
  font-size: 20px;
  padding: 10px;
  &:hover span::before,
  &:hover span::after {
    width: 50%;
  }
  .active::before,
  .active::after {
    width: 50%;
  }
  span {
    position: relative;
    &::before,
    &::after {
      border-bottom: 2px solid #bbbbbb;
      content: "";
      display: block;
      position: absolute;
      bottom: -3px;
      width: 0;
      transition: 0.2s all ease;
    }
    &::before {
      left: 50%;
    }
    &::after {
      right: 50%;
    }
  }
}
</style>
