import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    data: JSON.parse(localStorage.getItem('data')),
    position: 0,
    timeout: {
      getSignCode: ''
    }
  },
  mutations: {
    initData(state, data) {
      state.data = data;
    },
    savePosition(state, position) {
      state.position = position;
    }
  }
})

export default store
