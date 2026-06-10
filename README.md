# ccon

A command-line tool to convert color codes between different formats.

## Features

- Convert HEX to RGB
- Convert RGB to HSL

## Installation

```bash
git clone https://github.com/yourusername/ccon.git
cd ccon
go build -o ccon .
```

## Usage

```bash
ccon [command]
```

### Example

```go
r, g, b, err := cmd.ToRGB("CAF31D")
// R: 202, G: 243, B: 29

h, s, l := cmd.ToHSL([3]uint8{r, g, b})
// H: 72.00, S: 89.58, L: 53.33
```

## Supported Formats

| Format | Description              | Example           |
|--------|--------------------------|-------------------|
| HEX    | 6-digit hex color string | `CAF31D`, `#CAF31D` |
| RGB    | Red, Green, Blue (0–255) | `202, 243, 29`    |
| HSL    | Hue, Saturation, Lightness | `72°, 89.6%, 53.3%` |

## API

### `ToRGB(hex string) (r, g, b uint8, err error)`

Converts a hex color string (with or without `#`) to RGB values.

### `ToHSL(rgb [3]uint8) (h, s, l float64)`

Converts an RGB triplet to HSL values. Hue is in degrees (0–360), saturation and lightness are percentages (0–100).

## Requirements

- Go 1.21+

## Dependencies

- [cobra](https://github.com/spf13/cobra) — CLI framework

## License

MIT
