<template>
    <div class="system-button" :style="{  'color': color, }" @click="onClick">
        <div class="system-button-icon"><slot>ICON</slot></div><div class="system-button-text">{{text}}</div>
    </div>
</template>

<script>

export default {

    props: {
        text: {
          type: String,
          default: ''
        },
        color: {
          type: String,
          default: '#222'
        },
    },



    data() {
        return {
        }
    },
    methods: {
        onClick(e){
            this.$emit("click", e);
        }
    }
}
</script>

<style lang="scss" scoped>


.system-button {

    min-width: 32px;
    height:34px;
    line-height: 34px;
  
    padding: 5px;

    border-radius: 3px;

    margin-right: 5px;

    display: flex;
    justify-content: center;
    align-items: center;

  cursor: pointer;   
}

.system-button:hover {
  background-color: #ccc;
  color: #fff;
}


.system-button-icon{
    font-size: 24px;
}

.system-button-text{
    margin-left: 5px;
}


</style>
