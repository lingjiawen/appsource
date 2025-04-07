<template>
  <div class="app-list">
    <el-row :gutter="20">
      <el-col v-for="app in apps" :key="app.trackId" :xs="24" :sm="12" :md="8" :lg="8">
        <el-card shadow="hover" class="app-card">
          <img :src="getIconUrl(app)" :alt="app.trackName" class="app-icon" />
          <div class="app-name" :title="app.trackName">{{ app.trackName }}</div>
          <el-button type="primary" size="small" @click="selectIcon(getIconUrl(app))">选择图标</el-button>
        </el-card>
      </el-col>
    </el-row>
    <el-empty v-if="apps.length === 0" description="未找到相关应用" />
  </div>
</template>

<script setup>
import { ref, defineEmits } from 'vue'
import { ElMessage } from 'element-plus'

const apps = ref([])
const selectedIcon = ref(null)

const emit = defineEmits(['iconSelected']) // 定义事件

const displayResults = (appList) => {
  apps.value = appList

  if (appList.length === 0) {
    ElMessage.warning('未找到相关应用')
  }
}

const getIconUrl = (app) => {
  let iconUrl = app.artworkUrl512 || app.artworkUrl100
  return iconUrl.split('?')[0].replace(/\d+x\d+bb/, '512x512bb')
}

const selectIcon = (iconUrl) => {
  selectedIcon.value = iconUrl
  emit('iconSelected', iconUrl) // 触发事件通知父组件
  ElMessage.success(`已选择图标: ${iconUrl}`)
}

defineExpose({ displayResults }) // 仅暴露 `displayResults`
</script>


<style scoped>
.app-list {
  padding: 10px;
}

.app-card {
  text-align: center;
  border-radius: 12px;
  padding: 15px;
  margin-bottom: 20px; /* 让卡片之间有间距 */
}

.app-icon {
  width: 50px;
  height: 50px;
  border-radius: 22px;
  margin-bottom: 10px;
}

/* 处理过长的应用名称 */
.app-name {
  font-weight: bold;
  margin-bottom: 10px;
  font-size: 14px;
  max-width: 100%;
  white-space: nowrap;      /* 不换行 */
  overflow: hidden;         /* 超出部分隐藏 */
  text-overflow: ellipsis;  /* 超出显示省略号 */
  display: block;           /* 让其独占一行 */
}
</style>
