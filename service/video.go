package service

import (
	"server/global"

)

// Upload 上传食品
func Upload (upload chan string) error {
	dst := <- global.UPLOADQUEUE
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
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("upload success", c)
	}
}