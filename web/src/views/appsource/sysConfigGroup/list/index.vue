<template>
  <div class="appsource-sysConfigGroup-container">
    <el-card shadow="hover">
        <div class="appsource-sysConfigGroup-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="主键" prop="id">
                    <el-input
                        v-model="tableData.param.id"
                        placeholder="请输入主键"
                        clearable                        
                        @keyup.enter.native="sysConfigGroupList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="配置组名称" prop="groupName">
                    <el-input
                        v-model="tableData.param.groupName"
                        placeholder="请输入配置组名称"
                        clearable                        
                        @keyup.enter.native="sysConfigGroupList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="sysConfigGroupList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="配置组键名" prop="groupKey">
                    <el-input
                        v-model="tableData.param.groupKey"
                        placeholder="请输入配置组键名"
                        clearable                        
                        @keyup.enter.native="sysConfigGroupList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建者" prop="createBy">
                    <el-input
                        v-model="tableData.param.createBy"
                        placeholder="请输入创建者"
                        clearable                        
                        @keyup.enter.native="sysConfigGroupList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="更新者" prop="updateBy">
                    <el-input
                        v-model="tableData.param.updateBy"
                        placeholder="请输入更新者"
                        clearable                        
                        @keyup.enter.native="sysConfigGroupList"
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
                    <el-button type="primary"  @click="sysConfigGroupList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/appsource/sysConfigGroup/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/appsource/sysConfigGroup/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/appsource/sysConfigGroup/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>            
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="主键" align="center" prop="id"
            min-width="150px"            
             />          
          <el-table-column label="配置组名称" align="center" prop="groupName"
            min-width="150px"            
             />          
          <el-table-column label="配置组键名" align="center" prop="groupKey"
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
                v-auth="'api/v1/appsource/sysConfigGroup/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/appsource/sysConfigGroup/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="sysConfigGroupList"
        />
    </el-card>
    <ApiV1AppsourceSysConfigGroupEdit
       ref="editRef"       
       @sysConfigGroupList="sysConfigGroupList"
    ></ApiV1AppsourceSysConfigGroupEdit>
    <ApiV1AppsourceSysConfigGroupDetail
      ref="detailRef"      
      @sysConfigGroupList="sysConfigGroupList"
    ></ApiV1AppsourceSysConfigGroupDetail>    
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
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
} from "/@/views/appsource/sysConfigGroup/list/component/model"
import ApiV1AppsourceSysConfigGroupEdit from "/@/views/appsource/sysConfigGroup/list/component/edit.vue"
import ApiV1AppsourceSysConfigGroupDetail from "/@/views/appsource/sysConfigGroup/list/component/detail.vue"
defineOptions({ name: "apiV1AppsourceSysConfigGroupList"})
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
const state = reactive<SysConfigGroupTableDataState>({
    ids:[],
    tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
            pageNum: 1,
            pageSize: 10,            
            id: undefined,            
            groupName: undefined,            
            groupKey: undefined,            
            createBy: undefined,            
            updateBy: undefined,            
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
    sysConfigGroupList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    sysConfigGroupList()
};
// 获取列表数据
const sysConfigGroupList = ()=>{
  loading.value = true
  listSysConfigGroup(state.tableData.param).then((res:any)=>{
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
const handleSelectionChange = (selection:Array<SysConfigGroupInfoData>) => {
    state.ids = selection.map(item => item.id)
    single.value = selection.length!=1
    multiple.value = !selection.length
}
const handleAdd =  ()=>{
    editRef.value.openDialog()
}
const handleUpdate = (row: SysConfigGroupTableColumns|null) => {
    if(!row){
        row = state.tableData.data.find((item:SysConfigGroupTableColumns)=>{
            return item.id ===state.ids[0]
        }) as SysConfigGroupTableColumns
    }
    editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SysConfigGroupTableColumns|null) => {
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
            delSysConfigGroup(id).then(()=>{
                ElMessage.success('删除成功');
                sysConfigGroupList();
            })
        })
        .catch(() => {});
}
const handleView = (row:SysConfigGroupTableColumns)=>{
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