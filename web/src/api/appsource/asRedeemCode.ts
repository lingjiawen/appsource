import request from '/@/utils/request'
// 查询卡密管理列表
export function listAsRedeemCode(query:object) {
  return request({
    url: '/api/v1/appsource/asRedeemCode/list',
    method: 'get',
    params: query
  })
}
// 查询卡密管理详细
export function getAsRedeemCode(id:number) {
  return request({
    url: '/api/v1/appsource/asRedeemCode/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增卡密管理
export function addAsRedeemCode(data:object) {
  return request({
    url: '/api/v1/appsource/asRedeemCode/add',
    method: 'post',
    data: data
  })
}
// 修改卡密管理
export function updateAsRedeemCode(data:object) {
  return request({
    url: '/api/v1/appsource/asRedeemCode/edit',
    method: 'put',
    data: data
  })
}
// 删除卡密管理
export function delAsRedeemCode(ids:number[]) {
  return request({
    url: '/api/v1/appsource/asRedeemCode/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 卡密管理是否激活修改
export function changeAsRedeemCodeActive(id:number,active:number) {
  const data = {
    id,
    active
  }
  return request({
    url: '/api/v1/appsource/asRedeemCode/changeActive',
    method: 'put',
    data:data
  })
}
