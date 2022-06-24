<template>
  <div :class="classObj" class="app-wrapper">
    <div class="main-container">
      <div class="fixed-header">
        <navbar />
      </div>
      <div class="fixed-main">
        <app-main />
      </div>      
      <div class="fixed-footer">
        <Footer />
      </div>
    </div>
    <Dialogs></Dialogs>
  </div>
</template>

<script>
import { Navbar, Footer, AppMain, Dialogs } from './components'

export default {
  name: 'Layout',
  components: {
    Navbar,
    Footer,
    AppMain,
    Dialogs
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
    height: 32px;
    width: 100%;
    transition: width 0.28s;

    border-top-width: 1px;
    border-top-color: #ddd;
    border-top-style:solid;
  }


</style>
