<template>
  <div class="app-dialog-box" separator="/">
    
  
    <el-dialog title="PDF转换器" custom-class="login-dialog" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>

      <div class="login-title" >用户注册</div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-bottom: 50px;">
          <el-form-item label="用户名" prop="username">
            <el-input style="width: 300px;" placeholder="填写用户名（手机号码）" v-model.number="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <div style="display: flex;">
              <el-input style="width: 180px; margin-right: 50px;" maxlength=36 placeholder="请输入密码" type="text" v-model="ruleForm.password" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item label="确认密码" prop="checkpass">
            <div style="display: flex;">
              <el-input style="width: 180px; margin-right: 50px;" maxlength=36 placeholder="请输入密码" type="text" v-model="ruleForm.checkpass" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button style="width: 300px; height:40px; font-size: 12px;" type="primary" @click="submitForm('ruleForm')">立即注册</el-button>
          </el-form-item>
      </el-form>

      <el-divider></el-divider>
      <div class="tig-txt">
        <el-link type="primary" style="font-size:20px;" @click="onBtnToLogin"><b>已经有一个帐号，立即去登录</b></el-link>
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
          if (this.ruleForm.checkpass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
          callback();
        }
      };
      
      var validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请再次输入密码'));
        } else if (value !== this.ruleForm.password) {
          callback(new Error('两次输入密码不一致!'));
        } else {
          callback();
        }
      };


    return {
      dialogVisible:false,
      dialog_modal:false,
      click_dialog_modal_close:false,
      ruleForm: {
          password:'',
          checkpass:'',
          username: ''
      },
      rules: {
          password: [
            { validator: validatePass, trigger: 'blur' }
          ],
          checkpass: [
            { validator: validatePass2, trigger: 'blur' }
          ],
          username: [
            { validator: validateUsername, trigger: 'blur' }
          ]
      },
      sendsms_disabled:false,
      sendsms_text:"发送验证码1",
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
    async onBtnRegedit() {
      let username = this.ruleForm.username;
      let password = this.ruleForm.password

      ipcRenderer.once("regedit-ack", (event, ret) => {
        if ( ret.code == 20000 ) {
          this.$message({ message: "注册成功", type: 'success'});
          this.hide();
        }else{
          this.$message({ message: ret.message, type: 'error'});
        }
      });
      ipcRenderer.send("regedit", username, password);


    },

    submitForm(formName) {
      let _self = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          _self.onBtnRegedit();
        } else {
          console.log('error submit!!');
          return false;
        }
      });
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
    onBtnToLogin(){
      this.hide();
      ipcRenderer.send("user-to-login");      
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

.tig-txt {
  width: 100;
  text-align: center; 
  font-size: 16px;
  margin-bottom: 20px;
}


</style>
