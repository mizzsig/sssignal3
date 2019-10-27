<template>
    <div>
        <h1>水飴信号3 ~Starch Syrup Signal~</h1>
        <div class="keyboard-container" style="margin-bottom: -66px;">
            <div :id="'black' + index" v-for="(value, index) in this.keyboardsBlack" v-bind:key="index" @mousedown="playPiano(index, 'black')" @mouseup="pianoMouseUp(index, 'black')" @mouseleave="pianoMouseUp(index, 'black')" :style="isBlackMargin(index)" class="keyboard black">
            </div>
        </div>
        <div class="keyboard-container">
            <div :id="'white' + index" v-for="(value, index) in this.keyboardsWhite" v-bind:key="index" @mousedown="playPiano(index, 'white')" @mouseup="pianoMouseUp(index, 'white')" @mouseleave="pianoMouseUp(index, 'white')" class="keyboard white">
            </div>
        </div>
    </div>
</template>

<script>
const whiteKeys = ['z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '@', '['];
const blackKeys = ['s', 'd', 'g', 'h', 'j', 'l', ';', '2', '3', '4', '6', '7', '9', '0', '-'];

export default {
    data() {
        return {
            windowWidth: 0,
            // 表示しているキーボードを格納する配列、押されているときはtrue、押されてないときはfales
            keyboardsBlack: [],
            keyboardsWhite: [],
            hiddenBlackKey: [2, 5, 7, 10, 12, 15, 17, 20],
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
            const blackKeyBoards = [0, 0, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9, 10, 10, 11, 12, 12, 13, 14, 15, 15];
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
            if (index === 0) {
                return { marginLeft: '-23px' };
            } else if (this.hiddenBlackKey.indexOf(index) >= 0) {
                return { marginLeft: '29px' };
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
            if (whiteKeys.indexOf(event.key) >= 0 && !this.keyboardsWhite[event.key]) {
                this.keyboardsWhite[event.key] = true;
                event.preventDefault();
                this.playPiano(whiteKeys.indexOf(event.key), 'white');
            } else if (blackKeys.indexOf(event.key) >= 0 && !this.keyboardsBlack[event.key]) {
                this.keyboardsBlack[event.key] = true;
                event.preventDefault();
                this.playPiano(blackKeys.indexOf(event.key), 'black');
            }
        },
        // キーボードからピアノ鍵盤が離されたかを判定する
        pianoKeyUp(event) {
            if (whiteKeys.indexOf(event.key) >= 0) {
                const key = document.getElementById('white' + whiteKeys.indexOf(event.key));
                this.keyboardsWhite[event.key] = false;
                key.classList.remove('keydown');
            } else if (blackKeys.indexOf(event.key) >= 0) {
                const key = document.getElementById('black' + blackKeys.indexOf(event.key));
                this.keyboardsBlack[event.key] = false;
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