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
	board      [3][3]int
	missing    int
	difficulty int
}

func main() {
	playGame()
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
func playGame() {
	fmt.Println("-----------------------------------------")
	fmt.Println("WELCOME IN CROSS AND CIRCLE GAME")

	// infinite game loop
	for {
		lvl := chooseDifficulty()

		// init game board
		d := dataMembers{
			missing:    9,
			difficulty: lvl,
		}

		if d.difficulty > 2 {
			fmt.Println("At the moment only unbeatable's level is not available.\n")
			continue
		}

		// random selection of first turn (player or computer)
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		beginn := r.Intn(2)
		var turn int
		if beginn == 0 {
			turn = -1
			fmt.Println("Computer starts this round.")
		} else {
			turn = 1
			fmt.Println("You're lucky today! You start the game!")
		}

		for d.missing > 0 {
			var row, col int
			if turn == 1 {
				row, col = d.userInput()
			} else {
				fmt.Println("Computer's turn!")
				time.Sleep(5 * time.Second)
				if d.difficulty == 1 {
					row, col = d.computerMove()
				} else {
					win, field := d.preventLost(-1)
					if win {
						row = field / 3
						col = field % 3
					} else {
						defend, field := d.preventLost(1)
						if defend {
							row = field / 3
							col = field % 3
						} else {
							row, col = d.computerMove()
						}
					}
				}
				fmt.Println("Computer chose row no.", row, "and col no.", col)
			}

			d.board[row][col] = turn
			d.missing -= 1

			if isWin(d.board, turn) {
				d.printBoard()
				if turn == 1 {
					fmt.Println("\nCongratulations, you won!\n")
				} else {
					fmt.Println("\nYou lost!\n")
				}
				break
			}

			if turn == 1 {
				d.printBoard()
			}

			// change turn
			turn *= -1
		}
		if d.missing == 0 {
			fmt.Println("A draw!")
		}
	}
}

func isValidField(d *dataMembers, row int, col int) bool {
	if (*d).board[row][col] == 0 {
		return true
	} else {
		return false
	}
}

func isWin(d [3][3]int, turn int) bool {
	for i := 0; i < len(d); i++ {
		if d[i][0] == turn && d[i][1] == turn && d[i][2] == turn {
			return true
		}
		if d[0][i] == turn && d[1][i] == turn && d[2][i] == turn {
			return true
		}
	}

	if d[0][0] == turn && d[1][1] == turn && d[2][2] == turn {
		return true
	}
	if d[0][2] == turn && d[1][1] == turn && d[2][0] == turn {
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

func chooseDifficulty() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("There are three difficulty levels available: ")
		fmt.Println("	# beginner      - option 1")
		fmt.Println("	# medium        - option 2")
		fmt.Println("	# unbeatable    - option 3")
		for {
			fmt.Println("Please choose a dificulty level to (type 1, 2 or 3): ")
			fmt.Print("-> ")
			lvl, _ := reader.ReadString('\n')
			lvl = strings.Replace(lvl, "\n", "", -1)
			lvl = strings.ToUpper(lvl)
			if len(lvl) == 1 && lvl == "Q" {
				fmt.Println("Thank you very much for the game! Have a great day!")
				fmt.Println("Game will be closed in 5 seconds.")
				time.Sleep(5 * time.Second)
				os.Exit(0)
			}
			if len(lvl) != 1 || lvl < "1" || lvl > "3" {
				fmt.Println("Invalid input!")
				continue
			}

			lvlNumber, err := strconv.Atoi(lvl)
			if err != nil {
				fmt.Println("Invalid input!")
				continue
			}

			return lvlNumber
		}
	}
}

func (d *dataMembers) preventLost(numCheck int) (bool, int) {
	for i := 0; i < len(d.board); i++ {
		for j := 0; j < len(d.board[0]); j++ {
			if d.board[i][j] == 0 {
				var cp [3][3]int
				for k := 0; k < len(d.board); k++ {
					for l := 0; l < len(d.board[0]); l++ {
						cp[k][l] = d.board[k][l]
					}
				}
				cp[i][j] = numCheck
				if isWin(cp, numCheck) {
					return true, 3*i + j
				}
			}
		}
	}
	return false, -1
}
