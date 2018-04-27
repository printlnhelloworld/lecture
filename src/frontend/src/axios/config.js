import Vue from 'vue';
import axios from 'axios';
import router from '../router';
require('es6-promise').polyfill();
// 设置调试模式
console.log(localStorage.getItem('debugUrl'))
const debugUrl = localStorage.getItem('debugUrl');
let baseURL = window.location.protocol + '//' + window.location.host;
if (debugUrl) {
  baseURL = debugUrl;
}
localStorage.setItem('baseURL', baseURL);
const apiBaseUrl = baseURL + '/api/v1'
const axiosInstance = axios.create({
  baseURL: apiBaseUrl
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
// 添加响应拦截器
axiosInstance.interceptors.response.use(function (response) {
  // 对响应数据做点什么
  return response;
}, function (error) {
  Vue.prototype.$indicator.close();
  Vue.prototype.$toast(error.response.data.msg);
  // 对响应错误做点什么
  if (error.response.status === 401) {
    console.log('token过期');
    localStorage.setItem('redirect', window.location.href);
    localStorage.removeItem('auth');
    router.replace({
      path: '/login/error'
    })
  }
  return Promise.reject(error);
});
export default axiosInstance
