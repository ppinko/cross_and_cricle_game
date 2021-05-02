package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// game variables
type dataMembers struct {
	board   [3][3]int
	missing int
}

func main() {

	playGame()
	data := dataMembers{
		missing: 9,
	}
	data.printBoard()
	data.userInput()
}

func (d *dataMembers) printBoard() {
	fmt.Println()
	for i, row := range d.board {
		printRow(row)
		if i != 2 {
			fmt.Println("---------")
		}
	}
	fmt.Println()
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

func (d *dataMembers) userInput() (int, int) {
	// create an instance of buffered I/O
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("To quit the game press Q and press Enter.")
		fmt.Println("Please choose a row to mark: ")
		fmt.Print("-> ")
		row, _ := reader.ReadString('\n')
		row = strings.Replace(row, "\n", "", -1)
		row = strings.ToUpper(row)
		if len(row) == 1 && row == "Q" {
			fmt.Println("Thank you very much for the game! Have a great day!")
			os.Exit(0)
		}
		if len(row) != 1 || row < "0" || row > "2" {
			fmt.Println("Invalid input!")
			continue
		}

		fmt.Println("Please choose a column to mark: ")
		fmt.Print("-> ")
		col, _ := reader.ReadString('\n')

		// convert CRLF to LF
		col = strings.Replace(col, "\n", "", -1)
		col = strings.ToUpper(col)

		if len(col) != 1 || col < "0" || col > "2" {
			fmt.Println("Invalid input!")
			continue
		}

		rowNumber, err1 := strconv.Atoi(row)
		colNumber, err2 := strconv.Atoi(col)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input!")
			continue
		}
		return rowNumber, colNumber
	}
}

// playGame starts new minesweeper game.
func playGame() {
	fmt.Println("-----------------------------------------")
	fmt.Println("WELCOME IN CROSS AND CIRCLE GAME")
	fmt.Println("Please press any key to start a game ('Q' to quit):")
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToUpper(input)
	// validation of the input
	if input == "Q" {
		fmt.Println("Thank you very much for the game! Have a great day!")
		os.Exit(0)
	}
}
