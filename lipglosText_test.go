package lipglosstext_test

import (
	"fmt"
	"testing"

	. "github.com/Horryportier/lipgloss-text"
	"github.com/charmbracelet/lipgloss"
)

var (
   s1 = lipgloss.NewStyle().Foreground(lipgloss.Color("#ACABAC"))
   s2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAF200"))
   s3 = lipgloss.NewStyle().Foreground(lipgloss.Color("#AC5BAC"))
   s4 = lipgloss.NewStyle().Foreground(lipgloss.Color("#3CABFC"))

)

type TestInvalidType struct{
    x string
    y float32
}

func TestSpan(t *testing.T)  {
    if true {
        span := SpanFrom("test span", s3)
        t.Errorf("span: %s", span.Render(true))
    }
}

func TestLine(t *testing.T) {
    a := []any{
        "test string",
        []string{"test", "slice", "of", "strings"},
        []Span{SpanFrom("test", s2), SpanFrom("slice of Span", s3)},
        SpanFrom("test from span", s4),
        2,
        []int{2, 3},
        2.,
        []float32{2., 3.},
        2.,
        []float64{2., 3.},
        true,
        []bool{true, false},
    }

    if true {
        for _, v := range a {
            l := LineFrom(v, s1)
            fmt.Println(l.Render(Styled(true), Delimiter("::")))           
        }
        t.Error("end of line test")
    }
}
func TestText(t *testing.T) {
    a := []any{
        "test string",
        []string{"test", "slice", "of", "strings"},
        []Span{SpanFrom("test", s2), SpanFrom("slice of Span", s3)},
        SpanFrom("test from span", s4),
        LineFrom("test from line"),
        []Line{LineFrom("test", s1), LineFrom("from", s2), LineFrom("slice", s3), LineFrom("of Line", s4)},
        2,
        []int{2, 3},
        2.,
        []float32{2., 3.},
        2.,
        []float64{2., 3.},
        true,
        []bool{true, false},
        TestInvalidType{ x: "invalid type ", y: 3.14 },
    }

    if true {
        for _, v := range a {
            l := TextFrom(v, s1)
            fmt.Println(l.Render(Delimiter(";"), Styled(true)))           
        }
        t.Error("end of text test")
    }
}
