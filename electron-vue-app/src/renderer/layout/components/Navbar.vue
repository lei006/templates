<template>
  <div class="navbar-box ">

    <div class="navbar-logo-title titlebar">
      <div class="logo-box">
        <el-image style="width: 36px; height: 36px; margin-left:8px; margin-right:8px;" :src="logo_url" :fit="logo_fit"></el-image>海豚PDF转化器
      </div>
    </div>

    <div class="navbar-menu-control">
      <div class="app-menu-box" >
        <div class="menu-button-box" @click="onBtnLogin">
          <el-image class="menu-button-img" :src="user_url" :fit="logo_fit"></el-image>
          <div class="menu-button-text">用户登录</div>
        </div>
      </div>
    </div>

    <div class="navbar-system-control">
      <el-breadcrumb class="app-breadcrumb" separator="/">
        <div class="menu-item-txt" @click="onBtnMini">─</div>
        <div class="menu-item-txt" @click="onBtnExit">✖</div>
      </el-breadcrumb>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'

const {ipcRenderer, shell} = require('electron')

export default {
  components: {
    Breadcrumb,
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

.navbar-menu-control {
  height: 100%;
}




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
  background-color: #bbb;
  color: #fff;
}



</style>
