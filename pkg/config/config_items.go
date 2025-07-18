package config

import "github.com/spf13/viper"

// LogLevel gets log level
func LogLevel() string {
	return viper.GetString("log.level")
}

// Addr gets server address
func Addr() string {
	return viper.GetString("global.net.addr")
}

// LogPath gets log path
func LogPath() string {
	return viper.GetString("log.path")
}

// DBConfigs 获取所有数据库配置（支持多库）
// 返回 map[string]map[string]interface{}，自动类型断言
func DBConfigs() map[string]map[string]interface{} {
	res := make(map[string]map[string]interface{})
	for k, v := range viper.GetStringMap("dbs") {
		if m, ok := v.(map[string]interface{}); ok {
			res[k] = m
		}
	}
	return res
}

// DBConfig 获取指定数据库名的配置
func DBConfig(name string) map[string]interface{} {
	return viper.GetStringMap("dbs." + name)
}
