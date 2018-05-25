package blush

import (
	"fmt"
	"strconv"
	"strings"
)

// Some stock colours. There will be no colouring when NoColour is used.
var (
	NoColour  = Colour{-1, -1, -1}
	FgRed     = Colour{255, 0, 0}
	FgBlue    = Colour{0, 0, 255}
	FgGreen   = Colour{0, 255, 0}
	FgBlack   = Colour{0, 0, 0}
	FgWhite   = Colour{255, 255, 255}
	FgCyan    = Colour{0, 255, 255}
	FgMagenta = Colour{255, 0, 255}
	FgYellow  = Colour{255, 255, 0}
)

// DefaultColour is foreground blue.
var DefaultColour = FgBlue

// Colour is a RGB colour scheme. R, G and B should be between 0 and 255.
type Colour struct {
	R, G, B int
}

// Colourise wraps the input between colours defined in c for terminals.
func Colourise(input string, c Colour) string {
	return fmt.Sprintf("%s%s%s", format(c), input, unformat())
}

func format(c Colour) string {
	return fmt.Sprintf("\033[38;5;%dm", colour(c.R, c.G, c.B))
}

func unformat() string {
	return "\033[0m"
}

func colour(red, green, blue int) int {
	return 16 + baseColor(red, 36) + baseColor(green, 6) + baseColor(blue, 1)
}

func baseColor(value int, factor int) int {
	return int(6*float64(value)/256) * factor
}

func colorFromArg(colour string) Colour {
	if strings.HasPrefix(colour, "#") {
		return hexColour(colour)
	}
	return stockColour(colour)
}

func stockColour(colour string) Colour {
	c := DefaultColour
	switch colour {
	case "r", "red":
		c = FgRed
	case "b", "blue":
		c = FgBlue
	case "g", "green":
		c = FgGreen
	case "bl", "black":
		c = FgBlack
	case "w", "white":
		c = FgWhite
	case "cy", "cyan":
		c = FgCyan
	case "mg", "magenta":
		c = FgMagenta
	case "yl", "yellow":
		c = FgYellow
	case "no-colour", "no-color":
		c = NoColour
	}
	return c
}

func hexColour(colour string) Colour {
	var r, g, b int
	colour = strings.TrimPrefix(colour, "#")
	switch len(colour) {
	case 3:
		c := strings.Split(colour, "")
		r = getInt(c[0] + c[0])
		g = getInt(c[1] + c[1])
		b = getInt(c[2] + c[2])
	case 6:
		c := strings.Split(colour, "")
		r = getInt(c[0] + c[1])
		g = getInt(c[2] + c[3])
		b = getInt(c[4] + c[5])
	default:
		return DefaultColour
	}
	for _, n := range []int{r, g, b} {
		if n < 0 {
			return DefaultColour
		}
	}
	return Colour{R: r, G: g, B: b}
}

// getInt returns a number between 0-255 from a hex code. If the hex is not
// between 00 and ff, it returns -1.
func getInt(hex string) int {
	d, err := strconv.ParseInt("0x"+hex, 0, 64)
	if err != nil || d > 255 || d < 0 {
		return -99
	}
	return int(d)
}
