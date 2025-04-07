import request from '/@/utils/request'
// 查询设备管理列表
export function listSignDevice(query:object) {
  return request({
    url: '/api/v1/install/signDevice/list',
    method: 'get',
    params: query
  })
}
// 查询设备管理详细
export function getSignDevice(id:number) {
  return request({
    url: '/api/v1/install/signDevice/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增设备管理
export function addSignDevice(data:object) {
  return request({
    url: '/api/v1/install/signDevice/add',
    method: 'post',
    data: data
  })
}
// 修改设备管理
export function updateSignDevice(data:object) {
  return request({
    url: '/api/v1/install/signDevice/edit',
    method: 'put',
    data: data
  })
}
// 删除设备管理
export function delSignDevice(ids:number[]) {
  return request({
    url: '/api/v1/install/signDevice/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 设备管理禁用修改
export function changeSignDeviceActive(id:number,active:number) {
  const data = {
    id,
    active
  }
  return request({
    url: '/api/v1/install/signDevice/changeActive',
    method: 'put',
    data:data
  })
}
