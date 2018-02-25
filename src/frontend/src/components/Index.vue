<template>
  <div class="wrap">
    <div class="page-wrap">
      <mt-header :title="title"></mt-header>
      <div>
        <mt-cell class="page-part" title="当前选中" :value="selected"/>
      </div>
      <!-- tabcontainer -->
      <mt-tab-container class="page-tabbar-container" v-model="selected">
        <mt-tab-container-item id="list">
          <mt-loadmore
          :top-method="loadTop"
          :bottom-method="loadBottom"
          :bottom-all-loaded="allLoaded"
          auto-fill
          bottomPullText="上拉加载更多"
          bottomDropText="释放加载更多"
          bottomLoadingText="加载中"
          ref="loadmore">
            <div class="lectureList">
              <div v-for="item in list" class="lectureItem" :key="item.item">
                <span>{{ item.topic }}</span>
                <section>
                  <p>
                    <span>{{item.status}}</span>
                    <span>{{ getHM(item.startTimeStamp) }}</span>
                  </p>
                    <span>{{ getYMD(item.startTimeStamp) }}</span>
                </section>
              </div>
            </div>
            <!-- <ul class="page-loadmore-list">
              <li v-for="item in list" class="page-loadmore-listitem" :key="item.item">{{ item.item }}</li>
            </ul> -->
          </mt-loadmore>
        </mt-tab-container-item>
        <mt-tab-container-item id="mine">
          <div class="page-part">
           <!-- cell -->
            <mt-cell v-for="n in 12" :title="'我的 ' + n" :key="n"/>
          </div>
          <router-link to="/">
           <!-- button -->
            <mt-button type="danger" size="large">退出</mt-button>
          </router-link>
        </mt-tab-container-item>
      </mt-tab-container>
    </div>
    <mt-tabbar v-model="selected">
      <mt-tab-item id="list">
        <img slot="icon" src="../assets/icon/list.png">
        列表
      </mt-tab-item>
      <mt-tab-item id="mine">
        <img slot="icon" src="../assets/icon/mine.png">
        个人
      </mt-tab-item>
    </mt-tabbar>
  </div>
</template>

<script>
import { formatDateYMD, formatDateHM } from '../utils.js'
export default {
  data() {
    return {
      selected: 'list',
      list: [
        {
          'id': 1,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 2,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 3,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 4,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 5,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 6,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 7,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 8,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 9,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 10,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 11,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 12,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 13,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 14,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 15,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': '未开始',
          'startTimeStamp': 1519389118000
        },
        {
          'id': 16,
          'topic': 'xxx讲座',
          'type': 'xxx',
          'status': 'prepare',
          'startTimeStamp': 1519389118000
        }
      ],
      allLoaded: false,
      scrollMode: 'auto'
    }
  },
  computed: {
    title() {
      switch (this.selected) {
        case 'list':
          return '讲座列表';
        case 'mine':
          return '我的';
        default:
          return '';
      }
    }
  },
  methods: {
    // 上拉加载更多
    loadBottom() {
      console.log('bottom');
      this.allLoaded = true;// if all data are loaded
      this.$refs.loadmore.onBottomLoaded();
    },
    // 下拉刷新
    loadTop() {
      this.$refs.loadmore.onTopLoaded();
    },
    // 获取列表数据
    getData() {
      let _self = this;
      _self.$ajax({
        url: '/lectures',
        method: 'get'
      }).then(res => {
        let data = res.data;
        _self.list.push(...data.data)
        if (data.isLast === false) {
          // 数据加载完毕,无法上拉
          _self.allLoaded = true;
        }
      })
    },
    getYMD(time) {
      return formatDateYMD(time);
    },
    getHM(time) {
      return formatDateHM(time);
    }
  }
}
</script>

<style lang='scss' scoped>
.wrap{
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}
// .lectureList{
//   background-color: #bfd7e2;
// }
.page-wrap{
  // overflow: scroll;
  flex: 1;
  display: flex;
  flex-direction: column;
}
.page-tabbar-container{
  flex:1;
  padding-bottom: 55px;
  overflow: scroll;
}
.lectureList{
  height: 100%;
}
.lectureItem{
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 2.5rem;
  margin: 0 0.5rem 0 0.5rem;
  // font-size: 1rem;
  padding: 0 0.5rem 0 0.5rem;
  // border: 0.5rem black dotted;
  // border-top:none;
  border-radius: 0.5rem;
  background-color: white;
  >section {
    display: flex;
    flex-direction: column;
    font-size: 0.75rem;
    >p {
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>
