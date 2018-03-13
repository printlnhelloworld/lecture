// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import Mint from 'mint-ui'
import 'mint-ui/lib/style.css'
import './assets/css/my-mint.scss'
import store from './vuex/store'
require('es6-promise').polyfill();
Vue.use(Mint);
// 设置调试模式
console.log(localStorage.getItem('debug'))
const debug = localStorage.getItem('debug');
const baseURL = debug ? 'http://localhost:8080' : 'https://lecture.hduhelp.com';
localStorage.setItem('baseURL', baseURL);
const apiBaseUrl = baseURL + '/api/v1'
const axiosInstance = axios.create({
  baseURL: apiBaseUrl
})
Vue.prototype.$ajax = axiosInstance;
Vue.prototype.axios = axios;
Vue.prototype.$messageBox = Mint.MessageBox;
Vue.prototype.$toast = Mint.Toast;
Vue.prototype.$indicator = Mint.Indicator;
Vue.config.productionTip = false;
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
// 添加响应拦截器
axiosInstance.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  return response;
}, function (error) {
  Vue.$indicator.close();
  Vue.$toast(error.response.data.msg);
  // 对响应错误做点什么
  if (error.response.status === 401) {
    console.log('token过期');
    localStorage.removeItem('auth');
    router.replace({
      path: '/login/error'
    })
  }
  return Promise.reject(error);
});
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
