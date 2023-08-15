package lipglosstext

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Text  []Line
type Line  []Span
type Span  struct{ string; lipgloss.Style}

var (
    errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
)

/// takes string and style returns Span
func  SpanFrom(text string, style lipgloss.Style) Span {

      return Span{string: text, Style: style}
}

func (s Span) Render(styled bool) string  {
    if styled {
        return s.Style.Render(s.string)
    }
    return s.string
}

/// takes ([]Span, Span, string) and optional style returns Line
func LineFrom(t interface{}, style ...lipgloss.Style) Line {
    var final_style lipgloss.Style = lipgloss.NewStyle()
    for _,o := range style {
            final_style = o
    }

    switch t.(type) {
    case []Span:
        return t.([]Span)
    case Span:
        return []Span{t.(Span)}
    case string:
        return []Span{SpanFrom(t.(string), final_style)}
    case []string:
        var spans []Span
        for _,s := range t.([]string) {
            spans = append(spans, SpanFrom(s, final_style))
        }
        return spans
    case int:
        return []Span{SpanFrom(fmt.Sprint(t.(int)), final_style)}

    case []int:
        var spans []Span
        for _,s := range t.([]int) {
            spans = append(spans, SpanFrom(fmt.Sprint(s), final_style))
        }
        return spans

    case float32:
        return []Span{SpanFrom(fmt.Sprintf("%.2f",t.(float32)), final_style)}
    case []float32:
        var spans []Span
        for _,s := range t.([]float32) {
            spans = append(spans, SpanFrom(fmt.Sprintf("%.2f",s), final_style))
        }
        return spans

    case float64:
        return []Span{SpanFrom(fmt.Sprintf("%.2f",t.(float64)), final_style)}
    case []float64:
        var spans []Span
        for _,s := range t.([]float64) {
            spans = append(spans, SpanFrom(fmt.Sprintf("%.2f", s), final_style))
        }
        return spans

    case bool:
        return []Span{SpanFrom(fmt.Sprint(t.(bool)), final_style)}
    case []bool:
        var spans []Span
        for _,s := range t.([]bool) {
            spans = append(spans, SpanFrom(fmt.Sprint(s), final_style))
        }
        return spans
    default: 
        return []Span{SpanFrom(fmt.Sprintf("type (%T) not suported!", t), errStyle)}
    }
}

func (l *Line) Append(rhs ...interface{}) error {
    for _, v := range rhs {
        switch v.(type) {
        case Line:
            *l = append(*l, v.(Line)...)
        case []Line:
            for _, x := range v.([]Line) {
                *l = append(*l, x...)
            }
        default: 
            fmt.Errorf("wrong type [%T] for appendig to Line", v)
        }
    }
    return nil
}


func (l Line) Render(options ...Option) string{
    enable_style := true
    delimiter := " "
    for _,o := range options {
        switch c := o.Opt(); c.(type) {
        case Styled:
            enable_style = bool(c.(Styled))
        case Delimiter:
            delimiter = string(c.(Delimiter))
        }
    }

    var full []string 
    for _,line := range l {
        full = append(full, line.Render(enable_style))
    }
    return strings.Join(full, delimiter)
}

/// takes ([]Span, Span, string, Line, []Line) and optional style  returns Text
func TextFrom(t interface{}, style ...lipgloss.Style) Text {
    var final_style lipgloss.Style = lipgloss.NewStyle()
    for _,o := range style {
            final_style = o
    }
    switch t.(type) {
    case []Span:
        return []Line{LineFrom(t)}
    case Span:
        return []Line{LineFrom(t)}
    case string:
        return []Line{LineFrom(t, final_style)}
    case []string:
        return []Line{LineFrom(t, final_style)}
    case Line:
        return []Line{t.(Line)}
    case []Line:
        return t.([]Line)
    default: 
        return []Line{LineFrom(t, final_style)}
    }
}

func (t *Text) Append(rhs ...interface{}) error {
    for _, v := range rhs {
        switch v.(type) {
        case Text:
            *t = append(*t, v.(Text)...)
        case []Line:
            for _, x := range v.([]Text) {
                *t = append(*t, x...)
            }
        default: 
            fmt.Errorf("wrong type [%T] for appendig to Text", v)
        }
    }
    return nil
}

func (t Text) Render(options ...Option) string{
    enable_style := true
    delimiter := "\n"
    for _,s := range options {
        switch c := s.Opt(); c.(type) {
            case Styled: 
                enable_style = bool(c.(Styled))           
            case Delimiter: 
                delimiter = string(c.(Delimiter))
        }
        
    }
    var full []string 
    for _,line := range t {
        full = append(full, line.Render(Styled(enable_style)))
    }
    return strings.Join(full, delimiter)
}


