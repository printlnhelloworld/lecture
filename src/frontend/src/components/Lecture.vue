<template>
  <div>
    <mt-header :title="title">
      <div @click="goback"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
    </mt-header>
    <section>
      <ul class="detials">
        <li><span>主题：{{ lecture.topic }}</span></li>
        <li><span>时间：{{ getTime(lecture.startAt) }}</span></li>
        <li><span>地点：{{ lecture.location }}</span></li>
        <li><span>主办方：{{ lecture.host }}</span></li>
        <li><span>主讲人：{{ lecture.lecturer }}</span></li>
        <li><span>讲座类型：{{ lecture.type }}</span></li>
        <li><span>内容简介：{{ lecture.introduction }}</span></li>
      </ul>
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
      }
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
.detials{
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 2rem;
  >li{
    list-style: none;
    width: 60%;
    line-height: 1.5rem;
  }
}
</style>
