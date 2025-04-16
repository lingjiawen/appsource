<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import { useConfigStore } from "@/store/useConfigStore";
import {useRoute} from "vue-router";
const configStore = useConfigStore(); // 使用 store 实例


const helpList = ref([]);

const navbarTitle = ref("帮助中心");
const navbarSubtitle = ref("常见问题解答");

const currentOpenId = ref(1);

const toggleItem = item => {
  if (currentOpenId.value === item.id) {
    currentOpenId.value = null;
  } else {
    currentOpenId.value = item.id;
  }
};

// 获取配置
const getHelp = async () => {
  // 在 store 中获取数据
  helpList.value = configStore.getHelpList;
};

onMounted(() => {
  // 监听 loading 状态，直到数据加载完成
  watch(
      () => configStore.loading,
      loading => {
        if (!loading) {
          getHelp();
        }
      },
      {immediate: true} // 立即执行，监听 `loading` 状态
  );
});

watch(currentOpenId, (newId) => {
  helpList.value.forEach(item => {
    item.isExpand = item.id === newId;
  });
});
</script>

<template>
  <div class="help-container list-container">
    <nav-bar :title="navbarTitle" :subtitle="navbarSubtitle"/>
    <div class="help-content">
      <div class="help-list">
        <div
          v-for="item in helpList"
          :key="item.id"
          class="help-item"
          :class="{ active: item.id === currentOpenId }"
          @click="toggleItem(item)">
          <div class="help-header">
            <div class="help-title">
              <van-icon name="question-o" class="help-icon"/>
              <span>{{ item.title }}</span>
            </div>
            <van-icon
              :name="item.isExpand ? 'arrow-up' : 'arrow-down'"
              class="arrow-icon"
            />
          </div>
          <div
            class="help-detail"
            :class="{ 'is-open': item.isExpand }"
            v-html="item.content"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.help-container {
  padding: 0;
  background: var(--background-color);
  min-height: 100vh;
  padding-bottom: env(safe-area-inset-bottom);
}

.help-content {
  padding: 16px;
  padding-top: 12px;
}

.help-list {
  max-width: 800px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.help-item {
  background: var(--card-background);
  border-radius: var(--card-radius);
  box-shadow: var(--card-shadow);
  overflow: hidden;
  transition: all 0.3s ease;
  border: 1px solid var(--border-color);
}

.help-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  cursor: pointer;
  user-select: none;
  background-color: var(--van-background-2);
}

.help-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.help-icon {
  color: var(--primary-color);
  font-size: 20px;
  opacity: 0.9;
}

.arrow-icon {
  color: var(--text-secondary);
  transition: transform 0.3s ease;
  font-size: 16px;
}

.help-item.active .arrow-icon {
  transform: rotate(180deg);
}

.help-detail {
  padding: 0 16px;
  height: 0;
  opacity: 0;
  font-size: 14px;
  line-height: 1.6;
  color: var(--text-secondary);
  transition: all 0.3s ease;
  overflow: hidden;
  background-color: var(--van-background-2);
}

.help-detail.is-open {
  padding: 0 16px 16px;
  height: auto;
  opacity: 1;
}

/* 暗色模式优化 */
[data-theme='dark'] .help-item {
  background: var(--card-background);
  border-color: rgba(255, 255, 255, 0.08);
}

/* 响应式优化 */
@media (max-width: 480px) {
  .list-header {
    padding: 16px 16px 12px;
  }

  .list-title {
    font-size: 22px;
  }

  .list-subtitle {
    font-size: 13px;
  }

  .help-content {
    padding: 12px;
  }

  .help-list {
    gap: 8px;
  }

  .help-item {
    border-radius: 12px;
  }

  .help-header {
    padding: 14px 16px;
  }

  .help-title {
    font-size: 14px;
    gap: 8px;
  }

  .help-icon {
    font-size: 16px;
  }

  .help-detail {
    font-size: 13px;
    line-height: 1.5;
    padding: 0 16px;
  }

  .help-detail.is-open {
    padding: 0 16px 14px;
  }

  .arrow-icon {
    font-size: 14px;
  }
}

/* 增加中等尺寸屏幕的优化 */
@media (min-width: 481px) and (max-width: 768px) {
  .help-list {
    max-width: 600px;
  }

  .help-content {
    padding: 16px;
  }
}
</style>