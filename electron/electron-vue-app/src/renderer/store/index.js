import { defineStore } from 'pinia' //引入

const useStore = defineStore('storeId', {
  // arrow function recommended for full type inference
  state: () => {
    return {
      counter: 10000,
      name: 'Eduardo',
      isAdmin: true,
    }
  },
  // 相当于 vue 中的 computed 计算属性
  getters: {
	
	},
   // 相当于 vue 中的 methods 方法
  actions:{
    increment(){
        console.log("increment", this.counter)
        this.counter++
    }
  }
})
export default useStore //导出
