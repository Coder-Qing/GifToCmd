package request

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// GetGifDataByURL 获取图片数据
func GetGifDataByURL(gifURL string) (io.ReadCloser, error) {
	//本地图片转io.ReadCloser
	//判断是url还是本地地址
	if strings.HasPrefix(gifURL, "http") {
		resp, err := http.Get(gifURL)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	} else {
		file, err := os.Open(gifURL)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
}
