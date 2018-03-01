import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    data: JSON.parse(localStorage.getItem('data'))
  },
  mutations: {
    initData(state, data) {
      state.data = data;
    }
  }
})

export default store
