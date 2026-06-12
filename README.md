# ccon

A command-line tool to convert hex color codes into RGB, HSL, and OKLCH.

## Installation

```bash
git clone https://github.com/yourusername/ccon.git
cd ccon
go build -o ccon .
```

## Usage

```bash
ccon [hex] [--rgb] [--hsl] [--oklch]
```

| Flag | Output |
|------|--------|
| `--rgb` | `rgb(202, 243, 29)` |
| `--hsl` | `hsl(72, 89.6%, 53.3%)` |
| `--oklch` | `oklch(0.93, 0.22, 129.79)` |

No flag — all formats are shown.

### Examples

```bash
# Single format
ccon ff0000 --rgb
# rgb(255, 0, 0)

ccon ff0000 --hsl
# hsl(0, 100.0%, 50.0%)

# With # prefix
ccon "#ff0000" --oklch
#oklch(0.63, 0.26, 29.22)

```

## Supported Formats

| Format | Example |
|--------|---------|
| HEX input | `ff0000`, `#ff0000` |
| RGB | `rgb(255, 0, 0)` |
| HSL | `hsl(0, 100.0%, 50.0%)` |
| OKLCH | `oklch(0.63, 0.26, 29.23)` |

## Requirements

- Go 1.21+

## Dependencies

- [cobra](https://github.com/spf13/cobra) — CLI framework

## License

MIT
