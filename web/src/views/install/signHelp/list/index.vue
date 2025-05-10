<template>
  <div class="install-signHelp-container">
    <el-card shadow="hover">
        <div class="install-signHelp-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row class="search-fields-container">                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="ID" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入ID"
                        clearable                        
                        @keyup.enter.native="signHelpList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="标题" prop="title">
                    <el-input
                        v-model="tableData.param.title"
                        placeholder="请输入标题"
                        clearable                        
                        @keyup.enter.native="signHelpList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="signHelpList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="内容" prop="content">
                    <el-select filterable v-model="tableData.param.content" placeholder="请选择内容" clearable style="width:200px;">
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="默认展开" prop="isExpand">
                      <el-switch  v-model="tableData.param.isExpand" class="ml-2" />
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="权重" prop="weigh">
                    <el-input
                        v-model="tableData.param.weigh"
                        placeholder="请输入权重"
                        clearable                        
                        @keyup.enter.native="signHelpList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建日期" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择创建日期"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="signHelpList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
            <el-row :gutter="10" class="btn-container">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd"
                  v-auth="'api/v1/install/signHelp/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/install/signHelp/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/install/signHelp/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/install/signHelp/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/install/signHelp/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="标题" align="center" prop="title"
            min-width="150px"            
             />          
          <el-table-column label="内容" align="center" prop="content"
            min-width="150px"            
             />          
          <el-table-column label="默认展开" align="center" prop="isExpand"
          min-width="150px"          
          >
              <template #default="scope">
                  <el-switch  :active-value=1 :inactive-value=0 v-model="scope.row.isExpand" class="ml-2" @change="changeIsExpand(scope.row)"/>
              </template>
          </el-table-column>          
          <el-table-column label="权重" align="center" prop="weigh"
            min-width="150px"            
             />          
          <el-table-column label="创建日期" align="center" prop="createdAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="operation-container small-padding" min-width="80px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/install/signHelp/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/install/signHelp/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="signHelpList"
        />
    </el-card>
    <ApiV1InstallSignHelpEdit
       ref="editRef"       
       @signHelpList="signHelpList"
    ></ApiV1InstallSignHelpEdit>
    <ApiV1InstallSignHelpDetail
      ref="detailRef"      
      @signHelpList="signHelpList"
    ></ApiV1InstallSignHelpDetail>    
    <loadExcel ref="loadExcelSignHelpRef" @getList="signHelpList"
               upUrl="api/v1/install/signHelp/import"
               tplUrl="/api/v1/install/signHelp/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listSignHelp,
    getSignHelp,
    delSignHelp,
    addSignHelp,
    updateSignHelp,    
    changeSignHelpIsExpand,    
} from "/@/api/install/signHelp";
import {
    SignHelpTableColumns,
    SignHelpInfoData,
    SignHelpTableDataState,    
} from "/@/views/install/signHelp/list/component/model"
import ApiV1InstallSignHelpEdit from "/@/views/install/signHelp/list/component/edit.vue"
import ApiV1InstallSignHelpDetail from "/@/views/install/signHelp/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1InstallSignHelpList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelSignHelpRef = ref();
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
const state = reactive<SignHelpTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            title: undefined,            
            content: undefined,            
            isExpand: undefined,            
            weigh: undefined,            
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
    signHelpList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    signHelpList()
};
// 获取列表数据
const signHelpList = ()=>{
  loading.value = true
  listSignHelp(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
const changeIsExpand = (row:SignHelpTableColumns) => {
    changeSignHelpIsExpand(row.id,row.isExpand)
        .catch(()=>{
            setTimeout(()=>{
                row.isExpand = !row.isExpand
            },300)
        })
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<SignHelpInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: SignHelpTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:SignHelpTableColumns)=>{
            return item.id ===state.ids[0]
        }) as SignHelpTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SignHelpTableColumns|null) => {
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
            delSignHelp(id).then(()=>{
                ElMessage.success('删除成功');
                signHelpList();
            })
        })
        .catch(() => {});
}
const handleView = (row:SignHelpTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/install/signHelp/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelSignHelpRef.value.open()
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