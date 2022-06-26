<template>
  <div class="app-dialog-box" separator="/">
    
  
    <el-dialog title="用户登录" custom-class="login-dialog" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>

      <div class="login-title" >PDF文件转换</div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-bottom: 50px;">
          <el-form-item label="手机号" prop="username">
            <el-input style="width: 300px;" placeholder="请输入手机号码" v-model.number="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="验证码" prop="checkcode">
            <div style="display: flex;">
              <el-input style="width: 120px; margin-right: 50px;" maxlength=6 placeholder="请输入密码" type="text" v-model="ruleForm.checkcode" autocomplete="off"></el-input>
              <el-button type="primary" :disabled="sendsms_disabled" plain @click="onBtnSendSms">{{sendsms_text}}<span style="margin-left: 3px;">{{sendsms_space}}</span></el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button style="width: 300px; height:50px; font-size: 20px;" type="primary" @click="onBtnLogin('ruleForm')">立即登录</el-button>
          </el-form-item>
      </el-form>
      <el-divider>其它</el-divider>
    </el-dialog>

  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')

export default {
  data() {
    return {
      dialogVisible:false,
      dialog_modal:false,
      click_dialog_modal_close:false,
      ruleForm: {
          checkcode: '123243',
          username: '18115710898'
      },
      rules: {
          checkcode: [
            //  { validator: validatePass2, trigger: 'blur' }
          ],
          username: [
              //{ validator: checkAge, trigger: 'blur' }
          ]
      },      
      sendsms_disabled:false,
      sendsms_text:"发送验证码2",
      sendsms_space: "",
    }
  },

  created() {
  },


  
  methods: {
    handleLink(item) {

    },
    show(){
      this.dialogVisible = true;
    },
    hide(){
      this.dialogVisible = false;
    },
    async onBtnLogin() {
      let username = this.ruleForm.username;
      let checkcode = this.ruleForm.checkcode

      ipcRenderer.once("login-ack", (event, ret) => {
        if ( ret.code == 20000 ) {
          this.$message({ message: "登录成功", type: 'success'});
          this.hide();
        }else{
          this.$message({ message: ret.message, type: 'error'});
        }
      });
      ipcRenderer.send("login", username, checkcode);


    },

    sendsmsSuccess() {
      let _self = this;
      //发送成功，60秒后可以重试
      _self.sendsms_disabled = true;
      _self.sendsms_space = 60;
      let timer=setInterval(function(){

        _self.sendsms_space -= 1;

        if(_self.sendsms_space === 0) {
          _self.sendsms_space = "";
          _self.sendsms_disabled = false;
          clearTimeout(timer);
        }
      },1000);
    },

    async onBtnSendSms() {
      let username = this.ruleForm.username;
      ipcRenderer.once("send-sms-ack", (event, ret) => {
        if ( ret.code == 20000 ) {
          this.sendsmsSuccess();
        }else{
          this.$message({ message: ret.message, type: 'error'});
        }
      });
      ipcRenderer.send("send-sms", username);
   
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

.login-title {
    line-height: 40px; 
    font-size: 32px; 
    width: 100%; 
    text-align: center; 
    margin-bottom: 30px;

    -webkit-user-select: none;

}


</style>
