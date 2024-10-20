// Package colors provides a collection of color definitions for use with a spinner.
package colors

import "github.com/fatih/color"

// Color represents an item in the color map.
type Color int

type ColorT interface {
	Color | color.Attribute
}

const (
	// NoColor will bypass the color lookup causing characters to render
	// with the current default of the terminal.
	NoColor Color = iota

	// FgHiGreen is a foreground high intensity green color.
	FgHiGreen

	// FgHiYellow is a foreground high intensity yellow color.
	FgHiYellow

	// FgHiBlue is a foreground high intensity blue color.
	FgHiBlue

	// FgHiRed is a foreground high intensity red color.
	FgHiRed

	// FgWhite is a foreground white color.
	FgHiWhite
)

var lookup = map[Color]color.Attribute{
	FgHiGreen:  color.FgHiGreen,
	FgHiYellow: color.FgHiYellow,
	FgHiBlue:   color.FgHiBlue,
	FgHiRed:    color.FgHiRed,
	FgHiWhite:  color.FgHiWhite,
}

// GetColor returns a color.Color for the given color.
func GetColor[T any](c interface{}) *color.Color {
	if value, ok := c.(int); ok && value == 0 {
		return nil
	}

	if value, ok := c.(Color); ok {
		if val, ok := lookup[value]; ok {
			return color.New(val)
		}
	}

	if value, ok := c.(color.Attribute); ok {
		return color.New(value)
	}

	return nil
}
