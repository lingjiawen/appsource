<template>  
  <div class="install-signDevice-edit">
    <!-- 添加或修改设备管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.install-signDevice-edit .el-dialog', '.install-signDevice-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'设备管理'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="设备码" prop="udid">
          <el-input v-model="formData.udid" placeholder="请输入设备码" />
        </el-form-item>        
        <el-form-item label="证书名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入证书名" />
        </el-form-item>        
        <el-form-item label="证书标识" prop="certId">
          <el-input v-model="formData.certId" placeholder="请输入证书标识" />
        </el-form-item>        
        <el-form-item label="设备标识" prop="deviceId">
          <el-input v-model="formData.deviceId" placeholder="请输入设备标识" />
        </el-form-item>        
        <el-form-item label="证书文件" prop="p12">
          <el-input v-model="formData.p12" placeholder="请输入证书文件" />
        </el-form-item>        
        <el-form-item label="描述文件" prop="mobileprovision">
          <el-input v-model="formData.mobileprovision" placeholder="请输入描述文件" />
        </el-form-item>        
        <el-form-item label="添加时间" prop="addTime">
          <el-input v-model="formData.addTime" placeholder="请输入添加时间" />
        </el-form-item>        
        <el-form-item label="设备型号" prop="model">
          <el-input v-model="formData.model" placeholder="请输入设备型号" />
        </el-form-item>        
        <el-form-item label="过期时间" prop="expireTime">
          <el-input v-model="formData.expireTime" placeholder="请输入过期时间" />
        </el-form-item>        
        <el-form-item label="兑换卡密" prop="redeemCode">
          <el-input v-model="formData.redeemCode" placeholder="请输入兑换卡密" />
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
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio
              v-for="dict in statusOptions"
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
        <el-form-item label="对接售后类型" prop="apiWarrantyType">
          <el-radio-group v-model="formData.apiWarrantyType">
            <el-radio
              v-for="dict in apiWarrantyTypeOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="禁用" prop="active">
          <el-switch  v-model="formData.active" class="ml-2" />
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
  listSignDevice,
  getSignDevice,
  delSignDevice,
  addSignDevice,
  updateSignDevice,  
} from "/@/api/install/signDevice";
import {
  SignDeviceTableColumns,
  SignDeviceInfoData,
  SignDeviceTableDataState,
  SignDeviceEditState
} from "/@/views/install/signDevice/list/component/model"
defineOptions({ name: "ApiV1InstallSignDeviceEdit"})
const emit = defineEmits(['signDeviceList'])
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
    statusOptions:{
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
const state = reactive<SignDeviceEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    udid: undefined,    
    name: undefined,    
    certId: undefined,    
    deviceId: undefined,    
    p12: undefined,    
    mobileprovision: undefined,    
    addTime: undefined,    
    p12Password: undefined,    
    model: undefined,    
    expireTime: undefined,    
    serialNumber: undefined,    
    redeemCode: undefined,    
    accountType: undefined,    
    warrantyType: undefined,    
    deviceType: undefined,    
    status: undefined,    
    pool: undefined,    
    apiPlatform: undefined,    
    apiWarrantyType: undefined,    
    active: false ,    
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
    name : [
        { required: true, message: "证书名不能为空", trigger: "blur" }
    ],    
    status : [
        { required: true, message: "状态不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SignDeviceInfoData) => {
  resetForm();
  if(row) {
    getSignDevice(row.id!).then((res:any)=>{
      const data = res.data;      
      data.accountType = ''+data.accountType      
      data.warrantyType = ''+data.warrantyType      
      data.deviceType = ''+data.deviceType      
      data.status = ''+data.status      
      data.pool = ''+data.pool      
      data.apiPlatform = ''+data.apiPlatform      
      data.apiWarrantyType = ''+data.apiWarrantyType      
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
      addSignDevice(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('signDeviceList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSignDevice(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('signDeviceList')
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
    udid: undefined,    
    name: undefined,    
    certId: undefined,    
    deviceId: undefined,    
    p12: undefined,    
    mobileprovision: undefined,    
    addTime: undefined,    
    p12Password: undefined,    
    model: undefined,    
    expireTime: undefined,    
    serialNumber: undefined,    
    redeemCode: undefined,    
    accountType: '' ,    
    warrantyType: '' ,    
    deviceType: '' ,    
    status: '' ,    
    pool: '' ,    
    apiPlatform: '' ,    
    apiWarrantyType: '' ,    
    active: false ,    
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