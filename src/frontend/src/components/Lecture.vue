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
      <div class="buttonGroup" v-if="authority">
        <mt-button v-if="lecture.status === 'prepare'" type="primary" size="small" @click="changeStatus('runing')">开始讲座</mt-button>
        <mt-button v-if="lecture.status != 'ended'" type="primary" size="small" @click="$router.push({path: '/editLecture', query:{id: $route.query.id}})">编辑讲座</mt-button>
        <mt-button v-if="lecture.status === 'runing'" type="primary" size="small">签到管理</mt-button>
        <mt-button v-if="lecture.status != 'prepare'" type="primary" size="small">签到记录</mt-button>
        <mt-button v-if="lecture.status != 'ended'" type="danger" size="small" @click="deleteLecture">删除讲座</mt-button>
        <mt-button v-if="lecture.status === 'runing'" type="danger" size="small" @click="changeStatus('ended')">结束讲座</mt-button>
      </div>
      <div v-if="$store.state.data.type == 1">
        <div v-if="lecture.canSignin && !lecture.signin.isSigned" class="sign">
          <mt-field label="签到" placeholder="请输入签到码" v-model="signCode"></mt-field>
          <mt-button type="primary" size="small" @click="signIn">签到</mt-button>
        </div>
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
      signCode: '',
      pickerValue: null,
      title: '讲座详情',
      lecture: {
        id: 1,
        // creatorUserID: '15051342',
        creatorUserID: '',
        topic: '',
        location: '',
        introduction: '',
        startAt: 0,
        host: '',
        lecturer: '',
        type: '',
        status: 'runing',
        createAt: 0,
        finishedAt: 0,
        finished: false,
        canSignin: false,
        remark: '',
        signin: {
          isSigned: false, // 当前token的用户是否已经完成签到
          SignedAt: 111111111111111, // 签到时间
          type: 'byhand', // 签到类型
          remark: '' // 备注
        }
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
      console.log(this.$store.state.data.id === this.lecture.creatorUserID)
      return this.$store.state.data.id === this.lecture.creatorUserID
    }
  },
  methods: {
    getData() {
      let _self = this;
      return _self.$ajax({
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
      history.back();
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
    changeStatus(status) {
      let _self = this;
      let tips, msg;
      switch (status) {
        case 'runing' :
          msg = '讲座已开始';
          tips = '是否确认开始讲座？'
          break;
        case 'ended' :
          msg = '讲座已结束';
          tips = '结束讲座后将不能再编辑讲座,是否确认结束讲座？'
          break;
        default :
          msg = 'error';
          break;
      }
      _self.$messageBox.confirm(tips).then(action => {
        if (action === 'confirm') {
          _self.$indicator.open('Loading...');
          _self.$ajax({
            url: '/lectures/' + _self.lecture.id + '/status',
            method: 'patch',
            data: {
              status: status
            }.then(res => {
              _self.$indicator.close();
              let data = res.data;
              if (data.status === 'ok') {
                _self.$messageBox('提示', '操作成功');
              } else {
                _self.$toast(msg);
              }
            })
          })
        } else {
          console.log('concel');
        }
      });
    },
    // 删除讲座信息
    deleteLecture() {
      let _self = this;
      _self.$messageBox.confirm('是否确认删除讲座？').then(action => {
        if (action === 'confirm') {
          _self.$indicator.open('Loading...');
          _self.$ajax({
            url: '/lectures/' + _self.lecture.id,
            method: 'delete'
          }).then(res => {
            _self.$indicator.close();
            let data = res.data;
            if (data.status === 'ok') {
              _self.$messageBox.alert('删除成功').then(action => {
                _self.$router.go(-1);
              });
            } else {
              _self.$toast.alert(data.msg);
            }
          })
        } else {
          console.log('cancel');
        }
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
    },
    // 签到码签到
    signIn() {
      let _self = this;
      _self.$ajax({
        url: '',
        type: ''
      })
    }
  },
  created() {
    this.$indicator.open('Loading...');
  },
  mounted() {
    this.getData().then(() => {
      this.$indicator.close();
    });
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
.sign{
  margin-top: 1rem;
  display: flex;
  flex-direction: column;
}
</style>
