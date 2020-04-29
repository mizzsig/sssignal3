export const state = () => ({
  // trueの時に表示される
  isShow: false,
  // falseの時は広告を出す
  isCharacter: true,
  // ページや画像ホバーで出すコメント
  comment: "",
  // 表示するsceneのID
  scene: null
});

export const mutations = {
  setShow(state, isShow) {
    state.isShow = isShow;
  },
  setCharacter(state, isShow) {
    state.isCharacter = isShow;
  },
  setComment(state, comment) {
    state.comment = comment;
  },
  setScene(state, scene) {
    state.scene = scene;
  }
};
