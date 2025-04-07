import request from '/@/utils/request'
// 查询应用管理列表
export function listAsApplication(query:object) {
  return request({
    url: '/api/v1/appsource/asApplication/list',
    method: 'get',
    params: query
  })
}
// 查询应用管理详细
export function getAsApplication(id:number) {
  return request({
    url: '/api/v1/appsource/asApplication/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增应用管理
export function addAsApplication(data:object) {
  return request({
    url: '/api/v1/appsource/asApplication/add',
    method: 'post',
    data: data
  })
}

// 修改应用管理
export function updateAsApplication(data:object) {
  return request({
    url: '/api/v1/appsource/asApplication/edit',
    method: 'put',
    data: data
  })
}

// 搜索应用管理
export function searchIconAsApplication(data:object) {
  return request({
    url: '/api/v1/appsource/asApplication/search',
    method: 'post',
    data: data
  })
}

// 删除应用管理
export function delAsApplication(ids:number[]) {
  return request({
    url: '/api/v1/appsource/asApplication/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}

// 解析应用
export function extractAsApplication(data:object) {
  return request({
    url: '/api/v1/appsource/asApplication/extract',
    method: 'post',
    data: data,
    timeout: 0,
  })
}

// 导入软件源
export function sourceImportAsApplication(data:object) {
  return request({
    url: '/api/v1/appsource/asApplication/sourceImport',
    method: 'post',
    data: data,
  })
}

// 应用管理付费修改
export function changeAsApplicationLock(id:number,lock:number) {
  const data = {
    id,
    lock
  }
  return request({
    url: '/api/v1/appsource/asApplication/changeLock',
    method: 'put',
    data:data
  })
}
// 应用管理是否启用修改
export function changeAsApplicationActive(id:number,active:number) {
  const data = {
    id,
    active
  }
  return request({
    url: '/api/v1/appsource/asApplication/changeActive',
    method: 'put',
    data:data
  })
}
