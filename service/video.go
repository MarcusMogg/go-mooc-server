package service

import (
	"errors"
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

	if video.Format != "mp4" {
		param := []string{"-i", video.Path + "/" + video.VideoName + "." + video.Format, "-y", "-c:v", "libx264", "-strict", "-2", video.Path + "/" + video.VideoName + ".mp4"}
		cmd := exec.Command("ffmpeg", param...)
		if err := cmd.Run(); err != nil {
			return
		}
	}
	global.GDB.Model(&video).Update("format", "mp4")

	/* param := []string{"-i", video.Path + "/" + video.VideoName + ".mp4", "-b:v", "10000k", "-s", "640*480", video.Path + "/" + video.VideoName + ".ts"}
	cmd := exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		return
	}


	param = []string{"-y", "-i", video.Path + "/" + video.VideoName + ".mp4", "-vcodec", "copy", "-acodec", "copy", "-vbsf",
		"h264_mp4toannexb", video.Path + "/" + video.VideoName + ".ts"}
	cmd = exec.Command("ffmpeg", param...)
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
	global.GDB.Model(&video).Update("format", "m3u8")*/
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
