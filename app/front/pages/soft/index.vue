<template>
  <section class="container">
      <nuxt-link class="link" to="/soft/endless_typing">
          <h2>endless typing</h2>
          <div style="margin-bottom: 50px;">終わらないタイピングが始まる……！</div>
      </nuxt-link>
      <nuxt-link class="link" to="/soft/css_art">
          <h2>css art</h2>
          <div style="margin-bottom: 50px;">htmlとcssを組み合わせてアートを作ろう！</div>
      </nuxt-link>
      <nuxt-link class="link" to="/soft/midi_keyboard_realtime_score">
          <h2>midi keyboard realtime score</h2>
          <div style="margin-bottom: 50px;">MIDIキーボードで弾いた楽譜が出ます！カメラも使います！</div>
      </nuxt-link>
      <nuxt-link class="link" to="/soft/youtube_subscriber_list">
          <h2>Youtube subscriber list</h2>
          <div>Youtubeの登録チャンネルを整理するページ、管理人以外は使えません！</div>
      </nuxt-link>
  </section>
</template>

<script>
export default {
  head() {
    return {
      title: "Soft - Starch Syrup Signal3",
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
    };
  },
  mounted() {
    this.$store.commit('character/setShow', true);
  }
};
</script>

<style lang="scss" scoped>
.link {
  text-decoration: none;
}
</style>
