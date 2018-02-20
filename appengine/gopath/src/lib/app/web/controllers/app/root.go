package app

import (
	"image"
	"image/png"
	"lib/app/ico"
	"lib/app/web"
	"net/http"
	"strconv"

	"google.golang.org/appengine/log"

	"github.com/nfnt/resize"
)

// Root トップページ
func Root(w http.ResponseWriter, r *http.Request) {
	c := web.Context(r)

	c.HTML("views/index.html")
}

// Auto 自動生成
func Auto(w http.ResponseWriter, r *http.Request) {
	c := web.Context(r)

	imgs := []image.Image{}
	sizes := []uint{48, 32, 16}
	file, _, err := r.FormFile("152")
	if err != nil {
		c.Error(err)
		return
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		c.Error(err)
		return
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

	log.Debugf(c.Context, "i:%d", interpolationFunction)

	for _, size := range sizes {
		resizeImage := resize.Resize(size, 0, img, interpolationFunction)
		imgs = append(imgs, resizeImage)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=favicon.ico")
	ico.Encode(w, imgs...)
}

// Manual 各サイズ指定
func Manual(w http.ResponseWriter, r *http.Request) {
	c := web.Context(r)

	images := []image.Image{}
	sizes := []int{152, 48, 32, 16}

	for _, size := range sizes {
		file, _, err := r.FormFile(strconv.Itoa(size))
		if err != nil {
			c.Error(err)
			return
		}
		defer file.Close()
		image, err := png.Decode(file)
		if err != nil {
			c.Error(err)
			return
		}
		images = append(images, image)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=favicon.ico")
	ico.Encode(w, images...)
}
