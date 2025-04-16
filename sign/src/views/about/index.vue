<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {useConfigStore} from "@/store/useConfigStore";

const configStore = useConfigStore(); // 使用 store 实例
const goToUrl = url => {
  window.open(url);
};

const navbarTitle = ref("关于我们");
const navbarSubtitle = ref("平台介绍");

const aboutList = ref([[]]);

// 获取配置
const getAbout = async () => {
  // 在 store 中获取数据
  let abouts = configStore.getAboutList;
  // 处理数据
  abouts.forEach(item => {
    item.forEach(subItem => {
      subItem.isLink = (subItem.isLink == 1);
      subItem.url = subItem.url || "";
    });
  });
  aboutList.value = abouts;
};

onMounted(() => {
  // 监听 loading 状态，直到数据加载完成
  watch(
      () => configStore.loading,
      loading => {
        if (!loading) {
          getAbout();
        }
      },
      {immediate: true} // 立即执行，监听 `loading` 状态
  );
});
</script>

<template>
  <div class="about-container">
    <nav-bar :title="navbarTitle" :subtitle="navbarSubtitle"/>

    <div class="about-content">
      <van-cell-group v-for="(group, index) in aboutList" inset>
        <van-cell
            v-for="(item, itemIndex) in group"
            :key="itemIndex"
            :title="item.title"
            :icon="item.icon"
            :value="item.subtitle"
            :is-link="item.isLink"
            @click="goToUrl(item.url)"
        />
      </van-cell-group>
    </div>
  </div>
</template>

<style scoped>
.about-container {
  padding: 0;
  background: var(--background-color);
  min-height: 100vh;
}

.fixed-header {
  position: sticky;
  top: 0;
  z-index: 99;
  background: var(--nav-background);
  backdrop-filter: saturate(180%) blur(20px);
  -webkit-backdrop-filter: saturate(180%) blur(20px);
  border-bottom: 1px solid var(--border-color);
  padding-top: env(safe-area-inset-top);
}

.list-header {
  padding: 20px 16px 16px;
}

.list-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-primary);
  line-height: 1.3;
}

.list-subtitle {
  font-size: 14px;
  color: var(--text-secondary);
  margin-top: 4px;
  line-height: 1.4;
}

.about-content {
  max-width: 520px;
  margin: 0 auto;
  padding: 16px;
  padding-bottom: 100px;
}

:deep(.van-cell-group--inset) {
  margin: 0 0 12px;
  overflow: hidden;
  border-radius: 16px;
  background: var(--card-background);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

:deep(.van-cell) {
  padding: 16px;
  font-size: 15px;
  background: var(--card-background);
}

:deep(.van-cell__title) {
  color: var(--text-primary);
}

:deep(.van-cell__value) {
  color: var(--text-secondary);
}

:deep(.van-cell__left-icon) {
  font-size: 18px;
  color: var(--primary-color);
}

/* 暗色模式优化 */
[data-theme='dark'] :deep(.van-cell-group--inset) {
  background: var(--card-background);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
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

  .about-content {
    padding: 12px 12px 100px;
  }

  :deep(.van-cell) {
    padding: 14px;
    font-size: 14px;
  }
}
</style>