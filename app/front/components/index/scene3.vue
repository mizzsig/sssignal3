<template>
  <div class="container">
    <div class="movable-object" contenteditable="true">aaa</div>
    <div class="movable-object" contenteditable="true">bbb</div>
    <br />
    <div class="movable-object">ccc</div>
    <br />
    <div class="movable-object logo-string" contenteditable="true">水飴信号３</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      movingTarget: undefined,
      isDrag: false,
      mouseX: 0,
      mouseY: 0
    };
  },
  mounted() {
    window.addEventListener("mousedown", this.mousedown);
    window.addEventListener("mouseup", this.mouseup);
    window.addEventListener("mousemove", this.mousemove);
  },
  beforeDestroy() {
    window.removeEventListener("mousedown", this.mousedown);
    window.removeEventListener("mouseup", this.mouseup);
    window.removeEventListener("mousemove", this.mousemove);
  },
  methods: {
    mousedown(event) {
      this.movingTarget = event.target;
      this.isDrag = true;
      this.mouseX = event.clientX;
      this.mouseY = event.clientY;
      if (this.movingTarget.style.top === "") {
        this.movingTarget.style.top = "0px";
        console.log(this.movingTarget.style.top);
      }
      if (this.movingTarget.style.left === "") {
        this.movingTarget.style.left = "0px";
        console.log(this.movingTarget.style.left);
      }
    },
    mouseup() {
      this.movingTarget = undefined;
      this.isDrag = false;
    },
    mousemove(event) {
      console.log(event.target);
      if (
        this.isDrag &&
        this.movingTarget.className.indexOf("movable-object") > -1
      ) {
        this.movingTarget.style.left =
          String(
            parseInt(this.movingTarget.style.left) + event.clientX - this.mouseX
          ) + "px";
        this.movingTarget.style.top =
          String(
            parseInt(this.movingTarget.style.top) + event.clientY - this.mouseY
          ) + "px";
        this.mouseX = event.clientX;
        this.mouseY = event.clientY;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  position: relative;
  overflow: hidden;
  height: 10000px;

  .movable-object {
    cursor: move;
    position: relative;
    display: inline-block;
    user-select: none;

    // JSのイベントを起こすのはmovable-objectクラスだけにさせる
    div {
      pointer-events: none;
      user-select: none;
      display: inline-block;
    }
  }
}

.logo-string {
  font-size: 40px;
}
</style>