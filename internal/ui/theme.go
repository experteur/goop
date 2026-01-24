package ui

import (
	"github.com/gdamore/tcell/v2"
)

var Theme = struct {
	// Background colors
	BgPrimary   tcell.Color
	BgSecondary tcell.Color
	BgSelected  tcell.Color

	// Text colors
	TextPrimary   tcell.Color
	TextSecondary tcell.Color
	TextDim       tcell.Color

	// Status colors
	StatusActive   tcell.Color
	StatusInactive tcell.Color
	StatusArchived tcell.Color

	// Health colors
	HealthGreen  tcell.Color
	HealthYellow tcell.Color
	HealthRed    tcell.Color

	// UI element colors
	BorderColor   tcell.Color
	TitleColor    tcell.Color
	SelectedColor tcell.Color
}{
	// Background colors (Tokyo Night Storm)
	BgPrimary:   tcell.NewRGBColor(26, 27, 38), // #1a1b26
	BgSecondary: tcell.NewRGBColor(36, 40, 59), // #24283b
	BgSelected:  tcell.NewRGBColor(41, 46, 66), // #292e42

	// Text colors
	TextPrimary:   tcell.NewRGBColor(192, 202, 245), // #c0caf5
	TextSecondary: tcell.NewRGBColor(169, 177, 214), // #a9b1d6
	TextDim:       tcell.NewRGBColor(86, 95, 137),   // #565f89

	// Status colors
	StatusActive:   tcell.NewRGBColor(158, 206, 106), // #9ece6a (green)
	StatusInactive: tcell.NewRGBColor(86, 95, 137),   // dim blue-gray
	StatusArchived: tcell.NewRGBColor(65, 72, 104),   // darker muted

	// Health colors
	HealthGreen:  tcell.NewRGBColor(158, 206, 106), // #9ece6a
	HealthYellow: tcell.NewRGBColor(224, 175, 104), // #e0af68
	HealthRed:    tcell.NewRGBColor(247, 118, 142), // #f7768e

	// UI element colors
	BorderColor:   tcell.NewRGBColor(122, 162, 247), // #7aa2f7 (blue)
	TitleColor:    tcell.NewRGBColor(187, 154, 247), // #bb9af7 (purple)
	SelectedColor: tcell.NewRGBColor(26, 27, 38),
}
