// commands/renderer.go
package commands

import (
	"errors"

	"github.com/yigaue/linux/commands/formatter"
	"github.com/yigaue/linux/commands/pager"
)

var (
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrUnknownCommand   = errors.New("unknown command")
	ErrMissingTopic     = errors.New("missing topic")
	ErrMissingQuery     = errors.New("missing search query")
	ErrNoHelpFound      = errors.New("no help found")
	ErrNoResults        = errors.New("no results found")
)

type TLDRPage struct {
	Title string
	Description string
	Examples []CommandEntry
}

type CommandEntry struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type PageRenderer struct {
	formatter *formatter.Formatter
	pager     *pager.Pager
}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{
		formatter: formatter.NewFormatter(),
		pager:     pager.New(),
	}
}

func (r *PageRenderer) Render(page *TLDRPage) error {
	defer r.pager.Close()
	out := r.pager.Writer()

	r.formatter.PrintTitle(out, page.Title)
	
	if page.Description != "" {
		r.formatter.PrintDescription(out, page.Description)
	}

	for _, example := range page.Examples {
		r.formatter.PrintExample(out, example.Description, example.Command)
	}

	return nil
}

func PrintPage(page *TLDRPage) error {
	renderer := NewPageRenderer()
	return renderer.Render(page)
}