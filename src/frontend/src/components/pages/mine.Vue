<template>
  <div class="page-part">
    <div>
      <div class="info">
        <!-- <img src="../assets/logo.png" alt=""> -->
        <div>
          <span><label>姓名:</label>{{$store.state.data.name}}</span>
          <span><label>学号:</label>{{$store.state.data.id}}</span>
        </div>
        <div>
          <span><label>{{$store.state.data.id === 3 ? '归属单位' : '学院'}}:</label>{{$store.state.data.unitName}}</span>
        </div>
      </div>
      <div v-if="$store.state.data.type == 1">
        <mt-cell v-on:click.native="toogle(0)" class="account" :title="'已参与的学院专业讲座×' + mine.majorCount">
          <img class="arrow" src="../../assets/icon/show.png" v-if="show[0]"/>
          <img class="arrow" src="../../assets/icon/hidden.png" v-if="!show[0]"/>
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
          <img class="arrow" src="../../assets/icon/show.png" v-if="show[1]"/>
          <img class="arrow" src="../../assets/icon/hidden.png" v-if="!show[1]"/>
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
          <img class="arrow" src="../../assets/icon/show.png" v-if="show[2]"/>
          <img class="arrow" src="../../assets/icon/hidden.png" v-if="!show[2]"/>
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
</template>

<script>
import { formatDate } from '../../utils.js'
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
  methods: {
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
    getTime(time) {
      return formatDate(time);
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
          // 清除auth
          localStorage.removeItem('auth');
          window.open('http://i.hdu.edu.cn/dcp/logout0.jsp');
        } else {
          console.log(data.msg);
        }
      })
    }
  },
  computed: {
    permit() {
      // console.log(this.$store.state.data.permit)
      return this.$store.state.data.permit;
    }
  },
  mounted() {
    console.log(this.permit)
    console.log(this.lectures.list);
    this.getMineData();
    if (this.permit.lectureCreate) {
      this.getCreateList();
    }
  }
}
</script>

<style lang='scss' scoped>
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
.info {
  background-color: white;
  height: 4rem;
  padding: 0 1rem;
  display: flex;
  align-items: center;
  img {
    width: 2rem;
    height: 2rem;
  }
  div {
    display: flex;
    flex-direction: column;
    margin: 0 1rem;
  }
}
</style>
