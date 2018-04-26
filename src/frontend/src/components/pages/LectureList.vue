<template>
  <div class="loadmore_wrap">
    <div class="searchbar">
      <div>1</div>
      <div>2</div>
      <div>3</div>
    </div>
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
</template>

<script>
import { formatDate } from '../../utils.js'
export default {
  data() {
    return {
      lectures: {
        allLoaded: false,
        next: 0,
        list: []
      }
    }
  },
  methods: {
    wrapInit() {
      let wrapHeight = document.getElementsByClassName('loadmore_wrap')[0].offsetHeight;
      // document.getElementsByClassName('loadmore_wrap')[0].style.height = wrapHeight - searchBarHeight + 'px';
      document.getElementsByClassName('lectureList')[0].style.minHeight = wrapHeight + 'px';
      // document.getElementsByClassName('mint-loadmore')[0].style.height = wrapHeight + 'px';
    },
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
    }
  },
  mounted () {
    this.wrapInit();
    this.loadBottom();
  }
}
</script>

<style lang='scss' scoped>
$searchbarHeight: 2rem;
.lectureList{
  padding: 0.4rem 0;
  box-sizing: border-box;
}
.lectureList>:not(:last-child){
  /* border-top: 1px gainsboro solid; */
  margin-bottom: 0.4rem;
}
.loadmore_wrap{
  box-sizing: border-box;
  position: relative;
  padding-top: $searchbarHeight;
  overflow: scroll;
  height: 100%;
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
  padding: 0.4rem 0;
  box-sizing: border-box;
}
.searchbar{
  position: fixed;
  margin-top: -$searchbarHeight;
  display: flex;
  align-items: center;
  // justify-content: space-between;
  height: $searchbarHeight;
  width: 100%;
  background-color: white;
  z-index: 999;
  >div{
    flex: 1 0 auto;
  }
}
.lectureList>:not(:last-child){
  // border-top: 1px gainsboro solid;
  margin-bottom: 0.4rem;
}
</style>
