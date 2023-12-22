package asciiartweb

import (
	"image/color"
	"strings"
	"strconv"

)

func Colorize(hex string) (color.RGBA, error){
	hex = strings.TrimPrefix(hex, "#")

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}, nil
}