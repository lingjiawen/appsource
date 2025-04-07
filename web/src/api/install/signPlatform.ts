import request from '/@/utils/request'
// 查询平台列表
export function listSignPlatform(query:object) {
  return request({
    url: '/api/v1/install/signPlatform/list',
    method: 'get',
    params: query
  })
}
// 查询平台详细
export function getSignPlatform(id:number) {
  return request({
    url: '/api/v1/install/signPlatform/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增平台
export function addSignPlatform(data:object) {
  return request({
    url: '/api/v1/install/signPlatform/add',
    method: 'post',
    data: data
  })
}
// 修改平台
export function updateSignPlatform(data:object) {
  return request({
    url: '/api/v1/install/signPlatform/edit',
    method: 'put',
    data: data
  })
}
// 删除平台
export function delSignPlatform(ids:number[]) {
  return request({
    url: '/api/v1/install/signPlatform/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 平台开启SSL修改
export function changeSignPlatformOpenSsl(id:number,openSsl:number) {
  const data = {
    id,
    openSsl
  }
  return request({
    url: '/api/v1/install/signPlatform/changeOpenSsl',
    method: 'put',
    data:data
  })
}
// 平台启用修改
export function changeSignPlatformStatus(id:number,status:number) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/v1/install/signPlatform/changeStatus',
    method: 'put',
    data:data
  })
}
