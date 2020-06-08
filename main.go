package main

import (
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gocv.io/x/gocv"
)

const dataDir = "images"

func main() {
	extList := []string{
		"jpg",
		"png",
		"jpeg",
	}

	log.Print("Pattern started..")

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	lower := gocv.NewMatFromScalar(gocv.NewScalar(30.0, 150.0, 50.0, 0.0), gocv.MatTypeCV8UC3)
	upper := gocv.NewMatFromScalar(gocv.NewScalar(255.0, 255.0, 180.0, 0.0), gocv.MatTypeCV8UC3)

	for _, file := range files {
		extSplited := strings.Split(file.Name(), ".")
		ext := extSplited[len(extSplited)-1]

		if !stringInSlice(ext, extList) {
			return
		}

		img := filepath.Join(dataDir, file.Name())

		photo := gocv.IMRead(img, gocv.IMReadColor)

		imgCopy := photo.Clone()

		gocv.CvtColor(photo, &imgCopy, gocv.ColorBGRToHSV)

		mask := gocv.NewMat()

		gocv.InRange(imgCopy, lower, upper, &mask)

		gocv.IMWrite("dist/"+extSplited[0]+"-filtred."+ext, mask)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
