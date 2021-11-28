package main

import "fmt"

func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Println("Player", winningPlayer(), "Wins")
}

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Println("Turn", turn, "taken by player", currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
