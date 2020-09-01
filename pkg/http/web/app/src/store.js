import Vue from 'vue';
import Vuex from 'vuex';

// import APIClient from './apiClient';

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        apis: [],
    },
    mutations: {
        resetAPIs(state, apis) {
            state.apis = apis;
        },
    },
    getters: {
        apis(state) {
            return state.apis;
        },
    },
    actions: {
  
    }
});

export default store;