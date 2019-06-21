package player

import (
	"errors"
	"fmt"
)

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
func (p *Player) AcquireAsset(assetID, cost int) {
	p.assetsOwned = append(p.assetsOwned, assetID)
	p.currentBalance -= cost
	return
}

func (p *Player) UpdatePosition(newPos int) {
	p.currentPosition = newPos
}

func (p *Player) String() string {
	return fmt.Sprintf("Player id : %d is at %d with Balance %d and owns assets id %v", p.Id, p.currentPosition, p.currentBalance, p.assetsOwned)
}
