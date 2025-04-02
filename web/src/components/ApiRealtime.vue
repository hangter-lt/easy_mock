<template>
    <el-row style="height: 100vh; overflow: hidden; display: flex;">
        <el-col :span="4" class="scroll-container">
            <el-timeline style="max-width: 600px">
                <el-timeline-item v-for="(activity, index) in activities" :key="index" :icon="activity.icon"
                    :type="activity.type" :color="activity.color" :size="activity.size" :hollow="activity.hollow"
                    :timestamp="activity.timestamp">
                    {{ activity.content }}
                </el-timeline-item>
            </el-timeline>
        </el-col>

        <el-col :span="20" class="scroll-container">
            <el-card>
                <el-card></el-card>
                <el-card></el-card>

                <el-card></el-card>
                <el-card></el-card>
            </el-card>
        </el-col>
    </el-row>
</template>



<script lang="ts" setup>
import { MoreFilled } from '@element-plus/icons-vue'
import type { TimelineItemProps } from 'element-plus'
import { onMounted, onUnmounted, ref, reactive } from 'vue'

type realtime = {
    id: string
    req_method: string
    req_path: string
    req_time: number
}

interface ActivityType extends Partial<TimelineItemProps> {
    content: string
}

const activities = reactive<ActivityType[]>([])

const eventSource = ref<EventSource | null>();

const initSSE = () => {
    eventSource.value = new EventSource("/request/realtime");

    eventSource.value.onmessage = (event) => {
        console.log("收到消息内容是:", event);
        const data = JSON.parse(event.data) as realtime;
        console.log(data);
        activities.unshift({
            content: data.req_method + " " + data.req_path + " ",
            timestamp: new Date().toLocaleTimeString(),
            size: 'large',
            type: 'primary',
            icon: MoreFilled,
        });
        console.log(activities);
    };

    eventSource.value.onerror = (error) => {
        console.error("SSE 连接出错：", error);
        eventSource.value!.close();
    };
};

onMounted(() => {
    initSSE();
});

onUnmounted(() => {
    if (eventSource) {
        console.log("关闭 SSE 连接");
        eventSource.value?.close();
    }
});
</script>

<style scoped>
.scroll-container {
    height: 100vh;
    overflow-y: auto;
    padding: 16px;
    box-sizing: border-box;

    /* Hide scrollbar for Chrome/Safari/Edge */
    &::-webkit-scrollbar {
        display: none;
    }

    /* Hide scrollbar for Firefox */
    scrollbar-width: none;
}
</style>