<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import GridPatternDashed from "@/components/grid-pattern/grid-pattern-dashed.vue";
import Fa6SolidHeart from "@iconify-icons/fa6-solid/heart";
import { useConfigStore } from "@/store/useConfigStore";
import { useRoute } from "vue-router"; // 导入 store
import { deviceCheckApi, installApi, installPrivateApi } from "@/api/mock";
import {showFailToast, showToast, Toast} from "vant";
import ClipboardJS from "clipboard"; // 引入 ClipboardJS

const configStore = useConfigStore(); // 使用 store 实例

defineOptions({
  name: "Sign"
});

// 弹出框
const countdown = ref(3);
const showBottom = ref(false);

const startCountdown = () => {
  const timer = setInterval(() => {
    countdown.value -= 1;
    if (countdown.value === 0) {
      clearInterval(timer);
      showBottom.value = true;
    }
  }, 1000);
};

const showPopup = ref(false);

const handleNotShowIn24h = () => {
  localStorage.setItem(
    "hidePopupUntil",
    (Date.now() + 24 * 60 * 60 * 1000).toString()
  );
  showPopup.value = false;
};

// 检查是否需要显示弹窗
const checkPopupVisibility = () => {
  const hideUntil = localStorage.getItem("hidePopupUntil");
  showPopup.value = !(hideUntil && Date.now() < parseInt(hideUntil));
};

const udid = ref("");
const code = ref("");

const requestInstall = async () => {
  try {
    loadingText.value = "正在请求安装...";
    isLoading.value = true;
    const installData = await installApi({
      udid: udid.value,
      code: code.value
    }); // 请求安装
    isLoading.value = false;
    if (installData.url) {
      installURL.value = installData.url; // 获取安装链接
      showInstallPopup.value = true; // 显示安装提示弹窗
    } else {
      showFailToast("服务器返回错误");
    }
  } catch (error) {
    isLoading.value = false;
    showFailToast(error.message || error);
  }
};

const copyInstallURL = () => {
  if (installURL.value) {
    copyText(installURL.value); // 复制安装链接
  }
};

const openInstallURL = () => {
  if (installURL.value) {
    window.open(installURL.value);
  }
};

const getUDID = () => {
  const apiBaseUrl = import.meta.env.VITE_BASE_API.startsWith("http")
    ? import.meta.env.VITE_BASE_API
    : `${window.location.protocol}//${window.location.host}${import.meta.env.VITE_BASE_API}`;
  location.href = `${apiBaseUrl}/getudid`; // 拼接最终 URL
};

const copyText = text => {
  let tempButton = document.createElement("button");
  tempButton.style.position = "absolute";
  tempButton.style.opacity = "0";
  tempButton.style.pointerEvents = "none";
  tempButton.textContent = text;
  document.body.appendChild(tempButton);

  let clipboard = new ClipboardJS(tempButton, {
    text: function () {
      return text;
    }
  });

  clipboard.on("success", function (e) {
    e.clearSelection();
    showToast("🔗安装链接已复制到剪贴板");
    clipboard.destroy();
    tempButton.remove();
  });

  clipboard.on("error", function (e) {
    showFailToast("复制失败，请手动复制");
    clipboard.destroy();
    tempButton.remove();
  });
  tempButton.click();
};

const buyCode = () => {
  if (buyUrl.value) {
    window.open(buyUrl.value);
  }
};

const onConsumeSubmit = () => {
  requestInstall();
};

const base64p12 = ref("");
const base64mp = ref("");
const privatePassword = ref("1");

// 上传文件并获取Base64
// 上传文件并获取Base64（添加文件类型和大小限制）
function uploadFile(type: "p12" | "mp") {
  const input = document.createElement("input");
  input.type = "file";

  // 设置文件类型限制
  input.accept =
    type === "p12"
      ? ".p12, application/x-pkcs12"
      : ".mobileprovision, application/x-apple-aspen-config";

  input.addEventListener("change", e => {
    const target = e.target as HTMLInputElement;
    const file = target.files?.[0];

    if (!file) return;

    // 文件类型验证
    const isValidType =
      type === "p12"
        ? file.name.toLowerCase().endsWith(".p12")
        : file.name.toLowerCase().endsWith(".mobileprovision");

    if (!isValidType) {
      Toast.fail(
        `请选择${type === "p12" ? ".p12" : ".mobileprovision"} 格式文件`
      );
      input.value = ""; // 清空选择
      return;
    }

    // 文件大小验证（5MB）
    const MAX_SIZE = 5 * 1024 * 1024;
    if (file.size > MAX_SIZE) {
      Toast.fail("文件大小不能超过5MB");
      input.value = "";
      return;
    }

    // 读取文件
    const reader = new FileReader();
    reader.onload = event => {
      const result = event.target?.result;
      if (typeof result === "string") {
        const base64 = result.split(",")[1];
        if (type === "p12") {
          base64p12.value = base64;
        } else {
          base64mp.value = base64;
        }
      }
    };
    reader.readAsDataURL(file);
  });

  input.click();
}

