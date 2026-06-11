package main

import (
	"ccon/cmd"
	"fmt"
)

func main() {
	r, g, b, err := cmd.ToRGB("CAF31D")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("R: %d, G: %d, B: %d\n", r, g, b)

	h, s, l := cmd.ToHSL([3]uint8{r, g, b})
	fmt.Printf("H: %f, S: %f, L: %f\n", h, s, l)

	L, c, h, err := cmd.ToOKLCH("CAF31D")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("H: %f, S: %f, L: %f\n", L, c, h)
}
