package lipglosstext

type Option interface{
    Opt() interface{};
}

type Styled bool

func (t Styled) Opt() interface{} {
    return t
}

type Delimiter string

func (t Delimiter) Opt() interface{} {
    return t
}
