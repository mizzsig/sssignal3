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
        v-bind:key="image.imageUrl"
        target="_blank"
        :href="image.TweetUrl"
      >
        <div class="content">
          <img class="content-img" :src="image.ImageUrl" />
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
    // APIから画像一覧読み込み
    fetch("http://api.sssignal.com/gallery/images", {
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
  }
};
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  margin: 0px 10px;

  a {
    display: contents;
    .content {
      width: 300px;
      height: 200px;
      margin: 3px;
      overflow: hidden;

      .content-img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        filter: contrast(70%);
        transition: all 0.3s;

        &:hover {
          filter: none;
          transform: scale(1.1);
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
