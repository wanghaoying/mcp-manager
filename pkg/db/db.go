// Package db TODO
package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config defines the DB config
type Config struct {
	User        string `mapstructure:"user" json:"user"`
	Password    string `mapstructure:"password" json:"password"`
	Addr        string `mapstructure:"addr" json:"addr"`
	Name        string `mapstructure:"name" json:"name"`
	MaxOpenConn int    `mapstructure:"max_open_conn" json:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn" json:"max_idle_conn"`
	DebugLog    bool   `mapstructure:"debug_log" json:"debug_log"`
	Type        string `mapstructure:"type" json:"type"`
}

// DBManager defines the interface for database management
type DBManager interface {
	Connect() (*gorm.DB, error)
}

// MySQLManager implements DBManager for MySQL
type MySQLManager struct {
	config Config
	db     *gorm.DB
}

// NewMySQLManager initializes MySQLManager with connection pool
func NewMySQLManager(config Config) (*MySQLManager, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Addr, config.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	return &MySQLManager{config: config, db: db}, nil
}

// Connect returns the pooled *gorm.DB instance
func (m *MySQLManager) Connect() (*gorm.DB, error) {
	if m.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return m.db, nil
}

// DBFactory creates a DBManager based on the type in Config
func DBFactory(config Config) (DBManager, error) {
	switch config.Type {
	case "mysql":
		return NewMySQLManager(config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}
}

var dbManagerMap sync.Map = sync.Map{}

// RegisterDBManager 注册 dbName 与 DBManager 的映射（线程安全，使用 sync.Map）
func RegisterDBManager(dbName string, manager DBManager) {
	dbManagerMap.Store(dbName, manager)
}

// GetDBManagerByName 根据 dbName 获取对应的 DBManager（线程安全，使用 sync.Map）
func GetDBManagerByName(dbName string) (DBManager, bool) {
	value, ok := dbManagerMap.Load(dbName)
	if !ok {
		return nil, false
	}
	manager, ok := value.(DBManager)
	return manager, ok
}
