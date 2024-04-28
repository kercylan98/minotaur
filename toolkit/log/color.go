package log

import "github.com/fatih/color"

// Base attributes
const (
	ColorReset Attribute = iota
	ColorBold
	ColorFaint
	ColorItalic
	ColorUnderline
	ColorBlinkSlow
	ColorBlinkRapid
	ColorReverseVideo
	ColorConcealed
	ColorCrossedOut
)

const (
	ColorResetBold Attribute = iota + 22
	ColorResetItalic
	ColorResetUnderline
	ColorResetBlinking
	_
	ColorResetReversed
	ColorResetConcealed
	ColorResetCrossedOut
)

// Foreground text colors
const (
	ColorFgBlack Attribute = iota + 30
	ColorFgRed
	ColorFgGreen
	ColorFgYellow
	ColorFgBlue
	ColorFgMagenta
	ColorFgCyan
	ColorFgWhite
)

// Foreground Hi-Intensity text colors
const (
	ColorFgHiBlack Attribute = iota + 90
	ColorFgHiRed
	ColorFgHiGreen
	ColorFgHiYellow
	ColorFgHiBlue
	ColorFgHiMagenta
	ColorFgHiCyan
	ColorFgHiWhite
)

// Background text colors
const (
	ColorBgBlack Attribute = iota + 40
	ColorBgRed
	ColorBgGreen
	ColorBgYellow
	ColorBgBlue
	ColorBgMagenta
	ColorBgCyan
	ColorBgWhite
)

// Background Hi-Intensity text colors
const (
	ColorBgHiBlack Attribute = iota + 100
	ColorBgHiRed
	ColorBgHiGreen
	ColorBgHiYellow
	ColorBgHiBlue
	ColorBgHiMagenta
	ColorBgHiCyan
	ColorBgHiWhite
)

func NewColor(attributes ...Attribute) *Color {
	c := &Color{
		Color: color.New(attributes...),
		attrs: attributes,
	}
	c.Color.EnableColor()
	return c
}
