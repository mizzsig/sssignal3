<template>
    <div class="container">
        <div class="balloon" v-bind:class="{ 'balloon-character': isCharacter, 'balloon-advertisement': !isCharacter }">
            <div class="reload" @mouseenter="setReloadHover(true)" @mouseleave="setReloadHover(false)">
                <img v-show="!isReloadHover" src="~/assets/common/reload.gif">
                <img v-show="isReloadHover" src="~/assets/common/reload-hover.gif">
            </div>
            <div class="close" @click="$store.commit('character/setCharacter', !isCharacter); setCloseHover(false);" @mouseenter="setCloseHover(true)" @mouseleave="setCloseHover(false)">
                <img v-show="!isCloseHover" src="~/assets/common/close.gif">
                <img v-show="isCloseHover" src="~/assets/common/close-hover.gif">
            </div>
            ツイートと広告<br>(開いてるページへのコメント(できれば))
        </div>
        <div>
            <a href="https://twitter.com/mizzsig" target="_blank">
                <img src="~/assets/common/character.gif">
            </a>
        </div>
    </div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
    data() {
        return {
            isReloadHover: false,
            isCloseHover: false
        }
    },
    computed: {
        isCharacter() {
            return this.$store.state.character.isCharacter;
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
}
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    padding-right: 10px;
    box-sizing: border-box;
}

.reload, .close {
    position: absolute;
    width: 30px;
    height: 30px;
    border-radius: 4px;
    top: -15px;
    cursor: pointer;
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
}

.balloon-character {
    max-height: calc(100vh - 430px);
}

.balloon-advertisement {
    max-height: calc(100vh - 140px);
}

// 大きい
.balloon::before {
    content: '';
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
    content: '';
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
