<template>
  <div class="wrap">
    <div class="page-wrap">
      <mt-header :title="title">
        <router-link id="mine" to="/editLecture" slot="right" v-if="permit.lectureCreate === true">
          <mt-button>创建</mt-button>
        </router-link>
      </mt-header>
      <!-- tabcontainer -->
      <keep-alive>
        <router-view></router-view>
      </keep-alive>
    </div>
    <mt-tabbar v-model="selected">
      <mt-tab-item id="lectures">
        <img slot="icon" src="../assets/icon/list.png">
        列表
      </mt-tab-item>
      <mt-tab-item id="mine">
          <img slot="icon" src="../assets/icon/mine.png">
          个人
      </mt-tab-item>
    </mt-tabbar>
    <mt-popup
      class="announcement"
      closeOnClickModal="false"
      v-model="announcement.visible">
      <h3><span>{{_self.announcement.list[0].title}}</span></h3>
      <p class="message">{{_self.announcement.list[0].content}}</p>
      <p class="announcementTime">{{getTime(_self.announcement.list[0].createAt)}}</p>
      <p class="confirm" type="primary" @click="announcement.visible = !announcement.visible">确定</p>
    </mt-popup>
  </div>
</template>

<script>
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      announcement: {
        // 是否显示公告
        visible: false,
        // 已阅读的公告id数组
        idList: new Set(JSON.parse(localStorage.getItem('announcements')) || []),
        list: [{
          title: '',
          content: '',
          createAt: 0
        }]
      },
      show: [false, false, false],
      list2: [
        {
          'id': 1,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startAt': 1519389118000
        },
        {
          'id': 2,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startAt': 1519389118000
        },
        {
          'id': 3,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startAt': 1519389118000
        }
      ],
      scrollMode: 'auto',
      mine: {
        // 参与专业讲座数目
        majorCount: 0,
        // 参与校团委讲他做数目
        schoolCount: 0,
        createCount: 0,
        list: [
          // {
          //   // 讲座id
          //   id: 1,
          //   // 讲座主题
          //   topic: 'xxxx讲座',
          //   // 讲座类别 参考数字字典
          //   type: '校团委讲座',
          //   // 时间戳 秒级
          //   startAt: 1519389118000,
          //   signType: 'qcode'
          // }
        ],
        createList: []
      },
      options: [1, 2, 3],
      lectures: {
        allLoaded: false,
        next: 0,
        list: []
      }
    }
  },
  computed: {
    selected: {
      get: function () {
        return this.$route.name;
      },
      set: function (to) {
        this.$router.push({name: to});
      }
    },
    title() {
      switch (this.selected) {
        case 'lectures':
          return '讲座列表';
        case 'mine':
          return '我的';
        default:
          return '';
      }
    },
    permit() {
      // console.log(this.$store.state.data.permit)
      return this.$store.state.data.permit;
    }
  },
  methods: {
    getTime(time) {
      return formatDate(time);
    },
    getAnnouncement() {
      let _self = this;
      _self.$ajax({
        url: '/announcements',
        method: 'get'
      }).then(res => {
        let data = res.data;
        console.log(_self.announcement.idList)
        if (data.status === 'ok') {
          // 检查本地是否存储最新的公告id
          _self.announcement.list = data.data;
          if (!_self.announcement.idList.has(data.data[0].id)) {
            _self.announcement.idList.add(data.data[0].id);
            _self.announcement.visible = true;
            // 把公告id存储到本地 即已读
            localStorage.setItem('announcements', JSON.stringify(Array.from(_self.announcement.idList)));
          }
        }
      })
    }
    // getYMD(time) {
    //   return formatDateYMD(time);
    // },
    // getHM(time) {
    //   return formatDateHM(time);
    // }
  },
  mounted() {
    console.log(this.permit)
    console.log(this.lectures.list);
  }
}
</script>

<style lang='scss' scoped>
$searchbarHeight: 2rem;
.wrap{
  height: 100%;
  width: 100%;
}
.mint-cell{
  position:relative;
  top:0;
  left: 0;
}
.page-wrap{
  overflow: hidden;
  height: 100%;
  box-sizing: border-box;
  padding-bottom: 55px;
  display: flex;
  flex-direction: column;
}
.page-tabbar-container{
  flex:1;
  overflow: hidden;
}
#mine{
  overflow:scroll;
}
.announcement{
  width:20rem;
  padding: 0 1rem 0 1rem;
  >h3{
    height: 3rem;
    line-height: 3rem;
    text-align: center;
    border-bottom: 1px gainsboro solid;
  }
  .announcementTime{
    text-align: right;
  }
  .confirm{
    height: 3rem;
    line-height: 3rem;
    text-align: center;
    color: #8ACDFF;
    border-top: 1px gainsboro solid;
  }
  .message{
    margin: 1rem;
    text-indent: 2rem;
  }
}
</style>
