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
	if courseID <= 0 {
		response.FailWithMessage("错误的课程id", c)
		return
	}
	seq, _ := strconv.Atoi(c.PostForm("seq"))
	name := c.PostForm(c.PostForm("name"))
	ins := c.PostForm(c.PostForm("ins"))
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("file not exist", c)
	}
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", fmt.Sprintf("%v", courseID), "/", fmt.Sprintf("%v", seq)}, "")

	video := &entity.Video{VideoName: filename, Seq: seq, Format: format, Path: folder, Name: name, Ins: ins}

	err = service.SaveVideo(video, uint(courseID))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	dst := []string{folder, "/", filename, ".", format}
	os.MkdirAll(folder, os.ModePerm)
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("upload success", c)
		var QResult entity.CourseVideoResult
		global.GDB.Table("course_videos").Select("course_videos.course_id", "videos.seq", "course_videos.video_id").Joins("JOIN videos ON course_videos.video_id = videos.id").Where("videos.seq = ? AND course_videos.course_id = ?", video.Seq, courseID).Scan(&QResult)
		global.UPLOADQUEUE <- fmt.Sprintf("%v", QResult.VideoID)
	}

}

