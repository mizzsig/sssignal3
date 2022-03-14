<template>
    <div>
        <div style="display: flex; user-select: none;">
            <p style="margin-left: 10px;">midi_keyboard_realtime_score</p>
            <div style="margin-left: 20px;">
                <div>Octave</div>
                <div style="display: flex;">
                    <div @click="staffBaseValueDown()" style="cursor: pointer; margin: 3px;">-</div>
                    <div style="margin: 3px;">{{ staff.baseValue }}</div>
                    <div @click="staffBaseValueUp()" style="cursor: pointer; margin: 3px;">+</div>
                </div>
            </div>
            <div style="margin-left: 20px;">
                <div>Signature</div>
                <div style="display: flex;">
                    <div @click="staffBothSignatureDown()" style="cursor: pointer; margin: 3px;">-</div>
                    <div style="margin: 3px;">{{ staff.above.signature }}</div>
                    <div @click="staffBothSignatureUp()" style="cursor: pointer; margin: 3px;">+</div>
                </div>
            </div>
        </div>
        <div style="position: relative;">
            <!-- ここでv-forして音符を置いていく -->
            <div v-for="note in staff.notes" v-bind:key="note.key" v-bind:class="note.className" :style="note.style">
                <img width="24" :src="`/soft/midi_keyboard_realtime_score/note_top${note.index}.png`">
                <!-- 音符にシャープorフラットをつける場合はここでつける -->
                <div v-if="getSharpFlatMap()[note.pitch % 12] !== 0">
                    <img style="position: relative; width: 30px; left: -26px; top: -75px;" :src="`/soft/midi_keyboard_realtime_score/${getSharpFlatMap()[note.pitch % 12] > 0 ? 'sharp' : 'flat'}.png`">
                </div>
            </div>
            <div class="staff">
                <img v-for="index in [1, 2, 3, 4, 5]" v-bind:key="`staff1-${index}`" :src="`/soft/midi_keyboard_realtime_score/line${index}.png`" :style="{ marginTop: `${staff.lineMargin * (index - 1)}px`, minWidth: `${100}%`, height: `${10}px` }">
                <img :src="`/soft/midi_keyboard_realtime_score/${staff.above.clef}.png`" @click="toggleClef('above')" :style="{ cursor: 'pointer' }">
                <div :style="`position: relative; top: ${staff.above.signature > 0 ? staff.style.sharpTop[sig-1] : staff.style.flatTop[sig-1]}px; left: ${staff.above.signature > 0 ? staff.style.sharpLeft[sig-1] : staff.style.flatLeft[sig-1]}px;`" v-for="sig of Math.abs(staff.above.signature)" v-bind:key="`above-signature-${sig}`">
                    <img style="width: 60px;" :src="`/soft/midi_keyboard_realtime_score/${staff.above.signature > 0 ? 'sharp' : 'flat'}.png`">
                </div>
            </div>
            <div class="staff" :style="{ marginTop: `${staff.lineMargin * 6}px`}">
                <img v-for="index in [1, 2, 3, 4, 5]" v-bind:key="`staff2-${index}`" :src="`/soft/midi_keyboard_realtime_score/line${index}.png`" :style="{ marginTop: `${staff.lineMargin * (index - 1)}px`, minWidth: `${100}%`, height: `${10}px` }">
                <img :src="`/soft/midi_keyboard_realtime_score/${staff.below.clef}.png`" @click="toggleClef('below')" :style="{ cursor: 'pointer' }">
                <div :style="`position: relative; top: ${staff.below.signature > 0 ? staff.style.sharpTop[sig-1] : staff.style.flatTop[sig-1]}px; left: ${staff.below.signature > 0 ? staff.style.sharpLeft[sig-1] : staff.style.flatLeft[sig-1]}px;`" v-for="sig of Math.abs(staff.below.signature)" v-bind:key="`above-below-${sig}`">
                    <img style="width: 60px;" :src="`/soft/midi_keyboard_realtime_score/${staff.above.signature > 0 ? 'sharp' : 'flat'}.png`">
                </div>
            </div>
        </div>
        <video autoplay :style="{ marginTop: `${staff.lineMargin * 14}px`}" :srcObject.prop="video.srcObject" :width="video.width" :height="video.height"></video>
    </div>
</template>

