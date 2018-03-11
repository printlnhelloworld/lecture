import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '@/components/Login'
import Lecture from '@/components/Lecture'
import EditLecture from '@/components/EditLecture'
import SignManage from '@/components/SignManage'
import Tips from '@/components/loginStatus/Tips'
import Error from '@/components/loginStatus/Error'
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
    }
  ]
  // scrollBehavior(to, from, savedPosition) {
  //   if (savedPosition) {
  //     console.log(router.app.$store.state.position)
  //     return new Promise((resolve, reject) => {
  //       setTimeout(() => {
  //         window.scrollTo(0, 0)
  //       }, 500)
  //     })
  //   } else {
  //     return { x: 0, y: 0 };
  //   }
  // }
})
router.beforeEach((to, from, next) => {
  if (to.meta.requireAuth) {
    var auth = localStorage.getItem('auth');
    console.log(auth);
    if (auth) {
      next();
    } else {
      next({
        path: '/login'
      })
    }
  }
})
export default router
