<template>
  <div class="install-signRedeemCode-container">
    <el-card shadow="hover">
      <div class="install-signRedeemCode-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
          <el-row class="search-fields-container">

            <el-col :span="8">
              <el-form-item label="UDID" prop="udid">
                <el-input
                    v-model="tableData.param.udid"
                    placeholder="请输入UDID"
                    clearable
                    @keyup.enter.native="signRedeemCodeList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="时效类型" prop="accountType">
                <el-select filterable v-model="tableData.param.accountType" placeholder="秒出/审核" clearable
                           style="min-width:170px;">
                  <el-option
                      v-for="dict in apple_account_type"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8" class="colBlock">
              <el-form-item label="兑换码" prop="code">
                <el-input
                    v-model="tableData.param.code"
                    placeholder="请输入兑换码"
                    clearable
                    @keyup.enter.native="signRedeemCodeList"
                />
              </el-form-item>
            </el-col>

            <el-col :span="8">
              <el-form-item label="售后类型" prop="warrantyType">
                <el-select filterable v-model="tableData.param.warrantyType" placeholder="请选择售后类型" clearable
                           style="min-width:170px;">
                  <el-option
                      v-for="dict in apple_warranty_type"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="8">
              <el-form-item label="出书方式" prop="pool">
                <el-select filterable v-model="tableData.param.pool" placeholder="请选择出书方式" clearable
                           style="min-width:170px;">
                  <el-option
                      v-for="dict in apple_pool_type"
                      :key="dict.value"
                      :label="dict.label"
                      :value="dict.value"
                  />
                </el-select>
              </el-form-item>
            </el-col>

            <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
              <el-form-item class="search-container">
                <el-button type="primary" @click="signRedeemCodeList">
                  <el-icon>
                    <ele-Search/>
                  </el-icon>
                  搜索
                </el-button>
                <el-button @click="resetQuery(queryRef)">
                  <el-icon>
                    <ele-Refresh/>
                  </el-icon>
                  重置
                </el-button>
                <el-button type="primary" link @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll">
                    <ele-ArrowUp/>
                  </el-icon>
                  <el-icon v-show="!showAll">
                    <ele-ArrowDown/>
                  </el-icon>
                </el-button>
              </el-form-item>
            </el-col>

            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="证书标识" prop="certId">
                <el-input
                    v-model="tableData.param.certId"
                    placeholder="请输入证书标识"
                    clearable
                    @keyup.enter.native="signRedeemCodeList"
                />
              </el-form-item>
            </el-col>

            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="设备类型" prop="deviceType">
                <el-select filterable v-model="tableData.param.deviceType" placeholder="请选择设备类型" clearable
                           style="min-width:170px;">
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
              <el-form-item label="对接平台" prop="apiPlatform">
                <el-select filterable v-model="tableData.param.apiPlatform" placeholder="请选择对接平台" clearable
                           style="min-width:170px;">
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
                <el-select filterable v-model="tableData.param.apiWarrantyType" placeholder="请选择对接售后类型"
                           clearable style="min-width:170px;">
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
              <el-form-item label="备注" prop="note">
                <el-input
                    v-model="tableData.param.note"
                    placeholder="请输入备注"
                    clearable
                    @keyup.enter.native="signRedeemCodeList"
                />
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="禁用" prop="banned">
                <el-switch v-model="tableData.param.banned" class="ml-2"/>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="激活" prop="active">
                <el-switch v-model="tableData.param.active" class="ml-2"/>
              </el-form-item>
            </el-col>
            <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary" @click="signRedeemCodeList">
                  <el-icon>
                    <ele-Search/>
                  </el-icon>
                  搜索
                </el-button>
                <el-button @click="resetQuery(queryRef)">
                  <el-icon>
                    <ele-Refresh/>
                  </el-icon>
                  重置
                </el-button>
                <el-button type="primary" link @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll">
                    <ele-ArrowUp/>
                  </el-icon>
                  <el-icon v-show="!showAll">
                    <ele-ArrowDown/>
                  </el-icon>
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
                v-auth="'api/v1/install/signRedeemCode/add'"
            >
              <el-icon>
                <ele-Plus/>
              </el-icon>
              新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                :disabled="single"
                @click="handleUpdate(null)"
                v-auth="'api/v1/install/signRedeemCode/edit'"
            >
              <el-icon>
                <ele-Edit/>
              </el-icon>
              修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="danger"
                :disabled="multiple"
                @click="handleDelete(null)"
                v-auth="'api/v1/install/signRedeemCode/delete'"
            >
              <el-icon>
                <ele-Delete/>
              </el-icon>
              删除
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="warning"
                @click="handleExport()"
                v-auth="'api/v1/install/signRedeemCode/export'"
            >
              <el-icon>
                <ele-Download/>
              </el-icon>
              导出Excel
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
                type="success"
                @click="handleImport()"
                v-auth="'api/v1/install/signRedeemCode/import'"
            >
              <el-icon>
                <ele-Upload/>
              </el-icon>
              导入Excel
            </el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="兑换码" align="center" prop="code"
                         min-width="200px"
        />
        <el-table-column label="设备码" align="center" prop="udid"
                         min-width="250px"
        />
        <el-table-column label="证书标识" align="center" prop="certId"
                         min-width="100px"
        />
        <el-table-column label="时效类型" align="center" prop="accountType" :formatter="accountTypeFormat"
                         min-width="100px"
        />
        <el-table-column label="售后类型" align="center" prop="warrantyType" :formatter="warrantyTypeFormat"
                         min-width="100px"
        />
        <el-table-column label="设备类型" align="center" prop="deviceType" :formatter="deviceTypeFormat"
                         min-width="100px"
        />
        <el-table-column label="出书方式" align="center" prop="pool" :formatter="poolFormat"
                         min-width="100px"
        />
        <el-table-column label="对接平台" align="center" prop="apiPlatform" :formatter="apiPlatformFormat"
                         min-width="100px"
        />
        <el-table-column label="对接售后" align="center" prop="apiWarrantyType" :formatter="apiWarrantyTypeFormat"
                         min-width="100px"
        />
        <el-table-column label="备注" align="center" prop="note"
                         min-width="150px"
        />

        <el-table-column label="禁用" align="center" prop="banned"
                         min-width="150px"
        >
          <template #default="scope">
            <el-switch v-model="scope.row.banned"
                       class="ml-2"
                       :active-value=1
                       :inactive-value=0
                       @change="changeBanned(scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="激活" align="center" prop="active"
                         min-width="100px"
        >
          <template #default="scope">
            <el-switch v-model="scope.row.active"
                       :active-value=1
                       :inactive-value=0
                       class="ml-2"
                       @change="changeActive(scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="激活时间" align="center" prop="activeAt"
                         min-width="180px"
        >
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.activeAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" align="center" prop="createdAt"
                         min-width="180px"
        >
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" class-name="operation-container small-padding" min-width="80px"
                         fixed="right">
          <template #default="scope">
            <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/install/signRedeemCode/edit'"
            >
              <el-icon>
                <ele-EditPen/>
              </el-icon>
              修改
            </el-button>
            <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/install/signRedeemCode/delete'"
            >
              <el-icon>
                <ele-DeleteFilled/>
              </el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
          v-show="tableData.total>0"
          :total="tableData.total"
          v-model:page="tableData.param.pageNum"
          v-model:limit="tableData.param.pageSize"
          @pagination="signRedeemCodeList"
      />
    </el-card>
    <ApiV1InstallSignRedeemCodeAdd
        ref="addRef"
        :accountTypeOptions="apple_account_type"
        :warrantyTypeOptions="apple_warranty_type"
        :deviceTypeOptions="apple_device_type"
        :poolOptions="apple_pool_type"
        :apiPlatformOptions="api_platform_type"
        @signRedeemCodeList="signRedeemCodeList"
    ></ApiV1InstallSignRedeemCodeAdd>
    <ApiV1InstallSignRedeemCodeEdit
        ref="editRef"
        :accountTypeOptions="apple_account_type"
        :warrantyTypeOptions="apple_warranty_type"
        :deviceTypeOptions="apple_device_type"
        :poolOptions="apple_pool_type"
        :apiPlatformOptions="api_platform_type"
        :apiWarrantyTypeOptions="apple_warranty_type"
        @signRedeemCodeList="signRedeemCodeList"
    ></ApiV1InstallSignRedeemCodeEdit>
    <ApiV1InstallSignRedeemCodeDetail
        ref="detailRef"
        :accountTypeOptions="apple_account_type"
        :warrantyTypeOptions="apple_warranty_type"
        :deviceTypeOptions="apple_device_type"
        :poolOptions="apple_pool_type"
        :apiPlatformOptions="api_platform_type"
        :apiWarrantyTypeOptions="apple_warranty_type"
        @signRedeemCodeList="signRedeemCodeList"
    ></ApiV1InstallSignRedeemCodeDetail>
    <loadExcel ref="loadExcelSignRedeemCodeRef" @getList="signRedeemCodeList"
               upUrl="api/v1/install/signRedeemCode/import"
               tplUrl="/api/v1/install/signRedeemCode/excelTemplate"></loadExcel>
  </div>
