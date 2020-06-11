<template>
  <div>
    <header>
      <div class="header small">
        <div class="menu">
          <span>menu</span>
        </div>
      </div>
      <div class="toggle-menu" :class="{ 'toggle-show': toggleShow }">
        <nuxt-link class="menu" to="/">Top</nuxt-link>
        <nuxt-link class="menu" to="/about">about</nuxt-link>
        <nuxt-link class="menu" to="/gallery/images">gallery</nuxt-link>
        <nuxt-link class="menu" to="/column">column</nuxt-link>
      </div>
      <div class="header wide">
        <nuxt-link class="menu" to="/">
          <span v-bind:class="{ active: isActive('/') }"
            >Starch Syrup Signal3</span
          >
        </nuxt-link>
        <nuxt-link class="menu" to="/about">
          <span v-bind:class="{ active: isActive('/about') }">about</span>
        </nuxt-link>
        <nuxt-link class="menu" to="/gallery/images">
          <span v-bind:class="{ active: isActive('/gallery') }">gallery</span>
        </nuxt-link>
        <nuxt-link class="menu" to="/column">
          <span v-bind:class="{ active: isActive('/column') }">column</span>
        </nuxt-link>
      </div>
    </header>
    <div style="height: 50px;"></div>
    <nuxt
      v-bind:class="{
        'nuxt-show-character': isShow,
        'nuxt-notshow-character': !isShow
      }"
    />
    <character
      v-show="isShow"
      class="character"
      v-bind:class="{
        'character-show': isCharacter,
        'character-fade': !isCharacter
      }"
    ></character>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
import Character from "../components/common/character.vue";

export default {
  data() {
    return {
      toggleShow: false
    };
  },
  components: {
    Character
  },
  computed: {
    // 右側にキャラクターを出すかどうか
    isShow() {
      return this.$store.state.character.isShow;
    },
    // キャラクターが出てるか広告が出てるか
    isCharacter() {
      return this.$store.state.character.isCharacter;
    }
  },
  methods: {
    isActive(path) {
      if (path === "/") {
        return path === this.$route.path;
      }
      return this.$route.path.indexOf(path) >= 0;
    },
    toggleMenu(event) {
      // メニューが出ていない & 画面幅が小さい & メニューを押下したときにトグルメニューを出す
      if (!this.toggleShow && window.innerWidth <= 500 && event.clientY <= 50) {
        this.toggleShow = true;
      } else {
        this.toggleShow = false;
      }
    }
  },
  mounted() {
    // トップページはキャラ出さない
    if (this.$route.path !== "/") {
      this.$store.commit("character/setShow", true);
    }
    window.addEventListener("click", this.toggleMenu);
  },
  beforeDestroy() {
    window.removeEventListener("click", this.toggleMenu);
  }
};
</script>

<style lang="scss" scoped>
@import "@/assets/scss/defaultLayout";
</style>

<style lang="scss">
html {
  font-family: "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI",
    Roboto, "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}
html,
body {
  background-color: #000;
  height: 100%;
}
html,
body,
div {
  text-align: center;
  color: #bbbbbb;
  margin: 0px;
  padding: 0px;
}

a {
  color: rgb(54, 149, 222);

  &:hover {
    color: rgb(18, 111, 182);
  }
}
</style>
