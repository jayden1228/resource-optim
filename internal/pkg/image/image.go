package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"resource-optim/internal/pkg/pngquant"

	"github.com/gogf/gf/os/gfile"
)

const (
	optimSpeed     = "3"
	defaultQuality = 75
	pngSuffix = "png"
)

func OptimImage(inputPath, outputPath string) error {
	entryImg, err := LoadImage(inputPath)

	if err != nil {
		return err
	}
	log.Println("start compress: ", inputPath)
	outputByte, err := pngquant.Compress(entryImg, optimSpeed)
	if err != nil {
		return err
	}
	log.Println("finish compress: ", inputPath)

	return gfile.PutBytes(outputPath, outputByte)

}

// 加载图片
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// 存储jpeg
func SaveJpegImage(path string, m image.Image) error {
	var opt jpeg.Options
	opt.Quality = defaultQuality
	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}
	return jpeg.Encode(out, m, &opt)
}

// 存储png
func SavePngImage(path string, m image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}
	enc := png.Encoder{
		CompressionLevel: png.BestSpeed,
		BufferPool:       nil,
	}
	return enc.Encode(out, m)
}