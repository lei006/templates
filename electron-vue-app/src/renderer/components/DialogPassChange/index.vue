<template>
  <div class="app-dialog-box" separator="/">
    
  
    <el-dialog title="PDF转换器" custom-class="login-dialog" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>

      <div class="login-title" >修改密码</div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm" style="margin-bottom: 50px;">
          <el-form-item label="旧密码" prop="oldpassword">
            <div style="display: flex;">
              <el-input style="width: 280px; margin-right: 50px;" maxlength=36 placeholder="请输入旧密码" type="text" v-model="ruleForm.oldpassword" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <div style="display: flex;">
              <el-input style="width: 280px; margin-right: 50px;" maxlength=36 placeholder="请输入新密码" type="text" v-model="ruleForm.password" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item label="确认密码" prop="checkpass">
            <div style="display: flex;">
              <el-input style="width: 280px; margin-right: 50px;" maxlength=36 placeholder="请再输一次密码" type="text" v-model="ruleForm.checkpass" autocomplete="off"></el-input>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button style="width: 300px; height:40px; font-size: 12px;" type="primary" @click="submitForm('ruleForm')">立即修改</el-button>
          </el-form-item>
      </el-form>

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
      var validateOldPass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {
          callback();
        }
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
          oldpassword:'',
          password:'',
          checkpass:'',
      },
      rules: {
          oldpassword: [
            { validator: validateOldPass, trigger: 'blur' }
          ],
          password: [
            { validator: validatePass, trigger: 'blur' }
          ],
          checkpass: [
            { validator: validatePass2, trigger: 'blur' }
          ],
      },
      sendsms_disabled:false,
      sendsms_text:"发送验证码3",
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
    async onBtnChangePass() {
      let username = this.ruleForm.username;
      let oldpassword = this.ruleForm.oldpassword
      let password = this.ruleForm.password

      ipcRenderer.once("user-changepass-ack", (event, ret) => {
        if ( ret.code == 20000 ) {
          this.$message({ message: "修改成功", type: 'success'});
          this.hide();
        }else{
          this.$message({ message: ret.message, type: 'error'});
        }
      });
      ipcRenderer.send("user-changepass", oldpassword, password);

    },

    submitForm(formName) {
      let _self = this;
      this.$refs[formName].validate((valid) => {
        if (valid) {
          _self.onBtnChangePass();
        } else {
          console.log('error submit!!');
          return false;
        }
      });
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

.tig-txt {
  width: 100;
  text-align: center; 
  font-size: 16px;
  margin-bottom: 20px;
}


</style>
