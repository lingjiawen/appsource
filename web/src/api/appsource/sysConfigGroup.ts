import request from '/@/utils/request'
// 查询配置组列表
export function listSysConfigGroup(query:object) {
  return request({
    url: '/api/v1/appsource/sysConfigGroup/list',
    method: 'get',
    params: query
  })
}
// 查询配置组详细
export function getSysConfigGroup(id:number) {
  return request({
    url: '/api/v1/appsource/sysConfigGroup/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增配置组
export function addSysConfigGroup(data:object) {
  return request({
    url: '/api/v1/appsource/sysConfigGroup/add',
    method: 'post',
    data: data
  })
}
// 修改配置组
export function updateSysConfigGroup(data:object) {
  return request({
    url: '/api/v1/appsource/sysConfigGroup/edit',
    method: 'put',
    data: data
  })
}
// 删除配置组
export function delSysConfigGroup(ids:number[]) {
  return request({
    url: '/api/v1/appsource/sysConfigGroup/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
