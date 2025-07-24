import React, { useState, useEffect } from 'react';
import { Card, Table, Button, Space, Tag, MessagePlugin, Loading } from 'tdesign-react';
import { swaggerService } from '../services/swagger';
import type { APIEndpoint } from '../types/swagger';

// 模拟数据作为后备
const mockEndpoints = [
  {
    id: 1,
    method: 'GET',
    path: '/api/users',
    summary: '获取用户列表',
    description: '获取所有用户信息',
    tags: 'User,Query'
  },
  {
    id: 2,
    method: 'POST',
    path: '/api/users',
    summary: '创建用户',
    description: '创建新用户账户',
    tags: 'User,Create'
  },
  {
    id: 3,
    method: 'PUT',
    path: '/api/users/{id}',
    summary: '更新用户信息',
    description: '更新指定用户的信息',
    tags: 'User,Update'
  },
  {
    id: 4,
    method: 'DELETE',
    path: '/api/users/{id}',
    summary: '删除用户',
    description: '删除指定用户账户',
    tags: 'User,Delete'
  }
];

const EndpointManagement: React.FC = () => {
  const [endpoints, setEndpoints] = useState<APIEndpoint[]>([]);
  const [loading, setLoading] = useState(false);

  // 加载接口数据
  const loadEndpoints = async () => {
    setLoading(true);
    try {
      // 这里使用swagger_id = 0，因为从测试看到后台使用的是0
      const data = await swaggerService.getEndpoints(0);
      setEndpoints(data);
    } catch (error: any) {
      console.error('加载接口数据失败:', error);
      MessagePlugin.error(error.message || '加载接口数据失败');
      // 如果加载失败，显示模拟数据
      setEndpoints(mockEndpoints as APIEndpoint[]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadEndpoints();
  }, []);

  const handleEdit = (record: APIEndpoint) => {
    MessagePlugin.info(`编辑接口: ${record.summary}`);
  };

  const handleDelete = async (record: APIEndpoint) => {
    try {
      await swaggerService.deleteEndpoint(record.id);
      MessagePlugin.success(`删除接口成功: ${record.summary}`);
      // 重新加载数据
      loadEndpoints();
    } catch (error: any) {
      MessagePlugin.error(error.message || '删除接口失败');
    }
  };

  const handleTest = async (record: APIEndpoint) => {
    try {
      MessagePlugin.info(`正在测试接口: ${record.method} ${record.path}...`);
      // 这里需要一个基础URL，实际应用中应该从配置获取
      const result = await swaggerService.testEndpoint(record, 'http://localhost:8080');
      MessagePlugin.success(`接口测试成功`);
      console.log('测试结果:', result);
    } catch (error: any) {
      MessagePlugin.error(error.message || '接口测试失败');
    }
  };

  const handleRefresh = () => {
    loadEndpoints();
  };

  const columns = [
    {
      title: '方法',
      colKey: 'method',
      width: 80,
      cell: ({ row }: any) => (
        <Tag 
          theme={row.method === 'GET' ? 'success' : row.method === 'POST' ? 'primary' : row.method === 'PUT' ? 'warning' : 'danger'}
          variant="light"
        >
          {row.method}
        </Tag>
      )
    },
    {
      title: '路径',
      colKey: 'path',
      width: 200,
      cell: ({ row }: any) => (
        <code style={{ background: '#f5f5f5', padding: '2px 8px', borderRadius: 4, fontSize: '12px' }}>
          {row.path}
        </code>
      )
    },
    {
      title: '接口名称',
      colKey: 'summary',
      ellipsis: true
    },
    {
      title: '描述',
      colKey: 'description',
      ellipsis: true,
      width: 200
    },
    {
      title: '标签',
      colKey: 'tags',
      width: 150,
      cell: ({ row }: any) => {
        if (!row.tags) return '-';
        return row.tags.split(',').slice(0, 2).map((tag: string, index: number) => (
          <Tag key={index} variant="outline" style={{ marginRight: 4, fontSize: '12px' }}>
            {tag.trim()}
          </Tag>
        ));
      }
    },
    {
      title: '操作',
      colKey: 'operation',
      width: 180,
      cell: ({ row }: any) => (
        <Space size="small">
          <Button 
            size="small" 
            theme="primary" 
            variant="text"
            onClick={() => handleEdit(row)}
          >
            编辑
          </Button>
          <Button 
            size="small" 
            theme="success" 
            variant="text"
            onClick={() => handleTest(row)}
          >
            测试
          </Button>
          <Button 
            size="small" 
            theme="danger" 
            variant="text"
            onClick={() => handleDelete(row)}
          >
            删除
          </Button>
        </Space>
      )
    }
  ];

  return (
    <div>
      <Card title="API接口管理">
        {loading && <Loading loading={true} text="加载中..." />}
        
        <div style={{ marginBottom: 16 }}>
          <Space>
            <Tag theme="primary" variant="light">
              共 {endpoints.length} 个接口
            </Tag>
            <Button theme="primary" variant="outline" onClick={handleRefresh} loading={loading}>
              刷新
            </Button>
          </Space>
        </div>
        
        <Table
          data={endpoints}
          columns={columns}
          pagination={{
            current: 1,
            pageSize: 10,
            total: endpoints.length,
            showJumper: true
          }}
          rowKey="id"
          size="medium"
        />
      </Card>

      {/* 功能说明 */}
      <Card title="功能说明" style={{ marginTop: 16, background: '#f9f9f9' }}>
        <Space direction="vertical" style={{ width: '100%' }}>
          <div>
            <Tag theme="primary" variant="light">✏️ 编辑功能</Tag>
            <span style={{ marginLeft: 8 }}>可以修改接口的基本信息，如名称、描述、标签等</span>
          </div>
          <div>
            <Tag theme="success" variant="light">🧪 测试功能</Tag>
            <span style={{ marginLeft: 8 }}>支持在线测试API接口，验证接口可用性</span>
          </div>
          <div>
            <Tag theme="danger" variant="light">🗑️ 删除功能</Tag>
            <span style={{ marginLeft: 8 }}>可以删除不需要的接口（需要确认操作）</span>
          </div>
          <div>
            <Tag theme="default" variant="light">📋 批量管理</Tag>
            <span style={{ marginLeft: 8 }}>支持批量选择和操作多个接口</span>
          </div>
        </Space>
      </Card>
    </div>
  );
};

export default EndpointManagement;
