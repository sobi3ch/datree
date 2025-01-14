package printer

import (
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

type theme struct {
	Colors struct {
		Green  *color.Color
		Yellow *color.Color
		Red    *color.Color
		White  *color.Color
		Error  *color.Color
	}
	Spacing struct {
		Default string
	}
	Emoji struct {
		Error      string
		Suggestion string
	}
}

func createTheme() *theme {
	return &theme{
		Colors: struct {
			Green  *color.Color
			Yellow *color.Color
			Red    *color.Color
			White  *color.Color
			Error  *color.Color
		}{
			Green:  color.New(color.FgGreen),
			Yellow: color.New(color.FgYellow),
			Red:    color.New(color.FgHiRed, color.Bold),
			Error:  color.New(color.FgHiRed),
			White:  color.New(color.FgHiWhite),
		},
		Spacing: struct{ Default string }{
			Default: strings.Join([]string{" "}, ""),
		},
		Emoji: struct {
			Error      string
			Suggestion string
		}{
			Error:      emoji.Sprint(":cross_mark:"),
			Suggestion: emoji.Sprint(":light_bulb:"),
		},
	}
}
