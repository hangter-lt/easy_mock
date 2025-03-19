<template>
  <div style="display: flex">
    <div class="toli">
      <img class="titpoto" src="@/assets/yuo.png" alt="" />
      <div class="tolistway">
        <div
          class="quest"
          :class="{ active: activuBut == 'one' }"
          @click="OneBut()"
        >
          <div class="questword">实时请求</div>
        </div>
        <div
          class="quest"
          :class="{ active: activuBut == 'two' }"
          @click="TwoBut()"
        >
          <div class="questword">接口管理</div>
        </div>
        <div
          class="quest"
          :class="{ active: activuBut == 'three' }"
          @click="ThreeBut()"
        >
          <div class="questword">历史请求</div>
        </div>
      </div>
    </div>
    <ToRight v-if="onePage" />
    <ToManage v-if="twoPage" />
    <ToEdit v-if="threePage" />
  </div>
</template>

<script>
import ToRight from "./components/toright.vue";
import ToManage from "./components/tomanage.vue";
import ToEdit from "./components/toedit.vue";

export default {
  name: "MinIndex",
  components: {
    ToRight,
    ToManage,
    ToEdit,
  },
  data() {
    return {
      onePage: true,
      twoPage: false,
      threePage: false,
      activuBut: "one",
    };
  },
  methods: {
    OneBut() {
      this.onePage = true;
      this.twoPage = false;
      this.threePage = false;
      this.activuBut = "one";
    },
    TwoBut() {
      this.onePage = false;
      this.twoPage = true;
      this.threePage = false;
      this.activuBut = "two";
    },
    ThreeBut() {
      this.onePage = false;
      this.twoPage = false;
      this.threePage = true;
      this.activuBut = "three";
    },
  },
};
</script>

<style scoped>
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}

/* 左侧列表 */
.toli {
  height: 90vh;
  margin-top: 5vh;
  width: 12%;
  margin-left: 5%;
  background-color: #46475e;
  border-bottom-left-radius: 200px;
  border-top-left-radius: 200px;
  color: #7669d3;
}
.titpoto {
  width: 1000px;
  height: 1000px;
  border-radius: 50%;
  margin-top: 4vh;
}
.tolistway {
  display: flex;
  flex-direction: column;
  width: 100%;
  align-items: center;
}
.quest {
  width: 950px;
  height: 950px;
  background-color: #cbc9ff;
  border-radius: 100px;
  margin-top: 400px;
  cursor: pointer;
}
.questword {
  font-weight: 900;
  font-size: 320px;
}
.quest.active {
  color: #7660f6;
}
</style>
