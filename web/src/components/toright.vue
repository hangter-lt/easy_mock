<template>
  <div class="torig">
    <div class="tomess">
      <div class="toquonelist">
        <div class="tocri">
          <div class="toctop">
            <div class="totime">12:34:43</div>
            <div class="tostatus">200</div>
          </div>
          <div class="tocbot">GET/api/test</div>
        </div>
        <div class="tocri">
          <div class="toctop">
            <div class="totime">12:34:43</div>
            <div class="tostatus">200</div>
          </div>
          <div class="tocbot">GET/api/test</div>
        </div>
      </div>
    </div>
    <div class="torequest">
      <div class="requecri">
        <div class="requ">请求</div>
        <div class="requoe">2015/2/19 12:33:43</div>
        <div class="requtw">
          <div class="reget">GET</div>
          <div class="regetu">/api/v1/test</div>
        </div>
        <div class="rehead">
          <div class="rehename">header</div>
          <div class="recoment"></div>
        </div>
        <div class="reheada">
          <div class="rehename">params</div>
          <div class="recoment"></div>
        </div>
      </div>
    </div>
    <div class="torelevant">
      <div class="requecri">
        <div class="requ">响应</div>
        <div class="requtw">
          <div class="reget">200</div>
          <div class="regetu">application/json</div>
        </div>
        <div class="reheadb">
          <div class="rehename">body</div>
          <div class="recoment"></div>
        </div>
      </div>
    </div>
    <div></div>
  </div>
</template>

<script>
export default {
  name: 'ToRight',
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

/* 左侧列表 */
.torig {
  height: 90vh;
  margin-top: 5vh;
  width: 78%;
  background-color: #181825;
  background-color: #181825;
  display: flex;
  justify-content: center;
  border-top-right-radius: 200px;
  border-bottom-right-radius: 200px;
  color: #fff;
}
.tomess {
  width: 28%;
  height: 100%;
  border-right: 2px solid #333;
}
.torequest {
  width: 36%;
  height: 100%;
}
.torelevant {
  width: 36%;
  height: 100%;
}
/* 接口列表 */
.toquonelist {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}
.tocri {
  display: flex;
  flex-direction: column;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 60%;
  padding: 200px 200px;
  margin-top: 200px;
  color: #fff;
}
.toctop {
  display: flex;
  justify-content: space-between;
}
.tocbot {
  display: flex;
  margin-top: 200px;
}
/* 请求 */
.requecri {
  width: 75.8%;
  margin: 200px 7%;
  border-radius: 100px;
  border: 2px solid #FEFFFF;
  padding: 3% 5%;
  max-height: 90vh;
  background-color: #46475E;
  color: #fff;
}
.requ {
  display: flex;
}
.requoe {
  display: flex;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 70%;
  padding: 100px 6%;
  margin-top: 200px;
}
.requtw {
  display: flex;
  justify-content: space-between;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 70%;
  padding: 100px 6%;
  margin-top: 200px;
}
.rehead {
  display: flex;
  justify-content: space-between;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 87%;
  min-height:20vh;
  padding: 100px 6%;
  margin-top: 200px;
}
.reheada{
  display: flex;
  justify-content: space-between;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 87%;
  min-height: 35vh;
  padding: 100px 6%;
  margin-top: 200px;
}
.reheadb{
  display: flex;
  justify-content: space-between;
  background-color: #8B76FF;
  border: 1px solid #7669D3;
  border-radius: 100px;
  width: 87%;
  min-height: 67vh;
  padding: 100px 6%;
  margin-top: 200px;
}
</style>
