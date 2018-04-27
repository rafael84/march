package march_test

import (
	"strings"
	"testing"

	"github.com/rafael84/march"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok march.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: march.EOF},
		{s: `#`, tok: march.ILLEGAL, lit: `#`},
		{s: ` `, tok: march.WS, lit: " "},
		{s: "\t", tok: march.WS, lit: "\t"},
		{s: "\n", tok: march.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: march.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: march.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: march.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: march.FROM, lit: "FROM"},
		{s: `SELECT`, tok: march.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := march.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
