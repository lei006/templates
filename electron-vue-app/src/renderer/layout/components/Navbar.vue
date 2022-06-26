<template>
  <div class="navbar-box ">

    <div class="navbar-logo-title titlebar">
      <div class="logo-box">
        <el-image style="width: 36px; height: 36px; margin-left:8px; margin-right:8px;" :src="logo_url" :fit="logo_fit"></el-image>海豚PDF转化器
      </div>
    </div>

    <div class="navbar-system-control">
      <el-breadcrumb class="app-breadcrumb" separator="/">
        <SystemButton01 text="设置"><span class="icon iconfont icon-un-setup-o"></span></SystemButton01>
        <SystemButton01 text="登录"><span class="icon iconfont icon-un-setup-o"></span></SystemButton01>
        <SystemButton00 @click="onBtnMini"><span class="icon iconfont icon-un-setup-o"></span></SystemButton00>
        <SystemButton00 @click="onBtnMini">─</SystemButton00>
        <SystemButton00 @click="onBtnExit">✖</SystemButton00>
      </el-breadcrumb>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import {SystemButton00, SystemButton01} from '@/components/ButtonCustom'

const {ipcRenderer, shell} = require('electron')

export default {
  components: {
    Breadcrumb,SystemButton00,SystemButton01
  },
  computed: {
    ...mapGetters([
    ])
  },
  data() {
    return {
      userauth:false,
      userinfo: {},  //用户信息....
      userinfoShow:false,
      logo_url: require("@/assets/logo.png"),
      logo_fit:"cover",

      user_url: require("@/assets/img_navbar/user.png"),

    }
  },

  async mounted(){
    let _self = this;

  },

  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    onMenuClick(menu_cmd) {

    },

    onBtnLogin() {
      this.$emit("menuclick","login");
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


.logo-box {
  display: flex;
  justify-content: center;
  align-items: center;

  font-weight: 500;
  font-size: 20px;
}


.navbar-box {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: flex-end;
}

.navbar-logo-title {
  flex:1;
  height: 100%;

  display: flex;
  justify-content: flex-start;
}

.navbar-system-control {
  height: 100%;
}


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



</style>
