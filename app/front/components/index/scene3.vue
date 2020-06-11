<template>
  <div class="container">
    <div id="app"></div>
  </div>
</template>

<script>
let PIXI;
if (process.client) {
  PIXI = require("pixi.js");
}

const SPEED = 6;
const SLOW_SPEED_RATE = 2 / 6;

export default {
  data() {
    return {
      isPress: {},
      canPressKey: [
        "Z",
        "DOWN",
        "ARROWDOWN",
        "UP",
        "ARROWUP",
        "LEFT",
        "ARROWLEFT",
        "RIGHT",
        "ARROWRIGHT",
        "SHIFT"
      ],
      moveX: {
        DOWN: 0,
        ARROWDOWN: 0,
        UP: 0,
        ARROWUP: 0,
        LEFT: -1,
        ARROWLEFT: -1,
        RIGHT: 1,
        ARROWRIGHT: 1
      },
      moveY: {
        DOWN: 1,
        ARROWDOWN: 1,
        UP: -1,
        ARROWUP: -1,
        LEFT: 0,
        ARROWLEFT: 0,
        RIGHT: 0,
        ARROWRIGHT: 0
      }
    };
  },
  beforeDestroy() {
    if (this.$store.state.character.scene === 3) {
      window.removeEventListener("resize", this.resize);
      window.removeEventListener("keydown", this.keydown);
      window.removeEventListener("keyup", this.keyup);
    }
  },
  mounted() {
    if (this.$store.state.character.scene === 3) {
      this.canPressKey.forEach(value => {
        this.isPress[value] = false;
      });

      window.addEventListener("resize", this.resize);
      window.addEventListener("keydown", this.keydown);
      window.addEventListener("keyup", this.keyup);

      this.app = new PIXI.Application({
        width: window.innerWidth, // スクリーン(ビュー)横幅
        height: window.innerHeight - 53, // スクリーン(ビュー)縦幅
        backgroundColor: 0x202020 // 背景色 16進 0xRRGGBB
      });
      // HTMLの<div id="app"></div>の中に上で作ったPIXIアプリケーション(app)のビュー(canvas)を突っ込む
      let appElement = document.getElementById("app");
      appElement.appendChild(this.app.view);

      // テキストオブジェクトを作る
      //   this.textobj = new PIXI.Text("Hello World!", {
      //     font: "60pt Arial",
      //     fill: "white"
      //   });
      //   this.textobj.position.x = this.app.screen.width / 2;
      //   this.textobj.position.y = this.app.screen.height / 2;
      //   this.app.stage.addChild(this.textobj);

      // 図形描画
      //   this.ellipse = new PIXI.Graphics()
      //     .beginFill(0x800000)
      //     .drawEllipse(0, 0, 30, 20)
      //     .endFill();
      //   this.ellipse.pivot.x = 15;
      //   this.ellipse.pivot.y = 10;
      //   this.ellipse.x = 100;
      //   this.ellipse.y = 100;
      //   this.ellipse.rotation = Math.PI / 6;
      //   this.app.stage.addChild(this.ellipse);

      // 画像を読み込み、テクスチャにする
      let blackTexture = new PIXI.Texture.from("top/scene2/black1.png");
      // 読み込んだテクスチャから、スプライトを生成する
      this.blackSprite = new PIXI.Sprite(blackTexture);
      // 基準点を設定(%) 0.5はそれぞれの中心 位置・回転の基準になる
      this.blackSprite.anchor.x = 0.5;
      this.blackSprite.anchor.y = 0.5;
      // 位置決め
      this.blackSprite.x = this.app.screen.width * 0.3;
      this.blackSprite.y = this.app.screen.height * 0.5;
      // 表示領域に追加する
      this.app.stage.addChild(this.blackSprite);

      // ビームのグラフィック
      this.beamGraphic = new PIXI.Graphics();

      // アニメーション開始
      this.app.ticker.add(() => {
        // this.textobj.rotation += 0.003;

        // 自機の移動
        this.mineMove();

        // ビームを発射
        this.drawBeam();
      });
    }
  },
  methods: {
    keydown(event) {
      const key = event.key.toUpperCase();
      if (!this.isPress[key]) {
        this.isPress[key] = true;
      }
    },
    keyup(event) {
      const key = event.key.toUpperCase();
      if (this.isPress[key]) {
        this.isPress[key] = false;
      }
    },
    resize() {
      this.app.renderer.resize(window.innerWidth, window.innerHeight - 53);

      if (this.blackSprite.x > this.app.screen.width) {
        this.blackSprite.x = this.app.screen.width;
      }
      if (this.blackSprite.y > this.app.screen.height) {
        this.blackSprite.y = this.app.screen.height;
      }
    },
    drawBeam() {
      // zキーでビーム出す
      if (this.isPress["Z"]) {
        this.beamGraphic.clear();

        this.beamGraphic
          .lineStyle(2, 0xd0d0d0)
          .moveTo(
            this.blackSprite.x,
            this.blackSprite.y + Math.random() * 10 - 5
          )
          .lineTo(
            this.app.screen.width,
            this.blackSprite.y + Math.random() * 10 - 5
          )
          .moveTo(
            this.blackSprite.x,
            this.blackSprite.y + Math.random() * 10 - 5
          )
          .lineTo(
            this.app.screen.width,
            this.blackSprite.y + Math.random() * 10 - 5
          )
          .moveTo(
            this.blackSprite.x,
            this.blackSprite.y + Math.random() * 10 - 5
          )
          .lineTo(
            this.app.screen.width,
            this.blackSprite.y + Math.random() * 10 - 5
          );

        this.app.stage.addChild(this.beamGraphic);
      } else {
        this.app.stage.removeChild(this.beamGraphic);
      }
    },
    mineMove(key) {
      // shiftキー押した時は減速
      let speedRate = SPEED;
      if (this.isPress["SHIFT"]) {
        speedRate *= SLOW_SPEED_RATE;
      }

      // Arrowキーを押しているか確認
      Object.keys(this.moveX).forEach(key => {
        if (this.isPress[key]) {
          if (
            this.blackSprite.x + this.moveX[key] * speedRate >= 0 &&
            this.blackSprite.x + this.moveX[key] * speedRate <=
              this.app.screen.width
          ) {
            this.blackSprite.x += this.moveX[key] * speedRate;
          }
          if (
            this.blackSprite.y + this.moveY[key] * speedRate >= 0 &&
            this.blackSprite.y + this.moveY[key] * speedRate <=
              this.app.screen.height
          ) {
            this.blackSprite.y += this.moveY[key] * speedRate;
          }
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped></style>
