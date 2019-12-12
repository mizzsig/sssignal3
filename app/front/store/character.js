export const state = () => ({
    // trueの時に表示される
    isShow: false,
    // falseの時は広告を出す
    isCharacter: true,
});

export const mutations = {
    setShow(state, isShow) {
        state.isShow = isShow;
    },
    setCharacter(state, isShow) {
        state.isCharacter = isShow;
    }
};