<script>
export default {
    data() {
        return ({
            midiDevices: {
                inputs: {},
                outputs: {}
            },
            video: {
                srcObject: '',
                width: 360,
                height: 240
            },
            staff: {
                // 基準になるドの音（treble_clefなら一番下, base_clefなら一番上）12の倍数
                baseValue: 60,
                lineMargin: 24,
                // 設定された調に対して、どの鍵盤がドからどれだけ離れているかを設定している
                // ♭の場合はkeyが+10されている
                pitchMap: {
                    0: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    1: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    2: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    3: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    4: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    5: [0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6],
                    11: [0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6],
                    12: [0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6],
                    13: [0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6],
                    14: [0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6],
                    15: [0, 1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6]
                },
                // pitchMapに置いた音符にシャープ・フラットをつけるかを設定している
                // １:♯ , 0:none , -1:♭
                sharpFlatMap: {
                    0: [0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0],
                    1: [0, 1, 0, 1, 0, -1, 0, 0, 1, 0, 1, 0],
                    2: [-1, 0, 0, 1, 0, -1, 0, 0, 1, 0, 1, 0],
                    3: [-1, 0, 0, 1, 0, -1, 0, -1, 0, 0, 1, 0],
                    4: [-1, 0, -1, 0, 0, -1, 0, -1, 0, 0, 1, 0],
                    5: [-1, 0, -1, 0, 0, -1, 0, -1, 0, -1, 0, 0],
                    11: [0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1],
                    12: [0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1],
                    13: [0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1],
                    14: [0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1],
                    15: [0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1]
                },
                above: {
                    clef: 'treble_clef',
                    // ♯1個で+1, ♭1個で-1
                    signature: 0,
                    // major or minor
                    key: 'minor'
                },
                below: {
                    clef: 'base_clef',
                    // ♯1個で+1, ♭1個で-1
                    signature: 0,
                    // major or minor
                    key: 'major'
                },
                // 画面に表示する音符を入れる配列
                notes: [],
                style: {
                    // CSSに指定するスタイル
                    sharpTop: [-20, 10, -30, 0, 30],
                    sharpLeft: [80, 100, 120, 140, 160],
                    flatTop: [10, -20, 25, -5, 35],
                    flatLeft: [80, 100, 120, 140, 160]
                }
            }
        });
    },
    async mounted() {
        this.$store.commit("character/setShow", false);
        
        if (navigator.requestMIDIAccess === undefined) {
            console.log("not midi defined");
        } else {
            navigator.requestMIDIAccess()
                .then(midiAccess => {
                    // https://qiita.com/ShoheiOno/items/34ae96c2563586982490
                    const inputIterator = midiAccess.inputs.values();
                    for (let input = inputIterator.next(); !input.done; input = inputIterator.next()) {
                        this.midiDevices.inputs[input.value.name] = input.value;
                        this.midiDevices.inputs[input.value.name].addEventListener('midimessage', () => {
                            this.deleteOldNote();
                            this.addNote();
                        }, false);
                    }

                    // const outputIterator = midiAccess.outputs.values();
                    // for (let output = outputIterator.next(); !output.done; output = outputIterator.next()) {
                    //     this.midiDevices.output[output.value.name] = output.value;
                    // }
                });
        }

        // webカメラに接続
        navigator.mediaDevices.getUserMedia({
            audio: false,
            video: {
                width: { ideal: this.video.width },
                height: { ideal: this.video.height }
            }
        }).then(stream => {
            this.video.srcObject = stream;
        });
    },
    beforeDestroy() {
        const tracks = this.video.srcObject.getTracks();
        tracks.forEach(function(track) {
            track.stop();
        });
        this.video.srcObject = null;
    },
    methods: {
        toggleClef(type) {
            this.staff[type].clef = this.staff[type].clef === 'treble_clef' ? 'base_clef' : 'treble_clef';
        },
        addNote() {
            // キーが押された場合のみ処理を行う
            if (event.data[0] !== 159 && event.data[2] === 0) {
                return;
            }

            // 押されたキーの音を表示用の配列に追加
            const timeStamp = Date.now();
            this.staff.notes.push({
                key: `${timeStamp}-${event.data[1]}`,
                timeStamp: timeStamp,
                index: Math.floor(Math.random() * Math.floor(5)) + 1,
                pitch: event.data[1],
                className: {
                    note: true
                },
                style: {
                    top: this.getNoteCssTop(event.data[1])
                }
            });

            // console.log(event.data);
        },
        deleteOldNote() {
            if (this.staff.notes.length === 0) {
                return;
            }
            const timeStamp = Date.now();
            if (timeStamp > this.staff.notes[this.staff.notes.length - 1].timeStamp + 10000) {
                this.staff.notes = [];
            }
        },
        // 音符の高さに応じたCSSの高さを返す
        getNoteCssTop(pitch) {
            // 基準になるドの音の高さ
            const baseTop = this.staff.lineMargin * 5;
            // 入力したキーとbaseTopとの差分
            let inputDiff;
            // 現在のコードだと、どの鍵盤がどのy軸に出てくるかを表す
            const pitchMap = this.getPitchMap();
            if (pitch >= this.baseValue) {
                inputDiff = 
                    this.staff.lineMargin * 
                    (
                        (pitchMap[pitch % 12] / 2) + 
                        Math.floor((pitch - this.staff.baseValue - 1) / 12) * 3.5
                    );
            } else {
                inputDiff = 
                    this.staff.lineMargin * 
                    (
                        -(pitchMap[pitch % 12] / 2) + 
                        Math.floor((this.staff.baseValue - pitch + 11) / 12) * 3.5
                    ) * -1;
            }
            // オクターブの高さ - ドレミの高さ - baseValueの高さ
            return `${baseTop - inputDiff}px`;
        },
        getPitchMap() {
            return this.staff.pitchMap[this.staff.above.signature >= 0 ? this.staff.above.signature : Math.abs(this.staff.above.signature) + 10];
        },
        getSharpFlatMap() {
            return this.staff.sharpFlatMap[this.staff.above.signature >= 0 ? this.staff.above.signature : Math.abs(this.staff.above.signature) + 10];
        },
        staffBaseValueDown() {
            if (this.staff.baseValue > 0) {
                this.staff.baseValue -= 12;
            }
        },
        staffBaseValueUp() {
            if (this.staff.baseValue < 120) {
                this.staff.baseValue += 12;
            }
        },
        staffBothSignatureDown() {
            if (this.staff.above.signature > -5) {
                this.staff.above.signature--;
                this.staff.below.signature--;
            }
        },
        staffBothSignatureUp() {
            if (this.staff.above.signature < 5) {
                this.staff.above.signature++;
                this.staff.below.signature++;
            }
        }
    }
}
</script>

<style lang="scss" scoped>
.staff {
  position: absolute;
  width: 100%;
  padding-top: 7px;

  img {
    position: absolute;
    left: 0px;
  }
}

.note {
  position: absolute;
  right: 120%;
  animation: note-animation 8s linear;
}
@keyframes note-animation {
  0% {
    right: -5%;
  }

  100% {
    right: 120%;
  }
}

.note-under {
  transform: translateY(calc(-100% + 34px));
}
</style>