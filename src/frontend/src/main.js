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
  baseURL: 'https://lecture.hduhelp.com/api/v1'
})

// 添加请求拦截器
axiosInstance.interceptors.request.use(function (config) {
  // 在发送请求之前做些什么
  let auth = localStorage.getItem('auth');
  console.log(auth)
  if (auth) {
    config.headers.common['Authorization'] = auth;
  } else {
    router.replace({
      path: '/login'
    })
  }
  return config;
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error);
});
Vue.prototype.$ajax = axiosInstance;
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
