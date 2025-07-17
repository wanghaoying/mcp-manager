// Package main
package main

// @Title  main.go
// @Description  MCP Manager Application Entry Point
// @Author  socketwang  2025/5/26 13:01
// @Update  socketwang  2025/5/26 13:01

import (
	"fmt"
	"mcp-manager/internal/model"
	"mcp-manager/internal/router"
	"mcp-manager/pkg/config"
	"mcp-manager/pkg/logger"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("cfg", "c", "./cfg/cfg.yaml", "config file path.")
)

func onCfgChg(e fsnotify.Event) {
	log.Warnf("config changed, path: %s, event: %s", e.Name, e.Op)
	if e.Op != fsnotify.Write && e.Op != fsnotify.Create {
		return
	}
}

func init() {
	pflag.Parse()

	// load configs
	err := config.Init(*cfg, onCfgChg)
	if err != nil {
		fmt.Println("init config failed!")
		os.Exit(-2)
	}

	// init debug logger
	err = logger.InitDebugLogger(config.LogLevel(), config.LogPath())
	if err != nil {
		fmt.Println("init debug logger failed!")
		os.Exit(-3)
	}

	// init db
	err = model.InitDBs()
	if err != nil {
		log.Errorf("init db failed: %v", err)
		os.Exit(-4)
	}
}

func main() {
	r := gin.Default()
	router.RegisterRoutes(r)

	iDataServer := &http.Server{
		Addr:    config.Addr(),
		Handler: r,
	}

	err := gracehttp.Serve(iDataServer)
	if err != nil {
		fmt.Println(err)
	}
}
