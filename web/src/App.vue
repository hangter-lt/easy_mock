<script setup lang="ts">
import Menu from './components/Menu.vue';
</script>

<template>
  <div class="common-layout">
    <el-container style="height: 100vh;">
      <el-aside width="70px" style="border-right: 1px solid #eee;">
        <el-avatar shape="square" :size="60"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
        <div style="height: 50px;"></div>
        <Menu />
      </el-aside>
      <el-main>
        <RouterView />
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>

</style>
