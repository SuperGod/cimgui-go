//go:build !exclude_cimgui_glfw
// +build !exclude_cimgui_glfw

package cimgui

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"runtime"
)

type Texture struct {
	id     ImTextureID
	Width  int
	Height int
}

func NewTextureFromRgba(rgba *image.RGBA) *Texture {
	texID := CreateTextureRgba(rgba, rgba.Bounds().Dx(), rgba.Bounds().Dy())

	if texID == 0 {
		return nil
	}

	texture := Texture{
		id:     texID,
		Width:  rgba.Bounds().Dx(),
		Height: rgba.Bounds().Dy(),
	}

	// Set finalizer
	runtime.SetFinalizer(&texture, (*Texture).release)

	return &texture
}

func (t *Texture) release() {
	DeleteTexture(t.id)
}

func (t *Texture) ID() ImTextureID {
	return t.id
}

// ImageToRgba converts image.Image to *image.RGBA.
func ImageToRgba(img image.Image) *image.RGBA {
	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba
	}
}

// LoadImage loads image from file and returns *image.RGBA.
func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error opening image file %s: %w", imgPath, err)
	}

	defer func() {
		// nolint:govet // we want to reuse this err variable here
		if err := imgFile.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", imgPath))
		}
	}()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error decoding png image: %w", err)
	}

	return ImageToRgba(img), nil
}
