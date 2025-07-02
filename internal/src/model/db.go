package model

import (
	"mcp-manager/pkg/config"
	"mcp-manager/pkg/db"
)

// DBConfigFromMap 将 map[string]interface{} 转为 db.Config
func DBConfigFromMap(m map[string]interface{}) db.Config {
	cfg := db.Config{}
	if v, ok := m["user"].(string); ok {
		cfg.User = v
	}
	if v, ok := m["password"].(string); ok {
		cfg.Password = v
	}
	if v, ok := m["addr"].(string); ok {
		cfg.Addr = v
	}
	if v, ok := m["name"].(string); ok {
		cfg.Name = v
	}
	if v, ok := m["max_open_conn"].(int); ok {
		cfg.MaxOpenConn = v
	} else if v, ok := m["max_open_conn"].(float64); ok {
		cfg.MaxOpenConn = int(v)
	}
	if v, ok := m["max_idle_conn"].(int); ok {
		cfg.MaxIdleConn = v
	} else if v, ok := m["max_idle_conn"].(float64); ok {
		cfg.MaxIdleConn = int(v)
	}
	if v, ok := m["debug_log"].(bool); ok {
		cfg.DebugLog = v
	}
	if v, ok := m["type"].(string); ok {
		cfg.Type = v
	}
	return cfg
}

// InitAllDBs 初始化所有数据库连接
func InitAllDBs() (map[string]db.DBManager, error) {
	all := config.DBConfigs()
	result := make(map[string]db.DBManager)
	for name, m := range all {
		cfg := DBConfigFromMap(m)
		manager, err := db.DBFactory(cfg)
		if err != nil {
			return nil, err
		}
		result[name] = manager
		db.RegisterDBManager(name, manager)
	}
	return result, nil
}

// InitDBByName 初始化指定名称的数据库连接
func InitDBByName(name string) (db.DBManager, error) {
	m := config.DBConfig(name)
	cfg := DBConfigFromMap(m)
	return db.DBFactory(cfg)
}
