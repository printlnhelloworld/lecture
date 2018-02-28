<template>
  <div id="app">
    <keep-alive v-if="$route.meta.keepAlive">
      <router-view/>
    </keep-alive>
    <keep-alive v-if="!$route.meta.keepAlive">
      <router-view/>
    </keep-alive>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
*{
  padding: 0;
  margin: 0;
}
a {
    text-decoration:none;
    out-line: none;
    color: black;
}
html, body, #app{
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
}
/* #app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
