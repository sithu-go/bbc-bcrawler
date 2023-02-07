package utils

import (
	"strings"
)

func GetSrcFromImg(img string) (string, error) {

	src := img[strings.Index(img, "src="):]

	src = strings.Split(src, "\"")[1]

	return src, nil

}
