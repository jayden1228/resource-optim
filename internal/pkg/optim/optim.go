package optim

import (
	"resource-optim/internal/pkg/audio"
	"resource-optim/internal/pkg/image"
	"resource-optim/internal/pkg/video"

	"github.com/gogf/gf/os/gfile"
	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	mp4Suffix        = ".mp4"
	mp3Suffix        = ".mp3"
	pngSuffix        = ".png"
	jpgSuffix        = ".jpg"
	jpegSuffix       = ".jpeg"
)

const (
	audioType = "audio"
	videoType = "video"
	imageType = "image"
	allType = "all"
)

var TypeOptimMap = map[string]func(trans *transcoder.Transcoder, fPath, outPath string)error{}

func GetOptimTypes() []string {
	return []string{videoType, audioType, imageType, allType}
}

func init() {
	TypeOptimMap[audioType] =  OptimAudioType
	TypeOptimMap[videoType] = OptimVideoType
	TypeOptimMap[imageType] = OptimImageType
	TypeOptimMap[allType] = OptimAllType
}

// 处理所有类型选择处理
func OptimAllType(trans *transcoder.Transcoder, fPath, outPath string) (err error) {
	switch gfile.Ext(fPath) {
	case mp4Suffix:
		err = video.OptimVideoH264(trans, fPath, outPath)
	case mp3Suffix:
		err = audio.OptimAudio(trans, fPath, outPath)
	case pngSuffix, jpegSuffix, jpgSuffix:
		err = image.OptimImage(fPath, outPath)
	}
	return
}

// 处理图片类型
func OptimImageType(trans *transcoder.Transcoder, fPath, outPath string) (err error) {

	if gfile.Ext(fPath) == pngSuffix || gfile.Ext(fPath) == jpegSuffix || gfile.Ext(fPath) == jpgSuffix {
		return  image.OptimImage(fPath, outPath)
	}

	return nil
}

// 处理视频类型
func OptimVideoType(trans *transcoder.Transcoder, fPath, outPath string) (err error) {

	if gfile.Ext(fPath) == mp4Suffix {
		return  video.OptimVideoH264(trans, fPath, outPath)
	}

	return nil
}

// 处理音频类型
func OptimAudioType(trans *transcoder.Transcoder, fPath, outPath string) (err error) {

	if gfile.Ext(fPath) == mp3Suffix {
		return  audio.OptimAudio(trans, fPath, outPath)
	}

	return nil
}