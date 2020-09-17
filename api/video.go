package api

import (
	"fmt"
	"os"
	"server/global"
	"server/model/entity"
	"server/model/request"
	"server/model/response"
	"server/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	video := readFormData(c)
	t := c.PostForm("type")
	if t == "upload" {
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
		response.OkWithMessage("upload success", c)
		global.UPLOADQUEUE <- fmt.Sprintf("%v", video.ID)
	} else {
		name := c.PostForm("name")
		introduction := c.PostForm("introduction")
		tempID := c.PostForm("id") 
	}
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

	introduction := c.PostForm("introduction")
	video.Introduction = introduction

	return video
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		if err := service.DropVideo(id.ID); err != nil {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
			return
		}
		response.OkWithMessage("delete success", c)
	} else {
		response.FailValidate(c)
	}
}
