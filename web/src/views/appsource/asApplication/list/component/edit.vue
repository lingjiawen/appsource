<template>
  <div class="appsource-asApplication-edit">
    <!-- 添加或修改应用管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-asApplication-edit .el-dialog', '.appsource-asApplication-edit .el-dialog__header']">
          {{ (!formData.id || formData.id == 0 ? '添加' : '修改') + '应用管理' }}
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-space fill>
          <el-alert type="info" show-icon :closable="false">
            <p>1. 本地上传将存储文件/图标在本地, 需要大量的带宽消耗</p>
            <p>2. 解析文件需要临时下载软件包, 请勿解析超大网盘文件包</p>
          </el-alert>
          <el-form-item label="本地上传包">
            <upload-file :action="baseURL+'api/v1/system/upload/singleFile'"
                         :limit="1"
                         @update:modelValue="handleFileUploadSuccess">
            </upload-file>
            <div class="image-preview">
              <div class="el-form-item__label">图片预览</div>
              <el-image style="max-width:150px;max-height: 150px;"
                        :src="proxy.getUpFileUrl(formData.iconUrl)"
                        :zoom-rate="1.2"
                        :max-scale="7"
                        :min-scale="0.2"
                        :preview-src-list="[proxy.getUpFileUrl(formData.iconUrl)]"
                        :initial-index="0"
                        preview-teleported
                        hide-on-click-modal/>
            </div>
          </el-form-item>
        </el-space>
        <el-form-item label="软件地址" prop="fileUrl">
          <el-input v-model="formData.fileUrl" placeholder="请输入软件地址或上传">
            <template #append>
              <el-button @click="extractFile">解析文件</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="应用名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入应用名称"/>
        </el-form-item>
        <el-form-item label="图标" prop="iconUrl">
          <el-input v-model="formData.iconUrl" placeholder="请输入图标">
            <template #append>
              <upload-img class="app-image-upload"
                          :action="baseURL+'api/v1/system/upload/singleImg'"
                          @update:modelValue="handleImageUploadSuccess"
                          :limit="1">
              </upload-img>
            </template>
            <template #prepend>
              <el-button @click="iconSearchVisible = true">搜索</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="包名" prop="bundleId">
          <el-input v-model="formData.bundleId" placeholder="请输入包名"/>
        </el-form-item>
        <el-form-item label="版本号" prop="version">
          <el-input v-model="formData.version" placeholder="请输入版本号"/>
        </el-form-item>
        <el-form-item label="文件大小" prop="size">
          <el-input v-model="formData.showSize" placeholder="请输入文件大小"/>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select filterable clearable v-model="formData.type" placeholder="请选择类型">
            <el-option
                v-for="dict in typeOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input type="textarea"
                    :autosize="{ minRows: 4, maxRows: 10}"
                    v-model="formData.description" placeholder="请输入描述"/>
        </el-form-item>
        <el-form-item label="付费" prop="lock">
          <el-switch v-model="formData.lock" class="ml-2"/>
        </el-form-item>
        <el-form-item label="是否蓝奏云链接" prop="lanzou">
          <el-radio-group v-model="formData.lanzou">
            <el-radio
                v-for="dict in lanzouOptions"
                :key="dict.value"
                :value="dict.value"
            >{{ dict.label }}
            </el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="权重" prop="weigh">
          <el-input-number v-model="formData.weigh" placeholder="请输入权重"/>
        </el-form-item>
        <el-form-item label="是否启用" prop="active">
          <el-switch v-model="formData.active" class="ml-2"/>
        </el-form-item>
        <el-form-item label="备注" prop="note">
          <el-input v-model="formData.note" type="textarea" placeholder="请输入备注"/>
        </el-form-item>
      </el-form>
      <el-dialog
          v-model="iconSearchVisible"
          width="600"
          title="AppStore图标搜索"
          append-to-body
          :close-on-click-modal="false"
          :destroy-on-close="true"
          @opened="focusInput"
      >
        <div class="mt-4">
          <el-input
              ref="appSearchInputRef"
              v-model="searchedAppName"
              style="max-width: 600px"
              placeholder="输入应用名称"
              class="input-with-select"
              @keyup.enter="appStoreSearch"
          >
            <template #append>
              <el-button :icon="Search" @click="appStoreSearch"/>
            </template>
            <template #prepend>
              <el-select v-model="iconSearchCountrySelect" placeholder="选择地区" style="width: 115px">
                <el-option value="CN" label="中国"></el-option>
                <el-option value="US" label="美国"></el-option>
                <el-option value="JP" label="日本"></el-option>
                <el-option value="KR" label="韩国"></el-option>
                <el-option value="HK" label="香港"></el-option>
                <el-option value="TW" label="台湾"></el-option>
                <el-option value="GB" label="英国"></el-option>
                <el-option value="DE" label="德国"></el-option>
                <el-option value="FR" label="法国"></el-option>
              </el-select>
            </template>
          </el-input>
          <AppList ref="appListRef" @iconSelected="handleIconSelected"/>
        </div>
      </el-dialog>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="onSubmit" :disabled="loading">确 定</el-button>
          <el-button @click="onCancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script setup lang="ts">
