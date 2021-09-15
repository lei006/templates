<template>
  <div class="navbar-box ">
    <div class="navbar-logo-title titlebar">
      <hamburger class="hamburger-container" @toggleClick="toggleSideBar" />
    </div>
    <div class="navbar-menu-control">
      <MenuControl @menuclick="onMenuClick"/>
    </div>
    <div class="navbar-system-control">
      <SystemControl />
    </div>
    <!--
    <DialogLogin ref="dlgLogin"/>
    -->
    <DialogBuy ref="dlgBuy" />
    <DialogHeadsn ref="dlgHeadsn" />
    <DialogVipTip ref="dlgVipTip" />
    <DialogLoginRegedit ref="dlgRegedit" />
    <DialogLogin ref="dlgLogin" />
    <DialogPassChange ref="dlgPwdChg" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import SystemControl from '@/components/SystemControl'
import MenuControl from '@/components/MenuControl'
import DialogLogin from '@/components/DialogLogin'
import DialogBuy from '@/components/DialogBuy'
import DialogHeadsn from '@/components/DialogHeadsn'
import DialogVipTip from '@/components/DialogVipTip'

import DialogLoginRegedit from '@/components/DialogLoginRegedit'
import DialogPassChange from '@/components/DialogPassChange'


const {ipcRenderer, shell} = require('electron')


export default {
  components: {
    Breadcrumb,
    Hamburger,
    SystemControl,
    MenuControl,
    DialogLogin,
    DialogBuy,
    DialogHeadsn,
    DialogVipTip,
    DialogPassChange,
    DialogLoginRegedit,

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
    }
  },

  async mounted(){
    let _self = this;

    //用户信息更新....
    ipcRenderer.on("userinfo", (event, userinfo) => {
      if (userinfo == null) {
        _self.userinfo = {};
        _self.userauth = false;
        _self.userinfoShow = false;
      }else{
        _self.userinfo = userinfo;
        _self.userinfo.expired_at = (new Date(userinfo.expired_at)).Format("yyyy-MM-dd")
        _self.userauth = true;
      }
    });


    ipcRenderer.on("buy-vip-ack", (event) => {
        _self.$refs.dlgBuy.show();
    });
    ipcRenderer.on("show-user-login", (event) => {
        _self.$refs.dlgLogin.show();
    });

    ipcRenderer.on("show-user-regedit", (event) => {
        _self.$refs.dlgRegedit.show();
    });

    //程序启动--->要求后台自动登录
    ipcRenderer.send("autologin");
  },

  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    onMenuClick(menu_cmd) {

      //显示登录界面
      if (menu_cmd == "login" ) {
        this.$refs.dlgLogin.show();
        return;
      } 


      //显示硬件码
      if (menu_cmd == "hard" ) {
        this.$refs.dlgHeadsn.show();
        return;
      } 

      //显示密码修改
      if (menu_cmd == "pwdchg" ) {
        this.$refs.dlgPwdChg.show();
        return;
      } 

      //显示购买界面
      if (menu_cmd == "buy" ) {
        if (this.userauth == false) {
          this.$refs.dlgLogin.show();
        }else{
          this.$refs.dlgBuy.show();
        }
      } 

    },
  }
}
</script>

<style lang="scss" scoped>
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



.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);

  .hamburger-container {
    line-height: 50px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: background .3s;
    -webkit-tap-highlight-color:transparent;

    &:hover {
      background: rgba(0, 0, 0, .025)
    }
  }

  .breadcrumb-container {
    float: left;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background .3s;

        &:hover {
          background: rgba(0, 0, 0, .025)
        }
      }
    }

    .avatar-container {
      margin-right: 30px;

      .avatar-wrapper {
        margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}




</style>
