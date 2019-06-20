package box

type Ownable interface {
	ID() int
	Price() int
	IsOwned() bool
	OwnBy(playerId int)
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

func (h *hotel) ID() int {
	return h.id
}
func (h *hotel) Price() int {
	return h.cost
}

func (h *hotel) IsOwned() bool {
	return h.owner == 0
}

func (h *hotel) OwnBy(playerId int) {
	h.owner = playerId
}
