package env

import (
	"errors"
	"os/exec"
	"resource-optim/config"

	"github.com/gogf/gf/os/gfile"
)

const (
	ffmpegNotExist   = "ffmpeg tool for video and audio reduce is not found, please download and install from  https://ffmpeg.org/"
	pngquantNotExist = "pngyu tool for image reduce is not found, please download and install from https://nukesaq88.github.io/Pngyu/"
)

func CheckToolRequired() error {
	if !IsFfmpegExist() {
		return errors.New(ffmpegNotExist)
	}
	if !IsPngquantExist() {
		return errors.New(pngquantNotExist)
	}
	return nil
}

func IsPngquantExist() bool {
	return gfile.Exists(config.PngquantPath)
}

func IsFfmpegExist() bool {
	cmd := exec.Command("ffmpeg", "-h")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
