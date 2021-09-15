<template>
  <div :class="classObj" class="app-wrapper">
    <div class="main-container">
      <div class="fixed-main">
        <app-main />
      </div>
    </div>
  </div>
</template>

<script>
import { Navbar,Footer, AppMain } from './components'

export default {
  name: 'Layout',
  components: {
    Navbar,
    Footer,
    AppMain
  },
  computed: {

    fixedHeader() {
      return this.$store.state.settings.fixedHeader
    },
    classObj() {
      return {
        mobile: true,
      }
    }
  },
  methods: {
    handleClickOutside() {
      this.$store.dispatch('app/closeSideBar', { withoutAnimation: false })
    }
  }
}
</script>

<style lang="scss" scoped>
  @import "~@/styles/mixin.scss";
  @import "~@/styles/variables.scss";

  .app-wrapper {
    @include clearfix;
    position: relative;
    height: 100%;
    width: 100%;
  }
  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .main-container {
    width: 100%;
    height:100%;

    display: flex;
    flex-direction: column;
  }
  .fixed-header {
    width: 100%;
    height: 50px;
    transition: width 0.28s;
    background-color: #fcfcfc;

    border-bottom-width: 1px;
    border-bottom-color: #ddd;
    border-bottom-style:solid;
  }
  .fixed-main {
    width: 100%;
    flex: 1;    
  }

  .fixed-footer {
    height: 40px;
    width: 100%;
    transition: width 0.28s;

    border-top-width: 1px;
    border-top-color: #ddd;
    border-top-style:solid;
  }

  .mobile .fixed-header {
    width: 100%;
  }

</style>
