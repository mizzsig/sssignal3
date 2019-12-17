<template>
    <div>
        <div>
            <nuxt-link @click.native="$store.commit('character/setShow', true)" class="menu" to="/gallery/images">
                <span v-bind:class="{ 'active': isActive('/gallery/images') }">images</span>
            </nuxt-link>
            <nuxt-link @click.native="$store.commit('character/setShow', true)" class="menu" to="/gallery/movies">
                <span v-bind:class="{ 'active': isActive('/gallery/movies') }">movies</span>
            </nuxt-link>
        </div>
        ぎゃらりー！<br>
        画像のまとめページです<br>
        工事中
        <div class="container">
            <a target="_blank" href="https://twitter.com/mizzsig/status/1162932319826534401">
                <img src="https://pbs.twimg.com/media/ECOQxriUwAA8Xuo?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1134831855834783749">
                <img src="https://pbs.twimg.com/media/D7-7NEaUIAYrKwc?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1077168890788990978">
                <img src="https://pbs.twimg.com/media/DvLfZEVUYAI9I9l?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1063043963903307776">
                <img src="https://pbs.twimg.com/media/DsCweFBVAAA-gvh?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1052568869502476288">
                <img src="https://pbs.twimg.com/media/Dpt5UWTV4AIDN7h?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1043504856290357249">
                <img src="https://pbs.twimg.com/media/DntBfz5U8AE9XQ4?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1029658956438765569">
                <img src="https://pbs.twimg.com/media/DkoVWc4UUAEDT1C?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1024280928850341889">
                <img src="https://pbs.twimg.com/media/Djb3rlPU0AE9tok?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/1023515096180346880">
                <img src="https://pbs.twimg.com/media/DjRBiq3V4AAWbV0?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/997127576450551808">
                <img src="https://pbs.twimg.com/media/DdaCCitUQAET23m?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/982957041613340673">
                <img src="https://pbs.twimg.com/media/DaQpMKIV4AEegPn?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/982957041613340673">
                <img src="https://pbs.twimg.com/media/DaQpOKzVQAQUVbZ?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/972127884792836097">
                <img src="https://pbs.twimg.com/media/DX2xHPhVoAAvSON?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/967769801966829568">
                <img src="https://pbs.twimg.com/media/DW41HM8UMAAn6A3?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/963402717744918528">
                <img src="https://pbs.twimg.com/media/DV6xg_zVAAIhNn-?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/952563371076825088">
                <img src="https://pbs.twimg.com/media/DTgt9B0VAAIRcAX?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/947682476951867393">
                <img src="https://pbs.twimg.com/media/DSbX-xnVQAAMM5O?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/947425463957209092">
                <img src="https://pbs.twimg.com/media/DSXuS3yVQAA0cpV?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/932110288282730496">
                <img src="https://pbs.twimg.com/media/DO-EqaDVAAAzVUG?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/922444428849045504">
                <img src="https://pbs.twimg.com/media/DM0uUaUVoAExHd8?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/911765730659414016">
                <img src="https://pbs.twimg.com/media/DKc9NUjVYAAJuWK?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/874267420138946560">
                <img src="https://pbs.twimg.com/media/DCIE_BcUAAAJZRp?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/863729318056214532">
                <img src="https://pbs.twimg.com/media/C_yUfltUMAEXgwY?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/863729318056214532">
                <img src="https://pbs.twimg.com/media/C_yUgB4UMAE8YTh?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/863729318056214532">
                <img src="https://pbs.twimg.com/media/C_yUgoxVYAEiF4b?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/863306938712338433">
                <img src="https://pbs.twimg.com/media/C_sR7JKVYAAjW4W?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/861200324509351936">
                <img src="https://pbs.twimg.com/media/C_OZIW2UAAE1R-p?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/858667215343702016">
                <img src="https://pbs.twimg.com/media/C-qZR_ZUwAA397s?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/853544021020426245">
                <img src="https://pbs.twimg.com/media/C9hlwEZUwAA5luw?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/815493945757081600">
                <img src="https://pbs.twimg.com/media/C1E2fxPUQAAEQlr?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/815139718660665348">
                <img src="https://pbs.twimg.com/media/C0_1CxoUAActeiJ?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/815139718660665348">
                <img src="https://pbs.twimg.com/media/C0_1DciVQAA_K1a?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/809028758128599040">
                <img src="https://pbs.twimg.com/media/Czo9NZiUkAAqS8B?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/798161973523279872">
                <img src="https://pbs.twimg.com/media/CxOjk21VEAAdah3?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/798161973523279872">
                <img src="https://pbs.twimg.com/media/CxOjlbnUsAEzPC_?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/796001796988682240">
                <img src="https://pbs.twimg.com/media/Cwv2kH7UkAIQpxe?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/793784114578993152">
                <img src="https://pbs.twimg.com/media/CwQVzqGUsAAGpBj?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/793784114578993152">
                <img src="https://pbs.twimg.com/media/CwQV0dBVYAAFCEj?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/793784114578993152">
                <img src="https://pbs.twimg.com/media/CwQV1d5UkAAZzZs?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/793071965569359872">
                <img src="https://pbs.twimg.com/media/CwGOugcVIAEfVWv?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/793071965569359872">
                <img src="https://pbs.twimg.com/media/CwGOvVVUEAAjq9o?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/778180322529771521">
                <img src="https://pbs.twimg.com/media/Csym2-MWIAA-OX7?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/777513158126604288">
                <img src="https://pbs.twimg.com/media/CspGVm0UMAE-8oH?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/758301753259364352">
                <img src="https://pbs.twimg.com/media/CoYGlZWUAAEW5-0?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/711549205030920192">
                <img src="https://pbs.twimg.com/media/Cd_uOxbUkAAAj2B?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/694831789399912448">
                <img src="https://pbs.twimg.com/media/CaSJ0qeUMAARbmE?format=jpg">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/688703945263325184">
                <img src="https://pbs.twimg.com/media/CY7ElioUoAISmOB?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/688703945263325184">
                <img src="https://pbs.twimg.com/media/CY7EloUUEAA4x_t?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/688703945263325184">
                <img src="https://pbs.twimg.com/media/CY7EjviUkAAxVNV?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/688703945263325184">
                <img src="https://pbs.twimg.com/media/CY7Eju9UkAA8JMG?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/688326370317873154">
                <img src="https://pbs.twimg.com/media/CY1tL9BUkAAegbe?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/685835747660673026">
                <img src="https://pbs.twimg.com/media/CYST-lpUQAAfa-l?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/682725495872458752">
                <img src="https://pbs.twimg.com/media/CXmHOHFUMAAxsgM?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/681661436184023040">
                <img src="https://pbs.twimg.com/media/CXW_dpHUwAAIzH0?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/681661436184023040">
                <img src="https://pbs.twimg.com/media/CXW_dnAUkAEr9Dt?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/681661436184023040">
                <img src="https://pbs.twimg.com/media/CXW_dpBVAAARCs2?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/680392553116471296">
                <img src="https://pbs.twimg.com/media/CXE9a9gUMAEW82N?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/679495076200845312">
                <img src="https://pbs.twimg.com/media/CW4NK7pVAAAPrJF?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/679495076200845312">
                <img src="https://pbs.twimg.com/media/CW4NK7nUQAAaPqi?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/667354219138019332">
                <img src="https://pbs.twimg.com/media/CULrIClUEAAV-gG?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/667354219138019332">
                <img src="https://pbs.twimg.com/media/CULrIBfUEAA2ZrQ?format=png">
            </a>
            <a target="_blank" href="https://twitter.com/mizzsig/status/666174148348014592">
                <img src="https://pbs.twimg.com/media/CT6521DVEAAuE7E?format=jpg">
            </a>
        </div>
    </div>
</template>

<script>
import { mapMutations } from 'vuex';

export default {
    methods: {
        isActive(path) {
            return path === this.$route.path;
        }
    }
}
</script>

<style lang="scss" scoped>
img {
    max-width: 300px;
    max-height: 300px;
}

.menu {
    text-decoration: none;
    color: #cccccc;
    font-size: 20px;
    padding: 10px;
    &:hover span::before, &:hover span::after {
        width: 50%;
    }
    .active::before, .active::after {
        width: 50%;
    }
    span {
        position: relative;
        &::before, &::after {
            border-bottom: 2px solid #BBBBBB;
            content: "";
            display: block;
            position: absolute;
            bottom: -3px;
            width: 0;
            transition: 0.2s all ease;
        }
        &::before {
            left: 50%;
        }
        &::after {
            right: 50%;
        }
    }
}
</style>