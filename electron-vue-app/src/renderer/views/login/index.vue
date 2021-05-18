<template>

  <div class="login-container flex-center-center">

    <div class="login_box">
      <div class="login_title"><h1 class="title" style="margin:10px;">AnyCVR-监控流媒体平台</h1></div>
      <div class="login_content">

            <div class="login_content_left">
              <el-image :fit="loginFit" :src="login_left_png"></el-image>
            </div>
            <div class="login_content_right">
                <h2 class="title" style="margin-bottom:15px;">账号密码登录</h2>
                <el-form ref="loginForm" :model="loginForm" :rules="loginRules" label-width="0px" style="width:80%;">
                  <el-form-item prop="username">
                    <el-input v-model="loginForm.username"></el-input>
                  </el-form-item>
                  <el-form-item prop="password">
                    <el-input v-model="loginForm.password" type="password"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" style="width:100%" @click.native.prevent="handleLogin">立即创建</el-button>
                  </el-form-item>
                </el-form>
                <el-row style="margin-left:10px; width:264px;">
                    <p>默认用户:admin</p><p>默认密码:123456</p>
                </el-row>
            </div>
      </div>

    </div>

  </div>
</template>

<script>


export default {
  name: 'Login',
  data() {
    const validateUsername = (rule, value, callback) => {

        callback()
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 2) {
        callback(new Error('请输入有效密码'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      loginFit:"cover",
      loading: false,
      passwordType: 'password',
      redirect: undefined,
      //login_left_png: require("@/assets/img/login_left.png")

    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm).then(() => {
            this.$router.push({ path: this.redirect || '/' })
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
$bg:#fd3a4b;
$dark_gray:#889aa4;
$light_gray:#eee;



.flex-center-center {
  display: flex;
  flex-direction: row;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content: center;
  align-items: center;
}

.login_box {
  width:672px;
  height:452px;
  background-color: #fff;
  margin-bottom: 50px;

  border-radius: 5px;
  border-style:solid;
  border-width:5px;  
  border-color:#aaa;

  overflow: hidden;

  display: flex;
  flex-direction: column;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content:space-around;
  align-items: center;


}
.login_title {
  width:100%;
  height:72px;
  background-color: #f0f0f0;

  display: flex;
  flex-direction: row;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content:space-around;
  align-items: center;


}
.login_content {
  width:100%;
  flex:1;
  height:100%;

  display: flex;
  flex-direction: row;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content: center;
  align-items: center;

}

.login_content_left {
  flex:1;
  height:100%;
  background-color: #fff;

  display: flex;
  flex-direction: column;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content:space-around;
  align-items: center;

}

.login_content_right {
  flex:1;
  height:100%;

  display: flex;
  flex-direction: column;/*这里可以不写，flex布局默认方向横向即row*/
  justify-content:center;
  align-items: center;


}




.login-container {
  min-height: 100%;
  width: 100%;
  overflow: hidden;
  background-color: #f6f6f6;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }
}
</style>
