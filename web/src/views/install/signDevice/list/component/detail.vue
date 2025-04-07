<template>
  <!-- 设备管理详情抽屉 -->  
  <div class="install-signDevice-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>设备管理详情</h4>
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
                  设备码
                </div>
              </template>
              {{ formData.udid }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  证书名
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  证书标识
                </div>
              </template>
              {{ formData.certId }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  设备标识
                </div>
              </template>
              {{ formData.deviceId }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  证书文件
                </div>
              </template>
              {{ formData.p12 }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  描述文件
                </div>
              </template>
              {{ formData.mobileprovision }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  添加时间
                </div>
              </template>
              {{ formData.addTime }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  证书密码
                </div>
              </template>
              {{ formData.p12Password }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  设备型号
                </div>
              </template>
              {{ formData.model }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  过期时间
                </div>
              </template>
              {{ formData.expireTime }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  序列号
                </div>
              </template>
              {{ formData.serialNumber }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  兑换卡密
                </div>
              </template>
              {{ formData.redeemCode }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    时效类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.accountType, accountTypeOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    售后类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.warrantyType, warrantyTypeOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    设备类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.deviceType, deviceTypeOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    状态
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.status, statusOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    出书方式
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.pool, poolOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    对接平台
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.apiPlatform, apiPlatformOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">              
                <template #label>
                  <div class="cell-item">
                    对接售后类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.apiWarrantyType, apiWarrantyTypeOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                禁用
              </div>
            </template>
            <el-switch  v-model="formData.active" class="ml-2" disabled />
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
      </el-descriptions>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
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
  defineOptions({ name: "ApiV1InstallSignDeviceDetail"})  
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
      accountType: false ,      
      warrantyType: false ,      
      deviceType: false ,      
      status: false ,      
      pool: false ,      
      apiPlatform: false ,      
      apiWarrantyType: false ,      
      active: undefined,      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: SignDeviceInfoData) => {
    resetForm();
    if(row) {
      getSignDevice(row.id!).then((res:any)=>{
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
      accountType: false ,      
      warrantyType: false ,      
      deviceType: false ,      
      status: false ,      
      pool: false ,      
      apiPlatform: false ,      
      apiWarrantyType: false ,      
      active: undefined,      
      createdBy: undefined,      
      updatedBy: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
      deletedAt: undefined,      
    }
  };  
</script>
<style scoped>  
  .install-signDevice-detail :deep(.el-form-item--large .el-form-item__label){
    font-weight: bolder;
  }
  .pic-block{
    margin-right: 8px;
  }
  .file-block{
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
  .ml-2{margin-right: 5px;}
</style>