// 上传按钮处理函数
const uploadP12 = () => {
  uploadFile("p12"); // 限制.p12文件
};

const uploadMP = () => {
  uploadFile("mp"); // 限制.mobileprovision文件
};

const requestPrivateInstall = async () => {
  try {
    loadingText.value = "正在请求安装...";
    isLoading.value = true;
    const installPrivateData = await installPrivateApi({
      base64p12: base64p12.value,
      base64mp: base64mp.value,
      privatePassword: privatePassword.value
    }); // 请求安装
    isLoading.value = false;
    const url = installPrivateData.url;
    window.location.href = url;
  } catch (error) {
    isLoading.value = false;
    showFailToast(error.message || error);
  }
};

const onPrivateSubmit = () => {
  // 处理私有上传的提交逻辑
  // 例如，发送base64p12和base64mp到服务器
  requestPrivateInstall();
};

const searchValue = ref("");
const searchType = ref("udid");
const showDevicesPopup = ref(false);
const showInstallPopup = ref(false);
const installURL = ref("");
const searchDevices = ref([]);

const requestDeviceCheck = async () => {
  try {
    loadingText.value = "正在查询...";
    isLoading.value = true;
    const devicesData = await deviceCheckApi({
      searchType: searchType.value,
      searchValue: searchValue.value
    }); // 请求安装
    isLoading.value = false;
    const devices = devicesData.devices;
    showDevicesPopup.value = true;
    searchDevices.value = devices;
  } catch (error) {
    isLoading.value = false;
    showFailToast(error.message || error);
  }
};

const onSearchSubmit = () => {
  requestDeviceCheck();
};

const isLoading = ref(false);
const loadingText = ref("正在加载...");

const navbarTitle = ref("");
const navbarSubtitle = ref("一款IOS端免费的IPA签名工具");
const tabActive = ref(1);
const contactName = ref("");
const contactUrl = ref("");
const notice = ref("");
const popupHtml = ref("");
const buyUrl = ref("");

const freeOpen = ref(false);

const icp = ref("");
const vats = ref("");
const security = ref("");

// 获取配置
const getConfig = async () => {
  // 在 store 中获取数据
  navbarTitle.value = configStore.getNavbarTitle;
  contactName.value = configStore.getContactName;
  contactUrl.value = configStore.getContactUrl;
  notice.value = configStore.getNotice;
  popupHtml.value = configStore.getPopupHtml;
  buyUrl.value = configStore.getBuyUrl;

  freeOpen.value = configStore.isFreeOpen;
  icp.value = configStore.getIcp;
  vats.value = configStore.getVats;
  security.value = configStore.getSecurity;

  tabActive.value = 0;
};

onMounted(() => {
  // 监听 loading 状态，直到数据加载完成
  watch(
    () => configStore.loading,
    loading => {
      if (!loading) {
        getConfig();
      }
    },
    { immediate: true } // 立即执行，监听 `loading` 状态
  );

  const route = useRoute();
  const queryUdid = route.query.udid as string;
  if (queryUdid) {
    udid.value = queryUdid; // 填充到 udid 字段
  }

  checkPopupVisibility();
});
</script>

