<template>
    <div id="app">
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    name: 'app'
}
</script>

<style>
html, body {
    overflow-x: hidden;
}

a {
    text-decoration: none;
}

.line {
    height: 1px;
    background-color: #e0e6ed;
    margin: 10px -24px;
}

.text-right {
    text-align: right;
}

.text-center {
    text-align: center;
}

.full-width-button > button {
    width: 100%;
}

.tab-heading {
    padding-top: 10px;
    padding-bottom: 10px;
}

.tab-heading .tab-subject {
    font-weight: bold;
    font-size: 18px;
}

#app {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
}
</style>
