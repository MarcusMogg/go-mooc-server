package api

import (
	"fmt"
	"os/exec"
	"server/model/response"
	"strings"

	"github.com/gin-gonic/gin"

	"os"
)

func Upload(c *gin.Context) {
	course := c.PostForm("course")

	file, _ := c.FormFile("file")
	tmp := strings.Split(file.Filename, ".")
	filename, format := tmp[0], tmp[1]
	folder := strings.Join([]string{"video/", course, "/", filename}, "")
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	}
	dst := []string{folder, "/", filename, ".", format}
	dstLen := len(dst) - 1
	if err := c.SaveUploadedFile(file, strings.Join(dst, "")); err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("upload success", c)
	}

	if format != "mp4" {
		src := make([]string, dstLen+1)
		copy(src, dst)
		dst[dstLen] = "mp4"
		param := []string{"-i", strings.Join(src, ""), "-y", "-c:v", "libx264", "-strict", "-2", strings.Join(dst, "")}
		cmd := exec.Command("ffmpeg", param...)
		if err := cmd.Run(); err != nil {
			response.FailWithMessage(fmt.Sprintf("%v", err), c)
		} else {
			response.OkWithMessage("mp4 transform success", c)
		}
	}

	src := make([]string, dstLen+1)
	copy(src, dst)
	dst[dstLen] = "ts"
	param := []string{"-y", "-i", strings.Join(src, ""), "-vcodec", "copy", "-acodec", "copy", "-vbsf", "h264_mp4toannexb", strings.Join(dst, "")}
	cmd := exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		response.OkWithMessage(fmt.Sprintf("%v", err), c)
	}

	src[dstLen] = "ts"
	dst[dstLen] = "m3u8"
	param = []string{"-i", strings.Join(src, ""), "-c", "copy", "-map", "0", "-f", "segment", "-segment_list",
		strings.Join(dst, ""), "-segment_time", "5", strings.Join(dst[:dstLen-1], "") + "-%03d.ts"}
	cmd = exec.Command("ffmpeg", param...)
	if err := cmd.Run(); err != nil {
		response.OkWithMessage(fmt.Sprintf("%v", err), c)
	}

}
