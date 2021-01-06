package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

var path string

func init() {

	path = "static/imgs/"
}

func main() {

	list, err := ioutil.ReadDir(path)

	if err != nil {

		fmt.Println("error")

		return
	}

	for _, info := range list {

		temp := strings.Split(info.Name(), ".")

		suffix := temp[len(temp)-1]

		if suffix == "jpg" {

			jpgCompressing(info.Name())

			return
		}

		if suffix == "png" {

			return
		}

		fmt.Println("Unkown type")
	}
}

func jpgCompressing(fileName string) {

	file, err := os.Open(path + fileName)

	if err != nil {

		fmt.Println(err)
	}

	defer file.Close()

	fileCompressing, err := os.Create(path + "test.jpg")

	if err != nil {

		fmt.Println(err)
	}

	defer fileCompressing.Close()

	img, err := jpeg.Decode(file)

	if err != nil {

		fmt.Println(err)
	}

	jpeg.Encode(fileCompressing, img, &jpeg.Options{50})
}

func pngCompressing() {

}
