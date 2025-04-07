<template>  
  <div class="appsource-asApplication-edit">
    <!-- 添加或修改应用管理对话框 -->
    <el-dialog v-model="isShowDialog" width="800px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.appsource-asApplication-edit .el-dialog', '.appsource-asApplication-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'应用管理'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-space fill>
          <el-alert type="info" show-icon :closable="false">
            <p>1. 本地上传将存储文件/图标在本地, 需要大量的带宽消耗</p>
            <p>2. 解析文件需要临时下载软件包, 请勿解析超大网盘文件包</p>
          </el-alert>
          <el-form-item label="本地上传包" prop="fileUrl">
            <upload-file :action="baseURL+'api/v1/system/upload/singleFile'"
                         :limit="1"
                         @update:modelValue="handleFileUploadSuccess">
            </upload-file>
          </el-form-item>
        </el-space>
        <el-form-item label="文件地址" prop="fileUrl">
          <el-input v-model="formData.fileUrl" placeholder="请输入文件地址或上传">
            <template #append>
              <el-button>解析文件</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="应用名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入应用名称" />
        </el-form-item>        
        <el-form-item label="包名" prop="bundleId">
          <el-input v-model="formData.bundleId" placeholder="请输入包名" />
        </el-form-item>        
        <el-form-item label="版本号" prop="version">
          <el-input v-model="formData.version" placeholder="请输入版本号" />
        </el-form-item>        
        <el-form-item label="文件大小" prop="size">
          <el-input v-model="formData.size" placeholder="请输入文件大小" />
        </el-form-item>        
        <el-form-item label="类型" prop="type">
          <el-select filterable clearable v-model="formData.type" placeholder="请选择类型" >
            <el-option
                v-for="dict in typeOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
            ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="描述" prop="description">
          <el-input v-model="formData.description" placeholder="请输入描述" />
        </el-form-item>        
        <el-form-item label="图标" prop="iconUrl">
          <el-input v-model="formData.iconUrl" placeholder="请输入图标" />
        </el-form-item>        
        <el-form-item label="付费" prop="lock">
          <el-switch  v-model="formData.lock" class="ml-2" />
        </el-form-item>        
        <el-form-item label="是否蓝奏云链接" prop="lanzou">
          <el-radio-group v-model="formData.lanzou">
            <el-radio
              v-for="dict in lanzouOptions"
              :key="dict.value"
              :value="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="权重" prop="weigh">
          <el-input-number v-model="formData.weigh" placeholder="请输入权重"  />
        </el-form-item>        
        <el-form-item label="是否启用" prop="active">
          <el-switch  v-model="formData.active" class="ml-2" />
        </el-form-item>        
        <el-form-item label="备注" prop="note">
          <el-input v-model="formData.note" type="textarea" placeholder="请输入备注" />
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
  listAsApplication,
  getAsApplication,
  delAsApplication,
  addAsApplication,
  updateAsApplication,  
} from "/@/api/appsource/asApplication";
import uploadFile from "/@/components/uploadFile/index.vue"
import {
  AsApplicationTableColumns,
  AsApplicationInfoData,
  AsApplicationTableDataState,
  AsApplicationEditState
} from "/@/views/appsource/asApplication/list/component/model"
defineOptions({ name: "ApiV1AppsourceAsApplicationEdit"})
const emit = defineEmits(['asApplicationList'])
  const props = defineProps({
    typeOptions:{
      type:Array,
      default:()=>[]
    },
    lockOptions:{
      type:Array,
      default:()=>[]
    },    
    lanzouOptions:{
      type:Array,
      default:()=>[]
    },    
    activeOptions:{
      type:Array,
      default:()=>[]
    },    
  })
const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL
const {proxy} = <any>getCurrentInstance()
const formRef = ref<HTMLElement | null>(null);
const menuRef = ref();
const state = reactive<AsApplicationEditState>({
  loading:false,
  isShowDialog: false,
  formData: {    
    id: undefined,    
    fileUrl: [] ,    
    name: undefined,    
    bundleId: undefined,    
    version: undefined,    
    size: undefined,    
    type: undefined,    
    description: undefined,    
    iconUrl: undefined,    
    lock: false ,    
    lanzou: undefined,    
    weigh: undefined,    
    active: false ,    
    note: undefined,    
    createdBy: undefined,    
    updatedBy: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  },
  // 表单校验
  rules: {
    fileUrl : [
      { required: true, message: "文件地址不能为空", trigger: "blur" }
    ],
    name : [
        { required: true, message: "应用名称不能为空", trigger: "blur" }
    ],    
  }
});
const { loading,isShowDialog,formData,rules } = toRefs(state);
// 打开弹窗
const openDialog = (row?: AsApplicationInfoData) => {
  resetForm();
  if(row) {
    getAsApplication(row.id!).then((res:any)=>{
      const data = res.data;      
      data.fileUrl =data.fileUrl?JSON.parse(data.fileUrl) : []      
      data.type = parseInt(data.type)      
      data.lock = Boolean(data.lock)      
      data.lanzou = ''+data.lanzou      
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
      addAsApplication(state.formData).then(()=>{
          ElMessage.success('添加成功');
          closeDialog(); // 关闭弹窗
          emit('asApplicationList')
        }).finally(()=>{
          state.loading = false;
        })
      }else{
        //修改
      updateAsApplication(state.formData).then(()=>{
          ElMessage.success('修改成功');
          closeDialog(); // 关闭弹窗
          emit('asApplicationList')
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
    fileUrl: [] ,    
    name: undefined,    
    bundleId: undefined,    
    version: undefined,    
    size: undefined,    
    type: '' ,    
    description: undefined,    
    iconUrl: undefined,    
    lock: false ,    
    lanzou: '' ,    
    weigh: undefined,    
    active: false ,    
    note: undefined,    
    createdBy: undefined,    
    updatedBy: undefined,    
    createdAt: undefined,    
    updatedAt: undefined,    
  }  
};

// 处理文件上传成功的事件
const handleFileUploadSuccess = (fileList: any[]) => {
  if (fileList.length > 0) {
    const fileData = fileList[0]; // 假设只有一个文件上传
    console.log(fileData);
    state.formData.fileUrl = fileData.url;  // 存储文件URL
    state.formData.name = fileData.name;    // 存储文件名称

    // 处理文件大小转换
    let sizeInKB = fileData.size / 1024;  // 转为KB
    if (sizeInKB < 1024) {
      // 小于1MB
      state.formData.size = sizeInKB.toFixed(2) + " KB";
    } else {
      let sizeInMB = sizeInKB / 1024; // 转为MB
      if (sizeInMB < 1024) {
        // 小于1GB
        state.formData.size = sizeInMB.toFixed(2) + " MB";
      } else {
        let sizeInGB = sizeInMB / 1024; // 转为GB
        state.formData.size = sizeInGB.toFixed(2) + " GB";
      }
    }
  }
};
</script>
<style scoped>  
  .kv-label{margin-bottom: 15px;font-size: 14px;}
  .mini-btn i.el-icon{margin: unset;}
  .kv-row{margin-bottom: 12px;}
</style>