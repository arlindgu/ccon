package convert

import "math"

func ToOKLCH(hex string) (float64, float64, float64, error) {

	//hex to rgb
	colorR, colorG, colorB, err := ToRGB(hex)
	if err != nil {

	}

	//rgb to linear rgb
	rN := float64(colorR) / 255.0
	gN := float64(colorG) / 255.0
	bN := float64(colorB) / 255.0
	var linearR, linearB, linearG, x, y, z float64

	if rN < 0.04045 {
		linearR = rN / 12.92
	} else {
		linearR = math.Pow((rN+0.055)/1.055, 2.4)
	}

	if gN < 0.04045 {
		linearG = gN / 12.92
	} else {
		linearG = math.Pow((gN+0.055)/1.055, 2.4)
	}

	if bN < 0.04045 {
		linearB = bN / 12.92
	} else {
		linearB = math.Pow((bN+0.055)/1.055, 2.4)
	}

	//linear rgb to xyz
	x = linearR*0.4124 + linearG*0.3576 + linearB*0.1805
	y = linearR*0.2126 + linearG*0.7152 + linearB*0.0722
	z = linearR*0.0193 + linearG*0.1192 + linearB*0.9505

	//linear to OKLab
	l_ := math.Cbrt(0.8189*x + 0.3618*y - 0.1288*z)
	m_ := math.Cbrt(0.0329*x + 0.9293*y + 0.0361*z)
	s_ := math.Cbrt(0.0482*x + 0.2643*y + 0.6338*z)

	L := 0.2104*l_ + 0.7936*m_ - 0.0040*s_
	a := 1.9779*l_ - 2.4285*m_ + 0.4506*s_
	b := 0.0259*l_ + 0.7827*m_ - 0.8086*s_

	//OKLab to OKLCH
	c := math.Sqrt(a*a + b*b)
	h := math.Atan2(b, a) * (180 / math.Pi)

	if h < 0 {
		h += 360
	}

	return L, c, h, nil
}
