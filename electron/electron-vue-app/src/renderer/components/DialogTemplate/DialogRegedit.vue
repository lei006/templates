<template>
    <DialogBase01>
       <div style="width:100px; height:20px;"></div>
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" size="small" class="demo-ruleForm" style="margin-bottom: 50px;">
          <el-form-item label="手机号" prop="username">
            <el-input style="width: 240px;" placeholder="请输入手机号码" v-model.number="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="验证码" prop="checkcode">
            <div style="display: flex;">
              <el-input style="width: 120px; margin-right: 10px;" maxlength=6 placeholder="请输入密码" type="text" v-model="ruleForm.checkcode" autocomplete="off"></el-input>
              <el-button type="primary" :disabled="sendsms_disabled" plain @click="onBtnSendSms">{{sendsms_text}}<span style="margin-left: 3px;">{{sendsms_space}}</span></el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button  type="primary" @click="onBtnLogin('ruleForm')">立即登录</el-button>
          </el-form-item>
      </el-form>
    </DialogBase01>
</template>

<script>


import DialogBase01 from './DialogBase01.vue'

export default {
    components:{
        DialogBase01
    },
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

    }
}
</script>

<style lang="scss" scoped>


</style>
