<template>
  <el-breadcrumb class="app-breadcrumb" separator="/">
    <div class="menu-item-txt" @click="onBtnMini">─</div>
    <div class="menu-item-txt" @click="onBtnExit">✖</div>
  </el-breadcrumb>
</template>

<script>

const {ipcRenderer, shell} = require('electron')

export default {
  data() {
    return {
      levelList: null
    }
  },

  created() {
  },
  methods: {
    handleLink(item) {

    },
    onBtnMini() {
      ipcRenderer.send('app-mini') // prints "pong"
    },
    onBtnExit() {
      ipcRenderer.send('app-exit') // prints "pong"
    },
  }
}
</script>

<style lang="scss" scoped>
.app-breadcrumb.el-breadcrumb {
  font-size: 14px;
  height: 100%;

  padding-left: 5px;
  padding-right: 5px;
  margin-left: 0px;
  margin-right: 0px;

  display: flex;
  justify-content: center;
  align-items: center;  
}

.menu-item-txt {
  width:32px;
  height:32px;
  margin: 5px;

  border-radius: 3px;
  -webkit-user-select: none;

  display: flex;
  justify-content: center;
  align-items: center;

  cursor: pointer;   
}

.menu-item-txt:hover {
  background-color: #ddd;
  color: #fff;
}


</style>
