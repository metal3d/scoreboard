package components

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*EditableLabel)(nil)
var _ fyne.Tappable = (*EditableLabel)(nil)

// EditableLabel is a label that can be tapped to edit.
type EditableLabel struct {
	widget.BaseWidget
	label       *widget.Label
	entry       *numercialEntry
	edited      bool
	OnSubmitted func(score float64, editing bool)
}

// NewEditableLabel creates a new editable label.
func NewEditableLabel() *EditableLabel {
	editable := &EditableLabel{
		label: widget.NewLabel(""),
		entry: newNumericalEntry(),
	}
	editable.label.Hide()

	editable.entry.ActionItem = canvas.NewImageFromResource(theme.ErrorIcon())
	editable.entry.ActionItem.Hide()

	editable.entry.OnChanged = func(text string) {
		editable.entry.ActionItem.Hide()
		if text == "" {
			return
		}
		_, err := editable.getParsed()
		if err != nil {
			editable.entry.ActionItem.Show()
			return
		}
	}

	editable.entry.OnSubmitted = func(text string) {
		if text == "" {
			return
		}
		editable.entry.SetText(text)
		p, err := editable.getParsed()
		if err != nil {
			log.Println(err)
			return
		}
		editable.label.SetText(fmt.Sprintf("%v", p))
		if editable.OnSubmitted != nil {
			editable.OnSubmitted(p, editable.edited)
		}
		editable.label.Show()
		editable.entry.Hide()
	}
	editable.ExtendBaseWidget(editable)
	return editable
}

// CreateRenderer implements fyne.Widget interface.
func (e *EditableLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewMax(e.label, e.entry),
	)
}

// Tapped implements fyne.Tappable interface.
func (e *EditableLabel) Tapped(evt *fyne.PointEvent) {
	e.edited = true
	e.label.Hide()
	e.entry.Show()
	e.SetFocus()
}

func (e *EditableLabel) IsEditing() bool {
	return e.entry.Visible()
}

// SetOnSubmitted sets the function to be called when the user submits the entry.
func (e *EditableLabel) SetOnSubmitted(f func(float64, bool)) *EditableLabel {
	e.OnSubmitted = f
	return e
}

func (e *EditableLabel) getParsed() (float64, error) {
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, 0, e.entry.Text)
	if err != nil {
		return 0, err
	}
	text := tv.Value.String()
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, nil
	}
	p, _ := strconv.ParseFloat(text, 64)
	return p, nil
}

// Text returns the text of the label.
func (e *EditableLabel) Text() string {
	return e.label.Text
}

// SetFocus sets the focus on the entry.
func (e *EditableLabel) SetFocus() {
	// no-op
	//fyne.CurrentApp().Driver().CanvasForObject(e.entry).Focus(e.entry)
}

var (
	allowedRunes = map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'.': true,
		'/': true,
		'*': true,
		'+': true,
		'-': true,
		'(': true,
		')': true,
	}
)

var _ mobile.Keyboardable = (*numercialEntry)(nil)

// numercialEntry is an entry that only allows numerical input.
type numercialEntry struct {
	widget.Entry
}

// newNumericalEntry creates a new numerical entry.
func newNumericalEntry() *numercialEntry {
	n := &numercialEntry{}
	n.ExtendBaseWidget(n)
	return n
}

// Keyboard implements mobile.Keyboardable interface. We need to return NumberKeyboard to
// force the numerical keyboard to be shown (on mobile).
func (e *numercialEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

// TypedRune implements fyne.Focusable interface. We need to filter the input to only allow. Others
// methods from the interface are not needed, there are defined in the widget.Entry.
func (e *numercialEntry) TypedRune(r rune) {
	if !allowedRunes[r] {
		return
	}
	e.Entry.TypedRune(r)
}
