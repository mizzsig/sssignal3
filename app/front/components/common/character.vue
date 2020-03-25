<template>
  <div class="container">
    <div
      class="balloon"
      v-bind:class="{ 'balloon-character': isCharacter, 'balloon-advertisement': !isCharacter }"
    >
      <div
        class="reload"
        v-bind:class="{ 'pointer': isCharacter, 'nocursor': !isCharacter }"
        @click="tweetLoad"
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
      <div class="line" v-show="comment !== '' && isCharacter"></div>
      <div v-show="isCharacter" class="tweet">
        <a
          :href="tweet.TweetUrl"
          target="_blank"
          @mouseenter="$store.commit('character/setComment', 'Twitterもぜひ見て')"
        >
          <div class="tweet-header">
            <img
              class="tweet-icon"
              src="https://pbs.twimg.com/profile_images/950738808911233024/H357rT4H_400x400.jpg"
            />
            <div class="tweet-name">とえら＠mizzsig</div>
          </div>
          <!-- \nを<br>してツイートのテキスト出力 -->
          <div class="tweet-text" v-show="tweet.Text !== undefined" v-html="tweet.Text"></div>
          <!-- 画像 -->
          <div v-show="tweet.ImageUrl !== undefined">
            <img class="tweet-content" :src="tweet.ImageUrl" />
          </div>
          <!-- 動画 -->
          <div v-show="tweet.MovieUrl !== undefined">
            <video class="tweet-content" :src="tweet.MovieUrl" autoplay loop muted></video>
          </div>
          <div class="tweet-date">{{ tweet.Date }}</div>
        </a>
      </div>
      <div class="advertisement" v-show="!isCharacter">
        <div class="advertisement-banner">
          <!-- ここに広告入れる API化したい -->
        </div>
        <div class="advertisement-banner">
          <!-- ここに広告入れる API化したい -->
        </div>
      </div>
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
      isCloseHover: false,
      tweet: []
    };
  },
  mounted() {
    this.tweetLoad();
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
    },
    // 最初に表示する時、更新ボタンを押された時にAPIリクエスト投げてtweet取得
    tweetLoad() {
      fetch("http://api.sssignal.com/tweet", {
        method: "GET",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        }
      })
        .then(response => response.json())
        .then(resultJson => {
          this.tweet = resultJson;
          this.tweet.Text = this.tweet.Text.replace(/\r?\n/g, "<br>");
        });
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
  padding: 20px 10px 5px 10px;
  overflow-y: scroll;

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

  .tweet {
    transition: all 0.2s;
    border-radius: 4px;

    &:hover {
      background: rgb(206, 219, 221);
      .tweet-name {
        color: rgb(45, 56, 158);
      }
    }
    a {
      text-decoration: none;
    }
    div {
      color: rgb(0, 0, 0);
    }

    .tweet-header {
      display: flex;
      align-items: center;
      margin-bottom: 5px;
    }

    .tweet-icon {
      width: 30px;
      height: 30px;
      border-radius: 8px;
      margin-left: 15px;
      margin-right: 5px;
    }

    .tweet-name {
      display: inline-block;
    }

    .tweet-text {
      font-size: 14px;
    }

    .tweet-content {
      max-width: 100%;
      max-height: 150px;
    }

    .tweet-date {
      font-size: 12px;
      float: right;
    }
  }

  .advertisement {
    .advertisement-banner {
      padding: 5px 0px;

      img {
        max-width: 100%;
      }
    }
  }
}

.balloon-character {
  max-height: calc(100vh - 440px);
}

.balloon-advertisement {
  max-height: calc(100vh - 150px);
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
