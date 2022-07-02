<template>
  <div class="page-box">
    页面2
  </div>

</template>

<script>

export default {
  name: 'page02',

  data() {
    return {
      title: '主控面板',
    }
  },
  
  created() {

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


<style scoped>

.page-box {
  width: 100%;
  height: 100%;
}


</style>

