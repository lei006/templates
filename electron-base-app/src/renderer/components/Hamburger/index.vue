<template>
  <div class="logo-box" @click="toggleClick">
    <el-image style="width: 36px; height: 36px; margin-left:8px; margin-right:8px;" :src="logo_url" :fit="logo_fit"></el-image>基本APP
  </div>
</template>

<script>
export default {
  name: 'Hamburger',
  props: {
    isActive: {
      type: Boolean,
      default: false
    }
  },
  data: function() {
    return { 
      logo_url: require("@/assets/logo.png"),
      logo_fit:"cover",
    }
  },
  
  methods: {
    toggleClick() {
      this.$emit('toggleClick')
    }
  }
}
</script>

<style scoped>

.logo-box {
  display: flex;
  justify-content: center;
  align-items: center;

  font-weight: 500;
  font-size: 20px;
}

.hamburger {
  display: inline-block;
  vertical-align: middle;
  width: 20px;
  height: 20px;
}

.hamburger.is-active {
  transform: rotate(180deg);
}
</style>
