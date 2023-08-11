package lipglosstext





func SpanFrom[T Displayable](items ...interface{}) []Span[T] {  
    var accum []Span[T]

    opts, other := getOpts()
    for _,t := range other {
        switch t.(type) {
        case string:
            accum = append(accum, spanFrom(t.(T), getStyle(opts...)))
        case []string: 
            for _,v := range t.([]string) {
                accum = append(accum, spanFrom(v, getStyle(opts...)))
            }

        }
    }
    return accum 
}

