package board

import (
	"fmt"

	"github.com/onkarbanerjee/business-house-game/board/box"
	"github.com/pkg/errors"
)

// Board is the main entity on which all the boxes are laid out and players make their moves
type Board struct {
	name  string
	Boxes []*box.Box
}

// New takes in a layout and parses it to get appropriate boxes in the Board
func New(name string, layout []string) (*Board, error) {
	boxes, err := parse(layout)
	if err != nil {
		return nil, errors.Wrap(err, "Error in parsing the layout")
	}
	return &Board{
		name: name, Boxes: boxes,
	}, nil
}

// TODO: move validation out, refactor all the hard coded values
func parse(layout []string) ([]*box.Box, error) {
	ret := []*box.Box{}
	var hotelId int
	for _, each := range layout {
		typ, ok := box.Typs[each]
		if !ok {
			return nil, fmt.Errorf("Incorrect box type %d", typ)
		}
		eachBox := box.New(typ)
		switch typ {
		case box.JailTyp:
			deductible := box.NewJail("Central Jail", 150)
			eachBox.SetDeductible(deductible)
		case box.TreasureTyp:
			deductible := box.NewTreasure("Mosby's Treasure", 200)
			eachBox.SetDeductible(deductible)
		case box.HotelTyp:
			hotelId++
			ownable := box.NewHotel("Hotel Airport Link", hotelId, 150)
			eachBox.SetOwnable(ownable)
		}

		ret = append(ret, eachBox)
	}

	return ret, nil
}
