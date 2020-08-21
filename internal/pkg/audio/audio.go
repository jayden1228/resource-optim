package audio

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	audioBitRate = "160"
)

func OptimAudio(trans *transcoder.Transcoder, inputPath, outputPath string) error {
	err := trans.Initialize(inputPath, outputPath)

	if err != nil {
		return err
	}

	trans.MediaFile().SetAudioBitRate(audioBitRate)

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println(msg)
	}

	err = <-done

	return err
}
