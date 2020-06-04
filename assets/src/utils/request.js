import axios from 'axios';

// 创建axios实例
const service = axios.create({
  baseURL: "/file-manager", // api的base_url
  timeout: 1000 * 60                  // 请求超时时间
});

// respone拦截器
import {Message} from 'element-ui'

service.interceptors.response.use(
  response => {
    if (response.data.code !== 200) {
      if (response.data.code === 401) {
        window.location.href = '/login'
      } else {
        Message({
          message: response.data.error,
          type: 'error',
          duration: 5 * 1000,
        })
        return Promise.reject('error')
      }
    } else {
      return response.data.data
    }
  },
  error => {
    const response = JSON.parse(JSON.stringify(error)).response
    if (response.status === 401) {
      window.location.href = '/login'
    } else {
      Message({
        message: error.toString(),
        type: 'error',
        duration: 5 * 1000,
      })
      return Promise.reject('error')
    }
  },
)

export default service;
