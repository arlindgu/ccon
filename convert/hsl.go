package convert

import "math"

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
