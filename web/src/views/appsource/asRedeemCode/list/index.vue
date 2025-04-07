<template>
  <div class="appsource-asRedeemCode-container">
    <el-card shadow="hover">
        <div class="appsource-asRedeemCode-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>
              <el-col :span="8" class='colBlock'>
                <el-form-item label="设备码" prop="udid">
                  <el-input
                      v-model="tableData.param.udid"
                      placeholder="请输入设备码"
                      clearable
                      @keyup.enter.native="asRedeemCodeList"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="8" class="colBlock">
                  <el-form-item label="兑换码" prop="code">
                    <el-input
                        v-model="tableData.param.code"
                        placeholder="请输入兑换码"
                        clearable                        
                        @keyup.enter.native="asRedeemCodeList"
                    />                    
                  </el-form-item>
                </el-col>
              <el-col :span="8" class="colBlock">
                <el-form-item label="是否激活" prop="active">
                  <el-select v-model="tableData.param.active" placeholder="选择是否激活" style="width: 130px">
                    <el-option value="1" label="激活"></el-option>
                    <el-option value="0" label="未激活"></el-option>
                  </el-select>

                </el-form-item>
              </el-col>
              <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="asRedeemCodeList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                

                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="类型" prop="type">
                    <el-select filterable v-model="tableData.param.type" placeholder="请选择类型" clearable style="width:200px;">
                        <el-option
                            v-for="dict in appsource_redeem_code_type"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                

                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="激活时间" prop="activeAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.activeAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择激活时间"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="生成时间" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择生成时间"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="asRedeemCodeList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/appsource/asRedeemCode/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/appsource/asRedeemCode/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/appsource/asRedeemCode/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>             
             <el-col :span="1.5">
                <el-button
                        type="warning"
                        @click="handleExport()"
                        v-auth="'api/v1/appsource/asRedeemCode/export'"
                ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
             </el-col>            
                <el-col :span="1.5">
                    <el-button
                            type="success"
                            @click="handleImport()"
                            v-auth="'api/v1/appsource/asRedeemCode/import'"
                    ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
                </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="兑换码" align="center" prop="code"
            min-width="150px"            
             />          
          <el-table-column label="设备码" align="center" prop="udid"
            min-width="150px"            
             />          
          <el-table-column label="类型" align="center" prop="type" :formatter="typeFormat"
            min-width="150px"            
             />          
          <el-table-column label="是否激活" align="center" prop="active"
          min-width="150px"          
          >
              <template #default="scope">
                  <el-switch  v-model="scope.row.active"
                              :active-value=1
                              :inactive-value=0
                              class="ml-2"
                              @change="changeActive(scope.row)"/>
              </template>
          </el-table-column>          
          <el-table-column label="激活时间" align="center" prop="activeAt"
            min-width="150px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.activeAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>          
          <el-table-column label="生成时间" align="center" prop="createdAt"
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
                v-auth="'api/v1/appsource/asRedeemCode/get'"
              ><el-icon><ele-View /></el-icon>详情</el-button>              
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/appsource/asRedeemCode/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/appsource/asRedeemCode/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="asRedeemCodeList"
        />
    </el-card>
    <ApiV1AppsourceAsRedeemCodeEdit
       ref="editRef"       
       :typeOptions="appsource_redeem_code_type"       
       :activeOptions="sys_normal_disable"       
       @asRedeemCodeList="asRedeemCodeList"
    ></ApiV1AppsourceAsRedeemCodeEdit>
    <ApiV1AppsourceAsRedeemCodeAdd
        ref="addRef"
        :typeOptions="appsource_redeem_code_type"
        :activeOptions="sys_normal_disable"
        @asRedeemCodeList="asRedeemCodeList"
    ></ApiV1AppsourceAsRedeemCodeAdd>
    <ApiV1AppsourceAsRedeemCodeDetail
      ref="detailRef"      
      :typeOptions="appsource_redeem_code_type"      
      :activeOptions="sys_normal_disable"      
      @asRedeemCodeList="asRedeemCodeList"
    ></ApiV1AppsourceAsRedeemCodeDetail>    
    <loadExcel ref="loadExcelAsRedeemCodeRef" @getList="asRedeemCodeList"
               upUrl="api/v1/appsource/asRedeemCode/import"
               tplUrl="/api/v1/appsource/asRedeemCode/excelTemplate"></loadExcel>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listAsRedeemCode,
    getAsRedeemCode,
    delAsRedeemCode,
    addAsRedeemCode,
    updateAsRedeemCode,    
    changeAsRedeemCodeActive,    
} from "/@/api/appsource/asRedeemCode";
import {
    AsRedeemCodeTableColumns,
    AsRedeemCodeInfoData,
    AsRedeemCodeTableDataState,    
} from "/@/views/appsource/asRedeemCode/list/component/model"
import ApiV1AppsourceAsRedeemCodeAdd from "/@/views/appsource/asRedeemCode/list/component/add.vue"
import ApiV1AppsourceAsRedeemCodeEdit from "/@/views/appsource/asRedeemCode/list/component/edit.vue"
import ApiV1AppsourceAsRedeemCodeDetail from "/@/views/appsource/asRedeemCode/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1AppsourceAsRedeemCodeList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const addRef = ref();
const detailRef = ref();
const loadExcelAsRedeemCodeRef = ref();
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
    appsource_redeem_code_type,    
    sys_normal_disable,    
} = proxy.useDict(    
    'appsource_redeem_code_type',    
    'sys_normal_disable',    
)
const state = reactive<AsRedeemCodeTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            code: undefined,            
            udid: undefined,            
            type: undefined,            
            active: undefined,            
            activeAt: undefined,            
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
    asRedeemCodeList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    asRedeemCodeList()
};
// 获取列表数据
const asRedeemCodeList = ()=>{
  loading.value = true
  listAsRedeemCode(state.tableData.param).then((res:any)=>{
    let list = res.data.list??[];    
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
    showAll.value = !showAll.value;
}
// 类型字典翻译
const typeFormat = (row:AsRedeemCodeTableColumns) => {
    return proxy.selectDictLabel(appsource_redeem_code_type.value, row.type);
}
// 是否激活字典翻译
const activeFormat = (row:AsRedeemCodeTableColumns) => {
    return proxy.selectDictLabel(sys_normal_disable.value, row.active);
}

const changeActive = (row:AsRedeemCodeTableColumns) => {
    changeAsRedeemCodeActive(row.id,row.active)
        .catch(()=>{
            setTimeout(()=>{
                row.active = !row.active
            },300)
        })
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<AsRedeemCodeInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    addRef.value.openDialog()
}
const handleUpdate = (row: AsRedeemCodeTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:AsRedeemCodeTableColumns)=>{
            return item.id ===state.ids[0]
        }) as AsRedeemCodeTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: AsRedeemCodeTableColumns|null) => {
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
            delAsRedeemCode(id).then(()=>{
                ElMessage.success('删除成功');
                asRedeemCodeList();
            })
        })
        .catch(() => {});
}
const handleView = (row:AsRedeemCodeTableColumns)=>{
    detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
    downLoadXml('/api/v1/appsource/asRedeemCode/export',state.tableData.param,'get')
}
const handleImport=()=>{
    loadExcelAsRedeemCodeRef.value.open()
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