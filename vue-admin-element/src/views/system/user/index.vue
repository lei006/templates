<template>
  <div class="app-container">
    <div class="dashboard-text">用户管理</div>
    <el-row style="margin-bottom:5px;">
      <el-button type="primary" size="mini" @click="onShowDialogAddUser">增加用户</el-button>
    </el-row>
    <div class="container">

      <el-table
        v-loading="listLoading"
        :data="userlist"
        element-loading-text="Loading"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="ID" width="95">
          <template slot-scope="scope">
            {{ scope.$index }}
          </template>
        </el-table-column>
        <el-table-column prop="username" label="用户名"  width="180"></el-table-column>
        <el-table-column prop="avatar" label="角色" width="180" align="center"></el-table-column>
        <el-table-column prop="enable" label="是否启用" width="150" align="center">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.is_enable" active-color="#13ce66" inactive-color="#ff4949" @change="handleSaveUse(scope.row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="更新时间" width="160" align="center" :show-overflow-tooltip='true'>
          <template slot-scope="scope">
            {{scope.row.updated_at|timeDataFormat}}
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="160" align="center" :show-overflow-tooltip='true'>
          <template slot-scope="scope">
            {{scope.row.created_at|timeDataFormat}}
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="描述" width="240" align="center">
        </el-table-column>
        <el-table-column label="操作" fixed="right">
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="onShowDialogEditUser(scope.$index, scope.row)">编辑</el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>      
    </div>

    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleCloseDialog">

        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
          <el-form-item label="用户名" prop="username">
            <el-input :disabled="!dialogAddUser"  v-model="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input :type="passwordType" v-model="ruleForm.password"></el-input>
          </el-form-item>
          <el-form-item v-show="dialogAddUser" label="确认密码" prop="repassword">
            <el-input :type="passwordType" v-model="ruleForm.repassword"></el-input>
          </el-form-item>
          <el-form-item label="用户类型" prop="avatar">
            <el-select v-model="ruleForm.avatar" placeholder="请选择用户类型">
              <el-option label="管理员" value="admin"></el-option>
              <el-option label="一般用户" value="user"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="正常启用" prop="enable">
            <el-switch v-model="ruleForm.is_enable"></el-switch>
          </el-form-item>
          <el-form-item label="描述">
            <el-input  v-model="ruleForm.desc"></el-input>
          </el-form-item>

        </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onHandleAddUser()">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import { getList,addUser,delUser,setUser,enableUser } from '@/api/user'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    },
  },
  data() {
    return {
      list: [],
      userlist: [],
      passwordType:"password",
      listLoading: true,
      dialogVisible:false,
      dialogAddUser:true,
      ruleForm: {
          username: '',
          password: '',
          repassword: '',
          avatar: '',
          is_enable: true,
          desc:"",
        },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 1, max: 15, message: '长度在 1 到 5 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur' }
        ],
        repassword: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 1, max: 15, message: '长度在 1 到 15 个字符', trigger: 'blur' }
        ],
        avatar: [
          { required: true, message: '请选择活动区域', trigger: 'change' }
        ],
      }



    }
  },

  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.userlist = response.data.items
        this.listLoading = false
      }).catch(err=>{
        console.warn(err)
        this.listLoading = false
      })
    },
    async handleSaveUse(userinfo){
        let ret = await setUser(userinfo)
        if (ret.code !== 20000) {
          this.$message({message:ret.message,type: 'error'})
        }
    },

    async handleEnableUse(userinfo){
        let ret = await enableUser(userinfo.id, userinfo.is_enable)
        if (ret.code !== 20000) {
          this.$message({message:ret.message,type: 'error'})
        }
    },
    
    onShowDialogAddUser(){
      this.dialogAddUser = true;

      this.ruleForm.username = ""
      this.ruleForm.password = ""
      this.ruleForm.repassword = ""
      this.ruleForm.avatar = "user"
      this.ruleForm.is_enable = true
      this.ruleForm.desc = ""
      this.dialogVisible = true;
    },
    onShowDialogEditUser(index, userinfo){
      this.dialogAddUser = false;

      Object.assign(this.ruleForm, userinfo)

      this.dialogVisible = true;
    },
    async onHandleAddUser(){

      if (this.dialogAddUser == true) {

        if (this.ruleForm.password !== this.ruleForm.repassword) {
          this.$message("二次密码不同!")
          return;
        }
        if (this.ruleForm.username.length <=3 ) {
          this.$message("请输入用户名")
          return;
        }
        if (this.ruleForm.password.length <3 ) {
          this.$message("请输密码(最少三位)")
          return;
        }


        let ret = await addUser(this.ruleForm)
        if (ret.code === 20000) {
          this.dialogVisible = false;
          this.userlist.push(ret.data);
        }else{
          this.$message({message:ret.message,type: 'error'})
        }
      }else{

        let ret = await setUser(this.ruleForm)
        console.log(this.ruleForm.desc)
        if (ret.code === 20000) {

          for( let index in this.userlist) {
            if(this.userlist[index].username == this.ruleForm.username) {

              Object.assign(this.userlist[index], this.ruleForm)
            }
          }
          this.dialogVisible = false;
        }else{
          this.$message({message:ret.message,type: 'error'})
        }
      }

    },
    handleEdit(){
      this.dialogAddUser = false;
      this.dialogVisible = true;
    },
    async handleDelete(index, userinfo){

      if (userinfo.username.length <=1) {
        this.$message("用户无效")
        return;
      }

      let _self = this;
      this.$confirm('是否删除用户' + userinfo.username, '提示', {confirmButtonText: '确认删除',cancelButtonText: '取消',type: 'warning'}).then( async () => {

        let ret = await delUser(userinfo.id)
        if (ret.code === 20000) {
          for( index in _self.userlist) {
            if(_self.userlist[index].username == userinfo.username) {
              _self.userlist.splice(index,1)
            }
          }
        }else{
            _self.$message({message:ret.message,type: 'error'})
        }
      })

    },
    handleCloseDialog() {
      this.dialogVisible = false;
    }
  }
}
</script>

<style scoped>
  .dashboard-text{
    font-size: 30px;
    line-height: 46px;
    margin: 20px;;
  }


</style>
