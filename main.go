package main

import (
	"github.com/masterZSH/cxk/pkg/config"
	"github.com/masterZSH/cxk/pkg/gif"
	"github.com/masterZSH/cxk/pkg/request"
	"github.com/masterZSH/cxk/utils"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"strings"
)

func main() {
	// 0 1 and space
	sysConfig, colors := config.GetConfig()
	var file io.ReadCloser
	var err error
	file, err = request.GetGifDataByURL(sysConfig.GifURL)
	if err != nil {
		log.Fatal(err)
	}

	//chars := []string{"M", "8", "0", "V", "1", "i", ":", "*", "|", ".", " "}
	chars := strings.Split(sysConfig.Characters, "")
	bgColor, penColor := colors[sysConfig.BgColorType], colors[sysConfig.PenColorType]
	gif.Convert(file, chars, sysConfig.SubWidth, sysConfig.SubHeight, sysConfig.ImageOut, bgColor, penColor)
}

func init() {
	//初始化配置
	if err := utils.InitConfig(""); err != nil {
		//初始化错误了直接退出程序
		log.Fatalln("初始化配置文件出错", err.Error())
	}
}
