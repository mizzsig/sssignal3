<template>
  <div class="flex">
    <div class="half" v-html="column.Body"></div>
    <div class="half" style="padding-top: 30px; padding-bottom: 30px;">
      <div>column.Url</div>
      <input v-model="column.Url" style="width: 40%; background: #E0E0E0;" />
      <div>column.Title</div>
      <input v-model="column.Title" style="width: 40%; background: #E0E0E0;" />
      <div>column.Body</div>
      <div id="editor"></div>
      <div>column.Date</div>
      <input v-model="column.Date" style="width: 40%; background: #E0E0E0;" />
      <br />
      <button @click="setDate">↑now!</button>
      <div>column.CharacterComment</div>
      <input
        v-model="column.CharacterComment"
        style="width: 40%; background: #E0E0E0;"
      />
      <div>
        <button
          @click="postColumn"
          style="padding: 10px 20px; margin-top: 10px;"
        >
          submit!
        </button>
      </div>
    </div>
  </div>
</template>

<script>
// SSR時は呼ばないようにする
if (process.client) {
  require("ace-builds/src-min-noconflict/ace");
  require("ace-builds/webpack-resolver");
}

export default {
  layout: "columnEdit",
  data() {
    return {
      column: [],
      editor: null,
      setEditorId: null
    };
  },
  watch: {
    editorValue() {
      if (this.editor === null) {
        return "";
      } else {
        return this.editor.getValue();
      }
    }
  },
  mounted() {
    // APIから記事読み込み
    fetch(
      process.env.SSSIGNAL_API_DOMAIN + "/column/" + this.$route.params.url,
      {
        method: "GET",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        }
      }
    )
      .then(response => response.json())
      .then(resultJson => {
        this.column = resultJson;
        this.column.slicedDate = this.column.Date.slice(0, 10);
        this.column.Url = this.$route.params.url;

        // ace.jsを使う
        this.editor = ace.edit("editor");
        this.editor.setTheme("ace/theme/kuroir");
        this.editor.setFontSize(14);
        this.editor.getSession().setMode("ace/mode/html");
        this.editor.getSession().setTabSize(2);
        this.editor.session.setValue(this.column.Body);

        this.setEditorId = setInterval(() => {
          this.column.Body = this.editor.getValue();
        }, 100);
      });
  },
  beforeDestroy() {
    clearInterval(this.setEditorId);
  },
  methods: {
    setDate() {
      const padZero = num => (num < 10 ? "0" : "") + num.toString();
      const date = new Date();
      const year = date.getFullYear();
      const month = padZero(date.getMonth() + 1);
      const day = padZero(date.getDate());
      const hour = padZero(date.getHours());
      const minute = padZero(date.getMinutes());
      const second = padZero(date.getSeconds());
      this.column.Date = `${year}/${month}/${day} ${hour}:${minute}:${second}`;
    },
    postColumn() {
      fetch(
        process.env.SSSIGNAL_API_DOMAIN + "/column/" + this.$route.params.url,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            Url: this.column.Url,
            Title: this.column.Title,
            Body: this.column.Body,
            Date: this.column.Date,
            CharacterComment: this.column.CharacterComment
          })
        }
      ).then(response => {
        console.log(response);
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.flex {
  display: flex;
  .half {
    width: 50%;
    border: 3px #2f2f2f solid;
  }
}

#editor {
  width: 95%;
  height: 350px;
  margin: auto;
}

.column-container {
  margin: 30px;

  .column-title {
    display: inline-block;
    font-size: 30px;
    border-bottom: 2px #556b9e solid;
    margin-bottom: 8px;
    position: relative;

    .column-title-date {
      display: inline-block;
      font-size: 12px;
      position: absolute;
      left: -10px;
      top: -12px;
    }
  }
}
</style>
