<template>
    
    <el-dialog :title="title" custom-class="dialog-base-01" :width="width" :modal="dialog_modal" :close-on-click-modal="click_dialog_modal_close" :visible.sync="dialogVisible" :before-close="handleClose" v-draggable>
        <slot></slot>
    </el-dialog>

</template>

<script>

export default {

    props: {
        width: {
          type: String,
          default: '450px'
        },
        title: {
          type: String,
          default: '模板一'
        },
    },

    data() {
        return {
          sendsms_space: "",
          dialogVisible:true,
          dialog_modal:true,
          click_dialog_modal_close:false,
        }
    },
    methods: {
        show(){
          this.dialogVisible = true;
        },
        hide(){
          this.dialogVisible = false;
        },
        handleClose(){
          this.dialogVisible = false;
        },

    }
}
</script>

<style lang="scss" scoped>
.dialog-base-01{

}

</style>
