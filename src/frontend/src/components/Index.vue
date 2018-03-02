<template>
  <div class="wrap">
    <div class="page-wrap">
      <mt-header :title="title"></mt-header>
      <!-- tabcontainer -->
      <mt-tab-container class="page-tabbar-container" v-model="selected" ref="wrap">
        <mt-tab-container-item id="list">
          <div class="searchBar">
            <!-- <select2 :options="options" v-model="selected">
              <option disabled value="0">Select one</option>
            </select2> -->
          </div>
          <div class="loadmore_wrap" :style="{height: wrapHeight + 'px'}">
            <mt-loadmore
            :top-method="loadTop"
            :bottom-method="loadBottom"
            :bottom-all-loaded="lectures.allLoaded"
            auto-fill
            ref="loadmore">
              <div class="lectureList">
                <router-link v-for="item in lectures.list" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                  <span>{{ item.topic }}</span>
                  <section>
                    <p>
                      <mt-badge size="small" color="#888">{{item.type}}</mt-badge>
                      <span>{{item.status}}</span>
                    </p>
                      <span>{{ getTime(item.startAt) }}</span>
                  </section>
                </router-link>
              </div>
              <!-- <ul class="page-loadmore-list">
                <li v-for="item in list" class="page-loadmore-listitem" :key="item.item">{{ item.item }}</li>
              </ul> -->
            </mt-loadmore>
          </div>
        </mt-tab-container-item>
        <mt-tab-container-item id="mine">
          <div class="page-part">
            <mt-cell class="account" :title="'学院专业讲座×' + mine.marjorCount"/>
            <div class="lectureList">
              <router-link v-for="item in list2" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                <span>{{ item.topic }}</span>
                <section>
                  <p>
                    <mt-badge size="small" color="#888">{{item.type}}</mt-badge>
                    <span>{{item.status}}</span>
                  </p>
                    <span>{{ getTime(item.startAt) }}</span>
                </section>
              </router-link>
            </div>
            <mt-cell class="account" :title="'校团委讲座×' + mine.marjorCount"/>
            <div class="lectureList">
              <router-link v-for="item in list2" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                <span>{{ item.topic }}</span>
                <section>
                  <p>
                    <mt-badge size="small" color="#888">{{item.type}}</mt-badge>
                    <span>{{item.status}}</span>
                  </p>
                    <span>{{ getTime(item.startAt) }}</span>
                </section>
              </router-link>
            </div>
          </div>
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
import { formatDate } from '../utils.js'
export default {
  data() {
    return {
      selected: 'list',
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
        marjorCount: 0,
        // 参与校团委讲他做数目
        schoolCount: 0,
        list: [
          {
            // 讲座id
            id: 1,
            // 讲座主题
            topic: 'xxxx讲座',
            // 讲座类别 参考数字字典
            type: '校团委讲座',
            // 时间戳 秒级
            startAt: 1519389118000,
            signType: 'qcode'
          }
        ]
      },
      wrapHeight: '',
      options: [1, 2, 3],
      lectures: {
        allLoaded: false,
        next: 0,
        list: []
      }
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
    },
    type() {
      return this.$store.state.data.type;
    }
  },
  methods: {
    // 获取列表数据
    getData() {
      let _self = this;
      return _self.$ajax({
        url: '/lectures',
        method: 'get',
        params: {
          next: _self.lectures.next
        }
      })
    },
    // 上拉加载更多
    loadBottom() {
      let _self = this;
      console.log('bottom');
      _self.getData().then(res => {
        let data = res.data;
        _self.$refs.loadmore.onBottomLoaded();
        if (data.data.length === 0) {
          _self.lectures.allLoaded = true;
        } else {
          _self.lectures.list.push(...data.data);
          _self.lectures.next = data.next;
          console.log(_self.lectures.next);
          console.log(_self.lectures.list);
        }
      });
    },
    // 下拉刷新
    loadTop() {
      let _self = this;
      _self.$refs.loadmore.onTopLoaded();
      _self.lectures.allLoaded = false;
      _self.lectures.next = 0;
      console.log('top');
      _self.getData().then(res => {
        let data = res.data;
        _self.lectures.list = data.data;
        _self.lectures.next = data.next;
      });
    },
    getMineData() {

    },
    getTime(time) {
      return formatDate(time);
    },
    wrapInit() {
      let wrapHeight = document.getElementsByClassName('mint-tab-container')[0].offsetHeight;
      let searchBarHeight = document.getElementsByClassName('searchBar')[0].offsetHeight;
      document.getElementsByClassName('loadmore_wrap')[0].style.height = wrapHeight - searchBarHeight + 'px';
      // document.getElementsByClassName('lectureList')[0].style.height = wrapHeight - searchBarHeight + 'px';
      console.log(document.getElementsByClassName('loadmore_wrap')[0].style.height)
    }
    // getYMD(time) {
    //   return formatDateYMD(time);
    // },
    // getHM(time) {
    //   return formatDateHM(time);
    // }
  },
  mounted() {
    console.log('mounted')
    console.log(this.lectures.list);
    let _self = this;
    this.$nextTick(() => {
      _self.wrapInit();
    })
    this.loadTop();
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
.mint-cell{
  position:relative;
  top:0;
  left: 0;
}
.page-wrap{
  overflow: hidden;
  flex: 1;
  display: flex;
  flex-direction: column;
}
.loadmore_wrap{
  // flex:1;
  overflow: scroll;
  // height: 200px;
}
.page-tabbar-container{
  flex:1;
  margin-bottom: 55px;
  overflow: hidden;
}
#mine{
  overflow:scroll;
}
.lectureList{
  height: 100%;
  margin: 0.5rem 0 0.5rem  0
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
.mint-cell-wrapper{
  height:1rem;
  font-size: 0.75rem;
}
</style>
