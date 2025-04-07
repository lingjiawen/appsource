<template>  
  <div class="appsource-sysConfig-edit">
    <!-- 添加或修改配置对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-sysConfig-edit .el-dialog', '.appsource-sysConfig-edit .el-dialog__header']">{{(!formData.configId || formData.configId==0?'添加':'修改')+'配置'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="参数名称" prop="configName">
          <el-input v-model="formData.configName" placeholder="请输入参数名称" />
        </el-form-item>        
        <el-form-item label="参数键名" prop="configKey">
          <el-input v-model="formData.configKey" placeholder="请输入参数键名" />
        </el-form-item>        
        <el-form-item label="参数键值" prop="configValue">
          <el-input v-model="formData.configValue" placeholder="请输入参数键值" />
        </el-form-item>          
        <el-form-item label="系统内置（Y是 N否）" prop="configType">
          <el-select filterable clearable v-model="formData.configType" placeholder="请选择系统内置（Y是 N否）" >
            <el-option label="请选择字典生成" value="" />
          </el-select>
        </el-form-item>        
        <el-form-item label="参数分组" prop="group">
          <el-input v-model="formData.group" placeholder="请输入参数分组" />
        </el-form-item>        
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" placeholder="请输入备注" />
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
  listSysConfig,
  getSysConfig,
  delSysConfig,
  addSysConfig,
  updateSysConfig,  
} from "/@/api/appsource/sysConfig";
import {
  SysConfigTableColumns,
  SysConfigInfoData,
  SysConfigTableDataState,
  SysConfigEditState
} from "/@/views/appsource/sysConfig/list/component/model"
defineOptions({ name: "ApiV1AppsourceSysConfigEdit"})
const emit = defineEmits(['sysConfigList'])
  const props = defineProps({    
    groupOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SysConfigEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    configId: undefined,    
    configName: undefined,    
    configKey: undefined,    
    configValue: undefined,    
    configType: undefined,    
    group: undefined,    
    createBy: undefined,    
    updateBy: undefined,    
    remark: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    linkedSysConfigSysConfigGroup: {      
      groupKey:undefined,    // 配置组键名      
      groupName:undefined,    // 配置组名称      
    },    
  },
  // 表单校验
  rules: {    
    configId : [
        { required: true, message: "参数主键不能为空", trigger: "blur" }
    ],    
    configName : [
        { required: true, message: "参数名称不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SysConfigInfoData) => {
  resetForm();
  if(row) {
    getSysConfig(row.configId!).then((res:any)=>{
      const data = res.data;      
      data.configType = parseInt(data.configType)      
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
      if(!state.formData.configId || state.formData.configId===0){
        //添加
      addSysConfig(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('sysConfigList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSysConfig(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('sysConfigList')
        }).finally(()=>{
          state.loading = false;
        })
      }
    }
  });
};
const resetForm = ()=>{
  state.formData = {    
    configId: undefined,    
    configName: undefined,    
    configKey: undefined,    
    configValue: undefined,    
    configType: undefined,    
    group: undefined,    
    createBy: undefined,    
    updateBy: undefined,    
    remark: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
    linkedSysConfigSysConfigGroup: {      
      groupKey:undefined,    // 配置组键名      
      groupName:undefined,    // 配置组名称      
    },    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>