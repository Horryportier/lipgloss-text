package lipglosstext

import "golang.org/x/exp/constraints"


type Displayable interface {
    constraints.Float | constraints.Integer | string | bool | constraints.Complex
}

type Number interface {
	constraints.Integer | constraints.Float
}
