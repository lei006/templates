<template>
  <div class="dashboard-container">
    <div class="dashboard-text">name: {{ name }}</div>
    <div v-for="count in 136" :key="count">{{count}}==========</div>
  </div>
</template>

<script>

export default {
  name: 'Dashboard',
  data(){
    return {
      name:"",
      count:1,
    }
  }

}


</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
