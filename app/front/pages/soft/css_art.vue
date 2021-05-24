<template>
    <div>
        <div>CSS Paste Picture</div>
        <div class="main">
            <div class="sub-center">
                <div class="content sub-content-main" style="margin-bottom: 10px;">
                    <div v-for="(content, index) in contents" v-bind:key="`view-${content.key}`" :style="content" @mousedown="mouseDown(index, $event)"></div>
                </div>
                <textarea class="content" style="margin-bottom: 10px;" v-model="resultHtml" readonly></textarea>
            </div>
            <div class="sub-right" style="max-height: 100%;">
                <div class="content" style="padding: 10px; margin-bottom: 10px; display: flex; justify-content: space-between;">
                    <div class="icon" @mousedown="addDiv()">+</div>
                    <div class="icon" @mousedown="upDiv()">↑</div>
                    <div class="icon" @mousedown="downDiv()">↓</div>
                    <div class="icon" @mousedown="copyDiv()">copy</div>
                    <div class="icon" @mousedown="deleteDiv()">delete</div>
                </div>
                <div class="content" style="padding: 5px; margin-bottom: 10px;">
                    <div>[Parameter{{ drugIndex !== -1 ? ` : ${contents[drugIndex]["key"]}` : "" }}]</div>
                    <div v-if="drugIndex !== -1" style="max-height: 300px; overflow: scroll;">
                        <div v-for="(css, index) in contents[drugIndex]" v-bind:key="`content-css-${index}`">
                            <div v-if="!['position', 'key'].includes(index)" style="display: flex; overflow: hidden;">
                                <div style="margin-right: 5px;">{{ index }}</div>
                                <input type="text" v-model="contents[drugIndex][index]">
                            </div>
                        </div>
                        <div style="display: flex; overflow: hidden;">
                            <input type="text" v-model="addCssIndex" @keydown="addCssToDiv">
                            <input type="text" v-model="addCssValue" @keydown="addCssToDiv">
                        </div>
                    </div>
                </div>
                <div class="content" style="max-height: 200px; overflow: scroll;">
                    <div class="content" v-for="(content, index) in contents" v-bind:key="content.key" @mousedown="changeIndex(index)">
                        <div :style="(content === contents[drugIndex]) ? {background: '#4F4F8F'} : {}">{{ content.key }}</div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    head() {
        return {
            title: "CSS Paste Picture - Starch Syrup Signal3",
            // TODO : metaを直す
            meta: [
                { hid: "description", name: "description", content: "CSSのアートを作ろう！" },
                { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
                { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
                { hid: "og:description", property: "og:description", content: "CSSのアートを作ろう！" },
                { hid: "og:title", property: "og:title", content: "CSS Paste Picture - Starch Syrup Signal3" },
                { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/about/twitter_card.png" }
            ]
        };
    },
    data() {
        return ({
            contents: [],
            resultHtml: "",
            isDrug: false,
            addCssIndex: "",
            addCssValue: "",
            divIndex: 1,
            drugIndex: -1,
            drugBeforeX: 0,
            drugBeforeY: 0,
            reloadTimer: 0,
        });
    },
    mounted() {
        this.$store.commit("character/setShow", false);

        window.addEventListener("mousemove", this.mouseMove);
        window.addEventListener("mouseup", this.mouseUp);
        this.reloadTimer = setInterval(() => {
            const element = document.getElementsByClassName('sub-content-main')[0];
            this.resultHtml = '<div style="position: relative;">' + element.innerHTML + '</div>';
        }, 20);
    },
    beforeDestroy() {
        window.removeEventListener("mousemove", this.mouseMove);
        window.removeEventListener("mouseup", this.mouseUp);
        clearInterval(this.reloadTimer);
    },
    methods: {
        addDiv() {
            const template = {
                key: "",
                position: "absolute",
                top: "0px",
                left: "0px",
                width: "50px",
                height: "50px",
                background: "rgb(255, 255, 255)",
            };
            this.contents.push({
                ...template,
                key: `content-${this.divIndex++}`
            });
        },
        upDiv() {
            if (this.drugIndex <= 0) {
                return;
            }
            this.contents.splice(this.drugIndex - 1, 2, this.contents[this.drugIndex], this.contents[this.drugIndex - 1]);
            this.drugIndex--;
        },
        downDiv() {
            if (this.drugIndex < 0 || this.drugIndex == this.contents.length - 1) {
                return;
            }
            this.contents.splice(this.drugIndex, 2, this.contents[this.drugIndex + 1], this.contents[this.drugIndex]);
            this.drugIndex++;
        },
        addCssToDiv(event) {
            this.addCssIndex = this.addCssIndex.trim();
            this.addCssValue = this.addCssValue.trim();
            if (event.key !== "Enter" || this.drugIndex === -1 || this.addCssIndex === "" || this.addCssValue === "") {
                return;
            }
            if (!Object.keys(this.contents[this.drugIndex]).includes(this.addCssIndex)) {
                this.contents[this.drugIndex][this.addCssIndex] = this.addCssValue;
                this.addCssIndex = "";
                this.addCssValue = "";
            }
        },
        copyDiv() {
            if (this.drugIndex === -1) {
                return;
            }
            this.contents.push({
                ...this.contents[this.drugIndex],
                key: `content-${this.divIndex++}`,
                left: '0px',
                top: '0px'
            });
        },
        deleteDiv() {
            if (this.drugIndex === -1) {
                return;
            }
            this.contents = this.contents.filter(content => {
                return content.key !== this.contents[this.drugIndex].key;
            });
            this.drugIndex = -1;
        },
        mouseDown(index, event) {
            if (!this.isDrug) {
                this.isDrug = true;
                this.drugIndex = index;
                this.drugBeforeX = event.clientX;
                this.drugBeforeY = event.clientY;
            }
        },
        mouseMove(event) {
            if (this.isDrug) {
                const content = this.contents[this.drugIndex];
                this.contents[this.drugIndex]['left'] = `${parseInt(content['left']) + (event.clientX - this.drugBeforeX)}px`;
                this.contents[this.drugIndex]['top'] = `${parseInt(content['top']) + (event.clientY - this.drugBeforeY)}px`;
                this.drugBeforeX = event.clientX;
                this.drugBeforeY = event.clientY;
            }
        },
        mouseUp() {
            if (this.isDrug) {
                this.isDrug = false;
            }
        },
        changeIndex(index) {
            this.drugIndex = index;
        }
    }
}
</script>

<style lang="scss" scoped>
.main {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  margin: 10px auto;

  .sub {
    display: flex;
    margin: 10px;
    box-sizing: border-box;
    align-items: center;
    flex-direction: column;
    background: #4f4f4f;
  }

  .sub-center {
    @extend .sub;
    width: 60%;

    .sub-content-main {
      position: relative;
      overflow: hidden;
      width: 100%;
      &::before {
        content: "";
        display: block;
        padding-top: 100%;
      }
    }
    .content {
      box-sizing: border-box;
      width: 80%;
      max-width: 500px;
      border: 1px white solid;
      border-radius: 5px;
      resize: none;
    }
  }

  .sub-right {
    @extend .sub;
    width: 30%;

    .content {
      box-sizing: border-box;
      width: 100%;
      border: 1px white solid;
      border-radius: 5px;

      .icon {
        cursor: pointer;
      }

      input {
        width: 120px;
      }
    }
  }
}
</style>
