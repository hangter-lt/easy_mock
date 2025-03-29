<template>
    <el-menu default-active="1" style="width: auto; border-right: 0;" :collapse="true">
        <RouterLink to="/">
            <el-menu-item index="1">
                <el-icon><icon-menu /></el-icon>
                <template #title>实时请求</template>
            </el-menu-item>
        </RouterLink>
        <RouterLink to="/apilist">
            <el-menu-item index="2">
                <el-icon>
                    <document />
                </el-icon>
                <template #title>接口管理</template>
            </el-menu-item>
        </RouterLink>
        <el-menu-item index="3">
            <el-icon>
                <setting />
            </el-icon>
            <template #title>历史请求</template>
        </el-menu-item>
    </el-menu>
</template>

<script lang="ts" setup>
import {
    Document,
    Menu as IconMenu,
    Setting,
} from '@element-plus/icons-vue'

</script>

<style></style>