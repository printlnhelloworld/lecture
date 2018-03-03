<template>
  <div>
    <mt-header :title="title">
      <div @click="goback"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
    </mt-header>
    <section>
      <div class="edit">
        <mt-field label="主题" placeholder="主题名称" v-model="temp.topic"></mt-field>
        <mt-field label="时间" :placeholder="getTime(temp.startAt)" readonly v-on:click.native="openPicker"></mt-field>
        <mt-field label="地点" placeholder="地点" v-model="temp.location"></mt-field>
        <mt-field label="主办方" placeholder="主办方名称" v-model="temp.host"></mt-field>
        <mt-field label="主讲人" placeholder="主讲人姓名" v-model="temp.lecturer"></mt-field>
        <mt-field label="讲座类型" placeholder="请选择讲座类型" readonly v-model="temp.type" v-on:click.native="handleClick"></mt-field>
        <mt-field label="简介"  class="introduction" placeholder="简介" type="textarea" rows="8" v-model="temp.introduction"></mt-field>
        <mt-button type="primary" size="small" @click="submit" v-if="create">提交</mt-button>
        <mt-button type="primary" size="small" @click="submit" v-if="!create">保存</mt-button>
        <mt-button type="primary" size="small"  @click="$router.go(-1)">取消</mt-button>
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
      // 判断是创建还是修改
      create: !this.$route.query.id,
      title: this.$route.query.id ? '讲座编辑' : '创建讲座',
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
      popupVisible: false
      // 区别是创建还是修改 默认创建
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
      _self.$ajax({
        url: '/lectures/' + _self.$route.query.id,
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.temp = data.data;
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
      console.log(method)
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
  mounted() {
    if (!this.create) {
      this.getData();
    }
    console.log(!this.create)
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
