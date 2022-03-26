<template>
  <section class="container">
    <div class="header">
        <div class="title">わんわんわーるど２</div>
        <nuxt-link to="/">制作：水飴信号</nuxt-link>
        <div class="version">ver 0.1.0</div>
    </div>
    <div class="left">
        <img style="height: 600px; width: 200px;" src="/soft/wanwan_world2/web/densha2.png">
        <div style="height: 50px; width: 200px;">
            <div>思い出 2 / 9<br>進捗 ■■■■□□□□□□</div>
        </div>
    </div>
    <div class="center">
        <div id="unity-container" class="unity-desktop">
        <canvas id="unity-canvas" width=800 height=600></canvas>
        <div id="unity-loading-bar">
            <div id="unity-logo"></div>
            <div id="unity-progress-bar-empty">
                <div id="unity-progress-bar-full"></div>
            </div>
        </div>
        <div id="unity-warning"> </div>
    </div>       
    </div>
    <div v-show="unityValue != ''" class="right">
        <div style="height: 150px; width: 200px;">吹き出し</div>
        <img src="/soft/wanwan_world2/web/wanko2.png">
    </div>
    <div class="map">
        マップ
    </div>
    <div class="status">
        <div>わんこのHP</div>
        <div>てきのHP</div>
    </div>
  </section>
</template>

<script>
export default {
  layout: "plane",
  head() {
    return {
      title: "わんわんわーるど２",
      meta: [
        { hid: "description", name: "description", content: "水飴信号３" },
        { hid: "twitter:site", name: "twitter:site", content: "@mizzsig" },
        { hid: "twitter:card", name: "twitter:card", content: "summary_large_image" },
        { hid: "og:description", property: "og:description", content: "水飴信号３" },
        { hid: "og:title", property: "og:title", content: "Starch Syrup Signal3" },
        { hid: "og:image", property: "og:image", content: "https://ver3.sssignal.com/temporary_twitter_card.png" }
      ]
    };
  },
  data() {
    return {
        unityValue: ''
    };
  },
  methods: {
    getUnityValue(event) {
        this.unityValue = event.data;
    }
  },
  mounted() {
    window.addEventListener('message', this.getUnityValue);

    this.$store.commit('character/setShow', false);

    var container = document.querySelector("#unity-container");
    var canvas = document.querySelector("#unity-canvas");
    var loadingBar = document.querySelector("#unity-loading-bar");
    var progressBarFull = document.querySelector("#unity-progress-bar-full");
    var warningBanner = document.querySelector("#unity-warning");

    // Shows a temporary message banner/ribbon for a few seconds, or
    // a permanent error message on top of the canvas if type=='error'.
    // If type=='warning', a yellow highlight color is used.
    // Modify or remove this function to customize the visually presented
    // way that non-critical warnings and error messages are presented to the
    // user.
    function unityShowBanner(msg, type) {
        function updateBannerVisibility() {
            warningBanner.style.display = warningBanner.children.length ? 'block' : 'none';
        }
        var div = document.createElement('div');
        div.innerHTML = msg;
        warningBanner.appendChild(div);
        if (type == 'error') div.style = 'background: red; padding: 10px;';
        else {
            if (type == 'warning') div.style = 'background: yellow; padding: 10px;';
            setTimeout(function () {
                warningBanner.removeChild(div);
                updateBannerVisibility();
            }, 5000);
        }
        updateBannerVisibility();
    }

    var buildUrl = "Build";
    var loaderUrl = buildUrl + "/wanwan_world2.loader.js";
    var config = {
        dataUrl: buildUrl + "/wanwan_world2.data",
        frameworkUrl: buildUrl + "/wanwan_world2.framework.js",
        codeUrl: buildUrl + "/wanwan_world2.wasm",
        streamingAssetsUrl: "StreamingAssets",
        companyName: "sssignal.com",
        productName: "wanwan_world2",
        productVersion: "1.0",
        showBanner: unityShowBanner,
    };

    // By default Unity keeps WebGL canvas render target size matched with
    // the DOM size of the canvas element (scaled by window.devicePixelRatio)
    // Set this to false if you want to decouple this synchronization from
    // happening inside the engine, and you would instead like to size up
    // the canvas DOM size and WebGL render target sizes yourself.
    // config.matchWebGLToCanvasSize = false;

    if (/iPhone|iPad|iPod|Android/i.test(navigator.userAgent)) {
        container.className = "unity-mobile";
        // Avoid draining fillrate performance on mobile devices,
        // and default/override low DPI mode on mobile browsers.
        config.devicePixelRatio = 1;
        unityShowBanner('WebGL builds are not supported on mobile devices.');
    } else {
        canvas.style.width = "800px";
        canvas.style.height = "600px";
    }
    loadingBar.style.display = "block";

    var script = document.createElement("script");
    script.src = loaderUrl;
    script.onload = () => {
        createUnityInstance(canvas, config, (progress) => {
            progressBarFull.style.width = 100 * progress + "%";
        }).then((unityInstance) => {
            loadingBar.style.display = "none";
            // fullscreenButton.onclick = () => {
            //     unityInstance.SetFullscreen(1);
            // };
        }).catch((message) => {
            alert(message);
        });
    };
    document.body.appendChild(script);
  },
  beforeDestroy() {
    window.removeEventListener("message", this.getUnityValue);
  }
};
</script>

<style lang="scss" scoped>
.container {
  background-color: #333;
  display: grid;
  grid-template-rows: 50px 600px 100px;
  grid-template-columns: 200px 800px 200px;
  position: fixed;
  width: 1200px;
  left: 0px;
  right: 0px;
  bottom: 0px;
  margin: auto;

  .header {
    grid-row: 1 / 2;
    grid-column: 2 / 3;
    background-color: #444;
    .title {
      float: left;
      font-size: 28px;
      margin-top: 4px;
      margin-left: 30px;
    }
    .version {
      float: right;
      margin-top: 30px;
      margin-right: 20px;
    }
  }

  .left {
    background-color: #363;
    grid-row: 1 / 3;
    grid-column: 1 / 2;
  }

  .center {
    background-color: #633;
    grid-row: 2 / 3;
    grid-column: 2 / 3;
  }

  .right {
    background-color: #336;
    grid-row: 1 / 4;
    grid-column: 3 / 4;
  }

  .map {
    background-color: #66f;
    grid-row: 3 / 4;
    grid-column: 1 / 2;
  }

  .status {
    background-color: #f66;
    grid-row: 3 / 4;
    grid-column: 2 / 3;
  }

  .hidden {
    display: none;
  }
}
</style>
