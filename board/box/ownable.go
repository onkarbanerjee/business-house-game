package box

import "fmt"

type ownable interface {
	ownBy(playerId int) error
}

type hotel struct {
	id   int
	name string
	cost,
	owner int
}

func NewHotel(name string, id, cost int) *hotel {
	return &hotel{name: name, id: id, cost: cost}
}

func (h *hotel) ownBy(playerId int) error {
	if h.owner == 0 {
		return fmt.Errorf("Hotel %s is owned by player id %d", h.name, h.owner)
	}
	h.owner = playerId
	return nil
}
