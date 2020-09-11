package service

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"server/model/entity"
	"strconv"

	"gorm.io/gorm"
)

// SaveVideo 保存视频信息
func SaveVideo(v *entity.Video, courseID int) error {
	var tmp entity.Video
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("seq = ?", v.Seq).First(&tmp)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			var id int
			tx.Create(v)
			tx.Table("videos").Select("id").Where("seq = ?", v.Seq).Scan(&id)
			courseVideo := &entity.CourseVideo{CourseID: courseID, VideoID: id}
			return tx.Create(courseVideo).Error
		}
		tmp.Format = v.Format
		tx.Save(&tmp)
		return os.RemoveAll(tmp.Path)
	})
}

func upload(i int) {
	var video entity.Video
	global.GDB.First(&video, i)
	if video.Format != "mp4" {
		param := []string{"-i", video.Path + "/" + video.VideoName + "." + video.Format, "-y", "-c:v", "libx264", "-strict", "-2", video.Path + "/" + video.VideoName + ".mp4"}
		cmd := exec.Command("ffmpeg", param...)
		if err := cmd.Run(); err != nil {
			return
		}
	}
	global.GDB.Model(&video).Update("format", "mp4")
	param := []string{"-y", "-i", video.Path + "/" + video.VideoName + ".mp4", "-vcodec", "copy", "-acodec", "copy", "-vbsf",
		"h264_mp4toannexb", video.Path + "/" + video.VideoName + ".ts"}
	cmd := exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		return
	}
	global.GDB.Model(&video).Update("format", "ts")
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
	for i := range global.UPLOADQUEUE {
		id, _ := strconv.Atoi(i)
		go upload(id)
	}
}
