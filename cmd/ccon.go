package cmd

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func ToHSL(rgb [3]uint8) (h, s, l float64) {

	//normalize rgb values
	r := float64(rgb[0]) / 255.0
	g := float64(rgb[1]) / 255.0
	b := float64(rgb[2]) / 255.0

	cmax := max(r, g, b)
	cmin := min(r, g, b)
	delta := cmax - cmin

	// lightness
	l = (cmax + cmin) / 2

	//saturation
	s = delta / (1 - math.Abs(2*l-1))

	//hue
	if cmax == r {
		h = 60 * math.Mod((g-b)/delta, 6)
	}

	if cmax == g {
		h = 60 * ((b-r)/delta + 2)
	}

	if cmax == b {
		h = 60 * ((r-g)/delta + 4)
	}

	if h < 0 {
		h += 360
	}

	return h, s * 100, l * 100
}

func toOKLCH() {

}

func ToRGB(hex string) (r, g, b uint8, err error) {
	if len(hex) < 6 || len(hex) > 7 {
		err = errors.New("invalid hex string")
	}

	if len(hex) == 6 {
		var rgbarray [3]uint8
		val, err := strconv.ParseUint(hex[0:2], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[0] = uint8(val)
		val, err = strconv.ParseUint(hex[2:4], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[1] = uint8(val)
		val, err = strconv.ParseUint(hex[4:6], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[2] = uint8(val)

		r, g, b = rgbarray[0], rgbarray[1], rgbarray[2]
	}

	if len(hex) == 7 {
		strings.TrimPrefix(hex, "#")
		var rgbarray [3]uint8
		val, err := strconv.ParseUint(hex[0:2], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[0] = uint8(val)
		val, err = strconv.ParseUint(hex[2:4], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[1] = uint8(val)
		val, err = strconv.ParseUint(hex[4:6], 16, 8)
		if err != nil {
			err = errors.New("invalid hex string")
		}
		rgbarray[2] = uint8(val)

		r, g, b = rgbarray[0], rgbarray[1], rgbarray[2]
	}

	return r, g, b, err
}

func toHEX() {

}
