import request from '@/utils/request'


export function getList(params) {
  return request({
    url: '/users',
    method: 'get',
    params
  })
}



export function addUser(data) {
  return request({
    url: '/user',
    method: 'post',
    data
  })
}


export function setUser(data) {
  return request({
    url: `/user/${data.username}`,
    method: 'put',
    data
  })
}


export function delUser(id) {
  return request({
    url: `/user/${id}`,
    method: 'delete',
  })
}


