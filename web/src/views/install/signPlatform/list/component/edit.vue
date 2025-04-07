<template>  
  <div class="install-signPlatform-edit">
    <!-- 添加或修改平台对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.install-signPlatform-edit .el-dialog', '.install-signPlatform-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'平台'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="平台名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入平台名" />
        </el-form-item>        
        <el-form-item label="平台标识" prop="code">
          <el-input v-model="formData.code" placeholder="请输入平台标识" />
        </el-form-item>        
        <el-form-item label="域名" prop="baseUrl">
          <el-input v-model="formData.baseUrl" placeholder="请输入域名" />
        </el-form-item>        
        <el-form-item label="开启SSL" prop="openSsl">
          <el-switch  v-model="formData.openSsl" class="ml-2" />
        </el-form-item>        
        <el-form-item label="启用" prop="status">
          <el-switch  v-model="formData.status" class="ml-2" />
        </el-form-item>        
        <el-form-item label="对接Token" prop="token">
          <el-input v-model="formData.token" placeholder="请输入对接Token" />
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
import { reactive, toRefs, ref,unref,getCurrentInstance,computed } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listSignPlatform,
  getSignPlatform,
  delSignPlatform,
  addSignPlatform,
  updateSignPlatform,  
} from "/@/api/install/signPlatform";
import {
  SignPlatformTableColumns,
  SignPlatformInfoData,
  SignPlatformTableDataState,
  SignPlatformEditState
} from "/@/views/install/signPlatform/list/component/model"
defineOptions({ name: "ApiV1InstallSignPlatformEdit"})
const emit = defineEmits(['signPlatformList'])
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SignPlatformEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    name: undefined,    
    code: undefined,    
    baseUrl: undefined,    
    openSsl: false ,    
    status: false ,    
    token: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
    name : [
        { required: true, message: "平台名不能为空", trigger: "blur" }
    ],    
    baseUrl : [
        { required: true, message: "域名不能为空", trigger: "blur" }
    ],    
    status : [
        { required: true, message: "启用不能为空", trigger: "blur" }
    ],    
    token : [
        { required: true, message: "对接Token不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SignPlatformInfoData) => {
  resetForm();
  if(row) {
    getSignPlatform(row.id!).then((res:any)=>{
      const data = res.data;      
      data.openSsl = Boolean(data.openSsl)      
      data.status = Boolean(data.status)      
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
      if(!state.formData.id || state.formData.id===0){
        //添加
      addSignPlatform(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('signPlatformList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSignPlatform(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('signPlatformList')
        }).finally(()=>{
          state.loading = false;
        })
      }
    }
  });
};
const resetForm = ()=>{
  state.formData = {    
    id: undefined,    
    name: undefined,    
    code: undefined,    
    baseUrl: undefined,    
    openSsl: false ,    
    status: false ,    
    token: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>