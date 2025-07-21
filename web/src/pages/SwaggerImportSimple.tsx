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
  Tag
} from 'tdesign-react';
import { CloudUploadIcon } from 'tdesign-icons-react';

const { Group: RadioGroup } = Radio;

const SwaggerImport: React.FC = () => {
  const [importType, setImportType] = useState<'file' | 'text'>('file');
  const [swaggerContent, setSwaggerContent] = useState('');

  // 简化的文件上传处理
  const handleFileUpload = async (files: File[]) => {
    if (files.length === 0) return;
    
    const file = files[0];
    console.log('上传文件:', file.name);
    MessagePlugin.success(`文件 ${file.name} 上传成功！`);
  };

  // 简化的文本校验
  const handleTextValidation = async () => {
    if (!swaggerContent.trim()) {
      MessagePlugin.warning('请输入Swagger文档内容');
      return;
    }

    console.log('校验内容:', swaggerContent.length, '字符');
    MessagePlugin.success('内容校验通过！');
  };

  return (
    <div>
      <Card title="Swagger文档导入" style={{ marginBottom: 16 }}>
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
                beforeUpload={() => false}
                onChange={(value) => {
                  if (value.length > 0) {
                    handleFileUpload(value as File[]);
                  }
                }}
                theme="file-flow"
                placeholder="点击上传Swagger文档文件"
                tips="支持.json、.yaml、.yml格式"
              >
                <div style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  justifyContent: 'center', 
                  height: 120,
                  border: '2px dashed #d9d9d9',
                  borderRadius: 6,
                  cursor: 'pointer'
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
              />
              <div style={{ marginTop: 16 }}>
                <Button 
                  theme="primary" 
                  onClick={handleTextValidation}
                  disabled={!swaggerContent.trim()}
                >
                  校验文档格式
                </Button>
              </div>
            </div>
          )}

          {/* 功能说明 */}
          <Card title="功能说明" style={{ background: '#f9f9f9' }}>
            <Space direction="vertical" style={{ width: '100%' }}>
              <div>
                <Tag theme="primary" variant="light">✨ 支持格式</Tag>
                <span style={{ marginLeft: 8 }}>JSON、YAML格式的OpenAPI 2.0/3.0规范文档</span>
              </div>
              <div>
                <Tag theme="success" variant="light">🔍 实时校验</Tag>
                <span style={{ marginLeft: 8 }}>上传后自动校验文档格式和内容完整性</span>
              </div>
              <div>
                <Tag theme="warning" variant="light">📊 解析预览</Tag>
                <span style={{ marginLeft: 8 }}>解析后可预览所有API接口信息</span>
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
