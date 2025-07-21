import React, { useState } from 'react';
import {
  Card,
  Upload,
  Button,
  TextArea,
  Space,
  Divider,
  Radio,
  Message,
  Loading,
  Tag,
  List,
  Dialog
} from 'tdesign-react';
import { CloudUploadIcon, CheckCircleFilledIcon, ErrorCircleFilledIcon } from 'tdesign-icons-react';
import { swaggerService } from '../services/swagger';
import type { SwaggerValidationResult, SwaggerParseResult } from '../types/swagger';

const { RadioGroup } = Radio;

const SwaggerImport: React.FC = () => {
  const [importType, setImportType] = useState<'file' | 'text'>('file');
  const [swaggerContent, setSwaggerContent] = useState('');
  const [loading, setLoading] = useState(false);
  const [validationResult, setValidationResult] = useState<SwaggerValidationResult | null>(null);
  const [parseResult, setParseResult] = useState<SwaggerParseResult | null>(null);
  const [previewVisible, setPreviewVisible] = useState(false);

  // 文件上传处理
  const handleFileUpload = async (files: File[]) => {
    if (files.length === 0) return;
    
    const file = files[0];
    setLoading(true);
    
    try {
      const result = await swaggerService.validateFile(file);
      setValidationResult(result);
      
      if (result.valid) {
        Message.success('文件校验成功！');
      } else {
        Message.error('文件校验失败：' + result.message);
      }
    } catch (error) {
      Message.error('文件上传失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 文本内容校验
  const handleTextValidation = async () => {
    if (!swaggerContent.trim()) {
      Message.warning('请输入Swagger文档内容');
      return;
    }

    setLoading(true);
    try {
      const result = await swaggerService.validateText(swaggerContent);
      setValidationResult(result);
      
      if (result.valid) {
        Message.success('内容校验成功！');
      } else {
        Message.error('内容校验失败：' + result.message);
      }
    } catch (error) {
      Message.error('校验失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 解析并保存API接口
  const handleParseAndSave = async () => {
    if (!validationResult?.valid) {
      Message.warning('请先完成文档校验');
      return;
    }

    const content = importType === 'text' ? swaggerContent : validationResult.content;
    if (!content) {
      Message.warning('没有可解析的内容');
      return;
    }

    setLoading(true);
    try {
      const result = await swaggerService.parseAndSave(content);
      setParseResult(result);
      Message.success(`解析成功！共发现 ${result.endpoints.length} 个API接口`);
    } catch (error) {
      Message.error('解析失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 预览API接口
  const handlePreview = () => {
    if (!parseResult) return;
    setPreviewVisible(true);
  };

  const renderValidationStatus = () => {
    if (!validationResult) return null;

    return (
      <Card title="校验结果" style={{ marginTop: 16 }}>
        <Space direction="vertical" style={{ width: '100%' }}>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            {validationResult.valid ? (
              <CheckCircleFilledIcon style={{ color: '#52c41a', marginRight: 8 }} />
            ) : (
              <ErrorCircleFilledIcon style={{ color: '#ff4d4f', marginRight: 8 }} />
            )}
            <span>
              {validationResult.valid ? '文档格式正确' : '文档格式错误'}
            </span>
          </div>
          
          {validationResult.message && (
            <div>
              <strong>详细信息：</strong>
              <div style={{ background: '#f5f5f5', padding: 8, borderRadius: 4, marginTop: 4 }}>
                {validationResult.message}
              </div>
            </div>
          )}

          {validationResult.valid && (
            <Space>
              <Button type="primary" onClick={handleParseAndSave} loading={loading}>
                解析并保存到数据库
              </Button>
              {parseResult && (
                <Button theme="default" onClick={handlePreview}>
                  预览API接口 ({parseResult.endpoints.length})
                </Button>
              )}
            </Space>
          )}
        </Space>
      </Card>
    );
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
              onChange={setImportType}
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
                onChange={({ files }) => handleFileUpload(files)}
                theme="file-flow"
                placeholder="点击上传Swagger文档文件"
                tips="支持.json、.yaml、.yml格式"
              >
                <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', height: 120 }}>
                  <Space direction="vertical" align="center">
                    <CloudUploadIcon size="48" style={{ color: '#0052d9' }} />
                    <span>点击或拖拽文件到此区域上传</span>
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
              <TextArea
                placeholder="请粘贴完整的Swagger/OpenAPI文档内容（JSON或YAML格式）"
                value={swaggerContent}
                onChange={setSwaggerContent}
                autosize={{ minRows: 10, maxRows: 20 }}
                style={{ width: '100%' }}
              />
              <div style={{ marginTop: 16 }}>
                <Button 
                  type="primary" 
                  onClick={handleTextValidation}
                  loading={loading}
                  disabled={!swaggerContent.trim()}
                >
                  校验文档格式
                </Button>
              </div>
            </div>
          )}

          {/* 校验结果 */}
          {renderValidationStatus()}

        </Space>
      </Card>

      {/* API预览弹窗 */}
      <Dialog
        visible={previewVisible}
        onClose={() => setPreviewVisible(false)}
        header="API接口预览"
        width="80%"
        top="5vh"
      >
        {parseResult && (
          <div>
            <div style={{ marginBottom: 16 }}>
              <Tag theme="primary" variant="light">
                共 {parseResult.endpoints.length} 个接口
              </Tag>
            </div>
            <List
              data={parseResult.endpoints}
              renderItem={(item, index) => (
                <List.ListItem key={index}>
                  <div style={{ width: '100%' }}>
                    <div style={{ display: 'flex', alignItems: 'center', marginBottom: 8 }}>
                      <Tag 
                        theme={item.method === 'GET' ? 'success' : item.method === 'POST' ? 'primary' : 'warning'}
                        style={{ marginRight: 8, minWidth: 60, textAlign: 'center' }}
                      >
                        {item.method}
                      </Tag>
                      <code style={{ background: '#f5f5f5', padding: '2px 8px', borderRadius: 4 }}>
                        {item.path}
                      </code>
                    </div>
                    <div>
                      <strong>{item.summary || '未命名接口'}</strong>
                    </div>
                    {item.description && (
                      <div style={{ color: '#666', fontSize: '14px', marginTop: 4 }}>
                        {item.description}
                      </div>
                    )}
                    {item.tags && (
                      <div style={{ marginTop: 8 }}>
                        {item.tags.split(',').map((tag, i) => (
                          <Tag key={i} variant="outline" style={{ marginRight: 4 }}>
                            {tag.trim()}
                          </Tag>
                        ))}
                      </div>
                    )}
                  </div>
                </List.ListItem>
              )}
            />
          </div>
        )}
      </Dialog>

      {/* 全局loading */}
      <Loading 
        loading={loading} 
        text="处理中..." 
        size="medium"
        style={{ position: 'fixed' }}
      />
    </div>
  );
};

export default SwaggerImport;
