import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    selectedAPI: {},
  },
  mutations: {
    setSelectedAPI(state, api) {
      state.selectedAPI = api
    }
  },
  actions: {
  },
  methods:{
  
  },
  modules: {
  }
})
