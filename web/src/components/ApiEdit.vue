<template>

    <div class="menu">
        <el-button @click="btnCancel" text>返回</el-button>
        <el-button type="primary" v-if="route.params.id === 'create'" @click="btnCreate">创建</el-button>
        <el-button type="primary" v-else @click="btnUpdate">更新</el-button>
    </div>

    <el-form ref="formRef" :model="form" label-width="auto" :rules="rules">
        <el-card style="padding-top: 20px;">
            <el-row :gutter="20">
                <el-col :span="16">
                    <el-form-item label="名称" prop="name">
                        <el-input v-model="form.name" />
                    </el-form-item>
                </el-col>
                <el-col :span="8">
                    <el-form-item label="分组" prop="group">
                        <el-select v-model="form.group" filterable allow-create default-first-option placeholder="选择分组">
                            <el-option v-for="item in groups" :key="item" :label="item" :value="item" />
                        </el-select>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row :gutter="20">
                <el-col :span="16">
                    <el-form-item label="方法" prop="methods">
                        <el-checkbox-group v-model="form.methods">x
                            <el-checkbox-button v-for="method in methods" :key="method" :value="method">
                                {{ method }}
                            </el-checkbox-button>
                        </el-checkbox-group>
                    </el-form-item>
                </el-col>
                <el-col :span="8">
                    <el-form-item label="请求类型" prop="req_content_type">
                        <el-select v-model="form.req_content_type" placeholder="选择请求类型">
                            <el-option v-for="item in content_types" :key="item" :label="item" :value="item" />
                        </el-select>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-form-item label="路径" prop="path">
                <el-input v-model="form.path" placeholder="输入api地址">
                    <template #suffix>
                        <el-button @click="btnAddRoute" type="text">填充路由参数</el-button>
                    </template>
                </el-input>
            </el-form-item>

            <el-form-item label="描述">
                <el-input v-model="form.description" type="textarea" />
            </el-form-item>
        </el-card>
        <el-card v-for="(item, itemIndex) in form.params">
            <template #header>
                <div class="card-header">
                    <span>请求&响应</span>
                    <el-icon @click="removeParam(itemIndex)" style="float: right;">
                        <Delete />
                    </el-icon>
                </div>
            </template>
            <el-row :gutter="20">
                <el-col :span="10">
                    <el-form-item label="路由参数" v-if="item.route.length > 0">
                        <!-- 根据item.route动态生成input -->
                        <div v-for="(_, routeIndex) in item.route">
                            <el-form-item :prop="`params.${itemIndex}.route.${routeIndex}`"
                                :rules="[{ required: true, message: '请输入路由参数', trigger: 'blur' }]">
                                <el-input v-model="item.route[routeIndex]" :placeholder="'route' + routeIndex" />
                            </el-form-item>
                        </div>
                    </el-form-item>

                    <el-form-item label="请求参数">
                        <el-row :gutter="10" v-for="(reqData, reqDataIndex) in item.req_datas" :key="reqData.key">
                            <!-- gutter 用于设置列之间的间距 -->
                            <el-col :span="8">
                                <el-form-item :prop="`params.${itemIndex}.req_datas.${reqDataIndex}.key`"
                                    :rules="[{ required: true, message: '请输入key', trigger: 'blur' }]">
                                    <el-input v-model="reqData.key" placeholder="key" />
                                </el-form-item>
                            </el-col>
                            <el-col :span="14">
                                <el-form-item :prop="`params.${itemIndex}.req_datas.${reqDataIndex}.key`"
                                    :rules="[{ required: true, message: '请输入value', trigger: 'blur' }]">
                                    <el-input v-model="reqData.value" placeholder="value" />
                                </el-form-item>
                            </el-col>
                            <el-col :span="1">
                                <el-icon @click="addReqData(itemIndex, reqDataIndex)">
                                    <Plus />
                                </el-icon>
                            </el-col>
                            <el-col :span="1">
                                <el-icon @click="removeReqData(itemIndex, reqDataIndex)">
                                    <Delete />
                                </el-icon>
                            </el-col>
                        </el-row>
                        <el-col v-if="item.req_datas.length === 0">
                            <el-icon @click="addReqData(itemIndex, 0)">
                                <Plus />
                            </el-icon>
                        </el-col>
                    </el-form-item>
                </el-col>
                <el-col :span="14">
                    <el-form-item label="响应码" :prop="`params.${itemIndex}.res_code`"
                        :rules="[{ required: true, message: '响应码不能为空', trigger: 'blur' }]">
                        <el-input v-model="item.res_code" type="number" />
                    </el-form-item>
                    <el-form-item label="响应类型" :prop="`params.${itemIndex}.res_content_type`"
                        :rules="[{ required: true, message: '请选择响应类型', trigger: 'change' }]">
                        <el-select v-model="item.res_content_type" placeholder="选择响应类型">
                            <el-option v-for="item in content_types" :key="item" :label="item" :value="item" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="响应内容" :prop="`params.${itemIndex}.res_data`"
                        :rules="[{ required: true, message: '响应内容不能为空', trigger: 'blur' }]">
                        <el-input v-model="item.res_data" type="textarea" />
                    </el-form-item>
                </el-col>
            </el-row>
        </el-card>
        <el-button class="add-param-btn" @click="addParam">+</el-button>
    </el-form>