</template>
<script setup lang="ts">
import {computed, getCurrentInstance, onMounted, reactive, ref, toRaw, toRefs} from 'vue';
import {ElMessage, ElMessageBox, FormInstance} from 'element-plus';
import {
  changeSignRedeemCodeActive,
  changeSignRedeemCodeBanned,
  delSignRedeemCode,
  listSignRedeemCode,
} from "/@/api/install/signRedeemCode";
import {SignRedeemCodeInfoData, SignRedeemCodeTableColumns,} from "/@/views/install/signRedeemCode/list/component/model"
import ApiV1InstallSignRedeemCodeAdd from "/@/views/install/signRedeemCode/list/component/add.vue"
import ApiV1InstallSignRedeemCodeEdit from "/@/views/install/signRedeemCode/list/component/edit.vue"
import ApiV1InstallSignRedeemCodeDetail from "/@/views/install/signRedeemCode/list/component/detail.vue"
import {downLoadXml} from "/@/utils/zipdownload";
import loadExcel from "/@/components/loadExcel/index.vue"

defineOptions({name: "apiV1InstallSignRedeemCodeList"})
const {proxy} = <any>getCurrentInstance()
const loading = ref(false)
const queryRef = ref()
const editRef = ref();
const addRef = ref();
const detailRef = ref();
const loadExcelSignRedeemCodeRef = ref();
// 是否显示所有搜索选项
const showAll = ref(false)
// 非单个禁用
const single = ref(true)
// 非多个禁用
const multiple = ref(true)
const word = computed(() => {
  if (showAll.value === false) {
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
  apple_pool_type,
  api_platform_type,
} = proxy.useDict(
    'apple_account_type',
    'apple_warranty_type',
    'apple_device_type',
    'apple_pool_type',
    'api_platform_type',
)
const state = reactive<SignRedeemCodeAddState>({
  ids: [],
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
      certId: undefined,
      accountType: undefined,
      warrantyType: undefined,
      deviceType: undefined,
      pool: undefined,
      apiPlatform: undefined,
      apiWarrantyType: undefined,
      note: undefined,
      banned: undefined,
      active: undefined,
      activeAt: undefined,
      createdBy: undefined,
      createdAt: undefined,
      dateRange: []
    },
  },
});
const {tableData} = toRefs(state);
// 页面加载时
onMounted(() => {
  initTableData();
});
// 初始化表格数据
const initTableData = () => {
  signRedeemCodeList()
};
/** 重置按钮操作 */
const resetQuery = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
  signRedeemCodeList()
};
// 获取列表数据
const signRedeemCodeList = () => {
  loading.value = true
  listSignRedeemCode(state.tableData.param).then((res: any) => {
    let list = res.data.list ?? [];
    state.tableData.data = list;
    state.tableData.total = res.data.total;
    loading.value = false
  })
};
const toggleSearch = () => {
  showAll.value = !showAll.value;
}
// 时效类型字典翻译
const accountTypeFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(apple_account_type.value, row.accountType);
}
// 售后类型字典翻译
const warrantyTypeFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(apple_warranty_type.value, row.warrantyType);
}
// 设备类型字典翻译
const deviceTypeFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(apple_device_type.value, row.deviceType);
}
// 出书方式字典翻译
const poolFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(apple_pool_type.value, row.pool);
}
// 对接平台字典翻译
const apiPlatformFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(api_platform_type.value, row.apiPlatform);
}
// 对接售后类型字典翻译
const apiWarrantyTypeFormat = (row: SignRedeemCodeTableColumns) => {
  return proxy.selectDictLabel(apple_warranty_type.value, row.apiWarrantyType);
}
const changeBanned = (row: SignRedeemCodeTableColumns) => {
  changeSignRedeemCodeBanned(row.id, row.banned)
      .catch(() => {
        setTimeout(() => {
          row.banned = !row.banned
        }, 300)
      })
}
const changeActive = (row: SignRedeemCodeTableColumns) => {
  changeSignRedeemCodeActive(row.id, row.active)
      .catch(() => {
        setTimeout(() => {
          row.active = !row.active
        }, 300)
      })
}
// 多选框选中数据
const handleSelectionChange = (selection: Array<SignRedeemCodeInfoData>) => {
  state.ids = selection.map(item => item.id)
  single.value = selection.length != 1
  multiple.value = !selection.length
}
const handleAdd = () => {
  addRef.value.openDialog()
}
const handleUpdate = (row: SignRedeemCodeTableColumns | null) => {
  if (!row) {
    row = state.tableData.data.find((item: SignRedeemCodeTableColumns) => {
      return item.id === state.ids[0]
    }) as SignRedeemCodeTableColumns
  }
  editRef.value.openDialog(toRaw(row));
};
const handleDelete = (row: SignRedeemCodeTableColumns | null) => {
  let msg = '你确定要删除所选数据？';
  let id: number[] = [];
  if (row) {
    msg = `此操作将永久删除数据，是否继续?`
    id = [row.id]
  } else {
    id = state.ids
  }
  if (id.length === 0) {
    ElMessage.error('请选择要删除的数据。');
    return
  }
  ElMessageBox.confirm(msg, '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        delSignRedeemCode(id).then(() => {
          ElMessage.success('删除成功');
          signRedeemCodeList();
        })
      })
      .catch(() => {
      });
}
const handleView = (row: SignRedeemCodeTableColumns) => {
  detailRef.value.openDialog(toRaw(row));
}
//导出excel
const handleExport = () => {
  downLoadXml('/api/v1/install/signRedeemCode/export', state.tableData.param, 'get')
}
const handleImport = () => {
  loadExcelSignRedeemCodeRef.value.open()
}
</script>
<style lang="scss" scoped>
.colBlock {
  display: block;
}

