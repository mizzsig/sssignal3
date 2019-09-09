<template>
  <div>
    <img class="svg" src="~/assets/about/about.svg">
    <!-- <object type="image/svg+xml" data="~/assets/about/about.svg"></object> -->
    <div class="section-borderline"></div>
    <div class="section-container">
      <div class="section-title">このサイトについて！</div>
      <p>このサイトは、<b>とえら</b>が作った記事やその他のものが置かれてるサイトです！</p>
      <p>バージョン３です！</p>
    </div>

    <img class="svg" src="~/assets/about/history.svg">
    <div class="section-borderline"></div>
    <div class="section-container">
      <div class="section-title">歴史！！！！</div>
      <div class="old-versions">
        <div class="old-version">
          <h3>バージョン１</h3>
          <a href="http://sssignal.web.fc2.com">http://sssignal.web.fc2.com</a>
          <div class="date">2016/03/18 ~ 2017/06/28</div>
          <div>
            <img src="~/assets/about/sssignal1.png">
          </div>
          <p>FC2の無料サーバを借りて作ったサイトです。<br>一部でxserverの機能を借りたりしつつ、せっせとHTMLを書いておりました。</p>
        </div>

        <div class="old-version">
          <h3>バージョン２</h3>
          <a href="https://sssignal.com">http://sssignal.com</a>
          <div class="date">2018/01/28 ~ 2019/06/01</div>
          <div>
            <img src="~/assets/about/sssignal2.png">
          </div>
          <p>Conohaのレンタルサーバを借りて作ったLaravel製のサイトです。<br>part1はAngularで作っておりましたが、当時の自分の限界点を迎えたため今の形(part2)になりました。</p>
        </div>
      </div>
    </div>

    <img class="svg" src="~/assets/about/inquiry.svg">
    <div class="section-borderline"></div>
    <div class="section-container">
      <div class="section-title">ご意見フォーム！</div>
      <p>ご意見感想など、何かある方はこちらから<br>
      <a href="https://twitter.com/mizzsig">Twitter</a>からの連絡でも、よいです！</p>
      <form class="opinion-form">
        <input class="opinion-content address" name="address" placeholder="メールアドレス(任意)">
        <textarea class="opinion-content body" name="body" placeholder="本文(必須)"></textarea>
        <button class="opinion-content send">そうしん！</button>
      </form>
    </div>

    <img class="svg" src="~/assets/about/terms.svg">
    <div class="section-borderline"></div>
    <div class="section-container">
      <div class="section-title">規約！！！！</div>
      <p>このサイトはどのページもリンクフリーです。<br>
      画像の無断転載と直リンクはご遠慮ください。</p>
    </div>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: "About - Starch Syrup Signal3"
    };
  }
};
</script>

<style lang="scss" scoped>
img {
    max-width: 95%;
}

.svg {
    width: 400px;
    max-width: 100%;
}

.section-borderline {
    position: relative;
    top: 1em;
    border: 1px solid #bbbbbb;
    border-style: dashed;
    z-index: 0;
}

.section-container {
    position: relative;
    width: 90%;
    max-width: 800px;
    margin: 0px auto 20px auto;
    background-color: #000000;
    z-index: 5;

    .section-title {
        font-size: 20px;
        border: 2px solid #bbbbbb;
        border-radius: 3px;
        border-style: dashed;
    }
}

.old-versions {
    .old-version {
        max-width: 600px;
        margin: auto;
    }
}

.opinion-form {
    width: 90%;
    max-width: 600px;
    font-size: 20px;
    padding: 10px;
    border: 1px solid #bbbbbb;
    margin: auto;

    .opinion-content {
        display: inline-block;
        box-sizing: border-box;

        padding: 10px;
        border: 0px;
        margin: 0px 10px 5px 10px;
        width: calc(100% - 20px);

        background: rgba(255, 255, 255, 0.96);
        color: #101010;
    }

    .body {
        resize: none;
    }
}
</style>
