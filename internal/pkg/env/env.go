package env

import (
	"errors"
	"resource-optim/config"
	"runtime"
	"strings"

	"github.com/gogf/gf/os/gfile"
	"github.com/xfrr/goffmpeg/utils"
)

const (
	ffmpegCmd = "ffmpeg"
	pngquantCmd = "pngquant"
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
	if  gfile.Exists(config.PngquantPath) {
		return true
	}

	execPngquant := GetCmdExec(pngquantCmd)
	out, err := utils.TestCmd(execPngquant[0], execPngquant[1])
	if err != nil {
		return false
	}

	// 替换默认pngquant path路径
	config.PngquantPath = strings.Replace(strings.Split(out.String(), "\n")[0], utils.LineSeparator(), "", -1)

	return true
}

func IsFfmpegExist() bool {
	execFfmpeg:= GetCmdExec(ffmpegCmd)
	_, err := utils.TestCmd(execFfmpeg[0], execFfmpeg[1])
	if err != nil {
		return false
	}
	return true
}

func GetCmdExec(cmd string) []string {
	var platform = runtime.GOOS
	var command = []string{"", cmd}

	switch platform {
	case "windows":
		command[0] = "where"
		break
	default:
		command[0] = "which"
		break
	}

	return command
}
