<template>
  <div class="drag-area" ref="dragArea" separator="/">
        <slot></slot>
  </div>
</template>

<script>

const {ipcRenderer, shell} = require('electron')

export default {

    model: {
        prop: "FilterData",
        event: "eventFilterData"
    },
    props:["FilterData"],

    data() {
        return {
            dialogVisible:false,
            borderhover:false,
        }
    },
    async mounted(){
        let _self = this;
        let dragArea = this.$refs.dragArea;
        dragArea.addEventListener("drop",this.onEnentDrop,false)
        dragArea.addEventListener("dragleave",function (e) {
            e.stopPropagation();
            e.preventDefault();
            _self.borderhover =  false;
        })
        dragArea.addEventListener("dragenter",function (e) {
            e.stopPropagation();
            e.preventDefault();
            _self.borderhover =  true;
        })
        dragArea.addEventListener("dragover",function (e) {
            e.stopPropagation();
            e.preventDefault();
            _self.borderhover =  true
        })

    },

    beforeDestroy(){
    },

    methods: {
        
        //进入并松手...
        onEnentDrop: function(e){
            this.borderhover = false
            e.stopPropagation();
            e.preventDefault();  //必填字段
            let filelist = e.dataTransfer.files;

            let filters = this.FilterData;

            let filter_filelist = [];
            for (let index = 0; index < filelist.length; index++) {
                let file = filelist[index];

                if (this._allowFile(file.path, filters) == true) {
                    this.$emit("ondrag",file)
                    filter_filelist.push(file);
                }
            }
        },
        _allowFile(filename, filters){
            let ext_name = getFileExName(filename);
            let index_of = filters.indexOf(ext_name);
            if (index_of >= 0) {
                return true;
            }
            return false;
        },
    }
}



function getFileExName(fileNamePath) {
let tmp = fileNamePath.substr(fileNamePath.lastIndexOf('.') + 1);
return tmp;
}






</script>

<style lang="scss" scoped>

  .drag-area{
    width: 100%;
    height: 100%;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: #fff;

  }

   .drag-area:hover{
        border:1px dashed #409eff;
   }


</style>
