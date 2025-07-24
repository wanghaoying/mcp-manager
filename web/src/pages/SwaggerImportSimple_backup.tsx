import React, { useState } from 'react';
import {
  Card,
  Upload,
  Button,
  Textarea,
  Space,
  Divider,
  Radio,
  MessagePlugin,
  Tag,
  Loading
} from 'tdesign-react';
import { CloudUploadIcon } from 'tdesign-icons-react';
import { swaggerService } from '../services/swagger';{ useState } from 'react';
import {
  Card,
  Upload,
  Button,
  Textarea,
  Space,
  Divider,
  Radio,
  MessagePlugin,
  Tag,
  Loading
} from 'tdesign-react';
import { CloudUploadIcon } from 'tdesign-icons-react';
import { swaggerService } from '../services/swagger';

const { Group: RadioGroup } = Radio;

const SwaggerImport: React.FC = () => {
  const [importType, setImportType] = useState<'file' | 'text'>('file');
  const [swaggerContent, setSwaggerContent] = useState('');
  const [loading, setLoading] = useState(false);
  const [validationResult, setValidationResult] = useState<any>(null);
  const [parseResult, setParseResult] = useState<any>(null);

  // 文件上传处理 - 真正调用后台API
  const handleFileUpload = async (files: File[]) => {
    if (files.length === 0) return;
    
    const file = files[0];
    setLoading(true);
    
    try {
      // 先读取文件内容检查格式
      const fileContent = await file.text();
      try {
        const parsed = JSON.parse(fileContent);
        if (parsed.swagger && parsed.swagger.startsWith('2.')) {
          MessagePlugin.error('当前系统仅支持OpenAPI 3.0格式，请将Swagger 2.0文档转换为OpenAPI 3.0格式后再导入。\n您可以使用在线转换工具：https://converter.swagger.io/');
          setLoading(false);
          return;
        }
      } catch (e) {
        // 如果不是JSON格式，可能是YAML，继续处理
      }

      // 调用后台API进行文件校验
      const result = await swaggerService.validateFile(file);
      
      if (result.valid) {
        MessagePlugin.success(`文件 ${file.name} 校验成功！`);
        setValidationResult(result);
        
        // 如果校验成功，可以继续解析
        if (result.content) {
          await handleParseContent(result.content);
        }
      } else {
        MessagePlugin.error(result.message || '文件校验失败');
        setValidationResult(result);
      }
    } catch (error: any) {
      console.error('文件上传错误:', error);
      let errorMessage = error.message || '文件处理失败';
      
      // 如果是openapi格式错误，提供更友好的提示
      if (errorMessage.includes('openapi must be a non-empty string')) {
        errorMessage = '文档格式错误：当前系统仅支持OpenAPI 3.0格式。请确保文档中包含 "openapi": "3.0.0" 字段。';
      }
      
      MessagePlugin.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  // 文本校验处理 - 真正调用后台API
  const handleTextValidation = async () => {
    if (!swaggerContent.trim()) {
      MessagePlugin.warning('请输入Swagger文档内容');
      return;
    }

    setLoading(true);
    
    try {
      // 先检查格式是否为Swagger 2.0
      let content = swaggerContent;
      try {
        const parsed = JSON.parse(content);
        if (parsed.swagger && parsed.swagger.startsWith('2.')) {
          MessagePlugin.error('当前系统仅支持OpenAPI 3.0格式，请将Swagger 2.0文档转换为OpenAPI 3.0格式后再导入。\n您可以使用在线转换工具：https://converter.swagger.io/');
          setLoading(false);
          return;
        }
      } catch (e) {
        // 如果不是JSON格式，可能是YAML，继续处理
      }

      // 调用后台API进行文本校验
      const result = await swaggerService.validateText(content);
      
      if (result.valid) {
        MessagePlugin.success('内容校验通过！');
        setValidationResult(result);
        
        // 如果校验成功，可以继续解析
        await handleParseContent(content);
      } else {
        MessagePlugin.error(result.message || '内容校验失败');
        setValidationResult(result);
      }
    } catch (error: any) {
      console.error('文本校验错误:', error);
      let errorMessage = error.message || '内容校验失败';
      
      // 如果是openapi格式错误，提供更友好的提示
      if (errorMessage.includes('openapi must be a non-empty string')) {
        errorMessage = '文档格式错误：当前系统仅支持OpenAPI 3.0格式。请确保文档中包含 "openapi": "3.0.0" 字段。';
      }
      
      MessagePlugin.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  // 解析内容并保存到数据库
  const handleParseContent = async (content: string) => {
    try {
      const result = await swaggerService.parseAndSave(content);
      
      if (result.success) {
        MessagePlugin.success(`解析成功！共解析出 ${result.endpoints?.length || 0} 个API接口`);
        setParseResult(result);
      }
    } catch (error: any) {
      console.error('解析错误:', error);
      MessagePlugin.error(error.message || '解析失败');
    }
  };

  return (
    <div>
      <Card title="Swagger文档导入" style={{ marginBottom: 16 }}>
        {loading && <Loading loading={true} text="处理中..." />}
        
        <Space direction="vertical" style={{ width: '100%' }} size="large">
          {/* 导入方式选择 */}
          <div>
            <div style={{ marginBottom: 16 }}>
              <strong>选择导入方式：</strong>
            </div>
            <RadioGroup
              value={importType}
              onChange={(value) => setImportType(value as 'file' | 'text')}
              options={[
                { label: '文件上传', value: 'file' },
                { label: '文本粘贴', value: 'text' }
              ]}
              disabled={loading}
            />
          </div>

          <Divider />

          {/* 文件上传 */}
          {importType === 'file' && (
            <div>
              <div style={{ marginBottom: 16 }}>
                <strong>上传Swagger文档文件：</strong>
              </div>
              <Upload
                action=""
                accept=".json,.yaml,.yml"
                multiple={false}
                disabled={loading}
                beforeUpload={() => false}
                onChange={(files) => {
                  console.log('Upload onChange triggered:', files);
                  if (files && files.length > 0) {
                    // 直接处理文件，TDesign的Upload组件会传入FileList
                    const fileArray = Array.from(files as FileList);
                    handleFileUpload(fileArray);
                  }
                }}
                onRemove={() => {
                  // 清空结果
                  setValidationResult(null);
                  setParseResult(null);
                }}
                theme="file-flow"
                tips="支持.json、.yaml、.yml格式"
                placeholder="点击或拖拽文件到此区域上传"
              >
                <div style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  justifyContent: 'center', 
                  height: 120,
                  border: '2px dashed #d9d9d9',
                  borderRadius: 6,
                  cursor: loading ? 'not-allowed' : 'pointer',
                  opacity: loading ? 0.6 : 1
                }}>
                  <Space direction="vertical" align="center">
                    <CloudUploadIcon size="48" style={{ color: '#0052d9' }} />
                    <span>点击或拖拽文件到此区域上传</span>
                    <span style={{ color: '#999', fontSize: '12px' }}>支持 .json、.yaml、.yml 格式</span>
                  </Space>
                </div>
              </Upload>
            </div>
          )}

          {/* 文本粘贴 */}
          {importType === 'text' && (
            <div>
              <div style={{ marginBottom: 16 }}>
                <strong>粘贴Swagger文档内容：</strong>
              </div>
              <Textarea
                placeholder="请粘贴完整的Swagger/OpenAPI文档内容（JSON或YAML格式）"
                value={swaggerContent}
                onChange={setSwaggerContent}
                autosize={{ minRows: 10, maxRows: 20 }}
                style={{ width: '100%' }}
                disabled={loading}
              />
              <div style={{ marginTop: 16 }}>
                <Button 
                  theme="primary" 
                  onClick={handleTextValidation}
                  disabled={!swaggerContent.trim() || loading}
                  loading={loading}
                >
                  校验并解析文档
                </Button>
              </div>
            </div>
          )}

          {/* 校验结果显示 */}
          {validationResult && (
            <Card title="校验结果" style={{ marginTop: 16 }}>
              <div style={{ marginBottom: 8 }}>
                <Tag theme={validationResult.valid ? 'success' : 'danger'}>
                  {validationResult.valid ? '✅ 校验通过' : '❌ 校验失败'}
                </Tag>
              </div>
              <div>{validationResult.message}</div>
            </Card>
          )}

          {/* 解析结果显示 */}
          {parseResult && parseResult.success && (
            <Card title="解析结果" style={{ marginTop: 16 }}>
              <div style={{ marginBottom: 8 }}>
                <Tag theme="success">🎉 解析成功</Tag>
              </div>
              <div style={{ marginBottom: 8 }}>
                共解析出 <strong>{parseResult.endpoints?.length || 0}</strong> 个API接口
              </div>
              <div style={{ fontSize: '12px', color: '#666' }}>
                接口信息已保存到数据库，您可以在"API接口管理"页面查看详情
              </div>
            </Card>
          )}

          {/* 功能说明 */}
          <Card title="功能说明" style={{ background: '#f9f9f9' }}>
            <Space direction="vertical" style={{ width: '100%' }}>
              <div>
                <Tag theme="primary" variant="light">✨ 支持格式</Tag>
                <span style={{ marginLeft: 8 }}>仅支持OpenAPI 3.0格式的JSON、YAML文档（不支持Swagger 2.0）</span>
              </div>
              <div>
                <Tag theme="warning" variant="light">⚠️ 格式要求</Tag>
                <span style={{ marginLeft: 8 }}>文档必须包含 "openapi": "3.0.x" 字段，Swagger 2.0需要先转换</span>
              </div>
              <div>
                <Tag theme="success" variant="light">🔍 实时校验</Tag>
                <span style={{ marginLeft: 8 }}>上传后自动校验文档格式和内容完整性</span>
              </div>
              <div>
                <Tag theme="default" variant="light">� 格式转换</Tag>
                <span style={{ marginLeft: 8 }}>可使用 <a href="https://converter.swagger.io/" target="_blank" rel="noopener noreferrer">在线转换工具</a> 将Swagger 2.0转换为OpenAPI 3.0</span>
              </div>
              <div>
                <Tag theme="default" variant="light">💾 数据保存</Tag>
                <span style={{ marginLeft: 8 }}>解析成功的接口将保存到数据库</span>
              </div>
            </Space>
          </Card>
        </Space>
      </Card>
    </div>
  );
};

export default SwaggerImport;
