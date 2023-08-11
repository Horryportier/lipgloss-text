package lipglosstext

import "github.com/charmbracelet/lipgloss"


func getOpts(items ... interface{}) ([]Option, []interface{}) {
    var  opt []Option
    var  other []interface{}
    for _,t := range items {
        if v, ok := interface{}(t).(Option); ok {
            opt = append(opt, v)
        } else {
            other = append(other, v)
        }
    }
    return opt, other
}



func getStyle(opts ...Option) lipgloss.Style {
    for _,o := range opts {
        switch o.(type) { 
        case Style:
            return o.(Style).Opt().(lipgloss.Style)
        default:
            return defStyle
        }
    }
    return defStyle
}
