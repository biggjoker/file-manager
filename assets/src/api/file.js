import request from '@/utils/request'

export function getDic(path) {
  if (path === "/"){
    path = ''
  }
  return request({
    url: '/dir/dir',
    method: 'get',
    params: {
      path
    }
  })
}

export function createDic(dir_path) {
  return request({
    url: '/dir/dir',
    method: 'post',
    data: {
      dir_path
    }
  })
}

export function delDic(dir_path, force = false) {
  return request({
    url: '/dir/dir',
    method: 'delete',
    data: {
      dir_path,
      force
    }
  })
}

export function renameDic(newname, oldname) {
  return request({
    url: '/dir/dir',
    method: 'put',
    data: {
      newname,
      oldname
    }
  })
}

export function readFile(name) {
  return request({
    url: '/dir/file',
    method: 'get',
    params: {
      name
    }
  })
}

export function saveFile(name, content) {
  return request({
    url: '/dir/file',
    method: 'post',
    data: {
      name,
      content
    }
  })
}
