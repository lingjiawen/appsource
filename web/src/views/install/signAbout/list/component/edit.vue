<template>  
  <div class="install-signAbout-edit">
    <!-- 添加或修改关于我们对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.install-signAbout-edit .el-dialog', '.install-signAbout-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'关于我们'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">        
        <el-form-item label="图标" prop="icon">
          <el-input v-model="formData.icon" placeholder="请输入图标" />
        </el-form-item>        
        <el-form-item label="标题" prop="title">
          <el-input v-model="formData.title" placeholder="请输入标题" />
        </el-form-item>        
        <el-form-item label="内容" prop="subtitle">
          <el-input v-model="formData.subtitle" placeholder="请输入内容" />
        </el-form-item>        
        <el-form-item label="是否链接" prop="isLink">
          <el-switch  v-model="formData.isLink" class="ml-2" />
        </el-form-item>        
        <el-form-item label="链接" prop="url">
          <el-input v-model="formData.url" placeholder="请输入链接" />
        </el-form-item>        
        <el-form-item label="分组" prop="group">
          <el-radio-group v-model="formData.group">
            <el-radio
              v-for="dict in groupOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="权重" prop="weigh">
          <el-input v-model="formData.weigh" placeholder="请输入权重" />
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
  listSignAbout,
  getSignAbout,
  delSignAbout,
  addSignAbout,
  updateSignAbout,  
} from "/@/api/install/signAbout";
import {
  SignAboutTableColumns,
  SignAboutInfoData,
  SignAboutTableDataState,
  SignAboutEditState
} from "/@/views/install/signAbout/list/component/model"
defineOptions({ name: "ApiV1InstallSignAboutEdit"})
const emit = defineEmits(['signAboutList'])
  const props = defineProps({    
    groupOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<SignAboutEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    icon: undefined,    
    title: undefined,    
    subtitle: undefined,    
    isLink: false ,    
    url: undefined,    
    group: undefined,    
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
    subtitle : [
        { required: true, message: "内容不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: SignAboutInfoData) => {
  resetForm();
  if(row) {
    getSignAbout(row.id!).then((res:any)=>{
      const data = res.data;      
      data.isLink = Boolean(data.isLink)      
      data.group = ''+data.group      
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
      addSignAbout(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('signAboutList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateSignAbout(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('signAboutList')
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
    icon: undefined,    
    title: undefined,    
    subtitle: undefined,    
    isLink: false ,    
    url: undefined,    
    group: '' ,    
    weigh: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  }  
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>