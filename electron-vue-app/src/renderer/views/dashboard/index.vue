<template>
  <div class="page-box">

    <div class="pdf-util-box">
      22222222222222222222233333
        <span class="icon iconfont icon-un-setup-o icon-mini-font-size"></span>
    </div>

  </div>

</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Dashboard',

  data() {
    return {
      title: '主控面板',
    }
  },
  
  computed: {
    ...mapGetters([
      'name'
    ])
  },

  created() {

  },
  beforeDestroy() {
    clearTimeout(this.timer);
  },

  methods: {
    fetchData() {

    },
    onBtnJump(path_name, tool) {
      this.$router.push({ name: path_name, query: {tool:tool} })
    },
  }

}
</script>

<style lang="scss" scoped>



</style>

<style scoped>

.page-box {
  width: 100%;
  height: 100%;
}

.pdf-util-box {
  width: 100%;
  height: 100%;  

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

}


.main_care_box {
    display: flex;
}

.main_care {
    width: 456px;
    height: 300px;
    margin:10px;
    cursor: pointer;   
    border-radius: 10px;
}

.sub_care_box {
    display: flex;
}

.sub_care {
    width: 218px;
    height: 252px;
    margin:10px;
    cursor: pointer;   
}


.el-card {
    border: 1px solid #adadad;
}

.el-card:hover {
    background-color: aliceblue;
}

.card-shadow:hover {
  box-shadow: 0px 0px 6px 1px rgba(80,80,80,0.9);
  margin-top:6px;
}


.main_care_box .el-card__body {
  padding: 0px;
}

.img-sub {
  width: 176px;
  height: 210px;
}

.img-main {
  width: 100%;
  height: 256px;
}



.icon-mini-font-size{
  font-size: 48px;
}


</style>

