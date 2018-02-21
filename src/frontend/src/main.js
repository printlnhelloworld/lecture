// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
Vue.config.productionTip = false
require('es6-promise').polyfill();
const axiosInstance = axios.create({
  baseURL: 'http://lecture.hduhelp.com/',
  headers: { 'Content-Type': 'application/x-www-form-urlencoded/api/v1' }
})
Vue.prototype.$ajax = axiosInstance;
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
