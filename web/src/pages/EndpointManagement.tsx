import React, { useState, useEffect } from 'react';
import {
  Card,
  Table,
  Button,
  Space,
  Tag,
  Dialog,
  Form,
  Input,
  Select,
  Message,
  Popconfirm,
  Drawer,
  Descriptions
} from 'tdesign-react';
import { ViewIcon, EditIcon, DeleteIcon, PlayIcon } from 'tdesign-icons-react';
import { swaggerService } from '../services/swagger';
import type { APIEndpoint } from '../types/swagger';

const { FormItem } = Form;
const { Option } = Select;

const EndpointManagement: React.FC = () => {
  const [endpoints, setEndpoints] = useState<APIEndpoint[]>([]);
  const [loading, setLoading] = useState(false);
  const [editVisible, setEditVisible] = useState(false);
  const [detailVisible, setDetailVisible] = useState(false);
  const [testVisible, setTestVisible] = useState(false);
  const [currentEndpoint, setCurrentEndpoint] = useState<APIEndpoint | null>(null);
  const [form] = Form.useForm();
  const [testForm] = Form.useForm();
  const [testResult, setTestResult] = useState<any>(null);

  // 获取接口列表
  const fetchEndpoints = async (swaggerId?: number) => {
    setLoading(true);
    try {
      const result = await swaggerService.getEndpoints(swaggerId || 1); // 默认使用ID为1的swagger
      setEndpoints(result);
    } catch (error) {
      Message.error('获取接口列表失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEndpoints();
  }, []);

  // 编辑接口
  const handleEdit = (record: APIEndpoint) => {
    setCurrentEndpoint(record);
    form.setFields([
      { name: 'summary', value: record.summary },
      { name: 'description', value: record.description },
      { name: 'method', value: record.method },
      { name: 'path', value: record.path },
      { name: 'tags', value: record.tags }
    ]);
    setEditVisible(true);
  };

  // 保存编辑
  const handleSave = async (values: any) => {
    if (!currentEndpoint) return;
    
    setLoading(true);
    try {
      const updatedEndpoint = { ...currentEndpoint, ...values };
      await swaggerService.updateEndpoint(updatedEndpoint);
      Message.success('更新成功');
      setEditVisible(false);
      fetchEndpoints();
    } catch (error) {
      Message.error('更新失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 删除接口
  const handleDelete = async (id: number) => {
    setLoading(true);
    try {
      await swaggerService.deleteEndpoint(id);
      Message.success('删除成功');
      fetchEndpoints();
    } catch (error) {
      Message.error('删除失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  // 查看详情
  const handleViewDetail = (record: APIEndpoint) => {
    setCurrentEndpoint(record);
    setDetailVisible(true);
  };

  // 测试接口
  const handleTest = (record: APIEndpoint) => {
    setCurrentEndpoint(record);
    testForm.reset();
    setTestResult(null);
    setTestVisible(true);
  };

  // 执行接口测试
  const handleRunTest = async (values: any) => {
    if (!currentEndpoint) return;

    setLoading(true);
    try {
      const result = await swaggerService.testEndpoint(currentEndpoint, values.baseUrl);
      setTestResult(result);
      Message.success('测试完成');
    } catch (error) {
      Message.error('测试失败');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    {
      title: '方法',
      dataIndex: 'method',
      width: 80,
      render: (method: string) => (
        <Tag 
          theme={method === 'GET' ? 'success' : method === 'POST' ? 'primary' : 'warning'}
        >
          {method}
        </Tag>
      )
    },
    {
      title: '路径',
      dataIndex: 'path',
      width: 200,
      render: (path: string) => (
        <code style={{ background: '#f5f5f5', padding: '2px 8px', borderRadius: 4 }}>
          {path}
        </code>
      )
    },
    {
      title: '接口名称',
      dataIndex: 'summary',
      ellipsis: true
    },
    {
      title: '描述',
      dataIndex: 'description',
      ellipsis: true,
      width: 200
    },
    {
      title: '标签',
      dataIndex: 'tags',
      width: 150,
      render: (tags: string) => {
        if (!tags) return '-';
        return tags.split(',').slice(0, 2).map((tag, index) => (
          <Tag key={index} variant="outline" style={{ marginRight: 4 }}>
            {tag.trim()}
          </Tag>
        ));
      }
    },
    {
      title: '操作',
      width: 200,
      render: (_: any, record: APIEndpoint) => (
        <Space>
          <Button 
            size="small" 
            theme="default" 
            variant="text"
            onClick={() => handleViewDetail(record)}
          >
            <ViewIcon />
          </Button>
          <Button 
            size="small" 
            theme="primary" 
            variant="text"
            onClick={() => handleEdit(record)}
          >
            <EditIcon />
          </Button>
          <Button 
            size="small" 
            theme="success" 
            variant="text"
            onClick={() => handleTest(record)}
          >
            <PlayIcon />
          </Button>
          <Popconfirm
            content="确定要删除这个接口吗？"
            onConfirm={() => handleDelete(record.id)}
          >
            <Button 
              size="small" 
              theme="danger" 
              variant="text"
            >
              <DeleteIcon />
            </Button>
          </Popconfirm>
        </Space>
      )
    }
  ];

  return (
    <div>
      <Card title="API接口管理">
        <Table
          data={endpoints}
          columns={columns}
          loading={loading}
          pagination={{
            pageSize: 10,
            showJumper: true
          }}
          rowKey="id"
        />
      </Card>

      {/* 编辑接口弹窗 */}
      <Dialog
        visible={editVisible}
        onClose={() => setEditVisible(false)}
        header="编辑接口"
        confirmBtn="保存"
        onConfirm={() => form.submit()}
        width="600px"
      >
        <Form
          form={form}
          onSubmit={handleSave}
          labelWidth="100px"
        >
          <FormItem label="接口名称" name="summary" rules={[{ required: true, message: '请输入接口名称' }]}>
            <Input placeholder="请输入接口名称" />
          </FormItem>
          <FormItem label="接口描述" name="description">
            <Input placeholder="请输入接口描述" />
          </FormItem>
          <FormItem label="请求方法" name="method" rules={[{ required: true, message: '请选择请求方法' }]}>
            <Select placeholder="请选择请求方法">
              <Option value="GET">GET</Option>
              <Option value="POST">POST</Option>
              <Option value="PUT">PUT</Option>
              <Option value="DELETE">DELETE</Option>
              <Option value="PATCH">PATCH</Option>
            </Select>
          </FormItem>
          <FormItem label="接口路径" name="path" rules={[{ required: true, message: '请输入接口路径' }]}>
            <Input placeholder="请输入接口路径" />
          </FormItem>
          <FormItem label="标签" name="tags">
            <Input placeholder="多个标签用逗号分隔" />
          </FormItem>
        </Form>
      </Dialog>

      {/* 接口详情抽屉 */}
      <Drawer
        visible={detailVisible}
        onClose={() => setDetailVisible(false)}
        header="接口详情"
        size="large"
      >
        {currentEndpoint && (
          <Descriptions
            itemLayout="vertical"
            colon
            bordered
          >
            <Descriptions.DescriptionsItem label="接口名称" span={2}>
              {currentEndpoint.summary || '-'}
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="请求方法" span={1}>
              <Tag theme={currentEndpoint.method === 'GET' ? 'success' : 'primary'}>
                {currentEndpoint.method}
              </Tag>
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="接口路径" span={3}>
              <code style={{ background: '#f5f5f5', padding: '4px 8px', borderRadius: 4 }}>
                {currentEndpoint.path}
              </code>
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="描述" span={3}>
              {currentEndpoint.description || '-'}
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="标签" span={3}>
              {currentEndpoint.tags ? (
                currentEndpoint.tags.split(',').map((tag, index) => (
                  <Tag key={index} variant="outline" style={{ marginRight: 4 }}>
                    {tag.trim()}
                  </Tag>
                ))
              ) : '-'}
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="请求参数" span={3}>
              {currentEndpoint.parameters && currentEndpoint.parameters.length > 0 ? (
                <div>
                  {currentEndpoint.parameters.map((param, index) => (
                    <div key={index} style={{ marginBottom: 8 }}>
                      <Tag size="small">{param.in}</Tag>
                      <strong style={{ margin: '0 8px' }}>{param.name}</strong>
                      <span style={{ color: '#666' }}>({param.type})</span>
                      {param.required && <Tag theme="danger" size="small" style={{ marginLeft: 8 }}>必需</Tag>}
                    </div>
                  ))}
                </div>
              ) : '-'}
            </Descriptions.DescriptionsItem>
            <Descriptions.DescriptionsItem label="响应示例" span={3}>
              {currentEndpoint.responses ? (
                <pre style={{ background: '#f5f5f5', padding: 8, borderRadius: 4, overflow: 'auto' }}>
                  {currentEndpoint.responses}
                </pre>
              ) : '-'}
            </Descriptions.DescriptionsItem>
          </Descriptions>
        )}
      </Drawer>

      {/* 接口测试弹窗 */}
      <Dialog
        visible={testVisible}
        onClose={() => setTestVisible(false)}
        header="接口测试"
        width="800px"
        footer={null}
      >
        {currentEndpoint && (
          <Space direction="vertical" style={{ width: '100%' }} size="large">
            <div>
              <h4>接口信息</h4>
              <div style={{ background: '#f5f5f5', padding: 12, borderRadius: 4 }}>
                <Tag theme="primary" style={{ marginRight: 8 }}>{currentEndpoint.method}</Tag>
                <code>{currentEndpoint.path}</code>
                <div style={{ marginTop: 8, fontSize: '14px', color: '#666' }}>
                  {currentEndpoint.summary}
                </div>
              </div>
            </div>

            <Form
              form={testForm}
              onSubmit={handleRunTest}
              labelWidth="120px"
            >
              <FormItem 
                label="服务器基础URL" 
                name="baseUrl" 
                rules={[{ required: true, message: '请输入服务器基础URL' }]}
              >
                <Input 
                  placeholder="例如：https://api.example.com" 
                  defaultValue="http://localhost:8080"
                />
              </FormItem>
              
              <FormItem>
                <Space>
                  <Button type="primary" theme="primary" loading={loading}>
                    执行测试
                  </Button>
                  <Button onClick={() => setTestVisible(false)}>
                    取消
                  </Button>
                </Space>
              </FormItem>
            </Form>

            {testResult && (
              <div>
                <h4>测试结果</h4>
                <pre style={{ 
                  background: '#f5f5f5', 
                  padding: 12, 
                  borderRadius: 4, 
                  overflow: 'auto',
                  maxHeight: '300px'
                }}>
                  {JSON.stringify(testResult, null, 2)}
                </pre>
              </div>
            )}
          </Space>
        )}
      </Dialog>
    </div>
  );
};

export default EndpointManagement;
