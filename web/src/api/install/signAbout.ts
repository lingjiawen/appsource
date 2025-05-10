import request from '/@/utils/request'
// 查询关于我们列表
export function listSignAbout(query:object) {
  return request({
    url: '/api/v1/install/signAbout/list',
    method: 'get',
    params: query
  })
}
// 查询关于我们详细
export function getSignAbout(id:number) {
  return request({
    url: '/api/v1/install/signAbout/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增关于我们
export function addSignAbout(data:object) {
  return request({
    url: '/api/v1/install/signAbout/add',
    method: 'post',
    data: data
  })
}
// 修改关于我们
export function updateSignAbout(data:object) {
  return request({
    url: '/api/v1/install/signAbout/edit',
    method: 'put',
    data: data
  })
}
// 删除关于我们
export function delSignAbout(ids:number[]) {
  return request({
    url: '/api/v1/install/signAbout/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 关于我们是否链接修改
export function changeSignAboutIsLink(id:number,isLink:number) {
  const data = {
    id,
    isLink
  }
  return request({
    url: '/api/v1/install/signAbout/changeIsLink',
    method: 'put',
    data:data
  })
}
