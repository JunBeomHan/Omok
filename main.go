package main

import (
	"Omok/omok"
	"fmt"
)

func main() {
	board := omok.NewBoard()
	board.InitBoard()
	board.PrintBoard()

	playerOne := omok.NewStone(omok.ONE_PLAYER)
	playerTwo := omok.NewStone(omok.TWO_PLAYER)
	var switcher = omok.ONE_PLAYER

	for {
		switch switcher {
		case omok.ONE_PLAYER:
			fmt.Println("PlayerOne")
			fmt.Scan(&playerOne.Colum, &playerOne.Row)
			if !playerOne.PlaceMove(&board) {
				fmt.Println("제대로 하세요!")
				continue
			}
			switcher = omok.TWO_PLAYER
			if board.CheckWinner(&playerOne) {
				fmt.Println("PlyaerOne! WIN")
				board.PrintBoard()
				return
			}

		case omok.TWO_PLAYER:
			fmt.Println("PlayerTwo")
			fmt.Scan(&playerTwo.Colum, &playerTwo.Row)
			if !playerTwo.PlaceMove(&board) {
				fmt.Println("제대로 하세요!")
				continue
			}
			switcher = omok.ONE_PLAYER
			if board.CheckWinner(&playerTwo) {
				fmt.Println("PlayerTwo! WIN")
				board.PrintBoard()
				return
			}
		}
		board.PrintBoard()
	}
}
