import request from '/@/utils/request'
// 查询帮助中心列表
export function listSignHelp(query:object) {
  return request({
    url: '/api/v1/install/signHelp/list',
    method: 'get',
    params: query
  })
}
// 查询帮助中心详细
export function getSignHelp(id:number) {
  return request({
    url: '/api/v1/install/signHelp/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增帮助中心
export function addSignHelp(data:object) {
  return request({
    url: '/api/v1/install/signHelp/add',
    method: 'post',
    data: data
  })
}
// 修改帮助中心
export function updateSignHelp(data:object) {
  return request({
    url: '/api/v1/install/signHelp/edit',
    method: 'put',
    data: data
  })
}
// 删除帮助中心
export function delSignHelp(ids:number[]) {
  return request({
    url: '/api/v1/install/signHelp/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 帮助中心默认展开修改
export function changeSignHelpIsExpand(id:number,isExpand:number) {
  const data = {
    id,
    isExpand
  }
  return request({
    url: '/api/v1/install/signHelp/changeIsExpand',
    method: 'put',
    data:data
  })
}
