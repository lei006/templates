<template>
  <div class="footer-box text-no-select">
    <div class="footer-version flex-row-center-center ">
      <div v-if="cur_Version == last_version">当前版本: {{cur_Version}}</div>
      <div v-if="cur_Version != last_version" style="color:#ff0066;cursor: pointer;" @click="onBtnDownload">下载最新版本: {{last_version}}</div>
    </div>
    <div class="footer-iframe flex-row-center-center">
      <iframe class="frame-class" frameborder=0 name="showHere" scrolling="no" :src="url_bottom" ></iframe>      
    </div>
  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')


export default {
  components: {

  },

  data() {
    return {
      url_bottom:"",
      cur_Version:"2021.02.27",
      last_version:"2021.02.27",
    }
  },

  async mounted(){
    let _self = this;

    ipcRenderer.on("menu-cmd-ack", (event, cmd, data1) => {
      if (cmd == "url-bottom") {
        _self.url_bottom = data1;
      }
      if (cmd == "last-version") {
        _self.last_version = data1;
      }
      if (cmd == "url-download") {
        shell.openExternal(data1);
      }

    });
    ipcRenderer.send("menu-cmd", "url-bottom");
    ipcRenderer.send("menu-cmd", "last-version");

  },

  methods: {
    async logout() {

    },
    onBtnDownload(){
      ipcRenderer.send("menu-cmd", "url-download");
    },
  }
}
</script>

<style lang="scss" scoped>
  .footer-box {
    width: 100%;
    height: 100%;
    overflow: hidden;
    position: relative;
    background: #fff;
    box-shadow: 0 1px 4px rgba(0,21,41,.08);

    display: flex;
    justify-content: flex-start;
    align-items: center;
  }

  .footer-version {
    height: 100%;
    padding: 10px;

    font-size: 14px;
  }

  .footer-iframe {
    height: 100%;
    flex:1;
  }

  .frame-class{
    width:100%;
    height:100%;
  }


</style>
