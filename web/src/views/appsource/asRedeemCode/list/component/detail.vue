<template>
  <!-- 卡密管理详情抽屉 -->  
  <div class="appsource-asRedeemCode-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>卡密管理详情</h4>
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
                  兑换码
                </div>
              </template>
              {{ formData.code }}            
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
                    类型
                  </div>
                </template>
                {{ proxy.getOptionValue(formData.type, typeOptions,'value','label') }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                是否激活
              </div>
            </template>
            <el-switch  v-model="formData.active" class="ml-2" disabled />
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                激活时间
              </div>
            </template>
            {{ proxy.parseTime(formData.activeAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
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
                生成时间
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
  defineOptions({ name: "ApiV1AppsourceAsRedeemCodeDetail"})  
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
      type: false ,      
      active: undefined,      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: AsRedeemCodeInfoData) => {
    resetForm();
    if(row) {
      getAsRedeemCode(row.id!).then((res:any)=>{
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
      code: undefined,      
      udid: undefined,      
      type: false ,      
      active: undefined,      
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
  .appsource-asRedeemCode-detail :deep(.el-form-item--large .el-form-item__label){
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