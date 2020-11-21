<template>
 <div class="container">
    <div v-if="error.statusCode === 404">
        <h1>ページが見つかりません🙇</h1>
    </div>
    <h1 v-else>エラーが発生しました</h1>
    <div>status code : {{ error.statusCode }}</div>
  </div>
</template>

<script>
export default {
  props: ['error']
}
</script>