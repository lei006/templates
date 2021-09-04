<template>
  <div class="app-menu-box" >

    <div class="menu-button-box" @click="onBtnLogin">
      <el-image class="menu-button-img" :src="user_url" :fit="logo_fit"></el-image>
      <div class="menu-button-text">用户登录</div>
    </div>

  </div>
</template>

<script>
import pathToRegexp from 'path-to-regexp'
const {ipcRenderer, shell} = require('electron')


export default {
    data() {
      return {
        user_url: require("@/assets/img_navbar/user.png"),
      }
    },
  created() {
    this.getBreadcrumb()
  },

  async mounted(){

  },

  methods: {
    getBreadcrumb() {

    },
    onBtnLogin() {
      this.$emit("menuclick","login");
    },
  }
}
</script>

<style lang="scss" scoped>
.app-menu-box {
  font-size: 14px;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;  
}

.menu-button-box {
  height: 40px;
  margin-left: 5px;
  margin-right: 5px;

  padding-right:10px;

  cursor: pointer;   
  -webkit-user-select: none;

  display: flex;
  justify-content: space-between;
  align-items: center;
}

.menu-button-box:hover {
  background-color: #f8f8f8;
  color: #fff;
}


.menu-button-img {
  width: 32px; 
  height: 32px;

  border-radius: 16px;
  margin-left:8px; 
  margin-right:2px;
}

.menu-button-text {
  color: #000;
  font-size: 16px;
}

.login-dialog {

  width:24%;
}

.el-dialog__header {
  border-bottom-width: 2px;
  border-bottom-color: #ddd;
  border-bottom-style: solid;
}


</style>
