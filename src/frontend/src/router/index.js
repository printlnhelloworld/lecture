import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '@/components/Login'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      redirect: '/index'
    },
    {
      path: '/index',
      name: 'Index',
      component: Index
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    }
  ]
})
// router.beforeEach((to, from, next) => {
//   var auth = localStorage.getItem('auth');
//   if (auth || to.path === '/login') {
//     next();
//   } else {
//     next({
//       path: '/login'
//     })
//   }
// })
export default router
