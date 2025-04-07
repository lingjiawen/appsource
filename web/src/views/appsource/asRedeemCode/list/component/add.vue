<template>
  <div class="install-RedeemCode-add">
    <!-- 添加或修改卡密管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-asRedeemCode-add .el-dialog', '.appsource-asRedeemCode-add .el-dialog__header']">
          {{ (!formData.id || formData.id == 0 ? '添加' : '修改') + '卡密管理' }}
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-form-item label="兑换码前缀" prop="prefix">
          <el-input v-model="formData.prefix" placeholder="兑换码前缀"/>
        </el-form-item>
        <el-form-item label="生成数量" prop="quantity">
          <el-input ref="quantityInput" v-model="formData.quantity" placeholder="生成数量"/>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="formData.type">
            <el-radio
                v-for="dict in typeOptions"
                :key="dict.value"
                :value="dict.value"
            >{{ dict.label }}
            </el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
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
import {getCurrentInstance, onMounted, reactive, ref, toRefs, unref} from 'vue';
import {ElMessage} from 'element-plus';
import {addAsRedeemCode,} from "/@/api/appsource/asRedeemCode";
import {AsRedeemCodeAddState,} from "/@/views/appsource/asRedeemCode/list/component/model"

defineOptions({name: "ApiV1AppsourceAsRedeemCodeEdit"})
const emit = defineEmits(['asRedeemCodeList'])
const props = defineProps({
  typeOptions: {
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
const quantityInput = ref(null);
const state = reactive<AsRedeemCodeAddState>({
  loading: false,
  isShowDialog: false,
  formData: {
    id: undefined,
    quantity: undefined,
    prefix: undefined,
    type: undefined,
  }
});
const {loading, isShowDialog, formData, rules} = toRefs(state);
// 打开弹窗
const openDialog = () => {
  resetForm();
  state.isShowDialog = true;
  setTimeout(() => {
    if (quantityInput.value) {
      quantityInput.value.focus(); // 直接调用 el-input 的 focus 方法
    }
  }, 100);
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
        addAsRedeemCode(state.formData).then(() => {
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('asRedeemCodeList')
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
    quantity: 1,
    prefix: '',
    type: '3',
  }
};
</script>
<style scoped>
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