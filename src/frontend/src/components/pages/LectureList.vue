<template>
  <div class="page-part">
    <div class="searchbar">
      <div>1</div>
      <div>2</div>
      <div>3</div>
    </div>
    <div class="loadmore_wrap scroll">
      <mt-loadmore
      :top-method="loadTop"
      :bottom-all-loaded="lectures.allLoaded"
      :auto-fill="false"
      ref="loadmore">
      <ul
        class="lectureList"
        v-infinite-scroll="loadMore"
        infinite-scroll-disabled="lectures.allLoaded"
        infinite-scroll-distance="10">
          <router-link is="li" v-for="item in lectures.list" :to="{path:'/lecture',query:{id:item.id}}" class="lectureItem" :key="item.item">
            <div class="item-main">
              <p class="lecTopic">{{ item.topic }}</p>
              <p class="lecType"><img src="../../assets/icon/school2.png"/> {{ item.type }}</p>
              <p class="lecLoc"><img src="../../assets/icon/location.png"/> {{ item.location }}</p>
            </div>
            <section>
                <span>{{ getYMD(item.startAt) }}</span>
                <span>{{ getHM(item.startAt) }}</span>
            </section>
          </router-link>
      </ul>
        <!-- <ul class="page-loadmore-list">
          <li v-for="item in list" class="page-loadmore-listitem" :key="item.item">{{ item.item }}</li>
        </ul> -->
      </mt-loadmore>
    </div>
  </div>
</template>

<script>
import { formatDate, formatDateYMD, formatDateHM } from '../../utils.js'
export default {
  data() {
    return {
      lectures: {
        allLoaded: false,
        next: 0,
        list: []
      },
      loading: false
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
    // 加载更多
    loadMore() {
      let _self = this;
      _self.loading = true;
      _self.getData().then(res => {
        this.loading = false;
        let data = res.data;
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
    getYMD(time) {
      return formatDateYMD(time);
    },
    getHM(time) {
      return formatDateHM(time);
    }    
  },
  mounted () {
    this.wrapInit();
  }
}
</script>

<style lang='scss' scoped>
$searchbarHeight: 2rem;
.page-part{
  height: 100%;
  padding-top: $searchbarHeight;
  overflow: hidden;
}
.lectureList{
  padding: 0.4rem 0;
  box-sizing: border-box;
}
.lectureList>:not(:last-child){
  border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}
.loadmore_wrap{
  overflow: scroll;
  height: 100%;
}
.lectureItem{
  display: flex;
  justify-content: space-between;
  align-items: center;
  // height: 2.5rem;
  padding: 0.5rem 1rem;
  // border: 0.5rem black dotted;
  // border-top:none;
  background-color: white;
  .item-main{
    flex:1 1 auto;
    width: 5rem;
    .lecTopic{
      flex: 1 1 auto;
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
    }
    .lecLoc, .lecType{
      display: flex;
      align-items: center;
      color: #666;
      font-size: 14px;
      >img{
        width: 14px;
        height: 14px;
        margin-right: 0.2rem;
      }
    }
    .lecType{
      margin: 0.1rem 0;
    }
  }
  >section {
    display: flex;
    flex-direction: column;
    align-self: flex-start;
    font-size: 12px;
    flex:0 0 auto;
    >span {
      display: flex;
      align-items: center;
      justify-content: flex-end;
      height: 18px;
    }
  }
}
.lectureList{
  padding: 0.4rem 0;
  box-sizing: border-box;
}
.searchbar{
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
</style>
