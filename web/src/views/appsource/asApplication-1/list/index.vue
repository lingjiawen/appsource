<template>
  <div class="appsource-asApplication-container">
    <el-card shadow="hover">
      <div class="appsource-asApplication-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
          <el-row>
            <el-col :span="8" class="colBlock">
              <el-form-item label="ID" prop="id">
                <el-input
                    v-model="tableData.param.id"
                    placeholder="请输入ID"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" class="colBlock">
              <el-form-item label="文件" prop="fileUrl">
                <el-select filterable v-model="tableData.param.fileUrl" placeholder="请选择文件" clearable style="width:200px;">
                  <el-option label="请选择字典生成" value="" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary"  @click="asApplicationList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                <el-button type="primary" link  @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                  <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                </el-button>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="应用名称" prop="name">
                <el-input
                    v-model="tableData.param.name"
                    placeholder="请输入应用名称"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="包名" prop="bundleId">
                <el-input
                    v-model="tableData.param.bundleId"
                    placeholder="请输入包名"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="版本号" prop="version">
                <el-input
                    v-model="tableData.param.version"
                    placeholder="请输入版本号"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="文件大小" prop="size">
                <el-input
                    v-model="tableData.param.size"
                    placeholder="请输入文件大小"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="类型" prop="type">
                <el-select filterable v-model="tableData.param.type" placeholder="请选择类型" clearable style="width:200px;">
                  <el-option
                      v-for="dict in appsource_app_type"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="描述" prop="description">
                <el-input
                    v-model="tableData.param.description"
                    placeholder="请输入描述"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="图标" prop="iconUrl">
                <el-input
                    v-model="tableData.param.iconUrl"
                    placeholder="请输入图标"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="付费" prop="lock">
                <el-switch  v-model="tableData.param.lock" class="ml-2" />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="是否蓝奏云链接" prop="lanzou">
                <el-select filterable v-model="tableData.param.lanzou" placeholder="请选择是否蓝奏云链接" clearable style="width:200px;">
                  <el-option
                      v-for="dict in sys_yes_no"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="权重" prop="weigh">
                <el-input
                    v-model="tableData.param.weigh"
                    placeholder="请输入权重"
                    clearable
                    @keyup.enter.native="asApplicationList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="是否启用" prop="active">
                <el-switch  v-model="tableData.param.active" class="ml-2" />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="备注" prop="note">
                <el-input
                    v-model="tableData.param.note"
                    placeholder="请输入备注"
                    clearable
                    @keyup.enter.native="asApplicationList"
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
                <el-button type="primary"  @click="asApplicationList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                v-auth="'api/v1/appsource/asApplication/add'"
            ><el-icon><ele-Plus /></el-icon>新增</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                :disabled="single"
                @click="handleUpdate(null)"
                v-auth="'api/v1/appsource/asApplication/edit'"
            ><el-icon><ele-Edit /></el-icon>修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="danger"
                :disabled="multiple"
                @click="handleDelete(null)"
                v-auth="'api/v1/appsource/asApplication/delete'"
            ><el-icon><ele-Delete /></el-icon>删除</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="warning"
                @click="handleExport()"
                v-auth="'api/v1/appsource/asApplication/export'"
            ><el-icon><ele-Download /></el-icon>导出Excel</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                @click="handleImport()"
                v-auth="'api/v1/appsource/asApplication/import'"
            ><el-icon><ele-Upload /></el-icon>导入Excel</el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="ID" align="center" prop="id"
                         min-width="150px"
        />
        <el-table-column label="软件地址" align="center" prop="fileUrl"
                         min-width="150px"
        />
        <el-table-column label="应用名称" align="center" prop="name"
                         min-width="150px"
        />
        <el-table-column label="包名" align="center" prop="bundleId"
                         min-width="150px"
        />
        <el-table-column label="版本号" align="center" prop="version"
                         min-width="150px"
        />
        <el-table-column label="文件大小" align="center" prop="size"
                         min-width="150px"
        />
        <el-table-column label="类型" align="center" prop="type" :formatter="typeFormat"
                         min-width="150px"
        />
        <el-table-column label="描述" align="center" prop="description"
                         min-width="150px"
        />
        <el-table-column label="图标" align="center" prop="iconUrl"
                         min-width="150px"
        />
        <el-table-column label="付费" align="center" prop="lock"
                         min-width="150px"
        >
          <template #default="scope">
            <el-switch  v-model="scope.row.lock" class="ml-2" @change="changeLock(scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="是否蓝奏云链接" align="center" prop="lanzou" :formatter="lanzouFormat"
                         min-width="150px"
        />
        <el-table-column label="权重" align="center" prop="weigh"
                         min-width="150px"
        />
        <el-table-column label="是否启用" align="center" prop="active"
                         min-width="150px"
        >
          <template #default="scope">
            <el-switch  v-model="scope.row.active" class="ml-2" @change="changeActive(scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="备注" align="center" prop="note"
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
                v-auth="'api/v1/appsource/asApplication/get'"
            ><el-icon><ele-View /></el-icon>详情</el-button>
            <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/appsource/asApplication/edit'"
            ><el-icon><ele-EditPen /></el-icon>修改</el-button>
            <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/appsource/asApplication/delete'"
            ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="asApplicationList"
      />
    </el-card>
    <ApiV1AppsourceAsApplicationEdit
        ref="editRef"
        :typeOptions="appsource_app_type"
        :lockOptions="sys_normal_disable"
        :lanzouOptions="sys_yes_no"
        :activeOptions="sys_normal_disable"
        @asApplicationList="asApplicationList"
    ></ApiV1AppsourceAsApplicationEdit>
    <ApiV1AppsourceAsApplicationDetail
        ref="detailRef"
        :typeOptions="appsource_app_type"
        :lockOptions="sys_normal_disable"
        :lanzouOptions="sys_yes_no"
        :activeOptions="sys_normal_disable"
        @asApplicationList="asApplicationList"
    ></ApiV1AppsourceAsApplicationDetail>
    <loadExcel ref="loadExcelAsApplicationRef" @getList="asApplicationList"
               upUrl="api/v1/appsource/asApplication/import"
               tplUrl="/api/v1/appsource/asApplication/excelTemplate"></loadExcel>
  </div>
