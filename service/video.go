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
func SaveVideo(v *entity.Video) error {
	var tmp entity.Video
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("course_id = ? AND seq = ?", v.CourseID, v.Seq).First(&tmp)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tx.Create(v).Error
		}
		copyVideoInfo(&tmp, v)
		tx.Save(&tmp)
		v.ID = tmp.ID
		return nil
	})
}

// ModifyVideo 修改视频信息
func ModifyVideo(v *entity.Video) error {
	var tmp entity.Video
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("course_id = ? AND seq = ?", v.CourseID, v.Seq).First(&tmp)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		tmp.Name = v.Name
		tmp.Introduction = v.Introduction
		tx.Save(&tmp)
		return nil
	})
}

// DropVideo 删除视频
func DropVideo(vid uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		var v entity.Video
		err := tx.Where("id = ?", vid).First(&v).Error
		if err != nil {
			return err
		}
		os.RemoveAll(v.Path)
		if err := tx.Delete(&v).Error; err != nil {
			return err
		}
		return nil
	})
}

// copyVideoInfo 复制video信息
func copyVideoInfo(video *entity.Video, copiedVideo *entity.Video) {
	video.VideoName = copiedVideo.VideoName
	video.Format = copiedVideo.Format
	video.Name = copiedVideo.Name
	video.Introduction = copiedVideo.Introduction
}

// transcoding 视频转码
func transcoding(vid uint) {
	var video entity.Video
	global.GDB.First(&video, vid)

	if video.Format != ".mp4" {
		param := []string{"-i", video.Path + video.VideoName + video.Format, "-y", "-c:v", "libx264", "-strict", "-2", video.Path + video.VideoName + ".mp4"}
		cmd := exec.Command("ffmpeg", param...)
		if err := cmd.Run(); err != nil {
			return
		}
	}
	global.GDB.Model(&video).Update("format", ".mp4")

	param := []string{"-y", "-i", video.Path + video.VideoName + ".mp4", "-vcodec", "copy", "-acodec", "copy", "-vbsf",
		"h264_mp4toannexb", video.Path + video.VideoName + ".ts"}
	cmd := exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		return
	}
	global.GDB.Model(&video).Update("format", ".ts")

	param = []string{"-i", video.Path + video.VideoName + ".ts", "-c", "copy", "-map", "0", "-f", "segment", "-segment_list",
		video.Path + video.VideoName + ".m3u8", "-segment_time", "5", video.Path + video.VideoName + "-%03d.ts"}
	cmd = exec.Command("ffmpeg", param...)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out), err.Error())
	}
	global.GDB.Model(&video).Update("format", "m3u8")
}

// Transcoding 视频转码
func Transcoding() {
	for id := range global.UPLOADQUEUE {
		id, _ := strconv.Atoi(id)
		vid := uint(id)
		go transcoding(vid)
	}
}

// GeneratePath 生成视频路径
func GeneratePath(video *entity.Video) string {
	return video.Path + "/" + video.VideoName + video.Format
}
