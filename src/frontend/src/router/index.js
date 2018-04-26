import Vue from 'vue'
import Router from 'vue-router'
const Index = () => import('@/components/Index')
const Login = () => import('@/components/Login')
const Lecture = () => import('@/components/Lecture')
const EditLecture = () => import('@/components/EditLecture')
const SignManage = () => import('@/components/SignManage')
const SignRecord = () => import('@/components/SignRecord')
const Tips = () => import('@/components/loginStatus/Tips')
const Error = () => import('@/components/loginStatus/Error')
Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/index'
    },
    {
      path: '/index',
      name: 'index',
      component: Index,
      meta: {
        keepAlive: true,
        requireAuth: true
      }
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {
        keepAlive: false,
        savedPosition: 0
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
        keepAlive: false,
        requireAuth: true
      }
    },
    {
      path: '/editLecture',
      name: 'EditLecture',
      component: EditLecture,
      meta: {
        keepAlive: false,
        requireAuth: true
      }
    },
    {
      path: '/signManage',
      name: 'SignManage',
      component: SignManage,
      meta: {
        keepAlive: false,
        requireAuth: true
      }
    },
    {
      path: '/signRecord',
      name: 'SignRecord',
      component: SignRecord,
      meta: {
        keepAlive: false,
        requireAuth: true
      }
    }
  ]
  // scrollBehavior(to, from, savedPosition) {
  //   if (savedPosition) {
  //     return savedPosition
  //   } else {
  //     if (from.meta.keepAlive) {
  //       console.log('alive')
  //       from.meta.savedPosition = 30;
  //     }
  //     console.log(to.meta)
  //     return { x: 0, y: to.meta.savedPosition || 0 }
  //   }
  // }
})
router.beforeEach((to, from, next) => {
  console.log('defend')
  if (to.meta.requireAuth) {
    var auth = localStorage.getItem('auth');
    console.log(auth);
    if (auth) {
      console.log('has auth')
      next();
    } else {
      console.log('no auth')
      console.log(window.location.href);
      localStorage.setItem('redirect', window.location.href);
      next({
        path: '/login'
      })
    }
  } else {
    next();
  }
})
export default router
