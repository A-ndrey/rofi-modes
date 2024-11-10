package rofi

import "testing"

func TestOptions(t *testing.T) {
	modeOpts := []ModeOption{
		WithPrompt("Test prompt"),
		WithMessage("Test message"),
		WithUrgentMode("7:11"),
		WithMarkupRows(),
		WithActiveMode("1:2"),
		WithDelim(";"),
		WithCustom(func(args []string, env Env) error {
			return nil
		}),
		WithHotKeys(func(key int, env Env) error {
			return nil
		}),
		WithKeepSelection(),
		WithKeepFilter(),
		WithNewSelection(2),
		WithData("test data"),
		WithTheme("test theme"),
	}

	rowOpts := []RowOption{
		WithIcon("test icon"),
		WithDisplay("test display"),
		WithActiveRow(true),
		WithPermanent(),
		WithNonSelectable(),
		WithMeta("test meta"),
		WithInfo("test info"),
		WithUrgentRow(true),
	}

	opts := newOptions()

	opts.addMode(modeOpts)
	opts.addRow("test row", rowOpts)

	expected := "\000prompt\x1fTest prompt\x1fmessage\x1fTest message\x1furgent\x1f7:11\x1fmarkup-rows\x1ftrue\x1factive\x1f1:2\x1fdelim\x1f;\x1fno-custom\x1ffalse\x1fuse-hot-keys\x1ftrue\x1fkeep-selection\x1ftrue\x1fkeep-filter\x1ftrue\x1fnew-selection\x1f2\x1fdata\x1ftest data\x1ftheme\x1ftest theme;test row\000icon\x1ftest icon\x1fdisplay\x1ftest display\x1factive\x1ftrue\x1fpermanent\x1ftrue\x1fnonselectable\x1ftrue\x1fmeta\x1ftest meta\x1finfo\x1ftest info\x1furgent\x1ftrue;"
	actual := opts.String()
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
