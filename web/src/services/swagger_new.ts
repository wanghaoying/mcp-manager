import api from './api';
import type { 
  SwaggerValidationResult, 
  SwaggerParseResult, 
  SwaggerTextRequest, 
  APIEndpoint 
} from '../types/swagger';

export const swaggerService = {
  // 通过文件校验Swagger文档
  async validateFile(file: File): Promise<SwaggerValidationResult> {
    const formData = new FormData();
    formData.append('file', file);
    
    try {
      await api.post('/swagger/validate/file', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      
      return {
        valid: true,
        message: '校验成功',
        content: await file.text()
      };
    } catch (error: any) {
      return {
        valid: false,
        message: error.message || '文件校验失败'
      };
    }
  },

  // 通过文本校验Swagger文档
  async validateText(content: string): Promise<SwaggerValidationResult> {
    const request: SwaggerTextRequest = { content };
    
    try {
      await api.post('/swagger/validate/text', request);
      
      return {
        valid: true,
        message: '校验成功',
        content: content
      };
    } catch (error: any) {
      return {
        valid: false,
        message: error.message || '内容校验失败'
      };
    }
  },

  // 解析并保存Swagger接口
  async parseAndSave(content: string): Promise<SwaggerParseResult> {
    const request: SwaggerTextRequest = { content };
    
    try {
      const response = await api.post('/swagger/parse', request);
      
      return {
        success: true,
        message: '解析成功',
        endpoints: (response as unknown as APIEndpoint[]) || []
      };
    } catch (error: any) {
      throw new Error(error.message || '解析失败');
    }
  },

  // 获取API接口列表
  async getEndpoints(swaggerId: number): Promise<APIEndpoint[]> {
    try {
      const response = await api.get('/swagger/endpoints', {
        params: { swagger_id: swaggerId }
      });
      
      return (response as unknown as APIEndpoint[]) || [];
    } catch (error: any) {
      throw new Error(error.message || '获取接口列表失败');
    }
  },

  // 根据ID获取单个接口
  async getEndpointById(id: number): Promise<APIEndpoint> {
    try {
      const response = await api.get(`/swagger/endpoint/${id}`);
      return response as unknown as APIEndpoint;
    } catch (error: any) {
      throw new Error(error.message || '获取接口详情失败');
    }
  },

  // 更新接口
  async updateEndpoint(endpoint: APIEndpoint): Promise<APIEndpoint> {
    try {
      const response = await api.put('/swagger/endpoint', endpoint);
      return response as unknown as APIEndpoint;
    } catch (error: any) {
      throw new Error(error.message || '更新接口失败');
    }
  },

  // 删除接口
  async deleteEndpoint(id: number): Promise<void> {
    try {
      await api.delete(`/swagger/endpoint/${id}`);
    } catch (error: any) {
      throw new Error(error.message || '删除接口失败');
    }
  },

  // 测试接口
  async testEndpoint(endpoint: APIEndpoint, baseUrl: string): Promise<any> {
    try {
      const response = await api.post('/swagger/endpoint/test', endpoint, {
        params: { base_url: baseUrl }
      });
      
      return response;
    } catch (error: any) {
      throw new Error(error.message || '接口测试失败');
    }
  }
};
