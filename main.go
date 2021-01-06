package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

var path, compressPath string

func init() {

	// 图片源文件夹
	path = "static/imgs/"

	// 图片压缩生成文件夹
	compressPath = "static/imgs_compress/"

	// 创建文件夹
	os.MkdirAll(path, os.ModePerm)

	os.MkdirAll(compressPath, os.ModePerm)
}

func main() {

	// 读取列表
	list, err := ioutil.ReadDir(path)

	if err != nil {

		fmt.Println(err)

		return
	}

	// 遍历文件夹
	for _, info := range list {

		// 获取后缀名
		temp := strings.Split(info.Name(), ".")

		suffix := temp[len(temp)-1]

		// 过滤图片
		if suffix == "jpg" {

			fmt.Println(info.Name(), "compressing")

			jpgCompressing(info.Name())

			continue
		}

		if suffix == "png" {

			continue
		}

		fmt.Println("Unkown type")
	}
}

func jpgCompressing(fileName string) {

	// 读取文件
	file, err := os.Open(path + fileName)

	if err != nil {

		fmt.Println(err)

		return
	}

	defer file.Close()

	// 创建 20 质量文件
	fileCompressing20, err := os.Create(compressPath + "l_" + fileName)

	if err != nil {

		fmt.Println(err)

		return
	}

	// 创建 50 质量文件
	fileCompressing50, err := os.Create(compressPath + "g_" + fileName)

	if err != nil {

		fmt.Println(err)

		return
	}

	defer fileCompressing20.Close()

	defer fileCompressing50.Close()

	// 解码源文件
	img, err := jpeg.Decode(file)

	if err != nil {

		fmt.Println(err)

		return
	}

	// 重设质量
	jpeg.Encode(fileCompressing20, img, &jpeg.Options{20})

	jpeg.Encode(fileCompressing50, img, &jpeg.Options{50})
}

func pngCompressing() {

}

func gifCompressing() {

}
