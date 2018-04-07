// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axiosInstance from './axios/config.js'
import Mint from 'mint-ui'
import 'mint-ui/lib/style.css'
import './assets/css/my-mint.scss'
import store from './vuex/store'
Vue.use(Mint);
Vue.prototype.$ajax = axiosInstance;
Vue.prototype.$messageBox = Mint.MessageBox;
Vue.prototype.$toast = Mint.Toast;
Vue.prototype.$indicator = Mint.Indicator;
Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
