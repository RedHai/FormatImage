package main

import (
	"fmt"
	"github.com/RedHai/FormatImage/resize"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage : FormatImage [image dir]")
		return
	}

	generatePairImageWithDir(os.Args[1])

}

func generatePairImageWithDir(dirName string) {

	fileInfoArr, err := ioutil.ReadDir(dirName)

	if err != nil {
		panic(err)
	}

	for i := 0; i != len(fileInfoArr); i++ {

		if fileInfoArr[i].IsDir() {
			dirNewName := filepath.Join(dirName, fileInfoArr[i].Name())
			generatePairImageWithDir(dirNewName)
		} else {
			r, _ := regexp.Compile(".*\\.png")
			if r.MatchString(fileInfoArr[i].Name()) {
				imgPath := filepath.Join(dirName, fileInfoArr[i].Name())
				generatePairImage(imgPath)
			}
		}
	}

}

func generatePairImage(imgName string) {

	f, err := os.Open(imgName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer f.Close()

	m, err := png.Decode(f)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	newImg := resize.Resize(m, image.Rect(0, 0, m.Bounds().Max.X, m.Bounds().Max.Y),
		m.Bounds().Max.X/2, m.Bounds().Max.Y/2)

	stringArr := strings.Split(imgName, ".")

	var imgNewName string = ""

	if len(stringArr) == 2 {
		imgNewName = stringArr[0] + "@2x." + stringArr[1]
	}

	if imgNewName == "" {
		fmt.Println("inValid file name")
		return
	}

	os.Rename(imgName, imgNewName)

	outFile, fErr := os.Create(imgName)
	if fErr != nil {
		panic(fErr)
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(imgName, ",", imgNewName+"\n")
	}
}
