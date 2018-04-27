import Vue from 'vue'
import Router from 'vue-router'
import store from '../vuex/store'
const Index = () => import('@/components/Index')
const Login = () => import('@/components/Login')
const Lecture = () => import('@/components/Lecture')
const EditLecture = () => import('@/components/EditLecture')
const SignManage = () => import('@/components/SignManage')
const SignRecord = () => import('@/components/SignRecord')
const Tips = () => import('@/components/loginStatus/Tips')
const Error = () => import('@/components/loginStatus/Error')
const LectureList = () => import('@/components/pages/LectureList')
const Mine = () => import('@/components/pages/Mine')
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
      redirect: '/index/lectures',
      component: Index,
      meta: {
        requireAuth: true,
        keepAlive: true
      },
      children: [
        {
          path: 'lectures',
          name: 'lectures',
          component: LectureList,
          meta: {
            keepAlive: true
          }
        },
        {
          path: 'mine',
          name: 'mine',
          component: Mine,
          meta: {
            keepAlive: true
          }
        }
      ]
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {
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
        requireAuth: true
      }
    },
    {
      path: '/editLecture',
      name: 'EditLecture',
      component: EditLecture,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/signManage',
      name: 'SignManage',
      component: SignManage,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/signRecord',
      name: 'SignRecord',
      component: SignRecord,
      meta: {
        requireAuth: true
      }
    }
  ],
  scrollBehavior(to, from, savedPosition) {
    console.log(to.name)
    if (to.meta.keepAlive) {
      let scrollTop = store.state.common.scrollPos[to.name] || 0
      if (!scrollTop) {
        return;
      }
      // 对scroll元素进行设置
      new Promise((resolve, reject) => {
        resolve()
      }).then(() => {
        let documentElem = document.querySelector('.scroll')
        console.log(documentElem);
        if (documentElem) {
          try {
            documentElem.scrollTop = scrollTop;
          } catch (err) {
            console.log(err)
          }
        }
      })
    }
  }
})
router.beforeEach((to, from, next) => {
  // 记录上一个页面的scroll位置
  if (from.name) {
    let contentElem = document.querySelector('.scroll')
    let scrollTop = contentElem ? contentElem.scrollTop : '0'
    store.state.common.scrollPos[from.name] = scrollTop;
  }
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
