import Layout from "@/layout/index.vue";
import type { RouteRecordRaw } from "vue-router";
import Home from "@/views/home/index.vue";
import Help from "@/views/help/index.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "root",
    component: Layout,
    redirect: { name: "Sign" },
    children: [
      {
        path: "sign",
        name: "Sign",
        component: Home,
        meta: {
          title: "主页"
        }
      },
      {
        path: "help",
        name: "Help",
        component: () => import("@/views/help/index.vue"),
        meta: {
          title: "工具"
        }
      },
      {
        path: "about",
        name: "About",
        component: () => import("@/views/about/index.vue"),
        meta: {
          title: "关于",
          noCache: true
        }
      }
    ]
  }
];

export default routes;
