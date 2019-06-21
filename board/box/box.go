package box

import "fmt"

type typ int

const (
	Empty = iota
	TreasureTyp
	JailTyp
	HotelTyp
)

var Typs = map[string]typ{
	"E": Empty,
	"T": TreasureTyp,
	"J": JailTyp,
	"H": HotelTyp,
}

// Box denotes a single box on the board. It can have a typ, name and a deductible and ownable as well.
// In our current board game though, box can either be either of two ofcourse
type Box struct {
	typ
	Ownable
	Deductible
}

func New(boxTyp typ) *Box {
	return &Box{
		typ: boxTyp,
	}
}

func (b *Box) SetDeductible(deductible Deductible) {
	b.Deductible = deductible
}

func (b *Box) SetOwnable(ownable Ownable) {
	b.Ownable = ownable
}

func (b *Box) String() string {
	return fmt.Sprintf("Box is %d, is deductible? %t, is ownable? %t", b.typ, b.Deductible != nil, b.Ownable != nil)
}
