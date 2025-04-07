<template>
  <div class="appsource-sysConfig-container" v-loading="loading">
    <el-tabs v-model="state.activeTab" type="border-card" :before-leave="onBeforeLeave">
      <el-tab-pane class="config-tab-pane" v-for="(groupName, groupKey) in state.configGroup" :key="groupKey"
                   :name="groupKey"
                   :label="groupName">
        <div class="config-form-item" v-for="(item, idx) in (state.config[groupKey] || {})" :key="item.configKey">
          <el-form-item v-if="item.elementType == 'string'" :label-position="state.labelPosition"
                        :label="item.configName + ' (' + item.configKey +') '"
                        prop="item.configName">
            <el-input v-model="item.configValue" :placeholder="item.remark"/>
          </el-form-item>

          <el-form-item v-if="item.elementType == 'switch'" :label-position="state.labelPosition"
                        :label="item.configName + ' (' + item.configKey +') '"
                        prop="item.configName">
            <el-switch
                v-model="item.configValue"
                inline-prompt
                :active-action-icon="Hide"
                :inactive-action-icon="View"

                size="large"
            />
          </el-form-item>

          <el-form-item v-if="item.elementType == 'textarea'" :label-position="state.labelPosition"
                        :label="item.configName + ' (' + item.configKey +') '"
                        prop="item.configName">
            <el-input v-model="item.configValue" :placeholder="item.remark" type="textarea" :autosize="{ minRows: 4, maxRows: 10}"/>
          </el-form-item>

          <el-form-item v-if="item.elementType == 'image'" :label-position="state.labelPosition"
                        :label="item.configName + ' (' + item.configKey +') '"
                        prop="item.configName">
            <el-input v-model="item.configValue" :placeholder="item.remark">
              <template #append>
                <el-image :src="item.configValue" style="max-width:64px;"/>
              </template>
            </el-input>
          </el-form-item>
          <div class="del-config-form-item">
            <el-popconfirm
                v-if="item.allowDel == 1"
                confirmButtonText="删除"
                title="确定要删除该配置项吗?"
            >
              <template #reference>
                <el-icon color="red" size="20" class="el-icon--delete">
                  <ele-DeleteFilled/>
                </el-icon>
              </template>
            </el-popconfirm>
          </div>
        </div>
        <div class="footer">
          <el-button type="primary" @click="onSubmit()">保存</el-button>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import {onMounted, reactive, ref, unref} from 'vue';
import {ElMessage} from "element-plus";
import {addConfig, editConfig} from "/@/api/system/config";
import {listSysConfig, updateSysConfig} from "/@/api/appsource/sysConfig";
import { Hide, View } from '@element-plus/icons-vue'

interface DicState {
  isShowDialog: boolean;
  labelPosition: string;
  activeTab: string;
  config: anyObj;
  configGroup: anyObj;
}

defineOptions({name: "apiV1AppsourceConfigList"})
const props = defineProps({
  sysYesNoOptions: {
    type: Array,
    default: () => []
  }
})
const emit = defineEmits(['dataList'])
const formRef = ref<HTMLElement | null>(null);
const loading = ref(false)
const state = reactive<DicState>({
  isShowDialog: false,
  labelPosition: "top",
  config: [],
  configGroup: {},
  activeTab: "appsource",
});
// 页面加载时
onMounted(() => {
  dataList()
});

const dataList = () => {
  listSysConfig({}).then((res: any) => {
    let list = res.data.list ?? [];
    let groups = res.data.groups ?? [];
    state.config = list;
    state.configGroup = groups;

    // 遍历数据，转换 switch 类型的 configValue
    Object.keys(state.config).forEach((groupKey) => {
      state.config[groupKey] = state.config[groupKey].map((item: any) => ({
        ...item,
        configValue: item.elementType === "switch" ? item.configValue === "1" : item.configValue
      }));
    });

    console.log(list);
    console.log(groups);

    for (const key in state.configGroup) {
      state.activeTab = key
      break
    }
  });
};

// 关闭弹窗
const closeDialog = () => {
  state.isShowDialog = false;
};
// 取消
const onCancel = () => {
  closeDialog();
};
// 修改
const onSubmit = () => {
  if (!state.activeTab || !state.config[state.activeTab]) {
    ElMessage.warning("没有可提交的数据");
    return;
  }

  // 只提交当前选中的 activeTab 分组
  let postData = state.config[state.activeTab].map((item: any) => ({
    configId: item.configId,
    configKey: item.configKey,
    configValue: item.elementType === "switch" ? (item.configValue ? 1 : 0) : item.configValue,
    elementType: item.elementType,
    group: item.group,
  }));

  // 发送请求
  updateSysConfig(postData).then(() => {
    ElMessage.success("修改成功");
    dataList(); // 重新加载数据
  }).catch(err => {
    console.error("提交失败", err);
    ElMessage.error("修改失败");
  });
};

const onBeforeLeave = (newTabName: string | number) => {
  if (newTabName == 'add_config') {
    state.showAddForm = true
    return false
  }
}

</script>

<style scoped lang="scss">
.footer {
  text-align: right;
}

.config-form-item {
  display: flex;
  align-items: center;

  .el-form-item {
    flex: 13;
  }

  .config-form-item-name {
    opacity: 0;
    flex: 3;
    font-size: 13px;
    color: var(--el-text-color-disabled);
    padding-left: 20px;
  }

  .del-config-form-item {
    cursor: pointer;
    flex: 1;
    padding-left: 10px;
  }

  .close-icon {
    display: none;
  }

  &:hover {
    .config-form-item-name {
      opacity: 1;
    }

    .close-icon {
      display: block;
      color: var(--el-text-color-disabled) !important;
    }
  }
}
</style>