<template>  
  <div class="appsource-sysConfigGroup-edit">
    <!-- 添加或修改配置分组对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-sysConfigGroup-edit .el-dialog', '.appsource-sysConfigGroup-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'配置分组'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="配置组名称" prop="groupName">
          <el-input v-model="formData.groupName" placeholder="请输入配置组名称" />
        </el-form-item>        
        <el-form-item label="配置组键名" prop="groupKey">
          <el-input v-model="formData.groupKey" placeholder="请输入配置组键名" />
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
  listSysConfigGroup,
  getSysConfigGroup,
  delSysConfigGroup,
  addSysConfigGroup,
  updateSysConfigGroup,  
} from "/@/api/appsource/sysConfigGroup";
import {
  SysConfigGroupTableColumns,
  SysConfigGroupInfoData,
  SysConfigGroupTableDataState,
  SysConfigGroupEditState
} from "/@/views/appsource/sysConfigGroup/list/component/model"
defineOptions({ name: "ApiV1AppsourceSysConfigGroupEdit"})
const emit = defineEmits(['sysConfigGroupList'])
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SysConfigGroupEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    groupName: undefined,    
    groupKey: undefined,    
    createBy: undefined,    
    updateBy: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "主键不能为空", trigger: "blur" }
    ],    
    groupName : [
        { required: true, message: "配置组名称不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SysConfigGroupInfoData) => {
  resetForm();
  if(row) {
    getSysConfigGroup(row.id!).then((res:any)=>{
      const data = res.data;      
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
      addSysConfigGroup(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('sysConfigGroupList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSysConfigGroup(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('sysConfigGroupList')
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
    groupName: undefined,    
    groupKey: undefined,    
    createBy: undefined,    
    updateBy: undefined,    
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