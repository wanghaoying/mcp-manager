package controller

import (
	"io/ioutil"
	"mcp-manager/internal/parser"
	"mcp-manager/pkg/common"
	"mime/multipart"

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
	loader := parserObj.(*parser.DefaultSwaggerParser)
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
	loader := parserObj.(*parser.DefaultSwaggerParser)
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
