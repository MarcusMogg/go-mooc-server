package api

import (
	"fmt"
	"server/global"
	"server/model/entity"
	"server/model/response"
	"server/service"
	"strings"

	"github.com/gin-gonic/gin"

	"os"
)

// Upload 上传文件
func Upload(c *gin.Context) {

	uploader := c.PostForm("uploader")
	course := c.PostForm("course")
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("file not exist", c)
	}
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", course, "/", filename}, "")

	video := &entity.Video{VideoName: filename, Course: course, Uploader: uploader, Format: format, Path: folder}
	if err := service.SaveVideo(video); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	var id int
	global.GDB.Table("videos").Select("id").Where("video_name = ? AND course = ?", video.VideoName, video.Course).Scan(&id)
	global.UPLOADQUEUE <- fmt.Sprintf("%d", id)

	dst := []string{folder, "/", filename, ".", format}
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	response.OkWithMessage("upload success", c)
}
