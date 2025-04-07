<template>
  <!-- 应用管理详情抽屉 -->
  <div class="appsource-asApplication-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>应用管理详情</h4>
      </template>
      <el-descriptions
          class="margin-top"
          :column="3"
          border
          style="margin: 8px;"
      >
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              ID
            </div>
          </template>
          {{ formData.id }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              文件
            </div>
          </template>
          {{ formData.fileUrl }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              应用名称
            </div>
          </template>
          {{ formData.name }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              包名
            </div>
          </template>
          {{ formData.bundleId }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              版本号
            </div>
          </template>
          {{ formData.version }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              文件大小
            </div>
          </template>
          {{ formData.size }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              类型
            </div>
          </template>
          {{ proxy.getOptionValue(formData.type, typeOptions, 'value', 'label') }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              描述
            </div>
          </template>
          {{ formData.description }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              图标
            </div>
          </template>
          {{ formData.iconUrl }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              付费
            </div>
          </template>
          <el-switch v-model="formData.lock" class="ml-2" disabled/>
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              是否蓝奏云链接
            </div>
          </template>
          {{ proxy.getOptionValue(formData.lanzou, lanzouOptions, 'value', 'label') }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              是否启用
            </div>
          </template>
          <el-switch v-model="formData.active" class="ml-2" disabled/>
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              备注
            </div>
          </template>
          {{ formData.note }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              创建人
            </div>
          </template>
          {{ formData.createdBy }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              修改人
            </div>
          </template>
          {{ formData.updatedBy }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              创建时间
            </div>
          </template>
          {{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
        </el-descriptions-item>
        <el-descriptions-item :span="1">
          <template #label>
            <div class="cell-item">
              更新时间
            </div>
          </template>
          {{ proxy.parseTime(formData.updatedAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import {getCurrentInstance, reactive, ref, toRefs} from 'vue';
import {getAsApplication,} from "/@/api/appsource/asApplication";
import {AsApplicationEditState, AsApplicationInfoData} from "/@/views/appsource/asApplication/list/component/model"

defineOptions({name: "ApiV1AppsourceAsApplicationDetail"})
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
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<AsApplicationEditState>({
  loading: false,
  isShowDialog: false,
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
    lock: undefined,
    lanzou: false,
    weigh: undefined,
    active: undefined,
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
const {isShowDialog, formData} = toRefs(state);
// 打开弹窗
const openDialog = (row?: AsApplicationInfoData) => {
  resetForm();
  if (row) {
    getAsApplication(row.id!).then((res: any) => {
      const data = res.data;
      data.createdBy = data.createdUser?.userNickname
      data.updatedBy = data.updatedUser?.userNickname
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
const resetForm = () => {
  state.formData = {
    id: undefined,
    fileUrl: undefined,
    name: undefined,
    bundleId: undefined,
    version: undefined,
    size: undefined,
    type: undefined,
    description: undefined,
    iconUrl: undefined,
    lock: undefined,
    lanzou: false,
    weigh: undefined,
    active: undefined,
    note: undefined,
    createdBy: undefined,
    updatedBy: undefined,
    createdAt: undefined,
    updatedAt: undefined,
  }
};
</script>
<style scoped>
.appsource-asApplication-detail :deep(.el-form-item--large .el-form-item__label) {
  font-weight: bolder;
}

.pic-block {
  margin-right: 8px;
}

.file-block {
  width: 100%;
  border: 1px solid var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
  margin-bottom: 5px;
  padding: 3px 6px;
}

.ml-2 {
  margin-right: 5px;
}
</style>