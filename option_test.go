package lipglosstext_test

import (
	"testing"

	. "github.com/Horryportier/lipgloss-text"
	"github.com/charmbracelet/lipgloss"
)


func TestOtpion(t *testing.T) {
    var opts []Option  = []Option{
        Style(lipgloss.NewStyle()),
        Delimiter(";"),
        Raw(false),
    }
    var items []interface{} = []interface{}{
        "test",
        "hi",
        123,
        3.142,
        true,
    }
    
    items = append(items, opts)
    
    o, _ := GetOpts(items...)

    for i,j := range o {
        if j != opts[i] {
            t.Errorf("Option %v is not the same as %v at %v", j, opts[i], i)
    }
    }
    
}
