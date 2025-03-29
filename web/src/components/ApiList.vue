<template>
    <div class="addmenu">
        <RouterLink to="/apiedit/create">
            <el-button type="primary">创建</el-button>
        </RouterLink>
    </div>

    <div>
        <el-collapse v-model="activeNames">
            <el-collapse-item v-for="(group, groupIndex) in res.data" :title="groupIndex" :name="groupIndex">
                <div class="card-container">
                    <RouterLink :to="'/apiinfo/'+item.id" v-for="(item, itemIndex) in group" :key="itemIndex" >
                        <el-card class="card-item">
                            <el-text>{{ item.name }}</el-text>
                            <br />
                            <el-tag v-for="method in item.methods" type="primary">{{ method }}</el-tag>
                            <br />
                            <el-text>{{ item.path }}</el-text>
                            <br />
                            <el-text>{{ item.req_content_type }}</el-text>
                        </el-card>
                    </RouterLink>
                </div>
            </el-collapse-item>
        </el-collapse>
    </div>
</template>

<style>
.addmenu {
    width: 100%;
    text-align: right;
    margin-bottom: 20px;
}

.card-container {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 16px;
    width: 100%;
    overflow: visible;
    /* 确保边框不会被裁剪 */
}

.card-item {
    width: 100%;
    box-sizing: border-box;
    /* 确保边框包含在宽度内 */
    border: 1px solid #ccc;
    /* 示例边框样式 */
}

/* 媒体查询调整卡片宽度 */
@media (max-width: 768px) {
    .card-container {
        grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    }
}

@media (max-width: 480px) {
    .card-container {
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    }
}
</style>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { RouterLink } from 'vue-router';

interface ApiGroup {
    id: string;
    name: string;
    group: string;
    methods: string[];
    path: string;
    req_content_type: string;
}

interface ApiResponse {
    data: {
        [key: string]: ApiGroup[];
    };
    success: boolean;
}

const res = ref<ApiResponse>({
    data: {},
    success: false,
});



axios.get('/manages')
    .then(function (response) {
        res.value = response.data;
        activeNames.value = Object.keys(res.value.data);
    })
    .catch(function (error) {
        console.log(error);
    });



const activeNames = ref([""])
</script>