<template>
    
    <el-dialog title="PDF转换器" custom-class="login-dialog" width="450px" size="mini" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>
        <slot></slot>
    </el-dialog>

</template>

<script>

const {ipcRenderer, shell} = require('electron')

import {mapActions, mapState,mapGetters} from 'vuex' //注册 action 和 state


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
      dialogVisible:true,
      dialog_modal:true,
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
      sendsms_text:"发送验证码4",
      sendsms_space: "",
    }
  },
  computed: {
    //在这里映射 store.state.count，使用方法和 computed 里的其他属性一样
    ...mapState([
        'setup'
    ]),
    setup () {
        return this.$store.state.setup
    },
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
    },

    handleClose(){
      this.dialogVisible = false;
    },

    onBtnToRegedit(){
      this.hide();
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
