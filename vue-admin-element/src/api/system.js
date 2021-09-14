
import request from '@/utils/request'

export function getAbout() {
    return request({
      url: '/system/about',
      method: 'get',
    })
  }
   
  

  export function GetLicense() {
    return request({
        url: '/system/license',
        method: 'get',
    })
  }
  export function SetLicense(data) {
    return request({
        url: '/system/license',
        method: 'put',
        data
    })
  }
  
export function PostRestart() {
    return request({
        url: '/system/restart',
        method: 'post',
    })
}


