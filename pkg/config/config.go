// Package config provides configuration management for the application.
package config

// @Title  config.go
// @Description  Configuration management for the application.
// @Author socketwang  2025/6/25 12:55
// @Update socketwang  2025/6/25 12:55

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// CfgChgHandler defines a config change event handler
type CfgChgHandler func(e fsnotify.Event)

// Init initializes configurations from file
func Init(path string, handler CfgChgHandler) error {
	c := config{path: path, handler: handler}

	err := c.init()
	if err != nil {
		return err
	}

	c.watch()
	return nil
}

type config struct {
	path    string
	handler CfgChgHandler
}

func (c *config) init() error {
	if c.path != "" {
		viper.SetConfigFile(c.path) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("../cfg") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("cfg")
	}

	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *config) watch() {
	viper.WatchConfig()
	if c.handler != nil {
		viper.OnConfigChange(c.handler)
	}
}
