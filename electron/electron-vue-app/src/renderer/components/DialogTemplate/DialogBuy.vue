<template>
    <DialogBase01>
        <div style="width:100px; height:20px;"></div>
        <el-row style="background-color: #fff;">
            <el-col :span="12">
                <div class="buy_type_item" :class="buy_type_press=='person'?'buy_type_active':''" @click="buy_type_press='person'">
                    <div style="font-size: 18px;">个人会员</div>
                    <div style="color: #838c86; font-size: 12px;">个人使用高效便捷</div>
                </div>
            </el-col>
            <el-col :span="12">
                <div class="buy_type_item"  :class="buy_type_press=='business'?'buy_type_active':''" @click="buy_type_press='business'">
                    <div style="font-size: 18px;">企业会员</div>
                    <div style="color: #838c86; font-size: 12px;">企业使用多人共享</div>
                </div>
            </el-col>
        </el-row>
        <br>
        <el-row style="background-color: #fff;">

            <!--个人价格表-->
            <template v-if="buy_type_press=='person'">
                <el-col v-for="(item,i) in price_items" :key="'persion-' + i" :span="8">
                    <div class="price_box" v-if= "item.price_type == buy_type_press">
                        <div class="price_item" :class="cur_person_price_item.id==item.id?'person_price_item_active':''" @click="onBtnSelect(item)">
                            <div style="line-height: 40px; font-size: 16px;">{{item.title}}<span class="price_tig">{{item.hard_max_count}}台电脑</span></div>
                            <div class="money">￥<span style="font-size: 36px; ">{{item.price}}</span></div>
                            <div style="color: #999; line-height: 40px;">{{item.sub_title}}</div>
                        </div>
                    </div>
                </el-col>
            </template>

            <!--企业价格表-->
            <template v-if="buy_type_press=='business'">
                <el-col v-for="(item,i) in price_items" :key="'business-' + i" :span="8">
                    <div class="price_box"  v-if= "item.price_type == buy_type_press">
                        <div class="price_item" :class="cur_business_price_item.id==item.id?'business_price_item_active':''" @click="onBtnSelect(item)">
                            <div style="line-height: 40px; font-size: 16px;">{{item.title}}<span class="price_tig">{{item.hard_max_count}}台电脑</span></div>
                            <div class="money">￥<span style="font-size: 36px; ">{{item.price}}</span></div>
                            <div style="color: #999; line-height: 40px;">{{item.sub_title}}</div>
                        </div>
                    </div>
                </el-col>
            </template>
        </el-row>
        <el-row>
            <div v-if="buy_type_press=='person'" style="line-height: 120px; text-align: center; font-size: 12px;">应付金额:<span style="color: #8888CC; font-size: 32px;">{{cur_person_price_item.price}}</span>元</div>
            <div v-if="buy_type_press=='business'" style="line-height: 120px; text-align: center; font-size: 12px;">应付金额:<span style="color: #8888CC; font-size: 32px;">{{cur_business_price_item.price}}</span>元</div>
        </el-row>
        <el-row>
            <div style="display: flex; justify-content: space-around; margin-bottom: 20px;">
                <div class="order_box">
                    <el-button v-show="allow_alipay" type="primary" style="width: 180px; height: 50px; font-size:24px;" @click="onBtnBuy('alipay')">支付宝支付</el-button>
                    <el-button v-show="allow_wxpay" type="primary" style="width: 180px; height: 50px; font-size:24px;" @click="onBtnBuy('wxpay')">微信支付</el-button>
                </div>
            </div>
        </el-row>

        <el-row>
            <div style="line-height: 16px; text-align: center; margin-top: 10px;">
                <el-link v-show="buy_tig_show" type="primary" style="font-size: 2.0rem;">支付完成后，关闭本界面</el-link><br>
                <el-link type="info" style="font-size: 1.2rem; margin-top:30px;">已支付，但无反应，请连系客服</el-link>
            </div>
        </el-row>

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
