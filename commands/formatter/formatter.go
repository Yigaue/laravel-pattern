// commands/formatter/formatter.go
package formatter

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Formatter struct {
	title        *color.Color
	description  *color.Color
	exampleDesc  *color.Color
	command      *color.Color
	placeholder  *color.Color
	flag         *color.Color
	path         *color.Color
	url          *color.Color
	keyword      *color.Color
	syntaxRules  []syntaxRule
	keywordRegex *regexp.Regexp
	arrowColor	 *color.Color
}

type syntaxRule struct {
	pattern *regexp.Regexp
	color   *color.Color
}

func NewFormatter() *Formatter {
	f := &Formatter{
		title:        color.New(color.Bold, color.FgMagenta),
		description:  color.New(color.FgWhite),
		exampleDesc:  color.New(color.FgGreen),
		command:      color.New(color.FgBlue),
		placeholder:  color.New(color.FgYellow),
		flag:         color.New(color.FgHiMagenta, color.Bold),
		path:         color.New(color.FgRed, color.Bold),
		url:          color.New(color.FgRed),
		keyword:      color.New(color.FgHiYellow, color.Bold),
		arrowColor:		color.New(color.FgWhite, color.Bold),
	}

	f.syntaxRules = []syntaxRule{
		{regexp.MustCompile(`\{\{([^}]+)\}\}`), f.placeholder},
		{regexp.MustCompile(`(\s|^)(-\w+)\b`), f.flag},
		{regexp.MustCompile(`(\.?/)?[\w\-/]+(/[\w\-]+)+`), f.path},
		{regexp.MustCompile(`- >|>>|>`), f.arrowColor},
	}

	f.keywordRegex = regexp.MustCompile(`\b(stdin|stdout|file|path|output|input)\b`)
	return f
}

func (f *Formatter) PrintTitle(w io.Writer, text string) {
	fmt.Fprintf(w, "\n%s\n\n", f.title.Sprint(text))
}

func (f *Formatter) PrintDescription(w io.Writer, text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if strings.Contains(line, "http") {
			f.printWithURLs(w, line)
		} else {
			fmt.Fprintf(w, "  %s\n", f.description.Sprint(line))
		}
	}
	fmt.Fprintln(w)
}

func (f *Formatter) printWithURLs(w io.Writer, text string) {
	urlRegex := regexp.MustCompile(`(https?://\S+)`)
	parts := urlRegex.Split(text, -1)
	urls := urlRegex.FindAllString(text, -1)

	for i, part := range parts {
		if part != "" {
			fmt.Fprintf(w, "  %s", f.description.Sprint(part))
		}
		if i < len(urls) {
			fmt.Fprint(w, f.url.Sprint(urls[i]))
		}
	}
	fmt.Fprintln(w)
}

func (f *Formatter) PrintExample(w io.Writer, desc, cmd string) {
	fmt.Fprintf(w, "  %s\n\n", f.exampleDesc.Sprint(f.highlightKeywords(desc)))
	fmt.Fprintf(w, "    %s\n\n", f.formatCommand(cmd))
}

func (f *Formatter) highlightKeywords(text string) string {
	return f.keywordRegex.ReplaceAllStringFunc(text, func(m string) string {
		return f.keyword.Sprint(m)
	})
}

func (f *Formatter) formatCommand(cmd string) string {
	cleanCmd := strings.NewReplacer("{{", "", "}}", "").Replace(cmd)
	return f.applySyntaxHighlighting(cleanCmd)
}

func (f *Formatter) applySyntaxHighlighting(cmd string) string {
	var highlighted strings.Builder
	remaining := cmd

	for remaining != "" {
		var match struct {
			start, end int
			rule       *syntaxRule
		}
		match.end = len(remaining)

		for _, rule := range f.syntaxRules {
			loc := rule.pattern.FindStringIndex(remaining)
			if loc != nil && loc[0] < match.end {
				match.start = loc[0]
				match.end = loc[1]
				match.rule = &rule
			}
		}

		if match.rule == nil {
			highlighted.WriteString(f.command.Sprint(remaining))
			break
		}

		highlighted.WriteString(f.command.Sprint(remaining[:match.start]))
		highlighted.WriteString(match.rule.color.Sprint(remaining[match.start:match.end]))
		remaining = remaining[match.end:]
	}

	return highlighted.String()
}