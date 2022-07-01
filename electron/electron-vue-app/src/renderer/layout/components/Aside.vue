<template>
  <div class="aside-box ">

    <div class="aside-button-group">
      <SystemButton02 text="首页"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="传输"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="好友"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
    </div>

    <div class="aside-button-ex-group">
      <el-divider></el-divider>
      <SystemButton02 text="我试"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="我是测试"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
      <SystemButton02 text="我是试"><span class="icon iconfont icon-un-setup-o"></span></SystemButton02>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import {SystemButton02} from '@/components/ButtonCustom'

const {ipcRenderer, shell} = require('electron')

export default {
  components: {
    Breadcrumb,SystemButton02
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


.aside-box{
  width: 100%;
  height: 100%;

  padding-top:10px;
  padding-bottom: 20px;

  display: flex;
  flex-direction: column;
  justify-content:space-between;
  align-items: center;
}

.aside-button-group {
  flex:1;

  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}

.aside-button-ex-group {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
}



</style>
