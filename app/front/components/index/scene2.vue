<template>
  <div class="container">
    <h1>水飴信号3 ~Starch Syrup Signal~</h1>
    <div class="keyboard-container" style="margin-bottom: -66px;">
      <div
        :id="'black' + index"
        v-for="(value, index) in this.keyboardsBlack"
        v-bind:key="index"
        :style="isBlackMargin(index)"
        class="keyboard black"
        @mousedown="playPiano(index, 'black')"
        @mouseup="pianoMouseUp(index, 'black')"
        @mouseleave="pianoMouseUp(index, 'black')"
      ></div>
    </div>
    <div class="keyboard-container">
      <div
        :id="'white' + index"
        v-for="(value, index) in this.keyboardsWhite"
        v-bind:key="index"
        class="keyboard white"
        @mousedown="playPiano(index, 'white')"
        @mouseup="pianoMouseUp(index, 'white')"
        @mouseleave="pianoMouseUp(index, 'white')"
      ></div>
    </div>
    <div class="score-container">
      <div
        class="music-note"
        v-for="key in notesKey"
        :key="key"
        :style="{ left: notes[key].left + 'px', top: notes[key].top + 'px' }"
      >
        <img :src="notes[key].imageUrl" />
      </div>
      <div class="score-line-red"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line"></div>
      <div class="score-line"></div>
      <div class="score-line"></div>
      <div class="score-line"></div>
      <div class="score-line"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
      <div class="score-line-dark"></div>
    </div>
  </div>
</template>

<script>
const whiteKeys = [
  "z",
  "x",
  "c",
  "v",
  "b",
  "n",
  "m",
  ",",
  ".",
  "/",
  "q",
  "w",
  "e",
  "r",
  "t",
  "y",
  "u",
  "i",
  "o",
  "p",
  "@",
  "["
];
const blackKeys = [
  "s",
  "d",
  "g",
  "h",
  "j",
  "l",
  ";",
  "2",
  "3",
  "4",
  "6",
  "7",
  "9",
  "0",
  "-"
];

