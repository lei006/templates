<template>
  <div class="app-dialog-box" separator="/">
    <el-dialog title="购买VIP" custom-class="login-dialog" width="75%" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>

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

    </el-dialog>

    <el-dialog title="使用微信扫码支付" :visible.sync="wx_dialogVisible" width="30%" :before-close="wx_handleClose">
        <div id="qrCode" class="wx_qrcode" ref="qrCodeDiv"></div>
        <div class="wx_qrcode_txt">使用微信扫码支付</div>
    </el-dialog>


  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')
import QRCode from 'qrcodejs2';

export default {
  data() {
    return {
      dialogVisible:false,
      wx_dialogVisible:false,
      dialog_modal:false,
      click_dialog_modal_close:false,

      activeName:"first",
      buy_type_press:"person",
      price_items:[
          {id:"1",title:"半年VIP",hardsn_num:3,price:39,demo:"低至0.2元/天", price_type:"person"},
      ],
      cur_person_price_item:{},
      cur_business_price_item:{},
      cur_PriceItem:{},
      buy_tig_show:true,
      allow_wxpay:false,
      allow_alipay:false,
    }
  },
  async mounted(){
    let _self = this;
    ipcRenderer.on("price-list-ack", (event, price_items) => {
        _self.price_items = price_items;
        _self._price_for();
    });
  },

    beforeDestroy(){
        console.log("order beforeDestroy ")
    },

  methods: {
    handleLink(item) {

    },
    show() {
      this.buy_tig_show = true;
      this.allow_wxpay = false;
      this.allow_alipay = false;

      ipcRenderer.send("price-list");
      this.dialogVisible = true;
    },
    hide() {
        ipcRenderer.send("autologin");
        this.dialogVisible = false;
    },
  
    onButtonExit: function () {          
      ipcRenderer.send("win-close",'buy') // prints "pong"
    },
    onBtnLogin:function() {
      console.log("login")
      //this.dialogVisible = true;
    },
    _price_for() {
        for (let index = 0; index < this.price_items.length; index++) {
            let element = this.price_items[index];

            if (element.price < 100) {
                this.price_items[index].price = (this.price_items[index].price*0.01).toFixed(2)
            }else{
                this.price_items[index].price = (this.price_items[index].price*0.01).toFixed(0)
            }

            if (element.allow_wxpay == true) {
                this.allow_wxpay = true
            }
            if (element.allow_alipay == true) {
                this.allow_alipay = true
            }



            this.onBtnSelect(element)
        }
    },
    handleClick(){

    },
    showWxQRCocde(show_txt) {
        let _self = this;
        this.wx_dialogVisible = true;
        setTimeout(function(){
            _self.$refs.qrCodeDiv.innerHTML  = "";
            new QRCode(_self.$refs.qrCodeDiv, { text: show_txt,width: 200,height: 200, colorDark: "#333333", colorLight: "#ffffff", correctLevel: QRCode.CorrectLevel.L})

        }, 100);

    },
    async onBtnBuy(pay_type){
      let _self = this;

      // 1. 
      ipcRenderer.once("create-order-ack", (event, order) => {
        console.log("create-order-ack", order)
        if (order.pay_type == "alipay") {
            shell.openExternal(order.url);
        }else{
            _self.showWxQRCocde(order.url);
        }
      });

      // 2. 得到远中的价格 Item
      let price_item = this.cur_person_price_item;
      if (this.buy_type_press == "business") {
          price_item = this.cur_business_price_item;
      }
      // 3. 创建定单....
      ipcRenderer.send("create-order", pay_type, price_item.id);

    },
    onBtnSelect(element){
        if (element.price_type == "person") {
            this.cur_person_price_item = element
        }
        if (element.price_type == "business") {
            this.cur_business_price_item = element
        }
    },
    wx_handleClose() {
        this.wx_dialogVisible = false;
    },
    handleClose(){
        this.hide()
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

  .el-tabs__item{
      width: 50%;
      height: 80px;
      line-height: 80px;
  }
  .buy_type_item{
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      width: 100%; 
      height: 80px; 
      text-align: center;
      color: #999;
      background-color: #eee;
  }

  .buy_type_active {
      color: #222;
      background-color: #fff;
      font-weight: bold;
  }


  .price_box {
      width: 100%;
      height: 100%;

      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
  }

  .price_item{

      border:2px solid #aaa;
      border-radius: 5px;

      width: 216px; 
      height: 160px; 
      text-align: center;


      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
  }

  .person_price_item_active {
      font-weight:bold;
      border-color: rgb(52,122,235);
      color: #0033CC;
      background-color: rgb(235,245,253);
  }

  .business_price_item_active {
      font-weight: bold;
      border-color: rgb(201,151,88);
      color: rgb(201,151,88);
      background-color: rgb(255,253,247);
  }

  .price_tig {
      color: #fff;
      background-color: rgb(142,141,240);
      border-radius: 10px;
      padding: 4px;
      font-size: 12px;
      margin-left: 10px;
      padding: 5px;
      font-weight:unset;
  }

  .order_box {

  }

  .order_box .img {
      width:160px;
      height: 160px;

      border-radius: 10px;
      border-style:solid;
      border-width: 1px;
      border-color: cornflowerblue;
  }

  .order_box .txt {
      font-size: 20px;
      text-align: center;
  }
    .wx_qrcode {
        display: flex;
        justify-content: center;
    }

    .wx_qrcode_txt {
        width: 100%;
        text-align: center;
        font-size: 16px;
    }


</style>
