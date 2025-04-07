<template>
  <div class="appsource-sysConfig-container">
    <el-card shadow="hover">
        <div class="appsource-sysConfig-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="参数主键" prop="configId">
                    <el-input
                        v-model="tableData.param.configId"
                        placeholder="请输入参数主键"
                        clearable                        
                        @keyup.enter.native="sysConfigList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="参数名称" prop="configName">
                    <el-input
                        v-model="tableData.param.configName"
                        placeholder="请输入参数名称"
                        clearable                        
                        @keyup.enter.native="sysConfigList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="sysConfigList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="参数键名" prop="configKey">
                    <el-input
                        v-model="tableData.param.configKey"
                        placeholder="请输入参数键名"
                        clearable                        
                        @keyup.enter.native="sysConfigList"
                    />                    
                  </el-form-item>

                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="参数键值" prop="configValue">
                    <el-input
                        v-model="tableData.param.configValue"
                        placeholder="请输入参数键值"
                        clearable                        
                        @keyup.enter.native="sysConfigList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="系统内置（Y是 N否）" prop="configType">
                    <el-select filterable v-model="tableData.param.configType" placeholder="请选择系统内置（Y是 N否）" clearable style="width:200px;">
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="参数分组" prop="group">
                    <el-input
                        v-model="tableData.param.group"
                        placeholder="请输入参数分组"
                        clearable                        
                        @keyup.enter.native="sysConfigList"
                    />                    
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
                    <el-button type="primary"  @click="sysConfigList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/appsource/sysConfig/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/appsource/sysConfig/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/appsource/sysConfig/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="参数主键" align="center" prop="configId"
            min-width="150px"            
             />          
          <el-table-column label="参数名称" align="center" prop="configName"
            min-width="150px"            
             />          
          <el-table-column label="参数键名" align="center" prop="configKey"
            min-width="150px"            
             />          
          <el-table-column label="参数键值" align="center" prop="configValue"
            min-width="150px"            
             />          
          <el-table-column label="系统内置（Y是 N否）" align="center" prop="configType"
            min-width="150px"            
             />          
          <el-table-column label="参数分组" align="center" prop="linkedGroup.groupName"
            min-width="150px"            
             />          
          <el-table-column label="备注" align="center" prop="remark"
            min-width="150px"            
             />          
          <el-table-column label="创建时间" align="center" prop="createdAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/appsource/sysConfig/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/appsource/sysConfig/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="sysConfigList"
        />
    </el-card>
    <ApiV1AppsourceSysConfigEdit
       ref="editRef"       
       :groupOptions="groupOptions"       
       @sysConfigList="sysConfigList"
    ></ApiV1AppsourceSysConfigEdit>
    <ApiV1AppsourceSysConfigDetail
      ref="detailRef"      
      :groupOptions="groupOptions"      
      @sysConfigList="sysConfigList"
    ></ApiV1AppsourceSysConfigDetail>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listSysConfig,
    getSysConfig,
    delSysConfig,
    addSysConfig,
    updateSysConfig,    
    linkedDataSearch    
} from "/@/api/appsource/sysConfig";
import {
    SysConfigTableColumns,
    SysConfigInfoData,
    SysConfigTableDataState,    
    LinkedSysConfigSysConfigGroup,    
} from "/@/views/appsource/sysConfig/list/component/model"
import ApiV1AppsourceSysConfigEdit from "/@/views/appsource/sysConfig/list/component/edit.vue"
import ApiV1AppsourceSysConfigDetail from "/@/views/appsource/sysConfig/list/component/detail.vue"
defineOptions({ name: "apiV1AppsourceSysConfigList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
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
} = proxy.useDict(    
)
// groupOptions关联表数据
const groupOptions = ref<Array<ItemOptions>>([])
const state = reactive<SysConfigTableDataState>({
    configIds:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            configId: undefined,            
            configName: undefined,            
            configKey: undefined,            
            configValue: undefined,            
            configType: undefined,            
            group: undefined,            
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
    linkedData()    
    sysConfigList()
};
const linkedData = ()=>{
    linkedDataSearch().then((res:any)=>{        
        //关联sys_config_group表选项        
        groupOptions.value = proxy.setItems(res, 'groupKey', 'groupName','linkedSysConfigSysConfigGroup')        
    })
}
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    sysConfigList()
};
// 获取列表数据
const sysConfigList = ()=>{
  loading.value = true
  listSysConfig(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<SysConfigInfoData>) => {
    state.configIds = selection.map(item => item.configId)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: SysConfigTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:SysConfigTableColumns)=>{
            return item.configId ===state.configIds[0]
        }) as SysConfigTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SysConfigTableColumns|null) => {
    let msg = '你确定要删除所选数据？';
    let configId:number[] = [] ;
    if(row){
    msg = `此操作将永久删除数据，是否继续?`
    configId = [row.configId]
    }else{
    configId = state.configIds
    }
    if(configId.length===0){
        ElMessage.error('请选择要删除的数据。');
        return
    }
    ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
    })
        .then(() => {
            delSysConfig(configId).then(()=>{
                ElMessage.success('删除成功');
                sysConfigList();
            })
        })
        .catch(() => {});
}
const handleView = (row:SysConfigTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
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