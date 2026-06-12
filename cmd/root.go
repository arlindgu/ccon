package cmd

import (
	"ccon/convert"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var toRGB bool
var toHSL bool
var toOKLCH bool

var rootCmd = &cobra.Command{
	Use:   "ccon",
	Args:  cobra.ExactArgs(1),
	Short: "ccon is a CLI tool to convert color codes.",
	Long:  `ccon is a powerful command-line tool that helps you convert color codes between different formats.`,
	Run: func(cmd *cobra.Command, args []string) {
		hex := args[0]

		if toRGB {
			r, g, b, err := convert.ToRGB(hex)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("rgb(%d, %d, %d)\n", r, g, b)

		}

		if toHSL {
			r, g, b, err := convert.ToRGB(hex)
			if err != nil {
				fmt.Println(err)
				return
			}
			h, s, l := convert.ToHSL([3]uint8{r, g, b})
			fmt.Printf("hsl(%.0f, %.1f%%, %.1f%%)\n", h, s, l)
		}

		if toOKLCH {
			L, c, h, err := convert.ToOKLCH(hex)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("oklch(%.2f, %.2f, %.2f)\n", L, c, h)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVar(&toRGB, "rgb", false, "Output RGB")
	rootCmd.Flags().BoolVar(&toHSL, "hsl", false, "Output HSL")
	rootCmd.Flags().BoolVar(&toOKLCH, "oklch", false, "Output OKLCH")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
