package blog

import (
	"regexp"
	"strings"
)

type replacer struct {
	pattern   *regexp.Regexp
	replaceBy []byte
}

func NewTextLinter(s []string) []replacer {
	if len(s)%2 != 0 {
		panic("Len of replacers must be even")
	}
	rs := []replacer{}
	for i := 0; i < len(s)-1; i += 2 {
		rs = append(rs, replacer{
			regexp.MustCompile(s[i]),
			[]byte(s[i+1]),
		})
	}
	return rs
}

var textLinter = NewTextLinter([]string{
	`[\s\p{Zs}]*\.{3}`, "\u2026", // Elipis
	`\.+`, `.`,
	`\s*'\s*`, `'`,
	`\b\s*([.,])`, `$1`,
	`([.,])\b`, `$1 `,
	`[\s\p{Zs}]{2,}`, ` `, // 2+ consecutive spaces
	`\b[\s\p{Zs}]*([;:?!])[\s\p{Zs}]*\b`, "\u00A0$1 ",
	`\b[\s\p{Zs}]*([;:?!])[\s\p{Zs}]*$`, "\u00A0$1", // Insert non breakable space after ponctuation
})

var titleLinter = NewTextLinter([]string{
	`^[\p{Zs}]*|[\s\p{Zs}]*$`, ``,
})

func applyReplacers(s string, rs []replacer) string {
	b := []byte(s)
	for _, r := range rs {
		b = r.pattern.ReplaceAll(b, r.replaceBy)
	}
	return string(b)
}

func TextLinter(t string) string {
	return applyReplacers(t, textLinter)
}

func AllCapLinter(t string) string {
	t = applyReplacers(t, textLinter)
	t = applyReplacers(t, titleLinter)
	return strings.ToUpper(t)
}
