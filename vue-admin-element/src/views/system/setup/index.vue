<template>
  <div class="app-container">
    <div class="dashboard-text">配置管理</div>
    <div class="config_from_box"> 

      <el-form ref="form" label-width="120px">
        <el-form-item label="快照为缩略图:">
          <el-switch v-model="SnapshotToThumbnails" @change="handleChange_SnapshotToThumbnails"></el-switch>
        </el-form-item>        
        <el-form-item label="快照间隔时间:">
          <el-select v-model="SnapshotInterval" placeholder="请选择" @change="handleChange_SnapshotInterval">
            <el-option label="10秒" value=10></el-option>
            <el-option label="1分钟" value=60></el-option>
            <el-option label="10分钟" value=600></el-option>
            <el-option label="30分钟" value=1800></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="必需登录:">
          <el-switch v-model="CheckToken" @change="handleChange_CheckToken"></el-switch>
        </el-form-item>
        <el-form-item label="日志级别:">
          <el-select v-model="LogLevel" placeholder="请选择" @change="handleChange_LogLevel">
            <el-option label="详细日志" value=7></el-option>
            <el-option label="简要日志" value=5></el-option>
            <el-option label="关闭日志" value=1></el-option>
          </el-select>
        </el-form-item>
      </el-form>


    </div>

  </div>
</template>


<script>

import { getConfig,setConfig } from '@/api/config'


export default {
  data() {
    return {
      configs: [
        {id:11,name:"no_check_token",title:"不验证",desc:"不需要登录, 不再验证token,所有的api可以任意访问!",type:"bool",data:false},
        {id:12,name:"snapshot_span",title:"快照间隔时间(秒)",desc:"当设备在线时，每隔指定时间，刷新列表的缩略图!",type:"int",data:10*60},
        //{id:12,name:"不验证",title:"不验证",desc:"不需要登录, 不再验证token,所有的api可以任意访问!",type:"bool",data:true},
      ],
      dataLoading: false,
      CheckToken:false,
      SnapshotInterval:30,
      LogLevel:4,
      SnapshotToThumbnails:false,
    }
  },
  async mounted() {
    await this.onLoadAll()
  },

  methods: {
    onSubmit() {
      this.$message('submit!')
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    },
    async onLoadAll() {

      let ret = await getConfig("SnapshotInterval");
      this.SnapshotInterval = ret.data + ""

      ret = await getConfig("CheckToken");
      this.CheckToken = ret.data

      ret = await getConfig("LogLevel");
      this.LogLevel = ret.data + ""

      ret = await getConfig("SnapshotToThumbnails");
      this.SnapshotToThumbnails = ret.data
    },


    handleSaveConfig(configItem){

    },
    async handleChange_SnapshotInterval(){
      let ret = await setConfig("SnapshotInterval", this.SnapshotInterval)
      await this.onLoadAll()
    },    
    async handleChange_CheckToken(){
      let ret = await setConfig("CheckToken", this.CheckToken)
      await this.onLoadAll()
    },
    async handleChange_LogLevel(){
      let ret = await setConfig("LogLevel", this.LogLevel)
      await this.onLoadAll()
    },
    async handleChange_SnapshotToThumbnails(){
      let ret = await setConfig("SnapshotToThumbnails", this.SnapshotToThumbnails)
      await this.onLoadAll()
    },


  }
}
</script>

<style scoped>
  .dashboard-text{
    font-size: 30px;
    line-height: 46px;
  }
  .config_from_box{
    width: 700px;
    padding: 50px;;
    background-color: #fff;
    margin:auto auto;
  }


</style>
