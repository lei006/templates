<template>
  <div class="dashboard-container">
    <div class="container">

      <el-card class="box-card" style="width:500px;">
        <div slot="header" class="clearfix">
          <span> 版本信息</span>
          <el-button type="danger" size="mini" style="float: right;" @click="onHandleRestart">立即重启</el-button>
        </div>
        <div v-for="o in version" :key="o.key" class="text-item">
          <div class="label">{{o.key }}</div>
          <div class="data">{{o.val }}</div>
        </div>
      </el-card>
      <br>
      <el-card class="box-card" style="width:500px;">
        <div slot="header" class="clearfix">
          <span> 授权信息</span>
          <input type="hidden" value="male" id="fileInput" name="fileInput">
          <el-button type="danger" size="mini" style="float: right;" @click="onHandleLoadFile">更新授权</el-button>
        </div>
        <div v-for="o in license" :key="o.key" class="text-item">
          <div class="label">{{o.key }}</div>
          <div class="data">{{o.val }}</div>
        </div>
      </el-card>

    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

import { getAbout, PostRestart,GetLicense,SetLicense } from '@/api/system'


export default {
  name: 'About',
  data(){
    return{
      version:[],
      license:[],
      fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}, {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}],
    }
  },  
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  async created() {

    console.log("xxxxxxxx")
    let ret = await getAbout();
    if (ret.code === 20000) {
      this.version = ret.data;
    }
    console.log("xxxxxxxx")
    ret = await GetLicense();
    if (ret.code === 20000) {
      this.license = ret.data;
    }
    console.log("xxxxxxxx")

  },  
  methods: {
    handleSaveConfig(configItem){

    },
    onHandleLoadFile(){
      var input = document.createElement("input");
      input.type = "file";
      input.click();
      input.onchange = function(){
        var file = input.files[0];
        let file_reader = new FileReader();
        file_reader.onload = async () => {
            let licenseData = file_reader.result;
            let ret = await SetLicense(licenseData)
            if (ret.code === 20000) {
              location.reload();
            }
            console.log(ret)
        };
        file_reader.readAsText(file, 'UTF-8');            
      }

    },
    onHandleRestart(){
        this.$confirm('此操作将重启服务器, 是否继续?', '提示', {confirmButtonText: '确定',cancelButtonText: '取消',type: 'warning'}).then(() => {
          PostRestart()
        });
    }
  }


}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
    width:100%;
    height:100%;


  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>


<style scoped>
  .dashboard-text{
    font-size: 30px;
    line-height: 46px;
    margin-top: 30px;
    margin-bottom: 30px;
  }

  .container{
    display: flex;
    flex-flow: column;
    width:100%;

    display:flex;
    /*垂直排列*/
    flex-direction: column;
    align-items:center;/*由于flex-direction: column，因此align-items代表的是水平方向*/
    justify-content: center;/*由于flex-direction: column，因此justify-content代表的是垂直方向*/
  }

  .text-item {
    width: 100%;
    margin:10px;

    display:flex;
    flex-direction: row;
    align-items:left;/*由于flex-direction: column，因此align-items代表的是水平方向*/
  }
  .text-item .label{
    width: 150px;    
    background-color: #eee;
  }

  .text-item .data{
    margin-left: 10px;
    margin-top: 3px;

  }
  

</style>