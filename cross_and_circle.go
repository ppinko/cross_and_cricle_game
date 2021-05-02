package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// game variables
type dataMembers struct {
	board   [3][3]int
	missing int
}

func main() {

	data := dataMembers{
		missing: 9,
	}
	data.playGame()
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
		d.printBoard()
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

		if !isValidField(d, rowNumber, colNumber) {
			fmt.Println("Invalid input! This field has already been set.")
			continue
		}

		return rowNumber, colNumber
	}
}

// playGame starts new minesweeper game.
func (d *dataMembers) playGame() {
	fmt.Println("-----------------------------------------")
	fmt.Println("WELCOME IN CROSS AND CIRCLE GAME")
	fmt.Println("Please press an enter to start a game:")
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// -1 - computer, 1 - player
	turn := 1
	for d.missing > 0 {
		var row, col int
		if turn == 1 {
			row, col = d.userInput()
		} else {
			fmt.Println("Computer's turn!")
			row, col = d.computerMove()
		}

		(*d).board[row][col] = turn
		d.printBoard()
		(*d).missing -= 1

		if d.isWin(turn) {
			d.printBoard()
			if turn == 1 {
				fmt.Println("Congratulations, you won!")
			} else {
				fmt.Println("You lost!")
			}
			fmt.Println("Please press an enter to leave the game:")
			fmt.Print("-> ")
			reader.ReadString('\n')
		}

		// change turn
		turn *= -1
	}

	fmt.Println("A draw!")
	fmt.Println("Please press an enter to leave the game:")
	fmt.Print("-> ")
	reader.ReadString('\n')
}

func isValidField(d *dataMembers, row int, col int) bool {
	if (*d).board[row][col] == 0 {
		return true
	} else {
		return false
	}
}

func (d *dataMembers) isWin(turn int) bool {
	for i := 0; i < len(d.board); i++ {
		if d.board[i][0] == turn && d.board[i][1] == turn && d.board[i][2] == turn {
			return true
		}
		if d.board[0][i] == turn && d.board[1][i] == turn && d.board[2][i] == turn {
			return true
		}
	}

	if d.board[0][0] == turn && d.board[1][1] == turn && d.board[2][2] == turn {
		return true
	}
	if d.board[0][2] == turn && d.board[1][1] == turn && d.board[2][0] == turn {
		return true
	}
	return false
}

func (d *dataMembers) computerMove() (int, int) {
	var emptyFields []int
	for i := 0; i < len(d.board); i++ {
		for j := 0; j < len(d.board[0]); j++ {
			if d.board[i][j] == 0 {
				emptyFields = append(emptyFields, 3*i+j)
			}
		}
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	newPosition := r.Intn(len(emptyFields))
	val := emptyFields[newPosition]

	return val / 3, val % 3
}
