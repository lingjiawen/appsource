<template>
  <!-- 配置详情抽屉 -->  
  <div class="appsource-sysConfig-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>配置详情</h4>
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
                  参数主键
                </div>
              </template>
              {{ formData.configId }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  参数名称
                </div>
              </template>
              {{ formData.configName }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  参数键名
                </div>
              </template>
              {{ formData.configKey }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  参数键值
                </div>
              </template>
              {{ formData.configValue }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  系统内置（Y是 N否）
                </div>
              </template>
              {{ formData.configType }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">                  
                    <template #label>
                      <div class="cell-item">
                        参数分组
                      </div>
                    </template>
                    {{ formData.linkedGroup?formData.linkedGroup.groupName:'' }}            
          </el-descriptions-item>        
          <el-descriptions-item :span="1">            
              <template #label>
                <div class="cell-item">
                  备注
                </div>
              </template>
              {{ formData.remark }}            
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
  defineOptions({ name: "ApiV1AppsourceSysConfigDetail"})  
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
      linkedGroup:{groupKey:undefined,groupName:undefined },      
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
  const { isShowDialog,formData } = toRefs(state);
  // 打开弹窗
  const openDialog = (row?: SysConfigInfoData) => {
    resetForm();
    if(row) {
      getSysConfig(row.configId!).then((res:any)=>{
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
      configId: undefined,      
      configName: undefined,      
      configKey: undefined,      
      configValue: undefined,      
      configType: undefined,      
      group: undefined,      
      linkedGroup:{groupKey:undefined,groupName:undefined },      
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
  //关联sys_config_group表选项
  const getSysConfigGroupItemsGroup = () => {
    emit("getSysConfigGroupItemsGroup")
  }
  const getGroupOp = computed(()=>{
    getSysConfigGroupItemsGroup()
    return props.groupOptions
  })  
</script>
<style scoped>  
  .appsource-sysConfig-detail :deep(.el-form-item--large .el-form-item__label){
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