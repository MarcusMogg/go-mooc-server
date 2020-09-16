package api

import (
	"fmt"
	"os"
	"server/global"
	"server/model/entity"
	"server/model/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	video := readFormData(c)
	err := service.CourseExist(video.CourseID)
	if err != nil {
		response.FailWithMessage("课程id不存在", c)
		return
	}

	os.RemoveAll(video.Path)
	video.VideoName, video.Format, err = uploadFile(video.Path, c)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	if err = service.SaveVideo(video); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}

	response.OkWithMessage("视频上传成功", c)
	global.UPLOADQUEUE <- fmt.Sprintf("%v", video.ID)
}

func readFormData(c *gin.Context) *entity.Video {
	video := &entity.Video{}

	cidSnap, _ := strconv.Atoi(c.PostForm("courseId"))
	cid := uint(cidSnap)
	video.CourseID = cid

	seqSnap, _ := strconv.Atoi(c.PostForm("seq"))
	seq := uint(seqSnap)
	video.Seq = seq

	video.Path = "source/video/" + fmt.Sprintf("%v", cid) + "/" + fmt.Sprintf("%v", seq) + "/"

	name := c.PostForm("name")
	video.Name = name

	ins := c.PostForm("ins")
	video.Introduction = ins

	return video
}
