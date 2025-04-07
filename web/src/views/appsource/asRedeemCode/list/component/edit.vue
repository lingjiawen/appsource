<template>  
  <div class="appsource-asRedeemCode-edit">
    <!-- 添加或修改卡密管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-asRedeemCode-edit .el-dialog', '.appsource-asRedeemCode-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'卡密管理'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="兑换码" prop="code">
          <el-input v-model="formData.code" placeholder="请输入兑换码" />
        </el-form-item>        
        <el-form-item label="设备码" prop="udid">
          <el-input v-model="formData.udid" placeholder="请输入设备码" />
        </el-form-item>        
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="formData.type">
            <el-radio
              v-for="dict in typeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="是否激活" prop="active">
          <el-switch  v-model="formData.active" class="ml-2" />
        </el-form-item>        
        <el-form-item label="激活时间" prop="activeAt">
          <el-date-picker clearable  style="width: 200px"
            v-model="formData.activeAt"
            type="datetime"
            value-format="YYYY-MM-DD HH:mm:ss"
            placeholder="选择激活时间">
          </el-date-picker>
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
  listAsRedeemCode,
  getAsRedeemCode,
  delAsRedeemCode,
  addAsRedeemCode,
  updateAsRedeemCode,  
} from "/@/api/appsource/asRedeemCode";
import {
  AsRedeemCodeTableColumns,
  AsRedeemCodeInfoData,
  AsRedeemCodeTableDataState,
  AsRedeemCodeEditState
} from "/@/views/appsource/asRedeemCode/list/component/model"
defineOptions({ name: "ApiV1AppsourceAsRedeemCodeEdit"})
const emit = defineEmits(['asRedeemCodeList'])
  const props = defineProps({    
    typeOptions:{
      type:Array,
      default:()=>[]
    },    
    activeOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<AsRedeemCodeEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    code: undefined,    
    udid: undefined,    
    type: undefined,    
    active: false ,    
    activeAt: undefined,    
    createdBy: undefined,    
    updatedBy: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  },
  // 表单校验
  rules: {    
    id : [
        { required: true, message: "ID不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: AsRedeemCodeInfoData) => {
  resetForm();
  if(row) {
    getAsRedeemCode(row.id!).then((res:any)=>{
      const data = res.data;      
      data.type = ''+data.type      
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
      if(!state.formData.id || state.formData.id===0){
        //添加
      addAsRedeemCode(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('asRedeemCodeList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateAsRedeemCode(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('asRedeemCodeList')
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
    code: undefined,    
    udid: undefined,    
    type: '' ,    
    active: false ,    
    activeAt: undefined,    
    createdBy: undefined,    
    updatedBy: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    deletedAt: undefined,    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>