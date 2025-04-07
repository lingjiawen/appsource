import request from '/@/utils/request'
// 查询签名卡密列表
export function listSignRedeemCode(query:object) {
  return request({
    url: '/api/v1/install/signRedeemCode/list',
    method: 'get',
    params: query
  })
}
// 查询签名卡密详细
export function getSignRedeemCode(id:number) {
  return request({
    url: '/api/v1/install/signRedeemCode/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增签名卡密
export function addSignRedeemCode(data:object) {
  return request({
    url: '/api/v1/install/signRedeemCode/add',
    method: 'post',
    data: data
  })
}
// 修改签名卡密
export function updateSignRedeemCode(data:object) {
  return request({
    url: '/api/v1/install/signRedeemCode/edit',
    method: 'put',
    data: data
  })
}
// 删除签名卡密
export function delSignRedeemCode(ids:number[]) {
  return request({
    url: '/api/v1/install/signRedeemCode/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 签名卡密禁用修改
export function changeSignRedeemCodeBanned(id:number,banned:number) {
  const data = {
    id,
    banned
  }
  return request({
    url: '/api/v1/install/signRedeemCode/changeBanned',
    method: 'put',
    data:data
  })
}
// 签名卡密激活修改
export function changeSignRedeemCodeActive(id:number,active:number) {
  const data = {
    id,
    active
  }
  return request({
    url: '/api/v1/install/signRedeemCode/changeActive',
    method: 'put',
    data:data
  })
}
