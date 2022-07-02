<template>
  <div class="dialog-box ">
    <DialogTest v-if="dlg_test_show" ref="dlgTest" />
  </div>
</template>

<script>
import {DialogTest} from '@/components/DialogTemplate'


const {ipcRenderer, shell} = require('electron')


export default {
  components: {
    DialogTest,
  },
  data() {
    return {
      userauth:false,
      userinfo: {},  //用户信息....
      userinfoShow:false,
      dlg_test_show:false,
    }
  },

  async mounted(){
    let _self = this;

    //_self.$refs.dlgLogin.show();
  },

  methods: {

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
.dialog-box {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: flex-end;
}


</style>
