<template>
  <div class="wrap">
    <mt-header title="签到管理">
      <router-link :to="{path:'/lecture', query:{id: lecture.id}}" slot="left">
        <mt-button icon="back">返回</mt-button>
      </router-link>
      <mt-button icon="more" slot="right"></mt-button>
    </mt-header>
    <div class="signWrap">
      请选择发起签到的方式
      <div class="buttonGroup">
        <mt-button v-if="lecture.status === 'notsiging'" @click="changeStatus('signing')" :disabled="lecture.status === 'ended'" type="primary">手机展示二维码</mt-button>
        <mt-button v-if="lecture.status === 'notsiging'" @click="changeStatus('signing')" :disabled="lecture.status === 'ended'" type="primary">大屏幕展示二维码</mt-button>
        <mt-button v-if="lecture.status === 'siging'" @click="changeStatus('notsigning')" type="primary">暂停签到</mt-button>
      </div>
      <mt-field v-if="lecture.status === 'siging'" label="签到码网址" v-model="lectureCode" disabled></mt-field>
      <mt-field v-if="lecture.status === 'siging'" label="讲座认证号" v-model="signCode" disabled></mt-field>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lecture: {
        id: this.$route.query.id,
        status: 'notsiging'
      }
    }
  },
  methods: {
    getData() {
      let _self = this;
      return _self.$ajax({
        url: '/lectures/' + _self.id,
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.lecture = data.data;
        } else {
          _self.$toast(data.msg);
        }
      })
    }
  },
  computed: {
    permit() {
      // console.log(this.$store.state.data.permit)
      return this.$store.state.data.permit;
    }
  }
}
</script>

<style lang="scss" scoped>
.wrap{
  display: flex;
  flex-direction: column;
  height: 100%;
  .signWrap{
    flex: 1 1 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top:5rem;
    width: 100%;
    height: 100%;
    >button{
      width: 50%;
    }
  }
}
</style>
