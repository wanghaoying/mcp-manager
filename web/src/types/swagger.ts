// API参数类型
export interface APIParameter {
  name: string;
  in: string;
  type: string;
  required: boolean;
  value?: string;
}

// API端点类型
export interface APIEndpoint {
  id: number;
  swagger_id: number;
  operation_id: string;
  method: string;
  path: string;
  summary?: string;
  description?: string;
  tags?: string;
  parameters?: APIParameter[];
  headers?: { [key: string]: string };
  body?: string;
  responses?: string;
  created_at: string;
  updated_at: string;
}

// Swagger校验结果
export interface SwaggerValidationResult {
  valid: boolean;
  message?: string;
  content?: string;
}

// Swagger解析结果
export interface SwaggerParseResult {
  success: boolean;
  message: string;
  endpoints: APIEndpoint[];
}

// Swagger文本请求
export interface SwaggerTextRequest {
  content: string;
}
