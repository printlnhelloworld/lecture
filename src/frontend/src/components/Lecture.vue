<template>
  <div class="wrap">
    <mt-header :title="title">
      <div @click="goback"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
    </mt-header>
    <section>
        <mt-field label="主题" placeholder="主题名称" v-model="lecture.topic" readonly></mt-field>
        <mt-field label="时间" :placeholder="getTime(lecture.startAt)" readonly></mt-field>
        <mt-field label="地点" placeholder="地点" v-model="lecture.location" readonly></mt-field>
        <mt-field label="主办方" placeholder="主办方名称" v-model="lecture.host" readonly></mt-field>
        <mt-field label="主讲人" placeholder="主讲人姓名" v-model="lecture.lecturer" readonly></mt-field>
        <mt-field label="讲座类型" placeholder="请选择讲座类型" v-model="lecture.type" readonly></mt-field>
        <mt-field label="简介"  class="introduction" placeholder="简介" type="textarea" rows="8" v-model="lecture.introduction" readonly></mt-field>
      <div class="buttonGroup" v-if="authority">
        <mt-button @click="$router.push({path: '/signManage',query:{id: $route.query.id}})" type="primary">签到管理</mt-button>
        <!-- <mt-button v-if="lecture.status === 'runing'" type="primary">暂停签到</mt-button> -->
        <mt-button @click="$router.push({path: '/signRecord',query:{id: $route.query.id}})" type="primary">签到记录</mt-button>
        <mt-button v-if="lecture.status !== 'ended'" type="primary" @click="$router.push({path: '/editLecture', query:{id: $route.query.id}})">编辑讲座</mt-button>
        <mt-button v-if="lecture.status !== 'ended'" type="danger" @click="deleteLecture">删除讲座</mt-button>
        <mt-button v-if="lecture.status !== 'ended'" type="danger" @click="changeStatus('ended')">结束讲座</mt-button>
      </div>
      <div v-if="!authority">
        <div v-if="lecture.status === 'signing'" class="sign">
          <mt-field label="签到" placeholder="请输入签到码" v-model="signCode"></mt-field>
          <mt-button type="primary" size="small" @click="signIn()">签到</mt-button>
        </div>
        <!-- <div v-if="lecture.canSignin && !lecture.signin.isSigned" class="sign">
          <mt-field label="签到" placeholder="请输入签到码" v-model="signCode"></mt-field>
          <mt-button type="primary" size="small" @click="signIn">签到</mt-button>
        </div> -->
      </div>
    </section>
  </div>
</template>

<script>
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      flag: false,
      slots: [
        {
          values: ['', '校团委讲座', '机械工程学院', '计算机学院讲座', '数字媒体与艺术设计学院', '国际教育学院', '外国语学院', '经济学院', '理学院', '材料与环境工程学院']
        }
      ],
      signCode: this.$route.query.signCode ? this.$route.query.signCode : '',
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
        status: '',
        createAt: 0,
        finishedAt: 0,
        finished: false,
        canSignin: false,
        remark: '',
        signin: {
          isSigned: false, // 当前token的用户是否已经完成签到
          SignedAt: 0, // 签到时间
          type: '', // 签到类型
          remark: '' // 备注
        }
      }
    }
  },
  computed: {
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
          _self.$toast(data.msg);
        }
      })
    },
    goback() {
      // history.back();
      this.$router.push('/index');
    },
    getTime(time) {
      return formatDate(time);
    },
    changeStatus(status) {
      let _self = this;
      let tips, msg;
      switch (status) {
        case 'signing' :
          msg = '签到已开始';
          tips = '是否确认开始签到？'
          break;
        case 'notsigning' :
          msg = '签到已结束';
          tips = '是否确认结束签到？'
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
            method: 'put',
            data: {
              status: status
            }
          }).then(res => {
            _self.$indicator.close();
            let data = res.data;
            if (data.status === 'ok') {
              _self.$messageBox('提示', '操作成功');
            } else {
              _self.$toast(msg);
            }
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
    // 签到码签到
    signIn(qrType = 'code') {
      console.log(qrType)
      let _self = this;
      _self.$ajax({
        url: '/lectures/' + _self.lecture.id + '/users/' + qrType,
        method: 'post',
        data: {
          code: _self.signCode,
          type: qrType
        }
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.getData();
          _self.$toast('签到成功');
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    signStart() {
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
      if (this.signCode) {
        console.log('has signCode')
        this.signIn();
      }
    });
  }
}
</script>

<style lang="scss" scoped>
.wrap{
  height: 100%;
  overflow: height;
  display: flex;
  flex-direction: column;
}
.mint-header{
  flex: 0 0 auto;
}
section{
  padding:2rem;
  overflow: scroll;
  flex-basis: 1;
  flex: 0 1 auto;
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
  width: 100%;
  >button{
    margin: 0.2rem 0 0 0;
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
.modal{
    width: 100%;
    height: 100%;
    background-color: #fff;
}
.mask{
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top:10rem;
  width: 100%;
  height: 100%;
  >button{
    width: 50%;
  }
}
</style>
