<template>
  <div class="wrap">
    <img v-if="flag && loading && agree" src="../assets/icon/success.png" alt="">
    <img v-if="!flag && loading && agree" src="../assets/icon/warning.png" alt="">
    <div v-if="loading && agree">{{ flag ? '登陆成功' : '登录信息失效'}}</div>
    <div v-if="!agree" class="rule">
      <h4 class="title">使用本系统前,请先阅读《课外教育规定》</h4>
      <p>
        asddddddddddddddddddddddddasdddddddddd
        asdasdasddddddddd
      </p>
      <mt-button type="primary" >已阅读并同意使用本系统</mt-button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: true,
      flag: true,
      agree: true
    }
  },
  methods: {
    setAuth() {
      let _self = this;
      localStorage.setItem('auth', _self.$route.auth)
      _self.$ajax({
        url: '/users',
        method: 'get'
      }).then(res => {
        let data = res.data;
        _self.agree = data.agree;
        // 判断是否登录token是否有效
        if (data.status === 'ok') {
          _self.loading = false;
          // 判断是否阅读相关规则并同意使用本系统
          if (_self.agree === true) {
            setTimeout(() => {
              _self.$router.push('/index');
            }, 1000)
          }
        } else {
          _self.loading = false;
          _self.flag = false;
          setTimeout(() => {
            _self.gotoLogin();
          }, 1000)
        }
      })
    },
    gotoLogin() {
      let _self = this;
      _self.$ajax({
        url: 'https://lecture.hduhelp.com/api/v1/loginURL',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          window.location.href = data.loginURL;
        }
      })
    }
  },
  mounted() {
    let _self = this;
    if (_self.$route.params.auth) {
      _self.setAuth();
    }
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
