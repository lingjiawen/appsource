<template>
  <!-- 配置分组详情抽屉 -->  
  <div class="appsource-sysConfigGroup-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>配置分组详情</h4>
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
                  主键
                </div>
              </template>
              {{ formData.id }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  配置组名称
                </div>
              </template>
              {{ formData.groupName }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  配置组键名
                </div>
              </template>
              {{ formData.groupKey }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  创建者
                </div>
              </template>
              {{ formData.createBy }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  更新者
                </div>
              </template>
              {{ formData.updateBy }}            
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
  defineOptions({ name: "ApiV1AppsourceSysConfigGroupDetail"})  
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
  const { isShowDialog,formData } = toRefs(state);
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
  .appsource-sysConfigGroup-detail :deep(.el-form-item--large .el-form-item__label){
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