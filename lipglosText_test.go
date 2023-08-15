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

const (
    MAKE_IT_FAIL = false
)

type TestInvalidType struct{
    x string
    y float32
}

func TestSpan(t *testing.T)  {
    if MAKE_IT_FAIL {
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

    if MAKE_IT_FAIL {
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

    if MAKE_IT_FAIL {
        for _, v := range a {
            l := TextFrom(v, s1)
            fmt.Println(l.Render(Delimiter(";"), Styled(true)))           
        }
        t.Error("end of text test")
    }
}

func TestAppendText(t *testing.T)  { 
    var text Text 
    if len(text) != 0 {
        t.Error("how")
    }
    text2 := TextFrom("test", s2)
    text.Append(text2)
    if len(text) != 1 {
        t.Errorf("len of text is not 1 its [%v]", len(text))
    }
}


func TestAppendLine(t *testing.T)  { 
    var text Line 
    if len(text) != 0 {
        t.Error("how")
    }
    text2 := LineFrom("test", s2)
    text.Append(text2)
    if len(text) != 1 {
        t.Errorf("len of line is not 1 its [%v]", len(text))
    }
}