export default {
  data() {
    return {
      isDrawingScore: false,
      isNoteMoving: false,
      score: [],
      // 音符の位置を記録
      notes: {},
      notesKey: [],
      // 音符のAudioクラスの配列
      notesAudio: [],
      windowWidth: 0,
      // 表示しているキーボードを格納する配列、押されているときはtrue、押されてないときはfales
      keyboardsBlack: [],
      keyboardsWhite: [],
      hiddenBlackKey: [2, 5, 7, 10, 12, 15, 17, 20]
    };
  },
  mounted() {
    // mp3のプリロード
    whiteKeys.forEach((value, index) => {
      this.notesAudio[`white${index}`] = new Audio();
      this.notesAudio[`white${index}`].src = `/top/scene2/white${index}.mp3`;
    });
    blackKeys.forEach((value, index) => {
      this.notesAudio[`black${index}`] = new Audio();
      this.notesAudio[`black${index}`].src = `/top/scene2/black${index}.mp3`;
    });

    // 画面サイズに応じて表示する鍵盤の数を変える
    window.addEventListener("resize", this.onResize);
    this.onResize();

    window.addEventListener("keydown", this.pianoKeyDown);
    window.addEventListener("keyup", this.pianoKeyUp);

    // 赤線が右端に行った時の処理
    const redLine = document.querySelector(".score-line-red");
    redLine.addEventListener("transitionend", this.redLineTransitionEnd);
  },
  beforeDestroy() {
    if (!process.client) return;
    window.removeEventListener("resize", this.onResize);
    window.removeEventListener("keydown", this.pianoKeyDown);
    window.removeEventListener("keyup", this.pianoKeyUp);
    const redLine = document.querySelector(".score-line-red");
    redLine.removeEventListener("transitionend", this.redLineTransitionEnd);
  },
  methods: {
    redLineTransitionEnd() {
      if (!this.isNoteMoving) {
        this.isNoteMoving = true;
        // notesを全て移動させる
        const scoreContainer = document.getElementsByClassName(
          "score-container"
        )[0];
        Object.keys(this.notes).forEach(
          key => (this.notes[key].left -= scoreContainer.clientWidth + 50)
        );
      }
    },
    // 画面サイズに応じた鍵盤を表示する
    onResize() {
      const countWhite =
        parseInt(window.innerWidth / 26) > 22
          ? 22
          : parseInt(window.innerWidth / 26);
      const blackKeyBoards = [
        0,
        0,
        1,
        2,
        2,
        3,
        4,
        5,
        5,
        6,
        7,
        7,
        8,
        9,
        10,
        10,
        11,
        12,
        12,
        13,
        14,
        15,
        15
      ];
      const countBlack = blackKeyBoards[countWhite];
      if (countWhite > this.keyboardsWhite.length) {
        const tmp = Array(countWhite - this.keyboardsWhite.length).fill(false);
        this.keyboardsWhite = this.keyboardsWhite.concat(tmp);
      } else if (countWhite < this.keyboardsWhite.length) {
        this.keyboardsWhite = this.keyboardsWhite.splice(0, countWhite);
      }
      if (countBlack > this.keyboardsBlack.length) {
        const tmp = Array(countBlack - this.keyboardsBlack.length).fill(false);
        this.keyboardsBlack = this.keyboardsBlack.concat(tmp);
      } else if (countBlack < this.keyboardsBlack.length) {
        this.keyboardsBlack = this.keyboardsBlack.splice(0, countBlack);
      }
    },
    // 黒鍵盤の左側にマージンを入れるかどうかの判定
    isBlackMargin(index) {
      const noMarginWhite = [4, 8, 11, 15, 18, 22];
      if (
        index === 0 &&
        noMarginWhite.indexOf(this.keyboardsWhite.length) >= 0
      ) {
        return { marginLeft: "-23px" };
      } else if (this.hiddenBlackKey.indexOf(index) >= 0) {
        return { marginLeft: "29px" };
      }
      return {};
    },
    // ピアノの音を再生する
    playPiano(index, color) {
      const redLine = document.getElementsByClassName("score-line-red")[0];
      const scoreContainer = document.getElementsByClassName(
        "score-container"
      )[0];

      // 初めて鍵盤を押した時は、楽譜の記載を始める
      if (!this.isDrawingScore) {
        this.isDrawingScore = true;

        // 赤線を右端まで移動させる
        redLine.style.left = scoreContainer.clientWidth + "px";
      }

      const timeStamp = Date.now();

      // 古い音符を消す
      //   this.notesKey = this.notesKey.filter(function(key) {
      //     const index = key.indexOf("-");
      //     const stamp = Number(key.substring(0, index));
      //     return stamp + 4000 > timeStamp;
      //   });
      //   this.notes = this.notes.filter(note => note.timestamp + 4000 > timeStamp);

      // 音符を設置する
      this.notesKey.push(`${timeStamp}-${color}${index}`);
      // 連想配列は単純に追加すると動かない！ので$setを使う
      this.$set(this.notes, `${timeStamp}-${color}${index}`, {
        imageUrl:
          "/top/scene2/note" + (Math.floor(Math.random() * 3) + 1) + ".png",
        top: (21 - index) * 13,
        left:
          redLine.getBoundingClientRect().left -
          scoreContainer.getBoundingClientRect().left -
          13,
        timestamp: timeStamp
      });

      // 赤線が右端に到達してる時は動かす
      if (this.isNoteMoving) {
        setTimeout(() => {
          this.notes[`${timeStamp}-${color}${index}`].left -=
            scoreContainer.clientWidth + 50;
        }, 100);
      }

      // 押されたキーに色をつける
      const key = document.getElementById(color + index);
      key.classList.add("keydown");

      // 音の再生 : 読み込み済みのAudioクラスのクローンを作って、そこから音を再生させる
      const clone = this.notesAudio[`${color}${index}`].cloneNode();
      clone.play();
    },
    // キーボードからピアノ鍵盤が押されたかを判定する
    pianoKeyDown(event) {
      // firefoxのクイック検索対策
      if (event.key === "/") {
        event.preventDefault();
      }

      if (
        !this.keyboardsWhite[whiteKeys.indexOf(event.key)] &&
        whiteKeys.indexOf(event.key) >= 0 &&
        whiteKeys.indexOf(event.key) < this.keyboardsWhite.length
      ) {
        this.keyboardsWhite[whiteKeys.indexOf(event.key)] = true;
        this.playPiano(whiteKeys.indexOf(event.key), "white");
      } else if (
        !this.keyboardsBlack[blackKeys.indexOf(event.key)] &&
        blackKeys.indexOf(event.key) >= 0 &&
        blackKeys.indexOf(event.key) < this.keyboardsBlack.length
      ) {
        this.keyboardsBlack[blackKeys.indexOf(event.key)] = true;
        this.playPiano(blackKeys.indexOf(event.key), "black");
      }
    },
    // キーボードからピアノ鍵盤が離されたかを判定する
    pianoKeyUp(event) {
      if (
        whiteKeys.indexOf(event.key) >= 0 &&
        whiteKeys.indexOf(event.key) < this.keyboardsWhite.length
      ) {
        const key = document.getElementById(
          "white" + whiteKeys.indexOf(event.key)
        );
        this.keyboardsWhite[whiteKeys.indexOf(event.key)] = false;
        key.classList.remove("keydown");
      } else if (
        blackKeys.indexOf(event.key) >= 0 &&
        blackKeys.indexOf(event.key) < this.keyboardsBlack.length
      ) {
        const key = document.getElementById(
          "black" + blackKeys.indexOf(event.key)
        );
        this.keyboardsBlack[blackKeys.indexOf(event.key)] = false;
        key.classList.remove("keydown");
      }
    },
    pianoMouseUp(index, color) {
      const key = document.getElementById(color + index);
      key.classList.remove("keydown");
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  overflow: hidden;
}

.keyboard-container {
  display: flex;
  justify-content: center;
}

.keyboard {
  margin: 3px;
  width: 20px;
}

.white {
  background: rgb(210, 210, 210);
  height: 100px;
}

.black {
  background: rgb(72, 72, 72);
  z-index: 3;
  height: 60px;
}

.keydown {
  background: rgb(255, 137, 137);
  margin-top: 6px;
  margin-bottom: 0px;
}

.score-container {
  width: 95%;
  margin: 0px auto;
  position: relative;

  .music-note {
    position: absolute;
    width: 26px;
    height: 26px;
    transition: all 4s linear;

    img {
      width: 26px;
      height: 26px;
    }
  }

  .score-line {
    background: rgb(140, 140, 140);

    margin-top: 20px;
    height: 5px;
  }

  .score-line-dark {
    background: rgb(50, 50, 50);

    margin-top: 20px;
    height: 5px;
  }

  .score-line-red {
    position: absolute;
    background: rgb(255, 124, 124);

    left: 0px;
    width: 3px;
    height: 100%;

    transition: all 4s linear;
  }
}
</style>
