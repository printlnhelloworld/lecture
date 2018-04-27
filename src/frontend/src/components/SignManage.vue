<template>
  <div class="wrap">
    <mt-header title="签到管理">
      <div @click="$router.go('-1')"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
      <mt-button icon="more" slot="right"></mt-button>
    </mt-header>
    <div v-if="lecture.status === 'notsigning'" class="signWrap">
      请选择发起签到的方式
      <div class="buttonGroup">
        <mt-button @click="changeStatus('signing')" :disabled="lecture.status === 'ended'" type="primary">开始签到</mt-button>
        <!-- <mt-button @click="changeStatus('signing')" :disabled="lecture.status === 'ended'" type="primary">大屏幕展示二维码</mt-button> -->
      </div>
      <!-- <mt-field v-if="lecture.status === 'signing'" label="签到码网址" v-model="lectureCode" disabled></mt-field>
      <mt-field v-if="lecture.status === 'signing'" label="讲座认证号" v-model="signCode" disabled></mt-field> -->
    </div>
    <div v-show="lecture.status === 'signing'" class="qrWrap">
      <p><span>刷新时间:</span><span>{{ expireIn }}</span></p>
      <div id="qrcode"></div>
      <p><span>签到码:</span><span>{{signCode}}</span></p>
      <mt-button @click="changeStatus('notsigning')" type="primary">暂停签到</mt-button>
    </div>
  </div>
</template>

<script>
import QRCode from 'qrcodejs2'
let qrcode;
export default {
  data() {
    return {
      lecture: {
        id: this.$route.query.id,
        status: 'notsigning'
      },
      signCode: '',
      expireIn: 0
    }
  },
  computed: {
    timeOut: {
      set(val) {
        this.$store.state.timeout.getSignCode = val;
      },
      get() {
        return this.$store.state.timeout.getSignCode;
      }
    },
    permit() {
      // console.log(this.$store.state.data.permit)
      return this.$store.state.data.permit;
    }
  },
  watch: {
    signCode(cur, old) {
      let _self = this;
      if (_self.lecture.status === 'signing') {
        qrcode.makeCode(localStorage.getItem('baseURL') + '/lecture?id=' + _self.lecture.id + '&signCode=' + _self.signCode);
      }
    }
  },
  methods: {
    getData() {
      let _self = this;
      return _self.$ajax({
        url: '/lectures/' + _self.lecture.id,
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.lecture = data.data;
          if (_self.lecture.status === 'signing') {
            _self.getSignCode()
          } else if (_self.lecture.status === 'notsigning') {
            if (_self.timeOut) {
              clearTimeout(_self.timeOut);
            }
          }
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    // 初始化二维码对象
    qrcodeInit() {
      qrcode = new QRCode(document.getElementById('qrcode'), {
        width: 200,
        height: 200
      });
    },
    getSignCode () {
      let _self = this;
      // 由于设置定时器 必须检查调用方法时是否是当前组件
      if (_self.$route.path !== '/signManage') {
        console.log('not signManage')
        return;
      }
      _self.$ajax({
        url: '/lectures/' + _self.lecture.id + '/signinCode',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.expireIn = data.expireIn;
          console.log(data.expireIn)
          if (_self.signCode !== data.signinCode) {
            _self.signCode = data.signinCode;
          }
        } else {
          this.$toast(data.msg);
        }
        _self.timeOut = setTimeout(_self.getSignCode, 1000);
      }).catch(err => {
        if (err.response.data.status === 'Forbidden') {
          _self.getData();
        }
      })
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
            _self.getData();
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
    }
  },
  mounted() {
    // 初始化二维码对象
    this.qrcodeInit();
    // 保证定时器唯一
    if (this.timeOut) {
      clearTimeout(this.timeOut);
    }
    this.getData();
    console.log(this.signCode)
  }
}
</script>

<style lang="scss" scoped>
.wrap{
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
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
.mint-header{
  flex: 0 0 auto;
}
.qrWrap{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  >button {
    margin: 1rem;
  }
}
#qrcode{
  height: 60vmin;
  width: 60vmin;
  /deep/ img{
    width: 100%;
    height: 100%;
  }
}
</style>
