package gnureadline
import "testing"

func assert_true(b bool, msg string, t *testing.T) {
	if !b  {
		t.Errorf("failed %s\n", msg)
	}
}
	
func assert_false(b bool, msg string, t *testing.T) {
	if b  {
		t.Errorf("failed %s\n", msg)
	}
}

func assert_equal_int(expect, got int, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %d, got %d: %s\n", 
			expect, got, msg)
	}
}

func assert_equal_rune(expect, got rune, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %c, got %c: %s\n", 
			expect, got, msg)
	}
}

func assert_equal_string(expect, got string, msg string, t *testing.T) {
	if expect != got  {
		t.Errorf("failed expected %s, got %s: %s\n", 
			expect, got, msg)
	}
}

func TestReadline(t *testing.T) {
	assert_true(Rl_gnu_readline_p(), "Genuine GNU Readline", t)

	/* See that we can read all of these attributes */
	varlist := []string{
		"bell_style", "comment-begin", "completion-prefix-display-length",
		"completion-query-items", "editing-mode", "history-size",
		"isearch-terminators", "keymap",
		"bind-tty-special-chars",
		"blink-matching-paren",
		"byte-oriented",
		"completion-ignore-case",
		"convert-meta",
		"disable-completion",
		"enable-keypad",
		"expand-tilde",
		"history-preserve-point",
		"horizontal-scroll-mode",
		"input-meta",
		"mark-directories",
		"mark-modified-lines",
		"mark-symlinked-directories",
		"match-hidden-files",
		"meta-flag",
		"output-meta",
		"page-completions",
		"prefer-visible-bell",
		"print-completions-horizontally",
		"revert-all-at-newline",
		"show-all-if-ambiguous",
		"show-all-if-unmodified",
	}
	for _, variable := range(varlist) {
		assert_true(true, Rl_variable_value(variable), t)
				
	}

	Rl_editing_mode_set(Vi)
	assert_equal_int(int(Vi), int(Rl_editing_mode()),
		"Edit mode should be set to Vi", t)
	
	Rl_editing_mode_set(Emacs)
	assert_equal_int(int(Emacs), int(Rl_editing_mode()),
		"Edit mode should be set to Emacs", t)
}
