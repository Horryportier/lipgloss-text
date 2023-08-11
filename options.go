package lipglosstext

import (
	"github.com/charmbracelet/lipgloss"
)


type Option interface{
    Opt() interface{};
}


type Delimiter string

func (d Delimiter) Opt() interface{} {
    return string(d)
}

type Style lipgloss.Style

func (s Style) Opt() interface{} {
    return lipgloss.Style(s)
}

type Raw bool

func (r Raw) Opt() interface{} {
    return bool(r) 
}


