package image

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"resource-optim/internal/pkg/pngquant"
)

const (
	optimSpeed     = "3"
	defaultQuality = 75
)

func OptimImage(inputPath, outputPath string) error {
	entryImg, err := LoadImage(inputPath)

	if err != nil {
		return err
	}
	log.Println("start compress: ", inputPath)
	outputImg, err := pngquant.Compress(entryImg, optimSpeed)
	if err != nil {
		return err
	}
	log.Println("finish compress: ", inputPath)

	err = SaveImage(outputPath, outputImg)

	return err
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

// 存储
func SaveImage(path string, m image.Image) error {
	var opt jpeg.Options
	opt.Quality = defaultQuality
	out, err := os.Create(path)
	if err != nil {
		log.Printf("Error creating image file: %+v\n", err)
		return err
	}
	return jpeg.Encode(out, m, &opt)
}
