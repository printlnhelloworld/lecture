<template>
  <div class="wrap">
    <mt-header title="签到记录">
      <div @click="$router.go('-1')"  slot="left">
        <mt-button icon="back">返回</mt-button>
      </div>
      <mt-button @click="openSign" slot="right">手动签到</mt-button>
    </mt-header>
    <table>
      <thead>
        <tr>
          <th>共有{{total}}人签到</th>
        </tr>
        <tr>
          <td class="id">学号</td>
          <td class="name">姓名</td>
          <td class="time">签到时间</td>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in list" :key="item.userId">
          <td class="id">{{item.userId}}</td>
          <td class="name">{{item.name}}</td>
          <td class="time">{{getTime(item.signedAt)}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      lecture: {
        id: this.$route.query.id,
        status: 'notsigning'
      },
      signCode: '',
      total: 0,
      list: [
        // {
        //   userId: '15051342',
        //   name: '叶建武',
        //   signedAt: new Date().getTime() / 1000,
        //   type: 'qcode',
        //   remark: ''
        // }
      ]
    }
  },
  methods: {
    getData() {
      let _self = this;
      _self.$ajax({
        url: '/lectures/' + _self.lecture.id + '/users',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.total = data.total;
          _self.list = data.data;
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    signByHand(value) {
      let _self = this;
      _self.$ajax({
        url: '/lectures/' + _self.lecture.id + '/users/byhand',
        method: 'post',
        data: {
          id: value
        }
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.getData();
          if (data.data.name.lenth === 0) {
            _self.$messageBox('提示', '签到成功,但该学号未曾登录本系统,因此无法获取姓名,请确保该学号正确');
          } else {
            _self.$toast('签到成功');
          }
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    getTime(time) {
      return formatDate(time);
    },
    openSign() {
      let _self = this;
      _self.$messageBox.prompt('请输入学号').then(({ value, action }) => {
        if (action === 'confirm') {
          console.log('confirm');
          _self.signByHand(value);
        }
      });
    }
  },
  mounted () {
    this.getData();
  }
}
</script>

<style lang="scss" scoped>
.wrap{
  display: flex;
  flex-direction: column;
  height: 100%;
  overflow: hidden;
}
table{
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}
thead tr, tbody tr{
  display: flex;
  justify-content: center;
}
thead, tfoot{
  flex: 0 0 auto;
}
tbody{
  flex: 0 0 1;
  overflow: scroll;
}
td{
  padding: 0.5rem 0 0.5rem 0;
  text-align: center;
}
.id{
  flex-basis: 1;
  flex: 0 0 30%;
}
.name{
  flex-basis: 1;
  flex: 0 0 30%;
}
.time{
  flex-basis: 2;
  flex: 0 0 40%;
}
tr:nth-child(odd) td {
  background-color: #eeeeee;
}
</style>
