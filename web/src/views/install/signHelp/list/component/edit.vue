<template>  
  <div class="install-signHelp-edit">
    <!-- 添加或修改帮助中心对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.install-signHelp-edit .el-dialog', '.install-signHelp-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'帮助中心'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入标题" />
        </el-form-item>        
        <el-form-item label="内容">
          <gf-ueditor editorId="ueSignHelpContent" v-model="formData.content"></gf-ueditor>
        </el-form-item>        
        <el-form-item label="默认展开" prop="isExpand">
          <el-switch  v-model="formData.isExpand" class="ml-2" />
        </el-form-item>        
        <el-form-item label="权重" prop="weigh">
          <el-input v-model="formData.weigh" placeholder="请输入权重" />
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
  listSignHelp,
  getSignHelp,
  delSignHelp,
  addSignHelp,
  updateSignHelp,  
} from "/@/api/install/signHelp";
import GfUeditor from "/@/components/ueditor/index.vue"
import {
  SignHelpTableColumns,
  SignHelpInfoData,
  SignHelpTableDataState,
  SignHelpEditState
} from "/@/views/install/signHelp/list/component/model"
defineOptions({ name: "ApiV1InstallSignHelpEdit"})
const emit = defineEmits(['signHelpList'])
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SignHelpEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    title: undefined,    
    content: undefined,    
    isExpand: false ,    
    weigh: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
    title : [
        { required: true, message: "标题不能为空", trigger: "blur" }
    ],    
    content : [
        { required: true, message: "内容不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SignHelpInfoData) => {
  resetForm();
  if(row) {
    getSignHelp(row.id!).then((res:any)=>{
      const data = res.data;      
      data.isExpand = Boolean(data.isExpand)      
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
      addSignHelp(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('signHelpList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSignHelp(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('signHelpList')
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
    title: undefined,    
    content: undefined,    
    isExpand: false ,    
    weigh: undefined,    
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