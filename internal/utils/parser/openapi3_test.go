package parser

import (
	"testing"
)

func TestSwaggerParser(t *testing.T) {
	// 创建解析器
	p := NewOpenAPI3Parser()

	// 测试解析
	doc, err := p.Parse("/Users/wanghao/Desktop/project/go/mcp-manager/example/test.yaml")
	if err != nil {
		t.Fatalf("Failed to parse swagger file: %v", err)
	}

	// 验证基本信息（加强空指针检查）
	if doc.Info == nil {
		t.Fatal("Swagger document missing info section")
	}

	// 检查Info的必要字段
	if doc.Info.Title == "" {
		t.Error("Info.Title is required")
	}
	if doc.Info.Version == "" {
		t.Error("Info.Version is required")
	}

	if doc.Paths == nil {
		t.Fatal("Swagger document missing paths section")
	}

	// 安全访问路径信息
	var pathCount int
	if doc.Paths != nil {
		// 正确访问Paths中的实际路径数量（非Extensions）
		pathCount = len(doc.Paths.Map())
	}

	// 增强日志输出的安全性
	title := ""
	version := ""
	if doc.Info != nil {
		title = doc.Info.Title
		version = doc.Info.Version
	}

	t.Logf("API Title: %s", title)
	t.Logf("API Version: %s", version)
	t.Logf("Number of defined paths: %d", pathCount)
}
