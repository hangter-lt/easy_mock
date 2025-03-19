<template>
  <div class="torig">
    <div class="wtobtn">
      <div class="tomabut">
        <div class="tobuttext">创建</div>
      </div>
    </div>
    <div class="mablist">
      <div class="useman">用户管理</div>
      <div class="usemaun"></div>
      <div class="usmanblist">
        <div class="requecri">
          <div class="requ">创建</div>
          <div class="requtw">
            <div class="reget">GET</div>
            <div class="regetu">/api/v1/test</div>
          </div>
          <div class="requoe">application/json</div>
        </div>
        <div class="requecri">
          <div class="requ">更新</div>
          <div class="requtw">
            <div class="reget">GET</div>
            <div class="regetu">/api/v1/test</div>
          </div>
          <div class="requoe">application/json</div>
        </div>
        <div class="requecri">
          <div class="requ">删除</div>
          <div class="requtw">
            <div class="reget">GET</div>
            <div class="regetu">/api/v1/test</div>
          </div>
          <div class="requoe">application/json</div>
        </div>
      </div>
      <div class="useman">用户管理</div>
      <div class="usemaun"></div>
      <div class="usmanblist">
        <div class="requecri">
          <div class="requ">创建</div>
          <div class="requtw">
            <div class="reget">GET</div>
            <div class="regetu">/api/v1/test</div>
          </div>
          <div class="requoe">application/json</div>
        </div>
      </div>
      <div class="useman">用户管理</div>
      <div class="usemaun"></div>
      <div class="usmanblist">
        <div class="requecri">
          <div class="requ">创建</div>
          <div class="requtw">
            <div class="reget">GET</div>
            <div class="regetu">/api/v1/test</div>
          </div>
          <div class="requoe">application/json</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ToManage',
  props: {
    msg: String,
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
.mablist{
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  margin: 100px 0px;
}
.mablist::-webkit-scrollbar {
  width: 0;
  background: transparent;
}
.useman {
  display: flex;
  margin-top: 100px;
}
.usemaun {
  width: 100%;
  border-top: 2px solid #333;
  margin-top: 200px;
}
/* 用户列表 */
.usmanblist {
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 100px;
}
.requecri {
  width: 26.8%;
  margin: 200px 0px 200px 0px;
  border-radius: 100px;
  border: 2px solid #FEFFFF;
  padding: 200px 2%;
  margin-right: 3%;
  background-color: #46475E;
}
.requecri:nth-child(3n) {
  margin-right: 0px;
}
.requ {
  display: flex;
}
.requtw {
  display: flex;
  justify-content: space-between;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 89%;
  padding: 100px 5%;
  margin-top: 200px;
}
.requoe {
  display: flex;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 89%;
  padding: 100px 5%;
  margin-top: 200px;
}
</style>
