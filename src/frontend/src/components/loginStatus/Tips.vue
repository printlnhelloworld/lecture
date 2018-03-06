<template>
  <div>
    <div class="rule">
      <h4 class="title">使用本系统前,请先阅读《课外教育规定》</h4>
      <section>
        <p v-for="item of msg" :key="item">
          {{item}}
        </p>
      </section>
      <mt-button type="primary" @click="agreeToUse" >已阅读并同意使用本系统</mt-button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      msg: []
    }
  },
  methods: {
    getAgreement() {
      let _self = this;
      _self.$ajax({
        url: '/public/agreement',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.msg = data.data;
        } else {
          console.log(data.msg);
        }
      })
    },
    agreeToUse() {
      let _self = this;
      _self.$ajax({
        url: '/user/agree',
        method: 'post'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.$router.push('/index');
        } else {
          console.log(data.msg);
        }
      })
    }
  },
  mounted() {
    this.getAgreement();
  }
}
</script>
<style lang="scss" scoped>
.rule{
  display: flex;
  flex-direction: column;
  align-items: center;
  section{
    margin: 2rem;
    p{
      word-break: break-all;
      text-indent: 2rem;
    }
  }
}
</style>
