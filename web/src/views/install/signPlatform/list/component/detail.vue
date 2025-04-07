<template>
  <!-- 平台详情抽屉 -->  
  <div class="install-signPlatform-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>平台详情</h4>
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
                  平台名
                </div>
              </template>
              {{ formData.name }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  平台标识
                </div>
              </template>
              {{ formData.code }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  域名
                </div>
              </template>
              {{ formData.baseUrl }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                开启SSL
              </div>
            </template>
            <el-switch  v-model="formData.openSsl" class="ml-2" disabled />
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                启用
              </div>
            </template>
            <el-switch  v-model="formData.status" class="ml-2" disabled />
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  对接Token
                </div>
              </template>
              {{ formData.token }}            
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
  defineOptions({ name: "ApiV1InstallSignPlatformDetail"})  
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
      openSsl: undefined,      
      status: undefined,      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: SignPlatformInfoData) => {
    resetForm();
    if(row) {
      getSignPlatform(row.id!).then((res:any)=>{
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
  const resetForm = ()=>{
    state.formData = {      
      id: undefined,      
      name: undefined,      
      code: undefined,      
      baseUrl: undefined,      
      openSsl: undefined,      
      status: undefined,      
      token: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
    }
  };  
</script>
<style scoped>  
  .install-signPlatform-detail :deep(.el-form-item--large .el-form-item__label){
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