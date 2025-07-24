import axios from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    // 后台API返回格式为 {code, message, data}
    if (response.data && typeof response.data === 'object' && 'code' in response.data) {
      if (response.data.code === 0) {
        // 成功的情况下返回data字段
        return response.data.data || response.data;
      } else {
        // 错误的情况下抛出异常
        throw new Error(response.data.message || '请求失败');
      }
    }
    return response.data;
  },
  (error) => {
    console.error('API Error:', error);
    return Promise.reject(error);
  }
);

export default api;
