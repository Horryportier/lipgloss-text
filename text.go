package lipglosstext

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Text[T Displayable]  []Line[T]
type Line[T Displayable]  []Span[T]
type Span[T Displayable]  struct{t T; lipgloss.Style}

/// takes string and style returns Span
func  spanFrom[T Displayable](text T, style lipgloss.Style) Span[T] {
      return Span[T]{t: text, Style: style}
}

func (s Span[T]) Render(styled bool) string  {
    if styled {
        return s.Style.Render(fmt.Sprintf("%v", s.t))
    }
    return fmt.Sprintf("%v", s.t)
}

/// takes ([]Span, Span, string) and optional style returns Line
func lineFrom[T Displayable](t interface{}, options ...interface{}) Line[T] {
    var style lipgloss.Style = defStyle
    for _,o := range options {
        switch o.(type) {
        case lipgloss.Style:
            style = o.(lipgloss.Style)
        }
    }

    switch t.(type) {
    case []Span:
        return t.([]Span)
    case Span:
        return []Span{t.(Span)}
    case string:
        return []Span{spanFrom(t.(string), style)}
    case []string:
        var spans []Span
        for _,s := range t.([]string) {
            spans = append(spans, spanFrom(s, style))
        }
        return spans
    default: 
        return []Span{spanFrom(fmt.Sprintf("type (%T) not suported! for Line", t), errStyle)}
    }
}

func (l Line) Render(styled ...bool) string{
    var full []string 
    enable_style := true
    for _,s := range styled {
        if !s {
            enable_style = false
        }
    }
    for _,line := range l {
        full = append(full, line.Render(enable_style))
    }
    return strings.Join(full, " ")
}

/// takes ([]Span, Span, string, Line, []Line) and returns Text
func textFrom(t interface{}, options ...interface{}) Text {
    var style lipgloss.Style = defStyle
    for _,o := range options {
        switch o.(type) {
        case lipgloss.Style:
            style = o.(lipgloss.Style)
        }
    }
    switch t.(type) {
    case []Span:
        return []Line{lineFrom(t)}
    case Span:
        return []Line{lineFrom(t)}
    case string:
        return []Line{lineFrom(t, style)}
    case []string:
        return []Line{lineFrom(t, style)}
    case Line:
        return []Line{t.(Line)}
    case []Line:
        return t.([]Line)
    default: 
        return []Line{lineFrom(fmt.Sprintf("type (%T) not suported! for Text", t), errStyle)}
    }
}

func (t Text) Render(styled ...bool) string{
    enable_style := true
    for _,s := range styled {
        if !s {
            enable_style = false
        }
    }
    var full []string 
    for _,text := range t {
        full = append(full, text.Render(enable_style))
    }
    return strings.Join(full, "\n")
}


