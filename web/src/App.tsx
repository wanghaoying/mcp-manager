import React from 'react';
import { Layout, Menu } from 'tdesign-react';
import { ApiIcon, SwapIcon, ViewListIcon } from 'tdesign-icons-react';
import SwaggerImport from './pages/SwaggerImportSimple';
import EndpointManagement from './pages/EndpointManagementSimple';

const { Header, Content, Aside } = Layout;

const menuItems = [
  {
    value: 'swagger-import',
    label: 'Swagger文档导入',
    icon: <ApiIcon />
  },
  {
    value: 'endpoint-management',
    label: 'API接口管理',
    icon: <ViewListIcon />
  }
];

function App() {
  const [selectedMenu, setSelectedMenu] = React.useState('swagger-import');

  const handleMenuClick = (value: any) => {
    setSelectedMenu(value as string);
  };

  const renderContent = () => {
    switch (selectedMenu) {
      case 'swagger-import':
        return <SwaggerImport />;
      case 'endpoint-management':
        return <EndpointManagement />;
      default:
        return <SwaggerImport />;
    }
  };

  return (
    <div className="app">
      <Header className="header">
        <h1 style={{ margin: 0, fontSize: '20px', color: '#1f2937' }}>
          <SwapIcon style={{ marginRight: '8px', fontSize: '24px' }} />
          MCP工具市场管理平台
        </h1>
      </Header>
      
      <Layout style={{ flex: 1 }}>
        <Aside width="240px" style={{ background: '#fff', borderRight: '1px solid #e7e7e7' }}>
          <Menu
            value={selectedMenu}
            onChange={handleMenuClick}
            style={{ borderRight: 'none' }}
          >
            {menuItems.map(item => (
              <Menu.MenuItem key={item.value} value={item.value}>
                {item.icon}
                {item.label}
              </Menu.MenuItem>
            ))}
          </Menu>
        </Aside>
        
        <Content className="content">
          {renderContent()}
        </Content>
      </Layout>
    </div>
  );
}

export default App;
