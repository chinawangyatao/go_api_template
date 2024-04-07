package miniUpload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/utils"
	"strings"
)

// 文件上传

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.L.Error("no file!", err)
		fmt.Println("no file!")
		return
	}
	//获取文件保存并且返回文件访问路径
	type FileInfo struct {
		filePath string
	}
	fileInfo := &FileInfo{}

	// 获取文件的扩展名
	fileExt := getFileExt(file.Filename)
	// 根据扩展名判断文件类型
	fileType := getFileType(fileExt)
	if fileType == "images" {
		err = c.SaveUploadedFile(file, "./static/images/"+file.Filename)
		fileInfo.filePath = c.Request.Host + "/static/images/" + file.Filename // 构建文件的访问路径
	} else if fileType == "video" {
		err = c.SaveUploadedFile(file, "./static/videos/"+file.Filename)
		fileInfo.filePath = c.Request.Host + "/static/videos/" + file.Filename // 构建文件的访问路径
	}
	if err != nil {
		utils.L.Info("save file error :", err)
		return
	}

	utils.Success(c, map[string]interface{}{
		"images": fileInfo.filePath,
	})
}

func Uploads(c *gin.Context) {
	// 解析表单数据，限制文件大小为 32MB
	err := c.Request.ParseMultipartForm(64 << 20)
	if err != nil {
		utils.L.Error("failed to parse multipart form:", err)
		utils.Failed(c, 400, err.Error())
		return
	}

	// 获取上传的文件列表
	files := c.Request.MultipartForm.File["file"]

	var fileURLs []string

	// 遍历文件列表，逐个保存文件
	for _, fileHeader := range files {
		_, err := fileHeader.Open()
		if err != nil {
			utils.L.Error("failed to open file:", err)
			utils.Failed(c, 400, err.Error())
			return
		}

		// 获取文件的扩展名
		fileExt := getFileExt(fileHeader.Filename)
		// 根据扩展名判断文件类型
		fileType := getFileType(fileExt)

		var savePath string
		switch fileType {
		case "images":
			savePath = "./static/images/" + fileHeader.Filename
		case "video":
			savePath = "./static/videos/" + fileHeader.Filename
		default:
			savePath = "./static/unknown/" + fileHeader.Filename
		}

		// 保存文件到服务器
		if err := c.SaveUploadedFile(fileHeader, savePath); err != nil {
			utils.L.Error("failed to save file:", err)
			utils.Failed(c, 400, err.Error())
			return
		}

		// 构建文件的访问路径，并添加到 fileURLs 切片中
		fileURL := c.Request.Host + strings.TrimLeft(savePath, ".")
		fileURLs = append(fileURLs, fileURL)
	}

	utils.Success(c, map[string]interface{}{
		"files": fileURLs,
	})
}

func getFileExt(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

func getFileType(fileExt string) string {
	switch fileExt {
	case "jpg", "jpeg", "png", "gif":
		return "images"
	case "mp4", "avi", "mov", "mkv":
		return "video"
	case "pdf":
		return "pdf"
	default:
		return "unknown"
	}
}
