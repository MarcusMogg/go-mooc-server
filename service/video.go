package service

import (
	"errors"
	"fmt"
	"os/exec"
	"server/global"
	"server/model/entity"
	"strconv"

	"gorm.io/gorm"
)

// SaveVideo 保存视频信息
func SaveVideo(v *entity.Video) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("video_name = ?", v.VideoName).First(v)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tx.Create(v).Error
		}
		return errors.New("视频名已存在")
	})
}

func upload() {
	tmp := <-global.UPLOADQUEUE
	var video entity.Video
	id, _ := strconv.Atoi(tmp)
	global.GDB.Where("id = ?", id).Find(&video)
	param := []string{"-y", "-i", video.Path + "/" + video.VideoName + ".mp4", "-vcodec", "copy", "-acodec", "copy", "-vbsf",
		"h264_mp4toannexb", video.Path + "/" + video.VideoName + ".ts"}
	cmd := exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		return
	}
	param = []string{"-i", video.Path + "/" + video.VideoName + ".ts", "-c", "copy", "-map", "0", "-f", "segment", "-segment_list",
		video.Path + "/" + video.VideoName + ".m3u8", "-segment_time", "5", video.Path + "/" + video.VideoName + "-%03d.ts"}
	cmd = exec.Command("ffmpeg", param...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out), err.Error())
	}
	global.GDB.Model(&video).Update("format", "m3u8")
}

// Upload 视频转码
func Upload() {
	for {
		upload()
	}
}
