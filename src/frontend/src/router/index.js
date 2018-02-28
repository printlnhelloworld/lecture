import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '@/components/Login'
import Lecture from '@/components/Lecture'
import Tips from '@/components/loginStatus/Tips'
import Error from '@/components/loginStatus/Error'

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
      component: Index,
      meta: {
        keepAlive: true
      }
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {
        keepAlive: false
      },
      children: [
        {
          path: 'tips',
          name: 'Tips',
          component: Tips
        },
        {
          path: 'error',
          name: 'Error',
          component: Error
        }
      ]
    },
    {
      path: '/lecture',
      name: 'Lecture',
      component: Lecture,
      meta: {
        keepAlive: false
      }
    }
  ]
})
router.beforeEach((to, from, next) => {
  var auth = localStorage.getItem('auth');
  console.log(auth);
  if (auth || to.path === '/login') {
    next();
  } else {
    next({
      path: '/login'
    })
  }
})
export default router
