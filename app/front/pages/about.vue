<template>
  <div>
    <h1>このサイトについて！</h1>
    <p>このサイトは、<b>とえら</b>が作った記事やその他のものが置かれてるサイトです！</p>
    <p>バージョン３です！</p>

    <h2>歴史！！！！</h2>
    
    <h3>バージョン１</h3>
    <a href="http://sssignal.web.fc2.com">http://sssignal.web.fc2.com</a>
    <div class="date">2016/03/18 ~ 2017/06/28</div>
    <div>
      <img width="600" src="~/assets/sssignal_ver1.png">
    </div>

    <h3>バージョン２</h3>
    <a href="https://sssignal.com">http://sssignal.com</a>
    <div class="date">2018/01/28 ~ 2019/06/01</div>
    <div>
      <img width="600" src="~/assets/sssignal_ver2.png">
    </div>

    <h2>ご意見フォーム！</h2>
    <form class="opinion-form">
      <input class="opinion-content address" name="address" placeholder="メールアドレス(任意)">
      <textarea class="opinion-content body" name="body" placeholder="本文(必須)"></textarea>
      <button class="opinion-content send">そうしん！</button>
    </form>
    <h2>規約！！！！</h2>
    <p>このサイトはどのページもリンクフリーです。<br>
    画像の無断転載と直リンクはご遠慮ください。</p>
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

.opinion-form {
    width: 90%;
    max-width: 600px;
    font-size: 20px;
    padding: 10px;
    border: 1px solid #FFF;
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
