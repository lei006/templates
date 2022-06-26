<template>
    <DialogBase01>
        <div style="width:100px; height:20px;"></div>
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
    computed: {
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
