<template>
  <div :class="classObj" class="app-wrapper">
    <div class="main-container">
      <div class="fixed-header">
        <navbar />
      </div>
      <div class="fixed-main">
        <div class="container-row-box">
          <div class="fixed-Aside">
            <Aside></Aside>
          </div>
          <div class="fixed-app-main">
            <app-main />
          </div>
        </div>
      </div>
      <div class="fixed-footer">
        <Footer></Footer>
      </div>
    </div>
    <Dialogs></Dialogs>
  </div>
</template>

<script>
import { Navbar, Footer, AppMain, Dialogs, Aside } from './components'

export default {
  name: 'Layout',
  components: {
    Navbar,
    Footer,
    AppMain,
    Dialogs,
    Aside,
  },
  computed: {

    classObj() {
      return {
        mobile: true,
      }
    }
  },
  methods: {

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
    height: calc(100% - 50px - 32px);
  }

  .fixed-Aside{
    height:100%;
    width:68px;

    border-right-width: 1px;
    border-right-color: #ddd;
    border-right-style:solid;    
  }
  .fixed-app-main{
    height: 100%;
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

  .container-row-box{
    width: 100%;
    height: 100%;

    display: flex;
    
  }


</style>
