package config

import "github.com/spf13/viper"

// LogLevel gets log level
func LogLevel() string {
	return viper.GetString("log.level")
}

// Addr gets server address
func Addr() string {
	return viper.GetString("net.addr")
}

// LogPath gets log path
func LogPath() string {
	return viper.GetString("log.path")
}
