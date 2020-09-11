package service

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"server/model/entity"
	"strconv"

	"gorm.io/gorm"
)

// SaveVideo 保存视频信息
func SaveVideo(v *entity.Video, courseID uint) error {
	var tmp entity.Video
	var QResult entity.CourseVideoResult
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		tx.Table("course_videos").Select("course_videos.course_id", "videos.seq", "course_videos.video_id").Joins("JOIN videos ON course_videos.video_id = videos.id").Where("videos.seq = ? AND course_videos.course_id = ?", v.Seq, courseID).Scan(&QResult)
		if QResult.Seq == 0 {
			tx.Create(v)
			courseVideo := &entity.CourseVideo{CourseID: courseID, VideoID: v.ID}
			return tx.Create(courseVideo).Error
		}
		tx.First(&tmp, QResult.VideoID)
		tmp.Format = v.Format
		tmp.VideoName = v.VideoName
		tx.Save(&tmp)
		v = &tmp
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
