<template>
  <van-tabbar v-model="active" :placeholder="true" :route="true" fixed>
    <van-tabbar-item
      v-for="(item, index) in tabbarData"
      :key="index"
      :to="item.to"
    >
      <template #icon>
        <svg-icon v-if="item.useSvg" :name="item.svg as string ?? ''" />
        <van-icon v-if="!item.useSvg" :name="item.icon" />
      </template>
      {{ item.title }}
    </van-tabbar-item>
  </van-tabbar>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { useConfigStore } from "@/store/useConfigStore";

const configStore = useConfigStore(); // 使用 store 实例


const active = ref(0);
const tabbarData = reactive([
  {
    useSvg: true,
    svg: "appstore",
    icon: "wap-home-o",
    title: "签名安装",
    to: {
      name: "Sign"
    }
  },
  {
    useSvg: true,
    svg: "help",
    icon: "gem-o",
    title: "帮助中心",
    to: {
      name: "Help"
    }
  },
  {
    icon: "user-o",
    title: "关于我们",
    to: {
      name: "About"
    }
  }
]);

onMounted(async () => {
  configStore.setLoading(true);
  await configStore.fetchConfig(); // 在tabbar加载时获取config数据
  await configStore.fetchHelp(); // 在tabbar加载时获取config数据
  await configStore.fetchAbout(); // 在tabbar加载时获取config数据
  configStore.setLoading(false);
});
</script>
