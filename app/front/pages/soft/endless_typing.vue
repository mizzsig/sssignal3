<template>
    <div class="wrapper">
        <div v-show="stage === 'title'">
            <div style="font-size: 60px;">Endress typing</div>
            <div v-if="words.length === 0">now loading...</div>
            <div v-else>press space to play</div>
        </div>
        <div v-show="stage === 'play'">
            <div class="score" style="margin-bottom: 20px;">
                {{ typingTime }} sec<br>
                success: {{ typingData.success }}, failure: {{ typingData.failure }}<br>
                success rate: {{ successRate }}%<br>
                type rate: {{ typeRate }} / sec
            </div>
            <div style="position: relative; height: 70px;">
                <div class="scroll">
                    <div ref="scroll_left" :style="{maxWidth: '100%',  position: 'absolute', transform: 'translateX(-50%)', left: typingData.leftPosition, display: 'flex', gap: '0px 10px'}">
                        <div style="position: relative;" v-bind:class="{'typing-border': typingData.type.split(' ').length === 1}">
                            {{ showWords[0] }}
                            <div class="typing">
                                {{ typingData.type.split(' ')[0] }}
                            </div>
                        </div>
                        <div style="position: relative;" v-bind:class="{'typing-border': typingData.type.split(' ').length === 2}">
                            {{ showWords[1] }}
                            <div class="typing">
                                {{ typingData.type.split(' ')[1] }}
                            </div>
                        </div>
                    </div> 
                    <div ref="scroll_center" class="center" v-bind:class="{'typing-border': typingData.type.split(' ').length === 3}">
                        {{ showWords[2] }}
                        <div class="typing">
                            {{ typingData.type.split(' ')[2] }}
                        </div>
                    </div> 
                    <div ref="scroll_right" :style="{maxWidth: '100%', position: 'absolute', transform: 'translateX(-50%)', left: typingData.rightPosition, display: 'flex', gap: '0px 10px'}">
                        <div style="position: relative;">
                            {{ showWords[3] }}
                        </div>
                        <div style="position: relative;">
                            {{ showWords[4] }}
                        </div>
                    </div>
                </div>
            </div>
            <div>press space key between word and word</div>
        </div>
        <div v-show="stage === 'clear'">
            <div style="font-size: 20px;">Result</div>
            <div class="score" style="margin-bottom: 20px; font-size: 24px;">
                {{ typingTime }} sec<br>
                success: {{ typingData.success }}, failure: {{ typingData.failure }}<br>
                success rate: {{ successRate }}%<br>
                type rate: {{ typeRate }} / sec
            </div>
            <div>press space to title</div>
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
            // タイピングに使う全ての単語
            words: [],
            // 実際に画面に表示する単語
            showWords: [],
            // どの単語を入力中かを表す
            indexWords: 0,
            // 単語の何文字目を入力中かを表す
            indexInWord: 0,
            typingData: {
                // 経過時間(100ミリ秒単位)
                time: 0,
                timer: null,
                // 自分が入力した文字
                type: '',
                // 入力に成功した文字数
                success: 0,
                // 入力に失敗した文字数
                failure: 0,
                // 同じ文字で複数回失敗しても failure を増やさないためのフラグ
                failureFlag: false,
                // 入力文字の左・右の文字位置
                leftPosition: 'calc(50%)',
                rightPosition: 'calc(50%)',
            }
        });
    },
    computed: {
        typingTime() {
            return (this.typingData.time / 10).toFixed(1);
        },
        successRate() {
            if (this.typingData.success === 0) {
                return 0;
            } else {
                return ((this.typingData.success - this.typingData.failure) / this.typingData.success * 100).toFixed(2);
            }
        },
        typeRate() {
            if (this.typingData.time === 0) {
                return 0;
            } else {
                return (this.typingData.success / (this.typingData.time / 10)).toFixed(2);
            }
        }
    },
    mounted() {
        this.$store.commit('character/setShow', false);
        window.addEventListener('keydown', this.keydown);
        
        // APIから単語を取得
        fetch(`${process.env.SSSIGNAL_API_DOMAIN}/soft/endless_typing/all_words`, {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            } 
        })
        .then(response => response.json())
        .then(resultJson => {
            this.words = resultJson;
        })
    },
    beforeDestroy() {
        window.removeEventListener('keydown', this.keydown);
    },
    methods: {
        // 単語・スコアの初期化
        init() {
            // 単語は画面に5つ表示する
            this.showWords = [];
            for (let i = 0; i < 5; i++) {
                this.showWords.push(this.words[Math.floor(Math.random() * this.words.length)]);
            }
            this.indexWords = 0;
            this.indexInWord = 0;

            this.typingData.time = 0;
            this.typingData.type = '';
            this.typingData.success = 0;
            this.typingData.failure = 0;
            this.typingData.failureFlag = false;
            setTimeout(() => {
                this.calcWordPosition();
            }, 10); 

            // 経過時間
            clearInterval(this.typingData.timer);
            this.typingData.timer = setInterval(() => {
                this.typingData.time += 1;
            }, 100);
        },
        // 左右に表示する単語の位置を計算
        calcWordPosition() {
            this.typingData.leftPosition = `calc(50% - ${(this.$refs.scroll_left.offsetWidth / 2) + (this.$refs.scroll_center.offsetWidth / 2) + 10}px)`;
            this.typingData.rightPosition = `calc(50% + ${(this.$refs.scroll_right.offsetWidth / 2) + (this.$refs.scroll_center.offsetWidth / 2) + 10}px)`;
        },
        keydown(event) {
            if (this.stage === 'title') {
                if (event.key === ' ' && this.words.length > 0) {
                    this.stage = 'play';
                    this.init();
                }
            } else if (this.stage === 'play') {
                const key = event.key;
                if (key === 'Enter'){
                    clearInterval(this.typingData.timer);
                    this.stage = 'clear';
                    return;
                } else if (key === '/' || key.length > 1) {
                    return;
                } 
                
                // 入力中の単語
                const checkWord = this.showWords[this.indexWords];
                if (key === checkWord[this.indexInWord]) {
                    // 単語中の文字を打った時
                    this.indexInWord += 1;

                    this.typingData.type += key;
                    this.typingData.success += 1;
                    this.typingData.failureFlag = false;
                } else if (key === ' ' && this.indexInWord === checkWord.length) {
                    console.log('space');
                    // 単語と単語の間ではスペースを打つ
                    if (this.indexWords < 2) {
                        this.indexWords += 1;
                    } else {
                        // 先頭の単語を消して、末尾に単語を追加する
                        this.showWords.shift();
                        this.showWords.push(this.words[Math.floor(Math.random() * this.words.length)]);
                        // 入力した文字からも、先頭の単語を消す
                        this.typingData.type = this.typingData.type.split(' ');
                        this.typingData.type.shift();
                        this.typingData.type = this.typingData.type.join(' ');
                        setTimeout(() => {
                            this.calcWordPosition();
                        }, 1); 
                    }
                    this.indexInWord = 0;

                    this.typingData.type += key;
                    this.typingData.success += 1;
                    this.typingData.failureFlag = false;
                } else {
                    // 入力が誤っている時
                    if (!this.typingData.failureFlag) {
                        this.typingData.failure += 1;
                        this.typingData.failureFlag = true;
                    }
                }
                console.log(this.indexInWord, checkWord.length);
            } else if (this.stage === 'clear') {
                if (event.key === ' ') {
                    this.stage = 'title';
                    this.init();
                }
            }
        }
    }
}
</script>

<style lang="scss" scoped>
.wrapper {
  overflow: hidden;
  margin-top: 30px;
}

.scroll {
  font-size: 24px;
  font-family: monospace;
  position: absolute;
  user-select: none;
  width: 100%;

  .center {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
  }

  .typing {
    position: absolute;
    left: 0px;
    top: 0px;
    color: red;
  }

  .typing-border {
    padding-bottom: 3px;
    border-bottom: 1px solid red;
  }
}
</style>