</template>

<script lang="ts" setup>
import axios from 'axios'
import { reactive, ref, watchEffect } from 'vue'
import { Plus, Delete } from '@element-plus/icons-vue';
import type { Form } from './types';
import { useRoute, useRouter } from 'vue-router'
import { ElForm, ElMessage } from 'element-plus';
import type { FormRules } from 'element-plus'


const route = useRoute()
const router = useRouter()
let routeLength = 0

const rules = reactive<FormRules<Form>>({
    name: [{ required: true, message: '请输入接口名称', trigger: 'blur' }],
    group: [{ required: true, message: '请选择分组或输入新分组', trigger: 'blur' }],
    methods: [{ required: true, message: '请选择请求方法', trigger: 'change' }],
    req_content_type: [{ required: true, message: '请选择请求类型', trigger: 'blur' }],
    path: [{ required: true, message: '请输入接口路径', trigger: 'blur' }],
})


const formRef = ref<InstanceType<typeof ElForm> | null>(null);

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
            req_datas: [],
            res_code: 200,
            res_content_type: '',
            res_data: '',
        },
    ]
});

const addParam = () => {
    // 根据routeLength增加route内元素数量
    let route = []
    for (let i = 0; i < routeLength; i++) {
        route.push("")
    }
    form.params.push({
        route: route,
        req_datas: [{
            key: "",
            value: "",
        }],
        res_code: 200,
        res_content_type: '',
        res_data: '',
    })
}

const removeParam = (itemIndex: number) => {
    if (form.params.length === 1) {
        ElMessage({
            message: "留一个",
            type: 'warning',
        })
        return
    }
    form.params.splice(itemIndex, 1)
}

const addReqData = (itemIndex: number, reqDataIndex: number) => {
    // 实现每次向下插入一行
    form.params[itemIndex].req_datas.splice(reqDataIndex + 1, 0, {
        key: "",
        value: "",
    })
}

const removeReqData = (itemIndex: number, reqDataIndex: number) => {
    form.params[itemIndex].req_datas.splice(reqDataIndex, 1)
}

const groups = ref([])

const methods = ["GET", "POST", "DELETE", "PUT", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"]

const content_types = [
    'application/json',
    'multipart/form-data',
    'application/x-www-form-urlencoded',
    'application/xml',
]

const btnCreate = () => {
    formRef.value?.validate((valid) => {
        if (valid) {
            axios.post('/manages/create', form, {
                headers: {
                    'Content-Type': 'application/json'
                }
            })
                .then(function (response) {

                    if (response.data.success) {
                        ElMessage({
                            message: "成功",
                            type: 'success',
                        })
                        router.push("/apiinfo/" + response.data.data.id)
                    } else {
                        ElMessage({
                            message: response.data.message,
                            type: 'error',
                        })
                    }
                })
                .catch(function (error) {
                    if (error.response) {
                        ElMessage({
                            message: error.response.status + " " + error.response.data.message,
                            type: 'error',
                        })
                    } else {
                        ElMessage({
                            message: error.message,
                            type: 'error',
                        })
                    }

                })
        } else {
            ElMessage({
                message: "请检查表单",
                type: 'error',
            })
        }
    })
}

const btnUpdate = () => {

    formRef.value?.validate((valid) => {
        if (valid) {
            axios.post('/manages/update', form, {
                headers: {
                    'Content-Type': 'application/json'
                }
            })
                .then(function (response) {

                    if (response.data.success) {
                        ElMessage({
                            message: "成功",
                            type: 'success',
                        })
                        router.push("/apiinfo/" + form.id)
                    } else {
                        ElMessage({
                            message: response.data.message,
                            type: 'error',
                        })
                    }
                })
                .catch(function (error) {
                    if (error.response) {
                        ElMessage({
                            message: error.response.status + " " + error.response.data.message,
                            type: 'error',
                        })
                    } else {
                        ElMessage({
                            message: error.message,
                            type: 'error',
                        })
                    }

                })
        } else {
            ElMessage({
                message: '请检查表单',
                type: 'error',
            })
        }
    })
}
if (route.params.id != "create") {
    axios.get('/manages/' + route.params.id).then(res => {
        Object.assign(form, res.data.data)
    })
}

axios.get('/manages/groups').then(res => {
    console.log(res.data.data)
    groups.value = res.data.data
    // Object.assign(groups, res.data.data)
})

const btnCancel = () => {
    router.go(-1)
}

const btnAddRoute = () => {
    form.path += "/%s"
}

// 监听form.path变化，计算routeLength, 处理route数量
watchEffect(() => {
    let oldLength = routeLength
    routeLength = form.path.split("/%s").length - 1

    if (oldLength < routeLength) {
        for (let index = 0; index < routeLength - oldLength; index++) {
            form.params.forEach(item => {
                item.route.push("")
            })
        }
    } else if (oldLength > routeLength) {
        for (let index = 0; index < oldLength - routeLength; index++) {
            form.params.forEach(item => {
                item.route.pop()
            })
        }
    }
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
</style>