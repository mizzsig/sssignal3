<template>
  <div>
    <header>
      <div class="header small">
        <div class="menu">
          <span>Starch Syrup Signal3</span>
        </div>
      </div>
      <div class="toggle-menu" :class="{ 'toggle-show': toggleShow }">
        <nuxt-link
          @click.native="$store.commit('character/setShow', false)"
          class="menu"
          to="/"
          >Top</nuxt-link
        >
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/about"
          >about</nuxt-link
        >
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/gallery/images"
          >gallery</nuxt-link
        >
        <nuxt-link
          @click.native="$store.commit('character/setShow', true)"
          class="menu"
          to="/column"
          >column</nuxt-link
        >
      </div>
      <div class="header wide">
        <nuxt-link
          @click.native="$store.commit('character/setShow', false)"
          class="menu"
          to="/"
        >
          <span v-bind:class="{ active: isActive('/') }"
            >Starch Syrup Signal3</span
          >
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
      <div style="height: 50px;"></div>
    </header>
    <div style="display:none; width: 0px; height: 0px;">
        <input type="text" id="nuxt-link-url">
        <button id="nuxt-link-button" @click="locationHref"></button>
    </div>
    <nuxt />
  </div>
</template>

<script>
export default {
  data() {
    return {
      toggleShow: false
    };
  },
  components: {},
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
    },
    locationHref() {
        this.$router.push(document.getElementById('nuxt-link-url').value);
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

<style>
html {
  font-family: 'Avenir','Helvetica Neue','Helvetica','Arial','Hiragino Sans','ヒラギノ角ゴシック',YuGothic,'Yu Gothic','メイリオ', Meiryo,'ＭＳ Ｐゴシック','MS PGothic',sans-serif;
  font-size: 16px;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}
</style>