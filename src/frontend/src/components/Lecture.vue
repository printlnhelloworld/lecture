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
      <div class="buttonGroup" v-if="authority && !edit">
        <mt-button type="primary" size="small">开始讲座</mt-button>
        <mt-button type="primary" size="small" @click="edit = true, create = false, temp = lecture">编辑讲座</mt-button>
        <mt-button type="primary" size="small">签到管理</mt-button>
        <mt-button type="primary" size="small">签到列表</mt-button>
        <mt-button type="danger" size="small">结束讲座</mt-button>
      </div>
      <div class="edit" v-if="edit">
        <mt-field label="主题" placeholder="请输入用户名" v-model="temp.topic"></mt-field>
        <mt-field label="时间" :placeholder="getTime(temp.startAt)" readonly v-on:click.native="openPicker"></mt-field>
        <mt-field label="地点" placeholder="请输入密码" v-model="temp.location"></mt-field>
        <mt-field label="主办方" placeholder="请输入手机号" v-model="temp.host"></mt-field>
        <mt-field label="主讲人" placeholder="请输入网址" v-model="temp.lecturer"></mt-field>
        <mt-field label="讲座类型" placeholder="请选择讲座类型" readonly v-model="temp.type" v-on:click.native="handleClick"></mt-field>
        <mt-field label="简介"  class="introduction" placeholder="请输入简介" type="textarea" rows="8" v-model="temp.introduction"></mt-field>
        <mt-button type="primary" size="small" @click="submit" v-if="create">提交</mt-button>
        <mt-button type="primary" size="small" @click="submit" v-if="!create">保存</mt-button>
        <mt-button type="primary" size="small" @click="submit" v-if="!create">取消</mt-button>
      </div>
      <div v-if="lecture.canSignin">
        <mt-field label="签到" placeholder="请输入签到码" v-model="signCode"></mt-field>
      </div>
    </section>
    <mt-datetime-picker
      ref="picker"
      v-model="pickerTime"
      type="datetime"
      year-format="{value}"
      month-format="{value}"
      date-format="{value}"
      hourFormat="{value}"
      minuteFormat="{value}"
      @confirm="handleConfirm">
    </mt-datetime-picker>
    <mt-popup
      v-model="popupVisible"
      position="bottom">
        <mt-picker :slots="slots" @change="TypeChange"></mt-picker>
    </mt-popup>
  </div>
</template>

<script>
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      slots: [
        {
          values: ['', '校团委讲座', '机械工程学院', '计算机学院讲座', '数字媒体与艺术设计学院', '国际教育学院', '外国语学院', '经济学院', '理学院', '材料与环境工程学院']
        }
      ],
      signCode: 0,
      pickerValue: null,
      title: '讲座详情',
      lecture: {
        id: 1,
        creatorUserID: '',
        topic: '',
        location: '',
        introduction: '',
        startAt: 0,
        host: '',
        lecturer: '',
        type: '',
        status: '',
        createAt: 0,
        finishedAt: 0,
        finished: false,
        canSignin: true,
        remark: ''
      },
      temp: {
        topic: '',
        location: '',
        introduction: '',
        startAt: new Date(),
        host: '',
        lecturer: '',
        type: ''
      },
      // 进入编辑模式
      edit: false,
      // 编辑模式 时间控件临时参数
      pickerTime: new Date(),
      // 编辑模式 切换讲座类下选择的弹出层
      popupVisible: false,
      // 区别是创建还是修改 默认创建
      create: true
    }
  },
  computed: {
    type() {
      return this.$store.state.data.type;
    },
    authority() {
      return this.$store.state.data.id === this.lecture.creatorUserID
    }
  },
  methods: {
    getData() {
      let _self = this;
      _self.$ajax({
        url: '/lectures/' + _self.$route.query.id,
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
    },
    // 创建/修改讲座信息
    submit() {
      let _self = this;
      let url = _self.create ? '/lectures/' : '/lectures/' + _self.lecture.id
      let method = _self.create ? 'post' : 'patch'
      _self.$ajax({
        url: url,
        method: method,
        data: _self.temp
      })
    },
    openPicker() {
      this.$refs.picker.open();
    },
    handleConfirm() {
      this.temp.startAt = this.pickerTime;
    },
    // 讲座类型点击后显示选择的弹出层
    handleClick() {
      this.popupVisible = true;
    },
    TypeChange(picker, values) {
      this.temp.type = values[0]
    }
  },
  mounted() {
    console.log(this.lecture.list)
    this.getData();
    console.log(this.create)
    console.log(this.edit)
  }
}
</script>

<style lang="scss" scoped>
section{
  padding:2rem;
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
.introduction{
  height: 10rem;
}
.edit{
  display: flex;
  flex-direction: column;
}
</style>
