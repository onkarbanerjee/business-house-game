package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/onkarbanerjee/business-house-game/board"
	"github.com/onkarbanerjee/business-house-game/player"
)

func move(board *board.Board, p *player.Player, diceValue int) {
	newPos := p.CurrentPosition() + diceValue
	if newPos >= len(board.Boxes) {
		newPos = newPos % len(board.Boxes)
	}

	box := board.Boxes[newPos]
	if box.Deductible != nil {
		newBal := box.Deduct(p.CurrentBalance())
		p.SetBalance(newBal)
	} else if box.Ownable != nil {
		if !box.IsOwned() && p.CurrentBalance() > box.Price() {
			p.AcquireAsset(box.ID(), box.Price())
			box.OwnBy(p.Id)
		}
	}
	p.UpdatePosition(newPos)
}

func main() {
	layout := []string{"E", "E", "J", "E", "H", "T", "H", "E", "E", "T", "H", "E"}

	board, err := board.New("Business House Game", layout)
	if err != nil {
		log.Println("Could not get a Business House Game board", err)
		return
	}

	numberOfPlayers, numberOfRounds := 2, 10
	numberOfThrows := numberOfPlayers * numberOfRounds
	players := []*player.Player{}
	for i := 0; i < numberOfPlayers; i++ {
		players = append(players, player.New(i+1, 2000))
	}

	diceValues := []int{4, 4, 4, 6, 3, 9, 2, 12, 11, 10, 5, 6, 3, 8, 2, 2, 6, 10, 9, 7}
	if len(diceValues) < numberOfThrows {
		log.Println("Needed a minimum of", numberOfThrows, "diceValues but got", len(diceValues))
		return
	}
	diceValues = diceValues[:numberOfThrows]
	for idx, diceValue := range diceValues {
		turn := idx % numberOfPlayers
		move(board, players[turn], diceValue)
	}

	byBalance := player.ByBalance(players)
	sort.Stable(sort.Reverse(byBalance))
	for _, each := range byBalance {
		fmt.Println(each)
	}
}
