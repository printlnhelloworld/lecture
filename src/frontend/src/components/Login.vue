<template>
  <div class="wrap">
    <router-view/>
  </div>
</template>

<script>
export default {
  data() {
    return {
      auth: localStorage.getItem('auth')
    }
  },
  methods: {
    login() {
      let _self = this;
      let auth = _self.$route.query.auth;
      // 判断是否是url中是否带有auth(即是不是新登录)
      if (auth) {
        localStorage.setItem('auth', auth)
        _self.getUserInfo();
      } else if (_self.auth) {
      // 判断本地是否已存在auth (即是不是已登录)
        console.log('login already')
        _self.getUserInfo();
      } else {
        console.log('no login')
        _self.gotoLogin()
      }
    },
    getUserInfo() {
      let _self = this;
      _self.$ajax({
        url: '/user/userinfo',
        method: 'get'
      }).then(res => {
        let data = res.data;
        _self.agree = data.data.agree;
        // 判断是否登录token是否有效
        if (data.status === 'ok') {
          localStorage.setItem('data', JSON.stringify(data.data));
          _self.$store.commit('initData', JSON.parse(localStorage.getItem('data')));
          // console.log(_self.$store.state.data)
          // 判断是否阅读相关规则并同意使用本系统
          if (_self.agree === true) {
            _self.$router.push('/index');
          } else {
            _self.$router.push('/login/tips');
          }
        } else {
          localStorage.removeItem('auth');
          _self.loading = false;
          _self.flag = false;
          setTimeout(() => {
            _self.gotoLogin();
          }, 1000)
        }
      }).catch(err => {
        alert(err)
      })
    },
    gotoLogin() {
      console.log('gotologin')
      let _self = this;
      _self.$ajax({
        url: '/loginURL',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          window.location.href = data.loginURL;
        }
      })
    },
    editUserinfo() {
      let _self = this;
      _self.$ajax({
        url: '/user/userinfo',
        method: 'put',
        data: {
          agree: 'true'
        }
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.$router.push({
            path: '/index'
          })
        } else {
          alert(data.msg);
        }
      }).catch(err => {
        console.log(err);
      })
    }
  },
  mounted() {
    this.login();
  }
}
</script>
<style lang="scss" scoped>
.wrap{
  display: flex;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
  padding-top: 30%;
  height: 100%;
  img{
    width:10rem;
  }
}
.rule{
  display: flex;
  flex-direction: column;
  align-items: center;
  >p{
    width: 80%;
    word-break: break-all;
    margin: 2rem 2rem 2rem 2rem;
  }
}
</style>
