<template>
  <!-- 帮助中心详情抽屉 -->  
  <div class="install-signHelp-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>帮助中心详情</h4>
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
                  标题
                </div>
              </template>
              {{ formData.title }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                内容
              </div>
            </template>
            <div v-html="formData.content"></div>
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                默认展开
              </div>
            </template>
            <el-switch  v-model="formData.isExpand" class="ml-2" disabled />
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  权重
                </div>
              </template>
              {{ formData.weigh }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">
            <template #label>
              <div class="cell-item">
                创建日期
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
  defineOptions({ name: "ApiV1InstallSignHelpDetail"})  
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
      isExpand: undefined,      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: SignHelpInfoData) => {
    resetForm();
    if(row) {
      getSignHelp(row.id!).then((res:any)=>{
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
      title: undefined,      
      content: undefined,      
      isExpand: undefined,      
      weigh: undefined,      
      createdAt: undefined,      
      updatedAt: undefined,      
    }
  };  
  //富文本编辑器内容
  const setContentEditContent = (data:string) => {
    state.formData.content = data
  }  
</script>
<style scoped>  
  .install-signHelp-detail :deep(.el-form-item--large .el-form-item__label){
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