<template>
  <div class="wrap">
    <div class="page-wrap">
      <mt-header :title="title">
        <router-link id="mine" to="/editLecture" slot="right" v-if="permit.lectureCreate === true">
          <mt-button>创建</mt-button>
        </router-link>
      </mt-header>
      <!-- tabcontainer -->
      <mt-tab-container class="page-tabbar-container" v-model="selected" ref="wrap">
        <mt-tab-container-item id="list">
          <div class="loadmore_wrap">
            <mt-loadmore
            :top-method="loadTop"
            :bottom-method="loadBottom"
            :bottom-all-loaded="lectures.allLoaded"
            :auto-fill="false"
            ref="loadmore">
              <div class="lectureList">
                <router-link v-for="item in lectures.list" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                  <span><mt-badge :color="item.status === 'ended'? '#888' : (item.type === '校团委讲座' ? '#F44336' : '#26A2FF')">{{item.type === '校团委讲座' ? '校' : '院'}}</mt-badge>{{ item.topic }}</span>
                  <section>
                    <!-- <p>
                      <mt-badge size="small" color="#888">{{item.type === '校团委讲座' ? '校团委讲座' : '学院专业讲座'}}</mt-badge>
                    </p> -->
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
            <div>
              <div v-if="$store.state.data.type == 1">
                <mt-cell v-on:click.native="toogle(0)" class="account" :title="'已参与的学院专业讲座×' + mine.majorCount">
                  <img class="arrow" src="../assets/icon/show.png" v-if="show[0]"/>
                  <img class="arrow" src="../assets/icon/hidden.png" v-if="!show[0]"/>
                </mt-cell>
                <transition name="slide-fade">
                  <div class="lectureList" v-if="show[0]">
                    <router-link v-for="item in mine.list" v-if="item.type != '校团委讲座'" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                      <span><mt-badge :color="item.status === 'ended'? '#888' : (item.type === '校团委讲座' ? '#F44336' : '#26A2FF')">{{item.type === '校团委讲座' ? '校' : '院'}}</mt-badge>{{ item.topic }}</span>
                      <section>
                        <span>{{ getTime(item.startAt) }}</span>
                      </section>
                    </router-link>
                  </div>
                </transition>
                <mt-cell v-on:click.native="toogle(1)" class="account" :title="'已参与的校团委讲座×' + mine.schoolCount">
                  <img class="arrow" src="../assets/icon/show.png" v-if="show[1]"/>
                  <img class="arrow" src="../assets/icon/hidden.png" v-if="!show[1]"/>
                </mt-cell>
                <transition name="slide-fade">
                  <div class="lectureList" v-if="show[1]">
                    <router-link v-for="item in mine.list" v-if="item.type === '校团委讲座'" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                      <span><mt-badge :color="item.status === 'ended'? '#888' : (item.type === '校团委讲座' ? '#F44336' : '#26A2FF')">{{item.type === '校团委讲座' ? '校' : '院'}}</mt-badge>{{ item.topic }}</span>
                      <section>
                        <span>{{ getTime(item.startAt) }}</span>
                      </section>
                    </router-link>
                  </div>
                </transition>
              </div>
              <div v-if="permit.lectureCreate === true">
                <mt-cell v-on:click.native="toogle(2)" class="account" :title="'已创建讲座×' + mine.createCount">
                  <img class="arrow" src="../assets/icon/show.png" v-if="show[2]"/>
                  <img class="arrow" src="../assets/icon/hidden.png" v-if="!show[2]"/>
                </mt-cell>
                <transition name="slide-fade">
                  <div class="lectureList" v-if="show[2]">
                    <router-link v-for="item in mine.createList" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
                      <span><mt-badge :color="item.status === 'ended'? '#888' : (item.type === '校团委讲座' ? '#F44336' : '#26A2FF')">{{item.type === '校团委讲座' ? '校' : '院'}}</mt-badge>{{ item.topic }}</span>
                      <section>
                        <span>{{ getTime(item.startAt) }}</span>
                      </section>
                    </router-link>
                  </div>
                </transition>
              </div>
            </div>
            <mt-button @click="logout" type="danger">登出</mt-button>
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
      selected: 'list',
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
    permit() {
      // console.log(this.$store.state.data.permit)
      return this.$store.state.data.permit;
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
    getMineData(next) {
      let _self = this;
      _self.$ajax({
        url: '/user/lectures',
        method: 'get'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          _self.mine.majorCount = data.majorCount;
          _self.mine.schoolCount = data.schoolCount;
          _self.mine.list = data.list;
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    getCreateList(next = 0) {
      let _self = this;
      _self.$ajax({
        url: 'lectures',
        method: 'get',
        params: {
          owner: _self.$store.state.data.id,
          next: next
        }
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          if (data.data.length !== 0) {
            _self.mine.createList.push(...data.data);
            _self.mine.createCount = _self.mine.createList.length;
            _self.getCreateList(data.next);
          }
        } else {
          _self.$toast(data.msg);
        }
      })
    },
    // 上拉加载更多
    loadBottom() {
      let _self = this;
      if (_self.lectures.allLoaded) {
        _self.$refs.loadmore.onBottomLoaded();
        return;
      }
      console.log('bottom');
      _self.getData().then(res => {
        let data = res.data;
        _self.$refs.loadmore.onBottomLoaded();
        if (data.data.length === 0) {
          _self.lectures.allLoaded = true;
          _self.$toast({
            message: '已加载完毕',
            position: 'bottom'
          });
          console.log(_self.lectures.allLoaded)
        } else {
          _self.lectures.list.push(...data.data);
          _self.lectures.next = data.next;
        }
      });
    },
    // 下拉刷新
    loadTop() {
      let _self = this;
      _self.lectures.allLoaded = false;
      _self.lectures.next = 0;
      console.log('top');
      _self.getData().then(res => {
        _self.$refs.loadmore.onTopLoaded();
        let data = res.data;
        _self.lectures.list = data.data;
        _self.lectures.next = data.next;
      });
    },
    getTime(time) {
      return formatDate(time);
    },
    wrapInit() {
      let wrapHeight = document.getElementsByClassName('loadmore_wrap')[0].offsetHeight;
      // document.getElementsByClassName('loadmore_wrap')[0].style.height = wrapHeight - searchBarHeight + 'px';
      document.getElementsByClassName('lectureList')[0].style.minHeight = wrapHeight + 'px';
      // document.getElementsByClassName('mint-loadmore')[0].style.height = wrapHeight + 'px';
    },
    toogle(index) {
      this.show.splice(index, 1, !this.show[index]);
      console.log(this.show[index]);
    },
    logout() {
      let _self = this;
      _self.$ajax({
        url: '/user/tokens/self',
        method: 'delete'
      }).then(res => {
        let data = res.data;
        if (data.status === 'ok') {
          window.open('http://i.hdu.edu.cn/dcp/logout0.jsp');
        } else {
          console.log(data.msg);
        }
      })
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
    console.log(11)
    console.log('mounted')
    console.log(this.lectures.list);
    this.wrapInit();
    this.loadBottom();
    this.getMineData();
    this.getAnnouncement();
    if (this.permit.lectureCreate) {
      this.getCreateList();
    }
  },
  beforeRouteLeave(to, from, next) {
    let position = document.getElementsByClassName('loadmore_wrap')[0].scrollTop;
    console.log(position);
    this.$store.commit('savePosition', position); // 离开路由时把位置存起来
    next()
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
  height: 100%;
}
.page-tabbar-container{
  flex:1;
  margin-bottom: 55px;
  overflow: hidden;
}
#mine{
  overflow:scroll;
}
.lectureItem{
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 2.5rem;
  // font-size: 1rem;
  padding: 0 1rem 0 1rem;
  margin: 0 1rem 0 1rem;
  // border: 0.5rem black dotted;
  // border-top:none;
  border-radius: 0.5rem;
  font-size: 1rem;
  background-color: white;
  >span{
    >span{
      margin-right: 0.5rem;
    }
    flex: 0 1 auto;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }
  >section {
    display: flex;
    flex-direction: column;
    font-size: 1rem;
    flex-basis: 1;
    flex:0 0 auto;
    >p {
      display: flex;
      justify-content: flex-end;
    }
  }
}
.lectureList{
  padding-top: 0.4rem;
  box-sizing: border-box;
}
.lectureList>:not(:last-child){
  // border-top: 1px gainsboro solid;
  margin-bottom: 0.4rem;
}
.mint-cell-wrapper{
  height:1rem;
  font-size: 0.2rem;
}
.arrow{
  height: 1rem;
  width: 1rem;
}
.slide-fade-enter-active {
  transition: all .3s ease;
}
.slide-fade-leave-active {
  transition: all .3s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active for below version 2.1.8 */ {
  transform: translateX(10px);
  opacity: 0;
}
.page-part{
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  overflow: scroll;
  >button{
    flex: 0 0 auto;
  }
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
