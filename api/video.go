package api

import (
	"fmt"
	"os"
	"server/global"
	"server/model/entity"
	"server/model/response"
	"server/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件
func Upload(c *gin.Context) {

	courseID, _ := strconv.Atoi(c.PostForm("courseId"))
	seq, _ := strconv.Atoi(c.PostForm("seq"))
	name := c.PostForm(c.PostForm("name"))
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("file not exist", c)
	}
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", fmt.Sprintf("%v", courseID), "/", fmt.Sprintf("%v", seq), "/", filename}, "")

	video := &entity.Video{VideoName: filename, Seq: seq, Format: format, Path: folder, Name: name}

	if err := service.SaveVideo(video, courseID); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	dst := []string{folder, "/", filename, ".", format}
	os.MkdirAll(folder, os.ModePerm)
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("upload success", c)
		var id int
		global.GDB.Table("videos").Select("id").Where("seq = ?", video.Seq).Scan(&id)
		global.UPLOADQUEUE <- fmt.Sprintf("%v", id)
	}

}
