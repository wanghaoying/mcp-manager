package testutil

import (
	log "github.com/sirupsen/logrus"
	"mcp-manager/internal/model"
	"mcp-manager/pkg/config"
	"mcp-manager/pkg/logger"
)

// init 初始化测试环境，包括 config、logger、db
func init() {
	// 加载配置
	configPath := "/Users/wanghao/Desktop/github/go/mcp-manager/cfg/cfg_local.yaml"
	err := config.Init(configPath, nil)
	if err != nil {
		log.Fatalf("init config failed: %v", err)
	}

	// 初始化日志
	err = logger.InitDebugLogger(config.LogLevel(), config.LogPath())
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}

	// 初始化数据库（如需真实数据库）
	err = model.InitDBs()
	if err != nil {
		log.Fatalf("init db failed: %v", err)
	}
}