<template>
  <nav-bar :title="navbarTitle" :subtitle="navbarSubtitle" />
  <GridPatternDashed />
  <van-overlay :show="isLoading" class="flex justify-center items-center">
    <div class="wrapper">
      <van-loading vertical>
        <template #icon>
          <van-icon name="star-o" size="30" />
        </template>
        {{ loadingText }}
      </van-loading>
    </div>
  </van-overlay>
  <div class="demo-content px-[12px]">
    <div class="flex-container">
      <div class="max-container">
        <img
          class="block w-[80px] mx-auto mb-[10px] pt-[40px] dark:hidden"
          alt="Logo"
          src="~@/assets/logo.png"
        />
        <img
          class="block w-[80px] mx-auto mb-[10px] pt-[40px] hidden dark:block"
          alt="Logo"
          src="~@/assets/logo-dark.png"
        />
        <div
          class="text-[14px] py-[12px] px-[20px] rounded-[12px] bg-[var(--color-block-background)] mt-[14px]"
        >
          <div>
            <a class="flex items-center" :href="contactUrl" target="_blank">
              <i-icon
                :icon="Fa6SolidHeart"
                class="inline-block text-[20px] mr-3"
              />
              <h3 class="font-bold text-[18px] my-[4px]">{{ contactName }}</h3>
              <svg-icon class="text-[12px] ml-[5px]" name="link" />
            </a>
          </div>
          <p class="leading-[24px] my-[6px]">
            {{ notice }}
          </p>
        </div>

        <div class="mt-[14px] text-[24px]">
          <van-tabs
            v-model:active="tabActive"
            sticky
            class="bg-[var(--van-background, #1c1c1e)] rounded-[12px] overflow-hidden"
          >
            <van-tab title="兑换安装">
              <div class="tabs-content">
                <van-form @submit="onConsumeSubmit">
                  <van-cell-group class="rounded-b-[12px] overflow-hidden">
                    <van-field
                      v-model="udid"
                      label-width="3.5rem"
                      name="UDID"
                      label="UDID"
                      placeholder="UDID"
                      center
                      clearable
                    >
                      <template #button>
                        <van-button size="small" type="primary" @click="getUDID"
                          >获取
                        </van-button>
                      </template>
                    </van-field>
                    <van-field
                      v-model="code"
                      label-width="3.5rem"
                      name="兑换码"
                      label="兑换码"
                      center
                      placeholder="兑换码"
                    >
                      <template #button>
                        <van-button size="small" type="primary" @click="buyCode"
                          >购买
                        </van-button>
                      </template>
                    </van-field>
                  </van-cell-group>
                  <div style="margin: 16px">
                    <van-button round block type="primary" native-type="submit">
                      兑换安装
                    </van-button>
                  </div>
                  <div v-if="freeOpen" style="margin: 16px">
                    <van-button round block type="primary" native-type="submit">
                      安装体验版
                    </van-button>
                  </div>
                </van-form>
              </div>
            </van-tab>

            <van-tab title="自助上传">
              <div class="tabs-content">
                <van-form @submit="onPrivateSubmit">
                  <van-cell-group class="rounded-b-[12px] overflow-hidden">
                    <van-field
                      v-model="base64p12"
                      type="textarea"
                      name="base64p12"
                      label="证书文件"
                      label-align="top"
                      clearable
                      :autosize="{ maxHeight: 100, minHeight: 50 }"
                      placeholder="输入base64编码的证书文件"
                    >
                      <template #button>
                        <van-button
                          size="small"
                          type="primary"
                          @click="uploadP12"
                          >上传
                        </van-button>
                      </template>
                    </van-field>
                    <van-field
                      v-model="base64mp"
                      type="textarea"
                      name="base64mp"
                      label="描述文件"
                      label-align="top"
                      clearable
                      :autosize="{ maxHeight: 100, minHeight: 50 }"
                      placeholder="输入base64编码的描述文件"
                    >
                      <template #button>
                        <van-button
                          size="small"
                          type="primary"
                          @click="uploadMP"
                          >上传
                        </van-button>
                      </template>
                    </van-field>
                    <van-field
                      v-model="privatePassword"
                      name="证书密码"
                      label="证书密码"
                      label-align="top"
                      placeholder="输入证书密码"
                    />
                  </van-cell-group>
                  <div style="margin: 16px">
                    <van-button round block type="primary" native-type="submit">
                      安装
                    </van-button>
                  </div>
                </van-form>
              </div>
            </van-tab>

            <van-tab title="证书查询">
              <div class="tabs-content">
                <van-form @submit="onSearchSubmit">
                  <van-cell-group class="rounded-b-[12px] overflow-hidden">
                    <van-field
                      name="radio"
                      label="查询方式"
                      label-width="3.6rem"
                    >
                      <template #input>
                        <van-radio-group
                          v-model="searchType"
                          direction="horizontal"
                        >
                          <van-radio name="udid">UDID</van-radio>
                          <van-radio name="code">兑换码</van-radio>
                          <van-radio name="cert_id">证书编号</van-radio>
                        </van-radio-group>
                      </template>
                    </van-field>

                    <van-field
                      v-model="searchValue"
                      label-width="3.6rem"
                      name="searchValue"
                      label="搜索值"
                      placeholder="输入UDID/兑换码/证书编号"
                      clearable
                    />
                  </van-cell-group>
                  <div style="margin: 16px">
                    <van-button round block type="primary" native-type="submit">
                      查询
                    </van-button>
                  </div>
                </van-form>
              </div>
            </van-tab>
          </van-tabs>
        </div>
      </div>
      <div class="copyright">
        <div v-if="icp">
          <a
            href="https://beian.miit.gov.cn"
            target="_blank"
            style="text-decoration: none"
            >{{ icp }}</a
          >
        </div>
        <div v-if="vats">
          <span
            >增值电信业务经营许可证:
            <a
              href="https://beian.miit.gov.cn"
              target="_blank"
              style="text-decoration: none"
              >{{ vats }}</a
            ></span
          >
        </div>
        <div v-if="security">
          <img
            src="~@/assets/ga.png"
            alt="公网安备"
            style="width: 12px; height: 12px"
          /><a
            href="http://www.beian.gov.cn"
            target="_blank"
            style="text-decoration: none"
            >{{ security }}</a
          >
        </div>
      </div>
    </div>
  </div>
  <!-- 圆角弹窗（底部） -->
  <van-popup
    :show="showPopup"
    round
    position="bottom"
    class="notice-popup"
    @opened="startCountdown()"
  >
    <div class="notice-popup-wrapper">
      <div class="notice-popup-header">
        <div class="notice-popup-title">
          <van-icon name="info" class="notice-icon" />
          <span>使用须知</span>
        </div>
        <div v-show="!showBottom" class="notice-countdown">
          请仔细阅读 ({{ countdown }}s)
        </div>
      </div>
      <div class="notice-popup-content">
        <div class="notice-popup-text" v-html="popupHtml" />
      </div>
      <div class="notice-popup-footer">
        <div class="notice-buttons">
          <van-button
            round
            block
            :disabled="!showBottom"
            class="close-button"
            @click="showPopup = false"
          >
            <van-icon name="checked" class="button-icon" />
            <span>我知道了</span>
          </van-button>
          <van-button
            type="primary"
            :disabled="!showBottom"
            round
            block
            @click="handleNotShowIn24h"
          >
            <i
              class="van-badge__wrapper van-icon van-icon-clock-o button-icon"
            />
            <span>24小时内不再提示</span>
          </van-button>
        </div>
      </div>
    </div>
  </van-popup>

  <van-popup
    :show="showDevicesPopup"
    round
    position="bottom"
    class="notice-popup"
  >
    <div class="notice-popup-wrapper">
      <div class="notice-popup-header">
        <div class="notice-popup-title">
          <van-icon name="info" class="notice-icon" />
          <span>查询结果</span>
        </div>
      </div>
      <div class="notice-popup-content">
        <div class="notice-popup-text">
          <!--  如果为空显示没有设备的提示 -->
          <div v-if="searchDevices.length === 0">
            <p>没有查询到相关设备</p>
          </div>
          <!--  如果不为空显示设备列表 -->
          <div v-else>
            <van-cell-group
              v-for="device in searchDevices"
              :key="device.udid"
              inset
              class="cert-container"
            >
              <van-field v-model="device.cert_id" readonly label="证书编号" />
              <van-field v-model="device.udid" readonly label="UDID" />
              <van-field v-model="device.redeem_code" readonly label="兑换码" />
              <van-field v-model="device.name" readonly label="证书名称" />
              <van-field v-model="device.add_time" readonly label="添加时间" />
              <van-field
                v-model="device.account_type"
                readonly
                label="证书类型"
              />
              <van-field v-model="device.status" readonly label="状态" />
            </van-cell-group>
          </div>
        </div>
      </div>
      <div class="notice-popup-footer">
        <div class="notice-buttons">
          <van-button
            round
            block
            class="close-button"
            @click="showDevicesPopup = false"
          >
            <van-icon name="checked" class="button-icon" />
            <span>我知道了</span>
          </van-button>
        </div>
      </div>
    </div>
  </van-popup>

  <van-popup v-model:show="showInstallPopup" round>
    <div class="notice-popup-wrapper">
      <div class="notice-popup-header">
        <div class="notice-popup-title">
          <van-icon name="info" class="notice-icon" />
          <span>安装提示</span>
        </div>
        <van-icon
          name="close"
          style="font-size: 24px; color: var(--text-secondary)"
          @click="showInstallPopup = false"
        />
      </div>
      <div class="notice-popup-content">
        <div class="notice-popup-text">
          <p>您的APP`{{ navbarTitle }}`已签名成功</p>
        </div>
      </div>
      <div class="notice-popup-footer">
        <div class="notice-buttons">
          <van-button round block class="close-button" @click="copyInstallURL">
            <van-icon name="link-o" class="button-icon" />
            <span>复制安装链接</span>
          </van-button>
          <van-button round block class="close-button" @click="openInstallURL">
            <van-icon name="down" class="button-icon" />
            <span>立即安装</span>
          </van-button>
        </div>
      </div>
    </div>
  </van-popup>