import {computed, getCurrentInstance, nextTick, reactive, ref, toRefs, unref, watch} from 'vue';
import {ElLoading, ElMessage} from 'element-plus';
import AppList from './AppList.vue'
import {
  addAsApplication,
  extractAsApplication,
  getAsApplication,
  searchIconAsApplication,
  updateAsApplication,
} from "/@/api/appsource/asApplication";
import {AsApplicationEditState, AsApplicationInfoData} from "/@/views/appsource/asApplication/list/component/model"
import UploadFile from "/@/components/uploadFile/index.vue";
import UploadImg from "/@/components/uploadImg/index.vue";
import {Search} from "@element-plus/icons-vue";
import {appDownloadStore} from "/@/stores/appDownloadStore";

defineOptions({name: "ApiV1AppsourceAsApplicationEdit"})
const emit = defineEmits(['asApplicationList'])
const props = defineProps({
  typeOptions: {
    type: Array,
    default: () => []
  },
  lockOptions: {
    type: Array,
    default: () => []
  },
  lanzouOptions: {
    type: Array,
    default: () => []
  },
  activeOptions: {
    type: Array,
    default: () => []
  },
})
const baseURL: string | undefined | boolean = import.meta.env.VITE_API_URL
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const iconSearchVisible = ref(false);
const appSearchInputRef = ref(null) // 绑定输入框
const iconSearchCountrySelect = ref('CN')
const searchedAppName = ref('')
const state = reactive<AsApplicationEditState>({
  loading: false,
  isShowDialog: false,
  appDownloadMessage: "开始下载...",
  formData: {
    id: undefined,
    fileUrl: undefined,
    name: undefined,
    bundleId: undefined,
    version: undefined,
    size: undefined,
    type: undefined,
    description: undefined,
    iconUrl: undefined,
    lock: false,
    lanzou: undefined,
    weigh: undefined,
    active: false,
    note: undefined,
    createdBy: undefined,
    updatedBy: undefined,
    createdAt: undefined,
    updatedAt: undefined,
  },
  // 表单校验
  rules: {
    id: [
      {required: true, message: "ID不能为空", trigger: "blur"}
    ],
    name: [
      {required: true, message: "应用名称不能为空", trigger: "blur"}
    ],
  }
});
const {loading, isShowDialog, formData, rules, appDownloadMessage} = toRefs(state);
const appListRef = ref(null)
const appDownloadStoreAct = appDownloadStore()
const getProgressMessages = computed(() => {
  return appDownloadStoreAct.message;
});

const showDownloadProgress = (msg: any) => {
  state.appDownloadMessage = msg.msg
};

watch(() => state.formData.size, (newSize) => {
  let sizeInKB = newSize / 1024;  // 转为KB
  if (sizeInKB < 1024) {
    // 小于1MB
    state.formData.showSize = sizeInKB.toFixed(2) + " KB";
  } else {
    let sizeInMB = sizeInKB / 1024; // 转为MB
    if (sizeInMB < 1024) {
      // 小于1GB
      state.formData.showSize = sizeInMB.toFixed(2) + " MB";
    } else {
      let sizeInGB = sizeInMB / 1024; // 转为GB
      state.formData.showSize = sizeInGB.toFixed(2) + " GB";
    }
  }
});

// 监听APP下载进度
watch(getProgressMessages, (nv, ov) => {
  if (!nv || nv.type == undefined) {

    return;
  }
  showDownloadProgress(nv);
}, { immediate: true, deep: true });

const socketStatus = computed(() => {
  return appDownloadStoreAct.socketStatus; // 直接定义为计算属性
});

const focusInput = async () => {
  await nextTick() // 确保 DOM 更新
  appSearchInputRef.value?.focus() // 直接聚焦 input
}

