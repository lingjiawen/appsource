<template>  
  <div class="install-signRedeemCode-edit">
    <!-- 添加或修改签名卡密对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.install-signRedeemCode-edit .el-dialog', '.install-signRedeemCode-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'签名卡密'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="兑换码" prop="code">
          <el-input v-model="formData.code" placeholder="请输入兑换码" />
        </el-form-item>        
        <el-form-item label="UDID" prop="udid">
          <el-input v-model="formData.udid" placeholder="请输入UDID" />
        </el-form-item>        
        <el-form-item label="证书标识" prop="certId">
          <el-input v-model="formData.certId" placeholder="请输入证书标识" />
        </el-form-item>        
        <el-form-item label="时效类型" prop="accountType">
          <el-radio-group v-model="formData.accountType">
            <el-radio
              v-for="dict in accountTypeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="售后类型" prop="warrantyType">
          <el-radio-group v-model="formData.warrantyType">
            <el-radio
              v-for="dict in warrantyTypeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="设备类型" prop="deviceType">
          <el-radio-group v-model="formData.deviceType">
            <el-radio
              v-for="dict in deviceTypeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="出书方式" prop="pool">
          <el-radio-group v-model="formData.pool">
            <el-radio
              v-for="dict in poolOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="对接平台" prop="apiPlatform">
          <el-radio-group v-model="formData.apiPlatform">
            <el-radio
              v-for="dict in apiPlatformOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="备注" prop="note">
          <el-input v-model="formData.note" type="textarea" placeholder="请输入备注" />
        </el-form-item>        
        <el-form-item label="对接售后类型" prop="apiWarrantyType">
          <el-radio-group v-model="formData.apiWarrantyType">
            <el-radio
              v-for="dict in apiWarrantyTypeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="强制添加" prop="force">
          <el-switch  v-model="formData.force" class="ml-2" />
        </el-form-item>
        <el-form-item label="禁用" prop="banned">
          <el-switch  v-model="formData.banned" class="ml-2" />
        </el-form-item>        
        <el-form-item label="激活" prop="active">
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
  listSignRedeemCode,
  getSignRedeemCode,
  delSignRedeemCode,
  addSignRedeemCode,
  updateSignRedeemCode,  
} from "/@/api/install/signRedeemCode";
import {
  SignRedeemCodeTableColumns,
  SignRedeemCodeInfoData,
  SignRedeemCodeTableDataState,
  SignRedeemCodeEditState
} from "/@/views/install/signRedeemCode/list/component/model"
defineOptions({ name: "ApiV1InstallSignRedeemCodeEdit"})
const emit = defineEmits(['signRedeemCodeList'])
  const props = defineProps({    
    accountTypeOptions:{
      type:Array,
      default:()=>[]
    },    
    warrantyTypeOptions:{
      type:Array,
      default:()=>[]
    },    
    deviceTypeOptions:{
      type:Array,
      default:()=>[]
    },    
    poolOptions:{
      type:Array,
      default:()=>[]
    },    
    apiPlatformOptions:{
      type:Array,
      default:()=>[]
    },    
    apiWarrantyTypeOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SignRedeemCodeEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    code: undefined,    
    udid: undefined,    
    certId: undefined,    
    accountType: undefined,    
    warrantyType: undefined,    
    deviceType: undefined,    
    pool: undefined,    
    apiPlatform: undefined,    
    note: undefined,    
    apiWarrantyType: undefined,    
    banned: false,
    active: false,
    force: false,
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
const openDialog = (row?: SignRedeemCodeInfoData) => {
  resetForm();
  if(row) {
    getSignRedeemCode(row.id!).then((res:any)=>{
      const data = res.data;      
      data.accountType = ''+data.accountType      
      data.warrantyType = ''+data.warrantyType      
      data.deviceType = ''+data.deviceType      
      data.pool = ''+data.pool      
      data.apiPlatform = ''+data.apiPlatform      
      data.apiWarrantyType = ''+data.apiWarrantyType      
      data.banned = Boolean(data.banned)
      data.force = Boolean(data.force)
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
      addSignRedeemCode(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('signRedeemCodeList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSignRedeemCode(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('signRedeemCodeList')
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
    certId: undefined,    
    accountType: '' ,    
    warrantyType: '' ,    
    deviceType: '' ,    
    pool: '' ,    
    apiPlatform: '' ,    
    note: '',
    apiWarrantyType: '' ,    
    banned: false ,
    force: false,
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