</template>

<style lang="less" scoped>
.demo-content {
  display: flex;
  flex-direction: column;
  overflow-y: auto; /* 优先使用 auto 避免不必要的滚动条 */
  //-webkit-overflow-scrolling: touch; /* 启用 iOS 弹性滚动 */
  //height: 100vh; // 让整个页面填充满屏幕
  height: calc(100dvh - var(--header-height) - var(--van-tabbar-height));
  /* 处理刘海屏遮挡 */
  padding-top: env(safe-area-inset-top);
  padding-bottom: env(safe-area-inset-bottom);
  //padding-bottom: var(--van-tabbar-height); /* 避免内容被底部栏遮挡 */
  //margin-top: var(--van-nav-bar-height); /* 补偿顶部导航栏高度 */
  //overflow: scroll;

  .flex-container {
    flex: 1; /* 填充剩余空间 */
    display: flex;
    flex-direction: column;

    .max-container {
      flex-grow: 1; /* 内容区域自动扩展 */
      overflow-y: auto; /* 允许内容滚动 */
    }

    .copyright {
      flex-shrink: 0; /* 禁止底部收缩 */
      //padding: 12px 0;
    }
  }
}

.tabs-content {
  flex: 1;
  overflow-y: auto; // 允许内容滚动
  padding-bottom: 20px; // 防止滚动时内容过于贴近底部
}