.colNone {
  display: none;
}

.ml-2 {
  margin: 3px;
}

.operation-container {
  .el-button {
    margin: 0px;
  }
}

.btn-container {
  .el-button {
    margin-bottom: 10px;
  }
}

//:deep(.search-container) {
//  .el-form-item__content {
//    display: flex;
//    justify-content: center!important;
//    align-items: center;
//
//    .el-button {
//      margin: 5px;
//    }
//  }
//}

:deep(.search-fields-container) {
  .el-col {
    flex: 1 0 33%; /* 每个列占 33% */
    padding: 8px;

    .el-form-item {
      width: 100%;

      .el-form-item__content {
        display: flex;
        justify-content: center!important;
        align-items: center;

        .el-button {
          margin: 5px;
        }
      }
    }
  }
}

/* 手机屏幕下每个列占满一行 */
@media (min-width: 768px) and (max-width: 1199px) {
  :deep(.search-fields-container) {
    .el-col {
      flex: 1 0 50%; /* 每个列占满 100% */
      max-width: unset;
      padding: 0;

      .el-form-item {
        width: 100%;
        .el-form-item__content {
          display: flex;
          justify-content: right!important;
          align-items: center;

          .el-button {
            margin: 5px;
          }
        }
      }
    }

  }
}

/* 手机屏幕下每个列占满一行 */
@media (max-width: 767px) {
  :deep(.search-fields-container) {
    .el-col {
      flex: 1 0 100%; /* 每个列占满 100% */
      max-width: unset;
      padding: 0;

      .el-form-item {
        width: 100%;
        .el-form-item__content {
          display: flex;
          justify-content: right!important;
          align-items: center;

          .el-button {
            margin: 5px;
          }
        }
      }
    }

  }
}

</style>