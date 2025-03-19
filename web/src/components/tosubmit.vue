<template>
  <div class="torig">
    <div class="wtobtn">
      <div class="tomabut">
        <div class="tobuttext">提交</div>
      </div>
    </div>
    <div class="mablist">
      <div class="eddou">
        <div class="edman">名称:&nbsp;
          <el-input v-model="input" style="width: 240px;height: 30px; line-height: 30px;"  placeholder="请输入名称" />
        </div>
        <div class="edmang">分组:&nbsp;用户管理</div>
      </div>
      <div class="edmanway">请求方法:&nbsp;GET</div>
      <div class="edmantype">请求类型:&nbsp;application/json</div>
      <div class="elblocklist">
        <div class="edblock">
          <div class="edbllist">
            <div class="requa">请求参数:&nbsp;</div>
            <div class="requecri"></div>
          </div>
          <div class="edgray"></div>
          <div class="edbllista">
            <div class="requ">响应码:&nbsp;200</div>
            <div class="requ">响应类型:&nbsp;application/json</div>
            <div class="rebl">
              <div class="requ">响应内容:&nbsp;</div>
              <div class="requecria"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ToSubmit',
  props: {
    msg: String,
  },
  data(){
    return{
      input:'',
    }
  }
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

/* 右侧列表 */
.torig {
  height: 90vh;
  margin-top: 5vh;
  /* width: 1440px; */
  width: 66%;
  background-color: #181825;
  display: flex;
  flex-direction: column;
  border-top-right-radius: 200px;
  border-bottom-right-radius: 200px;
  padding: 0px 6%;
  color: #fff;
}
.wtobtn {
  width: 100%;
  display: flex;
  justify-content: end;
}
.tomabut {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 6vh;
  height: 4vh;
  padding: 100px 250px;
  margin-top: 200px;
}
.tobuttext {
  letter-spacing: 4px;
}
.mablist {
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  /* background-color: pink; */
  margin: 100px 0px;
}
.mablist::-webkit-scrollbar {
  width: 0;
  background: transparent;
}
.eddou {
  width: 100%;
  display: flex;
}
.edman,
.edmang {
  width: 50%;
  text-align: left;
  flex-wrap: wrap;
  margin-bottom: 200px;
}
.edmanway,
.edmantype {
  text-align: left;
  flex-wrap: wrap;
  margin-bottom: 200px;
}
.inname{
 width: 20vh;
 height: 5vh;
}
::v-deep .el-input__wrapper{
  background-color: transparent !important;
  box-shadow:none !important;
  border-bottom: 1px solid #333 !important;
  border-radius:0 !important;
}
::v-deep .el-input__inner::placeholder {
  color: #999 !important; /* 设置 placeholder 颜色 */
}
::v-deep .el-input__inner {
  font-size: 40px !important; /* 设置字体大小 */
  color: #333 !important; /* 设置文字颜色 */
  height: 20px;
}
/* ::v-deep .el-input__wrapper .el-input__inner{
  font-size: 80px !important;
} */
/* 历史请求列表 */
.elblocklist {
  overflow-y: auto;
}
.elblocklist::-webkit-scrollbar {
  width: 0;
  background: transparent;
}
.edblock {
  display: flex;
  border-radius: 300px;
  border: 2px solid #FEFFFF;
  padding: 200px 0px 200px 3%;
  margin-top: 100px;
  margin: 100px 0 300px;
  background-color: #46475E;
}
.edbllist {
  display: flex;
  flex-direction: column;
  width: 100%;
}
.requecri {
  width: 88%;
  margin: 200px 3% 0px 0px;
  border-radius: 100px;
  border: 2px solid #FEFFFF;
  padding: 200px 2%;
  min-height: 16vh;
}
.requecria {
  width: 76%;
  margin: 0px 0px 200px 0px;
  border-radius: 100px;
  border: 2px solid #FEFFFF;
  padding: 200px 2%;
  min-height: 10vh;
}
.rebl {
  display: flex;
  width: 100%;
}
.requa {
  display: flex;
}
.requ {
  display: flex;
  margin-bottom: 200px;
}
.edgray {
  width: 1px;
  height: 100%;
  border-right: 2px solid #333;
}
.edbllista {
  display: flex;
  flex-direction: column;
  width: 100%;
  margin: 0 200px 0 3%;
}
</style>
