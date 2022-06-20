<template>
  <div class="app-dialog-box" separator="/">
    <el-dialog title="硬件ID" custom-class="login-dialog" width="35%" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>
        <div class="heardsn-box">
            <div class="heardsn-text">购买提示</div>
            <div class="heardsn-text">{{heardsn}}</div>
            <div class="heardsn-text" style="color:#f00; font-size:12px;">{{tip}}&nbsp;</div>
            <el-button type="primary" @click="onBtnCopy">复制机器码</el-button>            
        </div>
    </el-dialog>
  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')
import copy from 'copy-to-clipboard';


export default {
  data() {
    return {
      dialogVisible:false,
      dialog_modal:false,
      click_dialog_modal_close:false,

      heardsn:"",
      tip:"",
    }
  },
  async mounted(){
    let _self = this;



  },
  methods: {

    show() {
      this.dialogVisible = true;
    },
    hide() {
      this.dialogVisible = false;
    },
    onBtnCopy() {
        copy(this.heardsn);
        this.tip = "已经复制到粘贴板"
    },
    handleClose(){
      this.dialogVisible = false;
    },
  }
}
</script>

<style lang="scss" scoped>

  .app-dialog-box{
    display: flex;
    justify-content: center;
    align-items: center;
  }

    .heardsn-text{
        margin: 6px;
        line-height: 20px;
        text-align: center;
        font-size: 16px;
    }
    .heardsn-box {
        width: 100%;
        display: flex;
        flex-direction: column;
        justify-content:space-around;
        align-items: center;
    }

</style>
