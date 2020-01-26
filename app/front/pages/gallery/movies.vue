<template>
  <div>
    <div class="gallery-menu">
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/images"
      >
        <span v-bind:class="{ 'active': isActive('/gallery/images') }">images</span>
      </nuxt-link>
      <nuxt-link
        @click.native="$store.commit('character/setShow', true)"
        class="menu"
        to="/gallery/movies"
      >
        <span v-bind:class="{ 'active': isActive('/gallery/movies') }">movies</span>
      </nuxt-link>
    </div>ぎゃらりー！
    <br />動画のまとめページです
    <br />工事中
    <div class="container">
      <a target="_blank" href="https://twitter.com/mizzsig/status/1011978515627905024">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/1011977212768088064/pu/vid/314x320/Xz_ZAj6iziXlYy--.mp4?tag=3"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/988796668492505089">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/988795633514430465/pu/vid/854x720/XJ2JYCuZqilhS9Oy.mp4?tag=3"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/976470789619728384">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/976469045372334080/pu/vid/526x640/yM7W7uLIZ5aaRoF1.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/876348107230855168">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/876348007758807040/pu/vid/480x360/3TqnL2Jg_vBMfT8b.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/871374292579082241">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/871373135332818944/pu/vid/1020x720/TY4JKoXU-2QxvDqq.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/869575942607323136">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/DBFabUUUIAA98PU.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/869219906461417472">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/DBAWRxOUIAEb7Y6.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/845977028562055168">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/845975028629876736/pu/vid/1140x720/QOaZCJ7K6ZqOF4YH.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/832570621435400193">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/832569328033689600/pu/vid/640x360/YjrouN4ofV9njW-I.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/805396312841592832">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/805395446621347840/pu/vid/638x360/nwGIouKh4Iywt4pq.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/796318274103353344">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/Cw0WmlrUQAEhwiT.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/795980723320590337">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/CwvjQzDUsAEGOlo.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/793783298971475968">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/CwQVS5aVIAACtgo.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/792308368492728320">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/Cv7X9lNVMAAjgXw.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/786587846052229120">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/CuqDTSzUIAAhEQp.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/784719310476283905">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/784716333267099648/pu/vid/480x640/3vtJUwL5OOfgOaPX.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/783630856870952960">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/783630618831708160/pu/vid/480x640/XhRZjEoT_npfwFkX.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/766603777231036417">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/766603013192507392/pu/vid/274x320/lnRN5nakEkFPdgI1.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/758300589553954818">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/CoYGSqrUIAAARxK.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/755040010504593412">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/tweet_video/CnpwCFTUAAAe3UY.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/730392899112730624">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/730388527729401856/pu/vid/478x360/_xNWlHrRjIKy9FBo.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/730017421768036354">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/730016754630414336/pu/vid/496x360/7wh4k6PQN4zsiZOL.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/729276539674628097">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/729276258337456129/pu/vid/384x360/KyN2Pz_ijtAwC6Ks.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/729276230483087360">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/729274014229635072/pu/vid/384x360/RXrnSYQg0jCOlOPR.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/704644649500061697">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/704631351035584512/pu/vid/440x360/PV9LbYwtXypbdivp.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/699580714996469760">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/699580371516559360/pu/vid/480x640/LawpdXd47rNICaaq.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/697786031035133953">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/697784652052824064/pu/vid/1280x720/WrmPLH1yAIedaktu.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/694884815187775489">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/694882097685618688/pu/vid/560x360/n2Qz7UCNduMyUhen.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/694524839877431300">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/694523263838351360/pu/vid/202x180/4lGbygc0F1YCBXx1.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/687250852436807680">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/687249817664290816/pu/vid/638x640/wvnHzoVjcuAEdKqq.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/687241207974526977">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/687240647426740229/pu/vid/240x240/Ca2quvrg_T6VWJbb.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/686555424703684608">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/686550270755487744/pu/vid/576x360/0C1887VRaWi7QRVd.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/684006698474815488">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/684006442957852672/pu/vid/480x360/O1Bv6lwGvA4hyKmn.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/680391523070951424">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/680389668345528320/pu/vid/480x480/PCkwc-60RPk1nktZ.mp4"
        ></video>
      </a>
      <a target="_blank" href="https://twitter.com/mizzsig/status/677478412961771520">
        <video
          class="content-video"
          @mouseenter="videoPlay"
          @mouseleave="videoPause"
          loop
          muted
          src="https://video.twimg.com/ext_tw_video/677477680669810689/pu/vid/286x162/RfIGtj2_gdtOZ1SI.mp4"
        ></video>
      </a>
    </div>
  </div>
</template>

<script>
import { mapMutations } from "vuex";

export default {
  mounted() {
    window.addEventListener("resize", this.onResize);
    // 画像読み込み後に初回リサイズ
    let videos = document.getElementsByClassName("content-video");
    Object.keys(videos).forEach(key => {
      videos[key].addEventListener("loadedmetadata", () => {
        this.onResize();
      });
    });
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.onResize);
  },
  methods: {
    onResize() {
      const gridAutoRows = 10;
      let videos = document.getElementsByClassName("content-video");
      Object.keys(videos).forEach(key => {
        videos[key].style.gridRowEnd = `span ${Math.ceil(
          videos[key].clientHeight / gridAutoRows
        )}`;
      });
    },
    isActive(path) {
      return path === this.$route.path;
    },
    videoPlay(e) {
      e.target.play();
    },
    videoPause(e) {
      e.target.pause();
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(290px, 1fr));
  -ms-grid-columns: repeat(auto-fit, minmax(290px, 1fr));
  grid-auto-rows: 10px;
  margin: 0px 10px;

  a {
    display: contents;
    .content-video {
      width: 100%;
    }
  }
}

.gallery-menu {
  margin-bottom: 10px;
}

.menu {
  text-decoration: none;
  color: #cccccc;
  font-size: 20px;
  padding: 10px;
  &:hover span::before,
  &:hover span::after {
    width: 50%;
  }
  .active::before,
  .active::after {
    width: 50%;
  }
  span {
    position: relative;
    &::before,
    &::after {
      border-bottom: 2px solid #bbbbbb;
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