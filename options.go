package rofi

import (
	"fmt"
	"strings"
)

const optionsDelimiter = '\x1f'

type options struct {
	customHandler CustomHandler
	hotKeyHandler HotKeysHandler
	delimiter     string
	sb            *strings.Builder
}

func newOptions() *options {
	o := options{
		sb:        new(strings.Builder),
		delimiter: "\n",
	}

	return &o
}

func (o *options) writeOptions(key, val string) {
	o.sb.WriteString(key)
	o.sb.WriteByte(optionsDelimiter)
	o.sb.WriteString(val)
}

func (o *options) addMode(opts []ModeOption) {
	if len(opts) > 0 {
		o.sb.WriteByte('\000')
	}

	for i, opt := range opts {
		opt(o)
		if i < len(opts)-1 {
			o.sb.WriteByte(optionsDelimiter)
		}
	}

	if len(opts) > 0 {
		o.sb.WriteString(o.delimiter)
	}
}

func (o *options) addRow(name string, opts []RowOption) {
	o.sb.WriteString(name)

	if len(opts) > 0 {
		o.sb.WriteByte('\000')
	}

	for i, opt := range opts {
		opt(o)
		if i < len(opts)-1 {
			o.sb.WriteByte(optionsDelimiter)
		}
	}

	o.sb.WriteString(o.delimiter)
}

func (o *options) String() string {
	return o.sb.String()
}

type ModeOption func(opt *options)

// WithPrompt updates the prompt text
func WithPrompt(prompt string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("prompt", prompt)
	}
}

// WithMessage updates the message text
func WithMessage(message string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("message", message)
	}
}

// WithMarkupRows renders markup in the rows if true
func WithMarkupRows() ModeOption {
	return func(opt *options) {
		opt.writeOptions("markup-rows", "true")
	}
}

// WithUrgentMode marks rows as urgent
func WithUrgentMode(urgent string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("urgent", urgent)
	}
}

// WithActiveMode marks rows as active
func WithActiveMode(active string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("active", active)
	}
}

// WithDelim sets the delimiter for next rows
func WithDelim(delim string) ModeOption {
	return func(opt *options) {
		opt.delimiter = delim
		opt.writeOptions("delim", delim)
	}
}

// WithNoCustom only accepts listed entries, ignoring custom input if set to 'true'
func WithNoCustom() ModeOption {
	return func(opt *options) {
		opt.writeOptions("no-custom", "true")
	}
}

// WithCustom enables custom keybindings for script
func WithCustom(customHandler CustomHandler) ModeOption {
	return func(opt *options) {
		opt.customHandler = customHandler
	}
}

// WithHotKeys enables custom keybindings for script if set to true
func WithHotKeys(hotKeysHandler HotKeysHandler) ModeOption {
	return func(opt *options) {
		opt.hotKeyHandler = hotKeysHandler
		opt.writeOptions("use-hot-keys", "true")
	}
}

// WithKeepSelection maintains the current position without moving to the first entry
func WithKeepSelection() ModeOption {
	return func(opt *options) {
		opt.writeOptions("keep-selection", "true")
	}
}

// WithKeepFilter keeps the filter from being cleared
func WithKeepFilter() ModeOption {
	return func(opt *options) {
		opt.writeOptions("keep-filter", "true")
	}
}

// WithNewSelection overrides the selected entry if keep-selection is set
func WithNewSelection(newSelection int) ModeOption {
	return func(opt *options) {
		opt.writeOptions("new-selection", fmt.Sprintf("%v", newSelection))
	}
}

// WithData passes data to the next execution of the script via ROFI_DATA
func WithData(data string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("data", data)
	}
}

// WithTheme changes the background color of a widget
func WithTheme(theme string) ModeOption {
	return func(opt *options) {
		opt.writeOptions("theme", theme)
	}
}

type RowOption func(opt *options)

// WithIcon sets the icon for that row
func WithIcon(icon string) RowOption {
	return func(opt *options) {
		opt.writeOptions("icon", icon)
	}
}

// WithDisplay replaces the displayed string, original string is used for filtering
func WithDisplay(display string) RowOption {
	return func(opt *options) {
		opt.writeOptions("display", display)
	}
}

// WithMeta specifies invisible search terms used for filtering
func WithMeta(meta string) RowOption {
	return func(opt *options) {
		opt.writeOptions("meta", meta)
	}
}

// WithNonSelectable prevents row activation if true
func WithNonSelectable() RowOption {
	return func(opt *options) {
		opt.writeOptions("nonselectable", "true")
	}
}

// WithPermanent keeps the row always visible, independent of filter
func WithPermanent() RowOption {
	return func(opt *options) {
		opt.writeOptions("permanent", "true")
	}
}

// WithInfo places info in the ROFI_INFO environment variable on selection
func WithInfo(info string) RowOption {
	return func(opt *options) {
		opt.writeOptions("info", info)
	}
}

// WithUrgentRow sets urgent flag on entry
func WithUrgentRow(urgent bool) RowOption {
	return func(opt *options) {
		opt.writeOptions("urgent", fmt.Sprintf("%v", urgent))
	}
}

// WithActiveRow sets active flag on entry
func WithActiveRow(active bool) RowOption {
	return func(opt *options) {
		opt.writeOptions("active", fmt.Sprintf("%v", active))
	}
}
