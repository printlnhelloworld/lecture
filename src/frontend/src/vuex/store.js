import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    data: {
      type: 1
    }
  },
  mutations: {
    initData(state, data) {
      state.data = data;
    }
  }
})

export default store
