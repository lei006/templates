const state = {
    version: "v0.0.1",
    hardsn: "hardsn-test",
  }
  
  const mutations = {
    DECREMENT_MAIN_COUNTER (state) {
      state.main--
    },
    INCREMENT_MAIN_COUNTER (state) {
      state.main++
    }
  }
  
  const actions = {
    someAsyncTask ({ commit }) {
      // do something async
      commit('INCREMENT_MAIN_COUNTER')
    }
  }
  
  export default {
    state,
    mutations,
    actions
  }
  