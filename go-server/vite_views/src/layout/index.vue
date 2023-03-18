<script setup lang="ts">
import PanelMain from './panel_main.vue'
import PanelLeft from './panel_left.vue'
import PanelRight from './panel_right.vue'
import TopbarMenu1 from './topbar_menu1.vue'
import TopbarMenu2 from './topbar_menu2.vue'
import TopbarMenu3 from './topbar_menu3.vue'
import TopbarSystem from './topbar_system.vue'

</script>

<template>
    <div class="layout">
      <el-container>
        <el-header height="30px" class="top-menubar">
          <TopbarMenu1></TopbarMenu1>
          <TopbarSystem></TopbarSystem>
        </el-header>
        <el-container>
          <el-aside width="200px">
            <PanelLeft></PanelLeft>
          </el-aside>
          <el-main>
            <PanelMain></PanelMain>
          </el-main>
          <el-aside width="300px">
            <PanelRight></PanelRight>            
          </el-aside>
        </el-container>
      </el-container>
    </div>
</template>


<script setup>

</script>

<style scoped>

.layout{
    width:100%;
    height: 100%;
}

.top-menubar{
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;

  padding: 0px;
  margin: 0px;

}

</style>
