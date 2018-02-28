<template>
  <div>
    <mt-header :title="title">
      <div @click="goback"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
    </mt-header>
    <section>
      <ul class="detials" v-if="!edit">
        <li><span>主题：{{ lecture.topic }}</span></li>
        <li><span>时间：{{ getTime(lecture.startAt) }}</span></li>
        <li><span>地点：{{ lecture.location }}</span></li>
        <li><span>主办方：{{ lecture.host }}</span></li>
        <li><span>主讲人：{{ lecture.lecturer }}</span></li>
        <li><span>讲座类型：{{ lecture.type }}</span></li>
        <li><span>内容简介：{{ lecture.introduction }}</span></li>
      </ul>
      <div class="buttonGroup">
        <mt-button type="primary">开始讲座</mt-button>
        <mt-button type="primary">编辑讲座</mt-button>
        <mt-button type="primary">签到管理</mt-button>
        <mt-button type="primary">签到列表</mt-button>
        <mt-button type="danger">结束讲座</mt-button>
      </div>
      <div v-if="edit">
        <mt-field label="主题" placeholder="请输入用户名" v-model="username"></mt-field>
        <mt-field label="时间" placeholder="请输入邮箱" type="email" v-model="email"></mt-field>
        <mt-field label="地点" placeholder="请输入密码" type="password" v-model="password"></mt-field>
        <mt-field label="主办方" placeholder="请输入手机号" type="tel" v-model="phone"></mt-field>
        <mt-field label="主讲人" placeholder="请输入网址" type="url" v-model="website"></mt-field>
        <mt-field label="讲座类型" placeholder="请输入数字" type="number" v-model="number"></mt-field>
        <mt-field label="内容简介" placeholder="自我介绍" type="textarea" rows="4" v-modal="introduction"></mt-field>
      </div>
    </section>
  </div>
</template>

<script>
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      title: '讲座详情',
      lecture: {
        id: 1,
        creatorUserID: '04xxx',
        topic: 'xxxx讲座',
        location: '6教南110',
        introduction: 'xxxxxxx',
        startAt: 1519389118000,
        host: 'xxx',
        lecturer: 'XXX',
        type: '校团委讲座',
        status: 'runing/ended/prepare',
        createAt: 1111111111,
        finishedAt: 1111111111,
        finished: true,
        canSignin: true,
        remark: '讲座自动完成'
      },
      edit: false
    }
  },
  computed: {
    type() {
      return this.$store.state.data.type;
    },
    createrFlag() {
      return this.$store.data.id === this.lecture.creatorUserID
    }
  },
  methods: {
    getData() {
      let _self = this;
      _self.$ajax({
        url: '/lectures',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.lecture = data.data;
        } else {
          alert(data.msg);
        }
      })
    },
    goback() {
      this.$router.go('-1');
    },
    getTime(time) {
      return formatDate(time);
    }
  }
}
</script>

<style lang="scss" scoped>
section{
  padding:10%;
}
.detials{
  display: flex;
  flex-direction: column;
  align-items: center;
  >li{
    list-style: none;
    width: 100%;
    line-height: 1.5rem;
  }
}
.buttonGroup{
  margin: 2rem 0 0 0;
  display:flex;
  flex-direction: column;
  align-items: center;
  >button{
    width: 80%;
  }
}
</style>
