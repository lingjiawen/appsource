<script setup lang="ts">
import tabbar from "@/components/tabbar/index.vue";
import NavBar from "@/components/nav-bar/index.vue";
import { useCachedViewStoreHook } from "@/store/modules/cached-view";
import { useDarkMode } from "@/composables/useToggleDarkMode";
import { computed, ref, watch } from "vue";
import { useRoute } from "vue-router";

const cachedViews = computed(() => {
  return useCachedViewStoreHook().cachedViewList;
});
const route = useRoute();
// const navbarTitle = ref(route.meta.title || "默认标题");
//
// // 监听路由变化，更新 navbarTitle
// watch(() => route.meta.title, (newTitle) => {
//   navbarTitle.value = newTitle || "默认标题";
// });
</script>

<template>
  <div class="app-wrapper">
    <van-config-provider :theme="useDarkMode() ? 'dark' : 'light'">
<!--      <nav-bar :title="navbarTitle" />-->
      <router-view v-slot="{ Component }">
        <keep-alive :include="cachedViews">
          <component :is="Component" />
        </keep-alive>
      </router-view>
      <tabbar />
    </van-config-provider>
  </div>
</template>

<style lang="less" scoped>
@import "@/styles/mixin.less";

.app-wrapper {
  .clearfix();
  position: relative;
  height: 100%;
  width: 100%;
}

// 子组件样式
:deep(.copyright) {
  position: absolute;
  bottom: var(--van-tabbar-height);
  left: 0;
  width: 100%;
  text-align: center;
  padding: 10px 0;
  font-size: 12px;
  color: #666;
  z-index: 999;
  * {
    white-space: nowrap;
  }
  div {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

:deep(.max-container) {
  max-width: 500px;
  margin: 0 auto;
}
</style>
