package main

import "fmt"

// game variables
type dataMembers struct {
	board [3][3]int
}

func main() {
	playGame()
	data := dataMembers{}
	data.printBoard()
}

func (b *dataMembers) printBoard() {
	fmt.Println()
	for i, row := range b.board {
		printRow(row)
		if i != 2 {
			fmt.Println("---------")
		}
	}
}

func printRow(row [3]int) {
	line := ""
	for i, val := range row {
		switch val {
		case 0:
			line += " "
		case 1:
			line += "X"
		case -1:
			line += "O"
		}
		if i != 2 {
			line += " | "
		}
	}
	fmt.Println(line)
}

// playGame starts new minesweeper game.
func playGame() {
	fmt.Println("-----------------------------------------")
	fmt.Println("WELCOME IN CROSS AND CIRCLE GAME")
}
