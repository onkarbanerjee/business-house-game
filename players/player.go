package player

import "errors"

var errInsufficientBalance = errors.New("Insufficient balance")

type Player struct {
	id,
	currentPosition,
	currentBalance int
	assetsOwned []int
}

func (p *Player) CurrentBalance() int {
	return p.currentBalance
}

func (p *Player) CurrentPosition() int {
	return p.currentPosition
}
func (p *Player) AcquireAsset(assetID int) {
	p.assetsOwned = append(p.assetsOwned, assetID)
	return
}

func (p *Player) Move(paces int) {
	p.currentPosition += paces
}
