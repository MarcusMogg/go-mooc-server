package api

import (
	"fmt"
	"os/exec"
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
	file, _ := c.FormFile("file")
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", course, "/", filename}, "")
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	dst := []string{folder, "/", filename, ".", format}
	dstLen := len(dst) - 1
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	if format != "mp4" {
		src := make([]string, dstLen+1)
		copy(src, dst)
		dst[dstLen] = "mp4"
		param := []string{"-i", strings.Join(src, ""), "-y", "-c:v", "libx264", "-strict", "-2", strings.Join(dst, "")}
		cmd := exec.Command("ffmpeg", param...)
		if err := cmd.Run(); err != nil {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		}
	} else {
		response.OkWithMessage("upload success", c)
	}
	video := &entity.Video{VideoName: filename, Course: course, Uploader: uploader, Format: "mp4", Path: folder}
	if err := service.SaveVideo(video); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	var id int
	global.GDB.Table("videos").Select("id").Where("video_name = ?", video.VideoName).Scan(&id)
	global.UPLOADQUEUE <- fmt.Sprintf("%d", id)
}
