package api

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// uploadFile 上传文件
func uploadFile(filePath string, c *gin.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileName := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filePath+fileName); err != nil {
		return err
	}

	return nil
}

// getFileInfo 拆分文件名和格式
func getFileInfo(fileName string) (string, string) {
	tmp := strings.Split(fileName, ".")
	tmpLen := len(tmp)
	return strings.Join(tmp[:tmpLen-1], "."), "." + tmp[tmpLen-1]
}
