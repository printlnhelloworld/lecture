// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import Mint from 'mint-ui'
import 'mint-ui/lib/style.css'
import './assets/css/my-mint.scss'
require('es6-promise').polyfill();
Vue.use(Mint);
const axiosInstance = axios.create({
  baseURL: 'http://lecture.hduhelp.com/api/v1',
  headers: { 'Authorization': 'x' }
})
Vue.prototype.$ajax = axiosInstance;
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
