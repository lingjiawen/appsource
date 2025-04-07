<template>
  <div class="install-signDevice-container">
    <el-card shadow="hover">
        <div class="install-signDevice-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="ID" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入ID"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="设备码" prop="udid">
                    <el-input
                        v-model="tableData.param.udid"
                        placeholder="请输入设备码"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="signDeviceList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="证书名" prop="name">
                    <el-input
                        v-model="tableData.param.name"
                        placeholder="请输入证书名"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="证书标识" prop="certId">
                    <el-input
                        v-model="tableData.param.certId"
                        placeholder="请输入证书标识"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="设备标识" prop="deviceId">
                    <el-input
                        v-model="tableData.param.deviceId"
                        placeholder="请输入设备标识"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="添加时间" prop="addTime">
                    <el-input
                        v-model="tableData.param.addTime"
                        placeholder="请输入添加时间"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="设备型号" prop="model">
                    <el-input
                        v-model="tableData.param.model"
                        placeholder="请输入设备型号"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="过期时间" prop="expireTime">
                    <el-input
                        v-model="tableData.param.expireTime"
                        placeholder="请输入过期时间"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="兑换卡密" prop="redeemCode">
                    <el-input
                        v-model="tableData.param.redeemCode"
                        placeholder="请输入兑换卡密"
                        clearable                        
                        @keyup.enter.native="signDeviceList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="时效类型" prop="accountType">
                    <el-select filterable v-model="tableData.param.accountType" placeholder="请选择时效类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in apple_account_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="售后类型" prop="warrantyType">
                    <el-select filterable v-model="tableData.param.warrantyType" placeholder="请选择售后类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in apple_warranty_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="设备类型" prop="deviceType">
                    <el-select filterable v-model="tableData.param.deviceType" placeholder="请选择设备类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in apple_device_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="状态" prop="status">
                    <el-select filterable v-model="tableData.param.status" placeholder="请选择状态" clearable style="width:200px;">
                        <el-option
                            v-for="dict in cert_status"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="出书方式" prop="pool">
                    <el-select filterable v-model="tableData.param.pool" placeholder="请选择出书方式" clearable style="width:200px;">
                        <el-option
                            v-for="dict in apple_pool_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="对接平台" prop="apiPlatform">
                    <el-select filterable v-model="tableData.param.apiPlatform" placeholder="请选择对接平台" clearable style="width:200px;">
                        <el-option
                            v-for="dict in api_platform_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="对接售后类型" prop="apiWarrantyType">
                    <el-select filterable v-model="tableData.param.apiWarrantyType" placeholder="请选择对接售后类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in apple_warranty_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="禁用" prop="active">
                      <el-switch  v-model="tableData.param.active" class="ml-2" />
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建时间" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择创建时间"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="signDeviceList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                        {{ word }}
                        <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                        <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>            
              </el-row>
            </el-form>
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd"
                  v-auth="'api/v1/install/signDevice/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/install/signDevice/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/install/signDevice/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/install/signDevice/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/install/signDevice/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="设备码" align="center" prop="udid"
            min-width="150px"            
             />          
          <el-table-column label="证书名" align="center" prop="name"
            min-width="150px"            
             />          
          <el-table-column label="证书标识" align="center" prop="certId"
            min-width="150px"            
             />          
          <el-table-column label="添加时间" align="center" prop="addTime"
            min-width="150px"            
             />          
          <el-table-column label="设备型号" align="center" prop="model"
            min-width="150px"            
             />          
          <el-table-column label="过期时间" align="center" prop="expireTime"
            min-width="150px"            
             />          
          <el-table-column label="兑换卡密" align="center" prop="redeemCode"
            min-width="150px"            
             />          
          <el-table-column label="时效类型" align="center" prop="accountType" :formatter="accountTypeFormat"
            min-width="150px"            
             />          
          <el-table-column label="售后类型" align="center" prop="warrantyType" :formatter="warrantyTypeFormat"
            min-width="150px"            
             />          
          <el-table-column label="设备类型" align="center" prop="deviceType" :formatter="deviceTypeFormat"
            min-width="150px"            
             />          
          <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat"
            min-width="150px"            
             />          
          <el-table-column label="出书方式" align="center" prop="pool" :formatter="poolFormat"
            min-width="150px"            
             />          
          <el-table-column label="对接平台" align="center" prop="apiPlatform" :formatter="apiPlatformFormat"
            min-width="150px"            
             />          
          <el-table-column label="对接售后类型" align="center" prop="apiWarrantyType" :formatter="apiWarrantyTypeFormat"
            min-width="150px"            
             />          
          <el-table-column label="禁用" align="center" prop="active"
          min-width="150px"          
          >
              <template #default="scope">
                <el-switch  v-model="scope.row.active"
                            class="ml-2"
                            :active-value=0
                            :inactive-value=1
                            @change="changeActive(scope.row)"/>
              </template>
          </el-table-column>          
          <el-table-column label="创建人" align="center" prop="createdBy"
            min-width="150px"            
             />          
          <el-table-column label="创建时间" align="center" prop="createdAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="200px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleView(scope.row)"
                v-auth="'api/v1/install/signDevice/get'"
              ><el-icon><ele-View /></el-icon>详情</el-button>              
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/install/signDevice/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/install/signDevice/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="signDeviceList"
        />
    </el-card>
    <ApiV1InstallSignDeviceEdit
       ref="editRef"       
       :accountTypeOptions="apple_account_type"       
       :warrantyTypeOptions="apple_warranty_type"       
       :deviceTypeOptions="apple_device_type"       
       :statusOptions="cert_status"       
       :poolOptions="apple_pool_type"       
       :apiPlatformOptions="api_platform_type"       
       :apiWarrantyTypeOptions="apple_warranty_type"       
       @signDeviceList="signDeviceList"
    ></ApiV1InstallSignDeviceEdit>
    <ApiV1InstallSignDeviceDetail
      ref="detailRef"      
      :accountTypeOptions="apple_account_type"      
      :warrantyTypeOptions="apple_warranty_type"      
      :deviceTypeOptions="apple_device_type"      
      :statusOptions="cert_status"      
      :poolOptions="apple_pool_type"      
      :apiPlatformOptions="api_platform_type"      
      :apiWarrantyTypeOptions="apple_warranty_type"      
      @signDeviceList="signDeviceList"
    ></ApiV1InstallSignDeviceDetail>    
    <loadExcel ref="loadExcelSignDeviceRef" @getList="signDeviceList"
               upUrl="api/v1/install/signDevice/import"
               tplUrl="/api/v1/install/signDevice/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listSignDevice,
    getSignDevice,
    delSignDevice,
    addSignDevice,
    updateSignDevice,    
    changeSignDeviceActive,    
} from "/@/api/install/signDevice";
import {
    SignDeviceTableColumns,
    SignDeviceInfoData,
    SignDeviceTableDataState,    
} from "/@/views/install/signDevice/list/component/model"
import ApiV1InstallSignDeviceEdit from "/@/views/install/signDevice/list/component/edit.vue"
import ApiV1InstallSignDeviceDetail from "/@/views/install/signDevice/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1InstallSignDeviceList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelSignDeviceRef = ref();
// 是否显示所有搜索选项
const showAll =  ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple =ref(true)
const word = computed(()=>{
    if(showAll.value === false) {
        //对文字进行处理
        return "展开搜索";
    } else {
        return "收起搜索";
    }
})
// 字典选项数据
const {    
    apple_account_type,    
    apple_warranty_type,    
    apple_device_type,    
    cert_status,    
    apple_pool_type,    
    api_platform_type,    
} = proxy.useDict(    
    'apple_account_type',    
    'apple_warranty_type',    
    'apple_device_type',    
    'cert_status',    
    'apple_pool_type',    
    'api_platform_type',    
)
const state = reactive<SignDeviceTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            udid: undefined,            
            name: undefined,            
            certId: undefined,            
            deviceId: undefined,            
            addTime: undefined,            
            model: undefined,            
            expireTime: undefined,            
            redeemCode: undefined,            
            accountType: undefined,            
            warrantyType: undefined,            
            deviceType: undefined,            
            status: undefined,            
            pool: undefined,            
            apiPlatform: undefined,            
            apiWarrantyType: undefined,            
            active: undefined,            
            createdBy: undefined,            
            createdAt: undefined,            
            dateRange: []
        },
    },
});
const { tableData } = toRefs(state);
// 页面加载时
onMounted(() => {
    initTableData();
});
// 初始化表格数据
const initTableData = () => {    
    signDeviceList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    signDeviceList()
};
// 获取列表数据
const signDeviceList = ()=>{
  loading.value = true
  listSignDevice(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    list.map((item:any)=>{        
        item.createdBy = item.createdUser?.userNickname        
    })    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 时效类型字典翻译
const accountTypeFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(apple_account_type.value, row.accountType);
}
// 售后类型字典翻译
const warrantyTypeFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(apple_warranty_type.value, row.warrantyType);
}
// 设备类型字典翻译
const deviceTypeFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(apple_device_type.value, row.deviceType);
}
// 状态字典翻译
const statusFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(cert_status.value, row.status);
}
// 出书方式字典翻译
const poolFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(apple_pool_type.value, row.pool);
}
// 对接平台字典翻译
const apiPlatformFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(api_platform_type.value, row.apiPlatform);
}
// 对接售后类型字典翻译
const apiWarrantyTypeFormat = (row:SignDeviceTableColumns) => {
    return proxy.selectDictLabel(apple_warranty_type.value, row.apiWarrantyType);
}
const changeActive = (row:SignDeviceTableColumns) => {
    changeSignDeviceActive(row.id,row.active)
        .catch(()=>{
            setTimeout(()=>{
                row.active = !row.active
            },300)
        })
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<SignDeviceInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: SignDeviceTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:SignDeviceTableColumns)=>{
            return item.id ===state.ids[0]
        }) as SignDeviceTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SignDeviceTableColumns|null) => {
    let msg = '你确定要删除所选数据？';
    let id:number[] = [] ;
    if(row){
    msg = `此操作将永久删除数据，是否继续?`
    id = [row.id]
    }else{
    id = state.ids
    }
    if(id.length===0){
        ElMessage.error('请选择要删除的数据。');
        return
    }
    ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            delSignDevice(id).then(()=>{
                ElMessage.success('删除成功');
                signDeviceList();
            })
        })
        .catch(() => {});
}
const handleView = (row:SignDeviceTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/install/signDevice/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelSignDeviceRef.value.open()
}
</script>
<style lang="scss" scoped>
    .colBlock {
        display: block;
    }
    .colNone {
        display: none;
    }
    .ml-2{margin: 3px;}
</style>