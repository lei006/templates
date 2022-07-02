<template>
  <div >
    页面一·----状态使用示例:{{ustoreId.counter}}------
  <div @click="onClick">add</div>
  </div>

</template>

<script>

import storeId from '@/store/index'  //storeId就是导出的那个值

export default {
  name: 'page01',

  data() {
    return {
      title: '主控面板',
      ustoreId:{},
    }
  },
  
  mounted(){
    this.ustoreId = storeId();
  },
  methods: {
    fetchData() {

    },
    onBtnJump(path_name, tool) {
      this.$router.push({ name: path_name, query: {tool:tool} })
    },
    onClick(){
      this.ustoreId.increment();
    }
  }

}
</script>


<style scoped>

.page-box {
  width: 100%;
  height: 100%;
  color: red;
}


</style>

