<template>
    <div class="wrapper">
        <div class="main">
            <div v-show="stage === 'title'">
                <div>Endress typing</div>
                <div>press space to start</div>
            </div>
        </div>
        <div class="sub">
            <div v-show="stage === 'play'">
                <div class="score"></div>
                <div class="scroll">
                    {{ type }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    head() {
        return {
            title: "Endress typing - Starch Syrup Signal3",
            meta: [
                { hid: "description", name: "description", content: "Endress typing!" },
                { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
                { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
                { hid: "og:description", property: "og:description", content: "Endress typing!" },
                { hid: "og:title", property: "og:title", content: "Endress typing - Starch Syrup Signal3" },
                { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/about/twitter_card.png" }
            ]
        };
    },
    data() {
        return ({
            // title, main, result
            stage: 'title',
            type: ''
        });
    },
    mounted() {
        window.addEventListener('keydown', this.keydown);
    },
    beforeDestroy() {
        window.removeEventListener('keydown', this.keydown);
    },
    methods: {
        keydown(event) {
            if (this.stage === 'title') {
                if (event.key === ' ') {
                    this.stage = 'play';
                }
            } else if (this.stage === 'play') {
                let key = event.key;
                if (key === '/' || key.length > 1) {
                    return;
                } 
                
                // 判定処理を入れる

                if (key === ' ') {
                    key = '🚀';
                }
                
                // 画面表示
                this.type = key;
            }
        }
    }
}
</script>

<style lang="scss" scoped>
.wrapper {
    overflow: hidden;
}
</style>
