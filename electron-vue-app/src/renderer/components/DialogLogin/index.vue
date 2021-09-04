<template>
  <div class="app-dialog-box" separator="/">
    
  
    <el-dialog title="PDF转换器" custom-class="login-dialog" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>

      <div class="login-title" >用户登录</div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="80px" class="demo-ruleForm" style="margin-bottom: 50px;">
          <el-form-item label="用户名" prop="username">
            <el-input style="width: 240px; margin-right: 150px;" placeholder="填写用户名（手机号码）" v-model.number="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <div style="display: flex;">
              <el-input style="margin-right: 150px;" maxlength=36 placeholder="请输入密码" type="text" v-model="ruleForm.password" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button style="width: 220px; height:40px; font-size: 12px;" type="primary" @click="submitForm('ruleForm')">立即登录</el-button>
          </el-form-item>
      </el-form>
      <el-divider></el-divider>
      <div class="tig-txt">
        <el-link type="primary" style="font-size:20px;" @click="onBtnToRegedit"><b>创建一个帐号，立即登录</b></el-link>
      </div>
    </el-dialog>

  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')

export default {
  data() {

      function isRegisterUserName(s)
      {
        var patrn=/^([a-zA-Z0-9]|[._@]){4,32}$/;
        if (!patrn.exec(s)) return false
        return true
      }

      var validateUsername = (rule, value, callback) => {
        if (isRegisterUserName(value) === false) {
          callback(new Error('4-32个字母,数字、“_”、“.”、“@”组成的字串'));
        }
        callback();
      };
      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          callback();
        }
      };
      
    return {
      dialogVisible:false,
      dialog_modal:false,
      click_dialog_modal_close:false,
      ruleForm: {
          password: '',
          username: ''
      },
      rules: {
          password: [
            { validator: validatePass, trigger: 'blur' }
          ],
          username: [
            { validator: validateUsername, trigger: 'blur' }
          ]
      },      
      sendsms_disabled:false,
      sendsms_text:"发送验证码",
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
      let password = this.ruleForm.password

      ipcRenderer.once("login-ack", (event, ret) => {
        if ( ret.code == 20000 ) {
          this.$message({ message: "登录成功", type: 'success'});
          this.hide();
        }else{
          this.$message({ message: ret.message, type: 'error'});
        }
      });
      ipcRenderer.send("login", username, password);


    },
    submitForm(formName) {
      let _self = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          _self.onBtnLogin();
        } else {
          console.log('error submit!!');
          return false;
        }
      });
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


    handleClose(){
      this.dialogVisible = false;
    },

    onBtnToRegedit(){
      this.hide();
      ipcRenderer.send("user-to-regedit");
    }

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

.tig-txt {
  width: 100;
  text-align: center; 
  font-size: 16px;
  margin-bottom: 20px;
}

.login-dialog {
  width: 400px;
}


</style>
