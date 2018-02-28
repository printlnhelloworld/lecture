<template>
  <div class="wrap">
    <img src="../../assets/icon/warning.png" alt="">
    <div>登录信息失效,请重新登录</div>
  </div>
</template>

<script>
export default {
  methods: {
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
    }
  },
  mounted() {
    this.gotoLogin();
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
</style>
