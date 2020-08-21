package video

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

func OptimVideoH264(trans *transcoder.Transcoder, inputPath, outputPath string) error {
	err := trans.Initialize(inputPath, outputPath)

	if err != nil {
		return err
	}

	trans.MediaFile().SetVideoCodec("libx264")

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println(msg)
	}

	err = <-done

	return err
}