// 打开弹窗
const openDialog = (row?: AsApplicationInfoData) => {
  resetForm();
  if (row) {
    getAsApplication(row.id!).then((res: any) => {
      const data = res.data;
      data.type = '' + data.type
      data.lock = Boolean(data.lock)
      data.lanzou = '' + data.lanzou
      data.active = Boolean(data.active)
      state.formData = data;
    })
  }
  state.isShowDialog = true;
};
// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
defineExpose({
  openDialog,
});
// 取消
const onCancel = () => {
  closeDialog();
};
// 提交
const onSubmit = () => {
  const formWrap = unref(formRef) as any;
  if (!formWrap) return;
  formWrap.validate((valid: boolean) => {
    if (valid) {
      state.loading = true;
      if (!state.formData.id || state.formData.id === 0) {
        //添加
        addAsApplication(state.formData).then(() => {
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('asApplicationList')
        }).finally(() => {
          state.loading = false;
        })
      } else {
        //修改
        updateAsApplication(state.formData).then(() => {
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('asApplicationList')
        }).finally(() => {
          state.loading = false;
        })
      }
    }
  });
};
const resetForm = () => {
  state.formData = {
    id: undefined,
    fileUrl: undefined,
    name: undefined,
    bundleId: undefined,
    version: undefined,
    size: undefined,
    showSize: "0 KB", // 用于提交的实际大小
    type: "0",
    description: undefined,
    iconUrl: "https://img0.baidu.com/it/u=104684862,1159259552&fm=253",
    lock: false,
    lanzou: "0",
    weigh: 0,
    active: true,
    note: undefined,
    createdBy: undefined,
    updatedBy: undefined,
    createdAt: undefined,
    updatedAt: undefined,
  }
};
const handleFileUploadSuccess = (fileList: any[]) => {
  if (fileList.length > 0) {
    const fileData = fileList[0]; // 只有一个文件上传
    state.formData.fileUrl = fileData.url;  // 存储文件URL
    state.formData.name = fileData.name;    // 存储文件名称
    state.formData.size = fileData.size;    // 存储文件名称

    // 处理文件大小转换
    let sizeInKB = fileData.size / 1024;  // 转为KB
    if (sizeInKB < 1024) {
      // 小于1MB
      state.formData.showSize = sizeInKB.toFixed(2) + " KB";
    } else {
      let sizeInMB = sizeInKB / 1024; // 转为MB
      if (sizeInMB < 1024) {
        // 小于1GB
        state.formData.showSize = sizeInMB.toFixed(2) + " MB";
      } else {
        let sizeInGB = sizeInMB / 1024; // 转为GB
        state.formData.showSize = sizeInGB.toFixed(2) + " GB";
      }
    }
  }
};

const handleImageUploadSuccess = (imageList: any[]) => {
  if (imageList.length > 0) {
    const imageData = imageList[0]; // 只有一个文件上传
    state.formData.iconUrl = imageData.url;  // 存储文件URL
  }
}

const handleIconSelected = (iconUrl) => {
  iconSearchVisible.value = false;
  state.formData.iconUrl = iconUrl;
}

const extractFile = () => {
  const loadingInstance = ElLoading.service({text: appDownloadMessage, fullscreen: true})
  extractAsApplication({
    "fileUrl": state.formData.fileUrl
  }).then((res: any) => {
    const data = res.data;
    state.formData.iconUrl = data.iconUrl;
    state.formData.name = data.name;
    state.formData.bundleId = data.bundleId;
    state.formData.version = data.version;
    state.formData.size = data.appSize;
    appDownloadMessage.value = "开始下载...";
    loadingInstance.close();
  }).finally(() => {
    loadingInstance.close();
  })
}

const appStoreSearch = async () => {
  const country = iconSearchCountrySelect.value
  const appName = searchedAppName.value

  if (!appName) {
    ElMessage.warning('请输入应用名称')
    return
  }

  // 启动 Element Plus 加载动画
  const loadingInstance = ElLoading.service({text: '搜索中...', fullscreen: true})

  searchIconAsApplication({
    "name": appName,
    "country": country,
  }).then((result) => {
    const data = result.data;
    if (data.results.length === 0) {
      ElMessage.warning('未找到相关应用')
    } else {
      appListRef.value.displayResults(data.results)
    }
  }).catch(() => {
    ElMessage.error('搜索服务暂时不可用，请稍后重试或尝试使用其他地区')
    loadingInstance.close()
  }).finally(() => {
    state.loading = false;
    loadingInstance.close()
  })
}

</script>
<style scoped>
.image-preview {
  display: flex;
  justify-content: start;
  align-items: flex-start;
  flex-direction: column;
  margin: 0 20px;
  height: 100%;
}

:global(.app-image-upload) {
  max-height: 100%;
}

:global(.app-image-upload .avatar-uploader) {
  justify-content: center;
  display: flex;
}

:global(.app-image-upload.up-img .el-icon.avatar-uploader-icon) {
  max-width: 32px;
  max-height: 32px;
}

:global(.app-image-upload.up-img img) {
  max-width: 32px;
  max-height: 32px;
}


.kv-label {
  margin-bottom: 15px;
  font-size: 14px;
}

.mini-btn i.el-icon {
  margin: unset;
}

.kv-row {
  margin-bottom: 12px;
}
</style>