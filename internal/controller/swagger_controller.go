package controller

import (
	"io/ioutil"
	"mcp-manager/internal/utils/parser"
	"mcp-manager/pkg/common"
	"mime/multipart"
	"strconv"

	"mcp-manager/internal/model"
	"mcp-manager/internal/service"

	"github.com/gin-gonic/gin"
)

// SwaggerFileRequest 用于文件上传方式的参数
// swagger:model SwaggerFileRequest
// 文件上传参数结构体
type SwaggerFileRequest struct {
	// swagger:file
	// 上传的Swagger文件
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// SwaggerTextRequest 用于文本粘贴方式的参数
// swagger:model SwaggerTextRequest
// 文本内容参数结构体
type SwaggerTextRequest struct {
	// 粘贴的Swagger内容字符串，必须为有效的OpenAPI内容
	Content string `json:"content" binding:"required"`
}

// ValidateSwaggerByFile godoc
// @Summary 通过文件上传校验Swagger文档
// @Description 上传Swagger文件并进行格式和内容校验
// @Tags Swagger
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Swagger文件"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/swagger/validate/file [post]
func ValidateSwaggerByFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.Error(c, 400, "file is required")
		return
	}
	tmpFile, err := file.Open()
	if err != nil {
		common.Error(c, 500, "failed to open file")
		return
	}
	defer tmpFile.Close()
	data, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		common.Error(c, 500, "failed to read file")
		return
	}
	parserObj := parser.NewSwaggerParser()
	loader := parserObj.(*parser.OpenAPI3Parser)
	openapiLoader := loader
	doc, err := openapiLoader.ParseFromData(data)
	if err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	if err := openapiLoader.Validate(doc); err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, gin.H{"message": "swagger validated successfully"})
}

// ValidateSwaggerByText godoc
// @Summary 通过文本内容校验Swagger文档
// @Description 粘贴Swagger内容并进行格式和内容校验
// @Tags Swagger
// @Accept json
// @Produce json
// @Param data body SwaggerTextRequest true "Swagger内容"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/swagger/validate/text [post]
func ValidateSwaggerByText(c *gin.Context) {
	var req SwaggerTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "content is required")
		return
	}
	parserObj := parser.NewSwaggerParser()
	loader := parserObj.(*parser.OpenAPI3Parser)
	openapiLoader := loader
	doc, err := openapiLoader.ParseFromData([]byte(req.Content))
	if err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	if err := openapiLoader.Validate(doc); err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, gin.H{"message": "swagger validated successfully"})
}

// SwaggerServiceHandler 提供对 SwaggerService 的 HTTP 封装
type SwaggerServiceHandler struct {
	Service service.SwaggerService
}

// NewSwaggerServiceHandler 构造函数
func NewSwaggerServiceHandler(s service.SwaggerService) *SwaggerServiceHandler {
	return &SwaggerServiceHandler{Service: s}
}

// ParseAndSave godoc
// @Summary 解析并保存Swagger接口
// @Description 上传Swagger内容并保存所有接口到数据库
// @Tags Swagger
// @Accept json
// @Produce json
// @Param data body SwaggerTextRequest true "Swagger内容"
// @Success 200 {array} model.APIEndpoint
// @Failure 400 {object} map[string]string
// @Router /api/swagger/parse [post]
func (h *SwaggerServiceHandler) ParseAndSave(c *gin.Context) {
	var req SwaggerTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "content is required")
		return
	}
	endpoints, err := h.Service.ParseAndSave(c.Request.Context(), []byte(req.Content))
	if err != nil {
		common.Error(c, 400, err.Error())
		return
	}
	common.Success(c, endpoints)
}

// ListAPIEndpoints godoc
// @Summary 查询指定swaggerID下所有APIEndpoint
// @Tags Swagger
// @Produce json
// @Param swagger_id query int true "SwaggerID"
// @Success 200 {array} model.APIEndpoint
// @Failure 400 {object} map[string]string
// @Router /api/swagger/endpoints [get]
func (h *SwaggerServiceHandler) ListAPIEndpoints(c *gin.Context) {
	swaggerIDStr := c.Query("swagger_id")
	swaggerID, err := strconv.ParseUint(swaggerIDStr, 10, 64)
	if err != nil {
		common.Error(c, 400, "invalid swagger_id")
		return
	}
	endpoints, err := h.Service.ListAPIEndpoints(c.Request.Context(), uint(swaggerID))
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, endpoints)
}

// GetAPIEndpointByID godoc
// @Summary 根据ID查询APIEndpoint
// @Tags Swagger
// @Produce json
// @Param id path int true "APIEndpoint ID"
// @Success 200 {object} model.APIEndpoint
// @Failure 400 {object} map[string]string
// @Router /api/swagger/endpoint/{id} [get]
func (h *SwaggerServiceHandler) GetAPIEndpointByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.Error(c, 400, "invalid id")
		return
	}
	endpoint, err := h.Service.GetAPIEndpointByID(c.Request.Context(), uint(id))
	if err != nil {
		common.Error(c, 404, err.Error())
		return
	}
	common.Success(c, endpoint)
}

// DeleteAPIEndpoint godoc
// @Summary 删除指定APIEndpoint
// @Tags Swagger
// @Produce json
// @Param id path int true "APIEndpoint ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/swagger/endpoint/{id} [delete]
func (h *SwaggerServiceHandler) DeleteAPIEndpoint(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.Error(c, 400, "invalid id")
		return
	}
	err = h.Service.DeleteAPIEndpoint(c.Request.Context(), uint(id))
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, gin.H{"message": "deleted"})
}

// UpdateAPIEndpoint godoc
// @Summary 更新APIEndpoint
// @Tags Swagger
// @Accept json
// @Produce json
// @Param data body model.APIEndpoint true "APIEndpoint数据"
// @Success 200 {object} model.APIEndpoint
// @Failure 400 {object} map[string]string
// @Router /api/swagger/endpoint [put]
func (h *SwaggerServiceHandler) UpdateAPIEndpoint(c *gin.Context) {
	var endpoint model.APIEndpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		common.Error(c, 400, "invalid body")
		return
	}
	err := h.Service.UpdateAPIEndpoint(c.Request.Context(), &endpoint)
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, endpoint)
}

// TestAPIEndpoint godoc
// @Summary 测试APIEndpoint
// @Tags Swagger
// @Accept json
// @Produce json
// @Param data body model.APIEndpoint true "APIEndpoint数据"
// @Param base_url query string true "服务器基础URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/swagger/endpoint/test [post]
func (h *SwaggerServiceHandler) TestAPIEndpoint(c *gin.Context) {
	var endpoint model.APIEndpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		common.Error(c, 400, "invalid body")
		return
	}
	baseURL := c.Query("base_url")
	if baseURL == "" {
		common.Error(c, 400, "base_url is required")
		return
	}
	resp, err := h.Service.TestAPIEndpoint(c.Request.Context(), &endpoint, baseURL)
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}
	common.Success(c, gin.H{"response": resp})
}
