
import request from '@/utils/request'


export function getConfig(name) {
    return request({
        url: `/config/${name}`,
        method: 'get',
})
}
  

export function setConfig(name, data) {
    let val = {data:data};
return request({
    url: `/config/${name}`,
    method: 'put',
    data:val
})
}

