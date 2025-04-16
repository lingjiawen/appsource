<script setup lang="ts">
import {
  useDarkMode,
  useToggleDarkMode
} from "@/composables/useToggleDarkMode";
import { ref, onMounted } from 'vue'

// 接受传递的 title 和 subtitle props
defineProps({
  title: {
    type: String,
    required: true,
  },
  subtitle: {
    type: String,
    required: true,
  },
});


const onClickRight = (event: TouchEvent | MouseEvent) => {
  useToggleDarkMode(event);
};
</script>


<template>
  <div class="fixed-header" @click-right="onClickRight">
    <div class="list-header">
      <div class="list-title">{{ title }}</div> <!-- 显示传递的 title -->
      <div class="list-subtitle">{{ subtitle }}</div> <!-- 显示传递的 subtitle -->
    </div>
  </div>
  <div class="theme-toggle" @click="onClickRight">
    <svg-icon class="text-[18px]" :name="useDarkMode() ? 'light' : 'dark'" />
  </div>
</template>


<style scoped>
.fixed-header {
  position: sticky;
  top: 0;
  z-index: 99;
  height: var(--header-height);
  background: var(--van-background-2);
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

.theme-toggle {
  position: fixed;
  top: calc((var(--header-height) - 36px) / 2); /* 动态定位 */
  //top: calc(env(safe-area-inset-top, 16px) + 16px);
  right: 16px;
  z-index: 1000;
  width: 36px;
  height: 36px;
  border-radius: 18px;
  //background-color: var(--van-background-2);
  backdrop-filter: saturate(180%) blur(25px);
  -webkit-backdrop-filter: saturate(180%) blur(25px);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 2px 8px var(--van-background);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  -webkit-tap-highlight-color: transparent;
}

.theme-toggle:active {
  transform: scale(0.92);
}
</style>