</template>
<script setup lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
  listAsApplication,
  getAsApplication,
  delAsApplication,
  addAsApplication,
  updateAsApplication,
  changeAsApplicationLock,
  changeAsApplicationActive,
} from "/@/api/appsource/asApplication";
import {
  AsApplicationTableColumns,
  AsApplicationInfoData,
  AsApplicationTableDataState,
} from "/@/views/appsource/asApplication/list/component/model"
import ApiV1AppsourceAsApplicationEdit from "/@/views/appsource/asApplication/list/component/edit.vue"
import ApiV1AppsourceAsApplicationDetail from "/@/views/appsource/asApplication/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"
defineOptions({ name: "apiV1AppsourceAsApplicationList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const detailRef = ref();
const loadExcelAsApplicationRef = ref();
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
  appsource_app_type,
  sys_normal_disable,
  sys_yes_no,
} = proxy.useDict(
    'appsource_app_type',
    'sys_normal_disable',
    'sys_yes_no',
)
const state = reactive<AsApplicationTableDataState>({
  ids:[],
  tableData: {
    data: [],
    total: 0,
    loading: false,
    param: {
      pageNum: 1,
      pageSize: 10,
      id: undefined,
      fileUrl: undefined,
      name: undefined,
      bundleId: undefined,
      version: undefined,
      size: undefined,
      type: undefined,
      description: undefined,
      iconUrl: undefined,
      lock: undefined,
      lanzou: undefined,
      weigh: undefined,
      active: undefined,
      note: undefined,
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
  asApplicationList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  asApplicationList()
};
// 获取列表数据
const asApplicationList = ()=>{
  loading.value = true
  listAsApplication(state.tableData.param).then((res:any)=>{
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
const typeFormat = (row:AsApplicationTableColumns) => {
  return proxy.selectDictLabel(appsource_app_type.value, row.type);
}
// 付费字典翻译
const lockFormat = (row:AsApplicationTableColumns) => {
  return proxy.selectDictLabel(sys_normal_disable.value, row.lock);
}
const changeLock = (row:AsApplicationTableColumns) => {
  changeAsApplicationLock(row.id,row.lock)
      .catch(()=>{
        setTimeout(()=>{
          row.lock = !row.lock
        },300)
      })
}
// 是否蓝奏云链接字典翻译
const lanzouFormat = (row:AsApplicationTableColumns) => {
  return proxy.selectDictLabel(sys_yes_no.value, row.lanzou);
}
// 是否启用字典翻译
const activeFormat = (row:AsApplicationTableColumns) => {
  return proxy.selectDictLabel(sys_normal_disable.value, row.active);
}
const changeActive = (row:AsApplicationTableColumns) => {
  changeAsApplicationActive(row.id,row.active)
      .catch(()=>{
        setTimeout(()=>{
          row.active = !row.active
        },300)
      })
}
// 多选框选中数据
const handleSelectionChange = (selection:Array<AsApplicationInfoData>) => {
  state.ids = selection.map(item => item.id)
  single.value = selection.length!=1
  multiple.value = !selection.length
}
const handleAdd =  ()=>{
  editRef.value.openDialog()
}
const handleUpdate = (row: AsApplicationTableColumns|null) => {
  if(!row){
    row = state.tableData.data.find((item:AsApplicationTableColumns)=>{
      return item.id ===state.ids[0]
    }) as AsApplicationTableColumns
  }
  editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: AsApplicationTableColumns|null) => {
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
        delAsApplication(id).then(()=>{
          ElMessage.success('删除成功');
          asApplicationList();
        })
      })
      .catch(() => {});
}
const handleView = (row:AsApplicationTableColumns)=>{
  detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = ()=>{
  downLoadXml('/api/v1/appsource/asApplication/export',state.tableData.param,'get')
}
const handleImport=()=>{
  loadExcelAsApplicationRef.value.open()
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