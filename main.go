package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("使用方法: %s <画像1> <画像2> <出力ファイル>", os.Args[0])
	}

	// 画像の読み込み
	img1, err := loadPNG(os.Args[1])
	if err != nil {
		log.Fatalf("画像1を読み込めませんでした: %v", err)
	}

	img2, err := loadPNG(os.Args[2])
	if err != nil {
		log.Fatalf("画像2を読み込めませんでした: %v", err)
	}

	// 画像の合成
	combined := combineImages(img1, img2)

	// 結果の保存
	file, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatalf("ファイルを開けませんでした: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, combined); err != nil {
		log.Fatalf("画像の書き込みに失敗しました: %v", err)
	}

	log.Println("画像が正常に保存されました:", os.Args[3])
}

// PNG画像の読み込み
func loadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// 2つの画像を合成
func combineImages(img1, img2 image.Image) *image.RGBA {
	bounds := img1.Bounds()
	combined := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel1 := img1.At(x, y)
			pixel2 := img2.At(x, y)

			// 明るさを比較して、より明るいピクセルを選択
			if brightness(pixel1) > brightness(pixel2) {
				combined.Set(x, y, pixel1)
			} else {
				combined.Set(x, y, pixel2)
			}
		}
	}

	return combined
}

// ピクセルの明るさを計算
func brightness(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	return (r + g + b) / 3
}
