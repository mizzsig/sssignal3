<template>
  <div class="container">
    <div
      class="balloon"
      v-bind:class="{ 'balloon-character': isCharacter, 'balloon-advertisement': !isCharacter }"
    >
      <div
        class="reload"
        v-bind:class="{ 'pointer': isCharacter, 'nocursor': !isCharacter }"
        @mouseenter="setReloadHover(true)"
        @mouseleave="setReloadHover(false)"
      >
        <img v-show="!isReloadHover && isCharacter" src="~/assets/common/reload.gif" />
        <img v-show="isReloadHover && isCharacter" src="~/assets/common/reload-hover.gif" />
      </div>
      <div
        class="close pointer"
        @click="$store.commit('character/setCharacter', !isCharacter); setCloseHover(false);"
        @mouseenter="setCloseHover(true)"
        @mouseleave="setCloseHover(false)"
      >
        <img v-show="!isCloseHover" src="~/assets/common/close.gif" />
        <img v-show="isCloseHover" src="~/assets/common/close-hover.gif" />
      </div>
      <div class="comment" v-show="isCharacter">{{ comment }}</div>
      <div class="line" v-show="comment != ''"></div>ツイートと広告
      <br />(開いてるページへのコメント(できれば))
    </div>
    <div>
      <a href="https://twitter.com/mizzsig" target="_blank">
        <img src="~/assets/common/character.gif" />
      </a>
    </div>
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  data() {
    return {
      isReloadHover: false,
      isCloseHover: false
    };
  },
  computed: {
    isCharacter() {
      return this.$store.state.character.isCharacter;
    },
    comment() {
      return this.$store.state.character.comment;
    }
  },
  methods: {
    setReloadHover(isReloadHover) {
      this.isReloadHover = isReloadHover;
    },
    setCloseHover(isCloseHover) {
      this.isCloseHover = isCloseHover;
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  padding-right: 10px;
  box-sizing: border-box;
}

.pointer {
  cursor: pointer;
}

.nocursor {
  cursor: default;
}

.reload,
.close {
  position: absolute;
  width: 30px;
  height: 30px;
  border-radius: 4px;
  top: -15px;
}

.reload {
  right: 50px;
}

.close {
  right: 5px;
}

.balloon {
  background: rgb(206, 206, 206);
  color: rgb(0, 0, 0);
  border-radius: 4px;
  padding: 20px 0px;

  .comment {
    color: rgb(0, 0, 0);
    margin: 0px 10px;
  }

  .line {
    width: 95%;
    height: 1px;
    margin: 5px auto;
    background-image: linear-gradient(
      to right,
      #606060,
      #606060 6px,
      transparent 6px,
      transparent 12px
    );
    background-size: 12px 1px;
    background-repeat: repeat-x;
  }
}

.balloon-character {
  max-height: calc(100vh - 430px);
}

.balloon-advertisement {
  max-height: calc(100vh - 140px);
}

// 大きい
.balloon::before {
  content: "";
  background: rgb(206, 206, 206);
  position: absolute;
  display: block;
  border-radius: 50%;
  width: 20px;
  height: 10px;
  left: 260px;
  bottom: 370px;
}

// 小さい
.balloon::after {
  content: "";
  background: rgb(206, 206, 206);
  position: absolute;
  display: block;
  border-radius: 50%;
  width: 8px;
  height: 8px;
  bottom: 360px;
  left: 250px;
}
</style>
