package top

import (
	"image"
	"image/png"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/nfnt/resize"
	"github.com/s-aska/favicon/contrlib/ico"
)

// Root /
func Root(c echo.Context) error {
	data := map[string]interface{}{}

	return c.Render(http.StatusOK, "index.html", data)
}

// Auto 自動生成
func Auto(c echo.Context) error {
	r := c.Request()
	w := c.Response()

	imgs := []image.Image{}
	sizes := []uint{48, 32, 16}
	file, _, err := r.FormFile("192")
	if err != nil {
		return err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	imgs = append(imgs, img)

	var interpolationFunction resize.InterpolationFunction
	algorithm := r.FormValue("algorithm")
	switch algorithm {
	case "Mitchell-Netravali":
		interpolationFunction = resize.MitchellNetravali
	case "Nearest-Neighbor":
		interpolationFunction = resize.NearestNeighbor
	case "Bilinear":
		interpolationFunction = resize.Bilinear
	case "Bicubic":
		interpolationFunction = resize.Bicubic
	case "Lanczos2":
		interpolationFunction = resize.Lanczos2
	case "Lanczos3":
		interpolationFunction = resize.Lanczos3
	}

	log.Infof("i:%d", interpolationFunction)

	for _, size := range sizes {
		resizeImage := resize.Resize(size, 0, img, interpolationFunction)
		imgs = append(imgs, resizeImage)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=favicon.ico")
	return ico.Encode(w, imgs...)
}

// Manual 各サイズ指定
func Manual(c echo.Context) error {
	r := c.Request()
	w := c.Response()

	images := []image.Image{}
	sizes := []int{152, 48, 32, 16}

	for _, size := range sizes {
		file, _, err := r.FormFile(strconv.Itoa(size))
		if err != nil {
			return err
		}
		defer file.Close()
		image, err := png.Decode(file)
		if err != nil {
			return err
		}
		images = append(images, image)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=favicon.ico")
	return ico.Encode(w, images...)
}
