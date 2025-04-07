import request from '/@/utils/request'
// 查询配置项列表
export function listSysConfig(query:object) {
  return request({
    url: '/api/v1/appsource/sysConfig/list',
    method: 'get',
    params: query
  })
}
// 查询配置项详细
export function getSysConfig(configId:number) {
  return request({
    url: '/api/v1/appsource/sysConfig/get',
    method: 'get',
    params: {
      configId: configId.toString()
    }
  })
}
// 新增配置项
export function addSysConfig(data:object) {
  return request({
    url: '/api/v1/appsource/sysConfig/add',
    method: 'post',
    data: data
  })
}
// 修改配置项
export function updateSysConfig(data:object) {
  return request({
    url: '/api/v1/appsource/sysConfig/edit',
    method: 'put',
    data: {
      "configs": data
    }
  })
}
// 删除配置项
export function delSysConfig(configIds:number[]) {
  return request({
    url: '/api/v1/appsource/sysConfig/delete',
    method: 'delete',
    data:{
      configIds:configIds
    }
  })
}
//相关连表查询数据
export function linkedDataSearch(){
  return request({
    url: '/api/v1/appsource/sysConfig/linkedData',
    method: 'get'
  })
}
