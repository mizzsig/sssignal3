<template>
  <div>
    <header>
      <nuxt-link class="menu" to="/"><span>Starch Syrup Signal3</span></nuxt-link>
      <nuxt-link class="menu" to="/about"><span>about</span></nuxt-link>
      <nuxt-link class="menu" to="/gallery"><span>gallery</span></nuxt-link>
      <nuxt-link class="menu" to="/column"><span>column</span></nuxt-link>
    </header>
    <div style="height: 50px;"></div>
    <nuxt class="nuxt"/>
  </div>
</template>

<style lang="scss" scoped>
$height: 50px;
header {
  display: flex;
  position: fixed;
  width: 100%;
  height: $height;
  z-index: 10;
  padding-left: 40px;
  background-color: rgba(30, 30, 30, 0.6);
  .menu {
    display: flex;
    align-items: center;
    text-decoration: none;
    color: #cccccc;
    font-size: 20px;
    padding: 0px 10px;
    height: 100%;
    &:hover span::before, &:hover span::after {
      width: 50%;
    }
    span {
      display: block;
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
}
.nuxt {
  height: calc(100% - 50px);
}
</style>

<style lang="scss">
html {
  /* font-family: "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", */
    /* Roboto, "Helvetica Neue", Arial, sans-serif; */
  font-size: 16px;
  /* word-spacing: 1px; */
  /* -ms-text-size-adjust: 100%; */
  /* -webkit-text-size-adjust: 100%; */
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}
html, body {
  background-color: #000;
  height: 100%;
}
html, body, div {
  text-align: center;
  color: #bbbbbb;
  margin: 0px;
  padding: 0px;
}

a {
    color: rgb(54, 149, 222);

    &:hover {
        color: rgb(18, 111, 182);
    }
}
</style>