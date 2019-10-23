<template>
    <div>
        <h1>水飴信号3 ~Starch Syrup Signal~</h1>
        <div class="keyboard-container" style="margin-bottom: -66px;">
            <div :id="'black' + index" v-for="(value, index) in this.keyboardsBlack" v-bind:key="index" @mousedown="playPiano(index, 'black')" @mouseup="pianoMouseUp(index, 'black')" @mouseleave="pianoMouseUp(index, 'black')" :style="isBlackShow(index)" class="keyboard black">
            </div>
        </div>
        <div class="keyboard-container">
            <div :id="'white' + index" v-for="(value, index) in this.keyboardsWhite" v-bind:key="index" @mousedown="playPiano(index, 'white')" @mouseup="pianoMouseUp(index, 'white')" @mouseleave="pianoMouseUp(index, 'white')" class="keyboard white">
            </div>
        </div>
    </div>
</template>

<script>
const whiteKeyCodes = [90, 88, 67, 86, 66, 78, 77, 188, 190, 191, 167, 81, 87, 69, 82, 84, 89, 85, 73, 79, 80, 64];
const blackKeyCodes = [83, 68, 70, 71, 72, 74, 75, 76, 59, 58, 221, 50, 51, 52, 53, 54, 55, 56, 57, 48, 173, 160];

export default {
    data() {
        return {
            windowWidth: 0,
            keyboardsBlack: [],
            keyboardsWhite: [],
            hiddenBlackKey: [2, 6, 9, 13, 16, 20],
        }
    },
    mounted() {
        // 画面サイズに応じて表示する鍵盤の数を変える
        window.addEventListener('resize', this.onResize);
        this.onResize();

        window.addEventListener('keydown', this.pianoKeyDown);
        window.addEventListener('keyup', this.pianoKeyUp);
    },
    methods: {
        // 画面サイズに応じた鍵盤を表示する
        onResize() {
            const countWhite = parseInt(window.innerWidth / 26) > 22 ? 22 : parseInt(window.innerWidth / 26);
            const countBlack = countWhite - 1;
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
        // 黒鍵盤を見せるかどうかの判定
        isBlackShow(index) {
            if (this.hiddenBlackKey.indexOf(index) >= 0) {
                return { opacity: 0 };
            }
            return {};
        },
        // ピアノの音を再生する
        playPiano(index, color) {
            // 押されたキーに色をつける
            const key = document.getElementById(color + index);
            key.classList.add('keydown');

            // 音の再生
            const audio = new Audio();
            audio.src = `/top/scene2/${color}${index}.mp3`;
            audio.play();
        },
        pianoMouseUp(index, color) {
            const key = document.getElementById(color + index);
            key.classList.remove('keydown');
        },
        // キーボードからピアノ鍵盤が押されたかを判定する
        pianoKeyDown(event) {
            if (whiteKeyCodes.indexOf(event.keyCode) >= 0) {
                this.playPiano(whiteKeyCodes.indexOf(event.keyCode), 'white');
            } else if (blackKeyCodes.indexOf(event.keyCode) >= 0) {
                this.playPiano(blackKeyCodes.indexOf(event.keyCode), 'black');
            }
        },
        // キーボードからピアノ鍵盤が離されたかを判定する
        pianoKeyUp(event) {
            if (whiteKeyCodes.indexOf(event.keyCode) >= 0) {
                const key = document.getElementById('white' + whiteKeyCodes.indexOf(event.keyCode));
                key.classList.remove('keydown');
            } else if (blackKeyCodes.indexOf(event.keyCode) >= 0) {
                const key = document.getElementById('black' + blackKeyCodes.indexOf(event.keyCode));
                key.classList.remove('keydown');
            }
        }
    }
}
</script>

<style lang="scss" scoped>
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
}
</style>