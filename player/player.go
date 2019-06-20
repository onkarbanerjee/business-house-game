package player

import "errors"

var errInsufficientBalance = errors.New("Insufficient balance")

type Player struct {
	Id,
	currentPosition,
	currentBalance int
	assetsOwned []int
}

func New(id, currentBalance int) *Player {
	return &Player{Id: id, currentPosition: -1, currentBalance: currentBalance}
}

func (p *Player) CurrentBalance() int {
	return p.currentBalance
}

func (p *Player) SetBalance(newBal int) {
	p.currentBalance = newBal
}

func (p *Player) CurrentPosition() int {
	return p.currentPosition
}
func (p *Player) AcquireAsset(assetID int) {
	p.assetsOwned = append(p.assetsOwned, assetID)
	return
}

func (p *Player) UpdatePosition(newPos int) {
	p.currentPosition = newPos
}
