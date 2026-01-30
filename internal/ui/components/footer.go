package components

import (
	"fmt"
	"strings"

	"github.com/experteur/goop/internal/ui"
	"github.com/rivo/tview"
)

type Footer struct {
	*tview.TextView
}

type Shortcut struct {
	Key   string
	Label string
}

func NewFooter(shortcuts []*Shortcut) *Footer {
	textView := tview.NewTextView()
	textView.SetBorder(true)
	textView.SetTitleColor(ui.Theme.TitleColor)
	textView.SetBorderColor(ui.Theme.BorderColor)
	f := &Footer{
		textView,
	}
    f.SetText(f.buildFooter(shortcuts))
	return f
}

func (f *Footer) buildFooter(shortcuts []*Shortcut) string {
	// `ProjectName   [ACTIVE]   14/22 (64%)   ⚠   @alex`
    var builder strings.Builder
    for _, shortcut := range shortcuts {
        pair := fmt.Sprintf("%s:%s\t", shortcut.Key, shortcut.Label)
        builder.WriteString(pair)
    }
    return builder.String()
}
