<template>
    <div class="menu">
        <RouterLink to="/apilist">
                <el-button text>返回</el-button>
            </RouterLink>
            <RouterLink :to="'/apiedit/' + form.id">
                <el-button type="primary">编辑</el-button>
            </RouterLink>
        </div>

        <el-form :model="form" label-width="auto">
            <el-card style="padding-top: 20px;">
                <el-row :gutter="20">
                    <el-col :span="16">
                        <el-form-item label="名称">
                            <el-input disabled v-model="form.name" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item label="分组">
                            <el-input disabled v-model="form.group" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="16">
                        <el-form-item label="方法">
                            <el-checkbox-group v-model="form.methods">
                                <el-checkbox-button disabled v-for="method in methods" :key="method" :value="method">
                                    {{ method }}
                                </el-checkbox-button>
                            </el-checkbox-group> 
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item label="请求类型">
                            <el-input disabled v-model="form.req_content_type" />
                        </el-form-item>
                    </el-col>
                </el-row>
            <el-form-item label="路径">
                <el-input disabled v-model="form.path" placeholder="输入api地址" />
            </el-form-item>

            <el-form-item label="描述">
                <el-input disabled v-model="form.description" type="textarea" />
            </el-form-item>
        </el-card>
        <el-card v-for="item in form.params">
            <template #header>
                <div class="card-header">
                    <span>请求&响应</span>
                </div>
            </template>
            <el-row :gutter="20">
                <el-col :span="10">
                    <el-form-item label="请求参数">
                        <el-row :gutter="10" v-for="reqData in item.req_datas" :key="reqData.key">
                            <!-- gutter 用于设置列之间的间距 -->
                            <el-col :span="8">
                                <el-input disabled v-model="reqData.key" placeholder="key" />
                            </el-col>
                            <el-col :span="14">
                                <el-input disabled v-model="reqData.value" placeholder="value" />
                            </el-col>
                        </el-row>
                    </el-form-item>
                </el-col>
                <el-col :span="14">
                    <el-form-item label="响应码">
                        <el-input disabled v-model="item.res_code" />
                    </el-form-item>
                    <el-form-item label="响应类型">
                        <el-input disabled v-model="item.res_content_type"></el-input>
                    </el-form-item>
                    <el-form-item label="响应内容">
                        <el-input disabled v-model="item.res_data" type="textarea" />
                    </el-form-item>
                </el-col>
            </el-row>
        </el-card>
    </el-form>
</template>

<script setup lang="ts">

import axios from 'axios'
import { reactive } from 'vue'
import { useRoute } from 'vue-router'
import type { Form } from './types'

const route = useRoute()

const methods = ["GET", "POST", "DELETE", "PUT", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"]

// 初始化 form
const form = reactive<Form>({
    id: '',
    name: '',
    group: '',
    methods: [],
    req_content_type: '',
    path: '',
    description: '',
    params: [
        {
            route: [],
            req_datas: [{
                key: "",
                value: "",
            }],
            res_code: 0,
            res_content_type: '',
            res_data: '',
        },
    ]
});

axios.get('/manages/' + route.params.id).then(res => {
    Object.assign(form, res.data.data)
})

</script>

<style scoped>
.menu {
    width: 100%;
    text-align: right;
    margin-bottom: 20px;
}

.el-form {
    margin: 0 100px;
}

.add-param-btn {
    width: 100%;
}

.el-card {
    margin: 10px 0;
}

:deep(.el-input.is-disabled .el-input__wrapper) {
    background-color: #ffffff;

}

:deep(.el-input.is-disabled .el-input__inner) {
    -webkit-text-fill-color: #606266;
}

:deep(.el-textarea.is-disabled .el-textarea__inner) {
    background-color: #ffffff;
}


:deep(.el-checkbox-button.is-disabled.is-checked .el-checkbox-button__inner) {
    background-color: var(--el-color-primary);
    border-color: var(--el-color-primary);
    color: #ffffff;
}

</style>