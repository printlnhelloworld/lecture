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
      </div>
      <div class="buttonGroup">
        <mt-switch v-model="temp.reviewed">是否经过相关部门审批</mt-switch>
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
import axios from 'axios'
export default {
  data() {
    return {
      slots: [
        {
          values: ['']
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
        reviewed: false,
        host: '',
        lecturer: '',
        type: ''
      },
      // 编辑模式 时间控件临时参数
      pickerTime: new Date(),
      // 编辑模式 切换讲座类下选择的弹出层
      popupVisible: false
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
  watch: {
    temp: {
      handler: function (newVal) {
        console.log(newVal);
        this.pickerTime = new Date(newVal.startAt * 1000);
      },
      deep: true
    }
  },
  methods: {
    getData() {
      let _self = this;
      axios.all([_self.getLectureData(), _self.getLectureType()])
        .then(axios.spread((res1, res2) => {
          let data1 = res1.data;
          let data2 = res2.data;
          if (data2.status === 'ok') {
            _self.slots[0].values.push(...data2.data);
          } else {
            _self.$toast(data2.msg);
          }
          if (data1.status === 'ok') {
            _self.temp = data1.data;
          } else {
            _self.$toast(data1.msg);
          }
        }))
    },
    getLectureData() {
      let _self = this;
      return _self.$ajax({
        url: '/lectures/' + _self.$route.query.id,
        method: 'get'
      })
    },
    getLectureType() {
      let _self = this;
      return _self.$ajax({
        url: '/public/lecture_type',
        method: 'get'
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
      let url = _self.create ? '/lectures' : '/lectures/' + _self.temp.id;
      let method = _self.create ? 'post' : 'put';
      let data = {};
      Object.assign(data, _self.temp);
      data.startAt = Date.parse(_self.temp.startAt) / 1000;
      _self.$indicator.open('Loading...');
      _self.$ajax({
        url: url,
        method: method,
        data: data
      }).then(res => {
        let data = res.data;
        _self.$indicator.close();
        if (data.status === 'ok') {
          _self.$toast({
            message: '操作成功'
          });
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    openPicker() {
      this.$refs.picker.open();
    },
    handleConfirm() {
      this.temp.startAt = this.pickerTime.getTime() / 1000;
    },
    // 讲座类型点击后显示选择的弹出层
    handleClick() {
      this.popupVisible = true;
    },
    TypeChange(picker, values) {
      console.log(picker, values)
      this.temp.type = values[0];
    }
  },
  mounted() {
    let _self = this;
    if (!this.create) {
      this.getData();
    } else {
      this.getLectureType()
        .then(res => {
          let data = res.data;
          if (data.status === 'ok') {
            _self.slots[0].values.push(...data.data);
          } else {
            _self.$toast(data.msg);
          }
        });
    }
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
.mint-switch{
  margin: 0.2rem;
}
</style>
