package convert

import (
	"errors"
	"strconv"
	"strings"
)

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
