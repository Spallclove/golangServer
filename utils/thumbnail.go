package utils

import (
	"os/exec"
)

// GenerateThumbnail 使用 ffmpeg 生成视频缩略图
func GenerateThumbnail(videoPath, thumbnailPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:05", "-vframes", "1", thumbnailPath)
	return cmd.Run()
}
