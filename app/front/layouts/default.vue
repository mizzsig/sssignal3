<template>
  <div>
    <header>
      <div class="header small">
        <div class="menu">
          <span>Starch Syrup Signal3</span>
        </div>
      </div>
      <div class="toggle-menu" :class="{ 'toggle-show': toggleShow }">
        <nuxt-link @click.native="$store.commit('character/setShow', false)" class="menu" to="/">Top</nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/about"
        >about</nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/gallery/images"
        >gallery</nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/column"
        >column</nuxt-link>
      </div>
      <div class="header wide">
        <nuxt-link @click.native="$store.commit('character/setShow', false)" class="menu" to="/">
          <span v-bind:class="{ active: isActive('/') }">Starch Syrup Signal3</span>
        </nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/about"
        >
          <span v-bind:class="{ active: isActive('/about') }">about</span>
        </nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/gallery/images"
        >
          <span v-bind:class="{ active: isActive('/gallery') }">gallery</span>
        </nuxt-link>
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/column"
        >
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
$height: 50px;
.toggle-menu {
  display: flex;
  flex-direction: column;
  position: fixed;
  z-index: 20;
  width: 100%;
  top: 4 * -40px;
  transition: all 0.3s;

  .menu {
    display: flex;
    align-items: center;
    justify-content: center;
    text-decoration: none;
    color: #cccccc;
    background: rgba(60, 60, 60, 0.95);
    height: 40px;
    align-items: center;
    transition: all 0.2s;
  }

  .menu:hover {
    background: rgba(80, 80, 80, 0.95);
  }
}

.toggle-show {
  top: 0px;
}

.header {
  display: flex;
  position: fixed;
  width: 100%;
  height: $height;
  z-index: 10;
  padding-left: 40px;
  background-color: rgba(30, 30, 30, 0.9);

  .menu {
    display: flex;
    align-items: center;
    text-decoration: none;
    color: #cccccc;
    font-size: 20px;
    padding: 0px 10px;
    height: 100%;
    &:hover span::before,
    &:hover span::after {
      width: 50%;
    }
    .active {
      color: #6d82b3;
    }
    .active::before,
    .active::after {
      border-bottom: 2px solid #6d82b3;
      width: 50%;
    }
    span {
      display: block;
      position: relative;
      &::before,
      &::after {
        border-bottom: 2px solid #bbbbbb;
        content: "";
        display: block;
        position: absolute;
        bottom: -5px;
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
}
.small {
  padding-left: 0px;
  cursor: pointer;
  justify-content: center;
}
@media screen and (max-width: 500px) {
  .wide {
    display: none;
  }
}
@media screen and (min-width: 501px) {
  .small {
    display: none;
  }
}

.nuxt-show-character {
  width: calc(100% - 300px);
  height: calc(100% - 50px);
}
.nuxt-notshow-character {
  width: 100%;
  height: calc(100% - 50px);
}
.character {
  position: fixed;
  right: 0px;
  transition: all 0.3s;
  width: 300px;
}
.character-show {
  bottom: -40px;
}
.character-fade {
  bottom: -330px;
}
@media screen and (max-width: 899px) {
  .nuxt-show-character {
    width: 100%;
  }
  .character,
  .character-fade {
    display: none;
  }
}
@media screen and (max-height: 499px) {
  .nuxt-show-character {
    width: 100%;
  }
  .character,
  .character-fade {
    display: none;
  }
}
</style>

<style lang="scss">
html {
  font-family: "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI",
    Roboto, "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  /* word-spacing: 1px; */
  /* -ms-text-size-adjust: 100%; */
  /* -webkit-text-size-adjust: 100%; */
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
