package api

import (
	"fmt"
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

	courseId := strconv.Atoic.PostForm("courseId")
	order, _ := strconv.Atoi(c.PostForm("order"))
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("file not exist", c)
	}
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", courseId, "/", filename}, "")

	video := &entity.Video{VideoName: filename, Order: order, Format: format, Path: folder}
	
	if err := service.SaveVideo(video); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	var id int
	global.GDB.Table("videos").Select("id").Where("video_name = ? AND course = ?", video.VideoName, video.Course).Scan(&id)
	global.UPLOADQUEUE <- id

	dst := []string{folder, "/", filename, ".", format}
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	response.OkWithMessage("upload success", c)
}