.notice-popup {
  max-height: 60vh;
}

.notice-popup-wrapper {
  display: flex;
  flex-direction: column;
  height: 100%;
  max-height: 60vh;
  background: var(--card-background);
  animation: slideUp 0.3s ease-out;
}

.notice-popup-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  align-items: center;
  justify-content: space-between;

  .notice-popup-title {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 18px;
    font-weight: 600;
    color: var(--text-primary);

    .notice-icon {
      color: var(--primary-color);
      font-size: 22px;
    }
  }
}

.notice-popup-content {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  -webkit-overflow-scrolling: touch;

  .notice-popup-text {
    font-size: 15px;
    line-height: 1.6;
    color: var(--text-secondary);
    white-space: pre-wrap;
  }
}

.notice-popup-footer {
  padding: 16px 20px calc(16px + env(safe-area-inset-bottom));
  border-top: 1px solid var(--border-color);
}

.button-icon {
  margin-right: 4px;
  font-size: 18px;
}

@media (max-width: 480px) {
  .notice-popup-header {
    padding: 16px;

    .notice-popup-title {
      font-size: 16px;

      .notice-icon {
        font-size: 20px;
      }
    }
  }

  .notice-popup-content {
    padding: 16px;

    .notice-popup-text {
      font-size: 14px;
    }
  }

  .notice-popup-footer {
    padding: 12px 16px calc(12px + env(safe-area-inset-bottom));
  }
}

@keyframes slideUp {
  0% {
    transform: translateY(20px);
    opacity: 0;
  }
  100% {
    transform: translateY(0);
    opacity: 1;
  }
}

.notice-buttons {
  display: flex;
  flex-direction: row;
  gap: 8px;
  white-space: nowrap;
}

.notice-buttons .van-button {
  margin: 0 !important;
  flex: 1;
}

.hide-button {
  background: var(--primary-color) !important;
  border-color: var(--primary-color) !important;
}

.close-button {
  color: var(--text-secondary) !important;
  border-color: var(--border-color) !important;
  background: transparent !important;
}

.cert-container {
  margin-left: 0 !important;
  margin-right: 0 !important;
  margin-bottom: 16px;
  margin-top: 16px;
  border: 1px solid var(--van-text-color-2);
  background-color: var(--van-text-color-2);
}

.dark .close-button {
  border-color: #ffffff1a !important;
}

@media (max-width: 480px) {
  .notice-buttons {
    gap: 6px;
  }

  .notice-buttons .van-button {
    height: 40px;
    font-size: 13px;
  }

  .button-icon {
    font-size: 16px;
  }
}

.notice-countdown {
  font-size: 14px;
  color: var(--primary-color);
  font-weight: 500;
  border-radius: 12px;
}
</style>
