package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	log "github.com/sirupsen/logrus"
)

// CheckAndConvertSVGToPNG checks if a PNG file exists for the given SVG,
// and if not, converts the SVG to PNG
func CheckAndConvertSVGToPNG(svgPath string) (string, error) {
	// Generate the output PNG path
	pngPath := strings.TrimSuffix(svgPath, filepath.Ext(svgPath)) + ".png"

	// Check if PNG already exists
	if _, err := os.Stat(pngPath); err == nil {
		// PNG exists, return its path
		return pngPath, nil
	}

	// PNG doesn't exist, we need to convert the SVG
	log.Infof("Converting SVG to PNG: %s -> %s", svgPath, pngPath)

	// Read the SVG file
	svgData, err := os.ReadFile(svgPath)
	if err != nil {
		return "", fmt.Errorf("failed to read SVG file: %v", err)
	}

	// Parse the SVG content and render to an image
	// Note: This is a simple placeholder implementation
	// A real SVG renderer would be much more complex
	img, err := renderSimpleSVG(svgData, 100, 100)
	if err != nil {
		return "", fmt.Errorf("failed to render SVG: %v", err)
	}

	// Create the PNG file
	outFile, err := os.Create(pngPath)
	if err != nil {
		return "", fmt.Errorf("failed to create PNG file: %v", err)
	}
	defer outFile.Close()

	// Encode the image as PNG
	if err := png.Encode(outFile, img); err != nil {
		return "", fmt.Errorf("failed to encode PNG: %v", err)
	}

	return pngPath, nil
}

// renderSimpleSVG is a placeholder function that creates a simple lander image
// In a real application, you would use a proper SVG parser and renderer
func renderSimpleSVG(svgData []byte, width, height int) (image.Image, error) {
	// Create a new RGBA image
	img := ebiten.NewImage(width, height)
	img.Fill(color.RGBA{0, 0, 0, 0}) // Transparent background

	// Draw a simplified lander shape
	// This is a very basic representation and doesn't actually parse the SVG

	// Draw the main body (rectangle)
	drawRect(img, 35, 40, 30, 30, color.RGBA{209, 213, 219, 255})

	// Draw the top part
	drawRect(img, 45, 30, 10, 10, color.RGBA{156, 163, 175, 255})
	drawRect(img, 47, 25, 6, 5, color.RGBA{156, 163, 175, 255})

	// Draw windows
	drawRect(img, 42, 45, 6, 6, color.RGBA{96, 165, 250, 255})
	drawRect(img, 52, 45, 6, 6, color.RGBA{96, 165, 250, 255})

	// Draw landing legs
	drawRect(img, 30, 70, 5, 15, color.RGBA{156, 163, 175, 255})
	drawRect(img, 20, 85, 15, 5, color.RGBA{156, 163, 175, 255})
	drawRect(img, 65, 70, 5, 15, color.RGBA{156, 163, 175, 255})
	drawRect(img, 65, 85, 15, 5, color.RGBA{156, 163, 175, 255})

	// Draw antennas
	drawRect(img, 31, 35, 4, 8, color.RGBA{156, 163, 175, 255})
	drawRect(img, 25, 33, 6, 2, color.RGBA{156, 163, 175, 255})
	drawRect(img, 65, 35, 4, 8, color.RGBA{156, 163, 175, 255})
	drawRect(img, 69, 33, 6, 2, color.RGBA{156, 163, 175, 255})

	// Draw thruster
	drawRect(img, 47, 70, 6, 5, color.RGBA{156, 163, 175, 255})

	// Draw control panels
	drawRect(img, 36, 55, 28, 3, color.RGBA{75, 85, 99, 255})
	drawRect(img, 36, 60, 28, 2, color.RGBA{75, 85, 99, 255})

	// Convert ebiten.Image to standard image.Image
	buf := &bytes.Buffer{}
	png.Encode(buf, img)
	pngImg, _, _ := image.Decode(buf)

	return pngImg, nil
}

// drawRect draws a filled rectangle on the image
func drawRect(img *ebiten.Image, x, y, width, height float32, clr color.RGBA) {
	vector.DrawFilledRect(img, x, y, width, height, clr, true)
}
