package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	playGame()
}

// game variables
type dataMembers struct {
	board      [3][3]int
	missing    int
	difficulty int
}

// Print current state of the game
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

// Print one row of the game board
//
// Function arguments:
// row - slice of data to print
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

// Ask player for input and validate the input.
// Return choose row and column to mark, or exit the game when 'Q' is pressed.
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

// Main game loop
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
				} else if d.difficulty == 2 {
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
				} else {
					if d.missing >= 5 {
						move := d.bestMove()
						row = move / 3
						col = move % 3
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

// Validates if given field is already occupied.
// Return true if the field is not occupied, otherwise false.
//
// Function arguments:
// d - instance of game variables with board to check
// row - number of row to check
// col - number of column to check
func isValidField(d *dataMembers, row int, col int) bool {
	if (*d).board[row][col] == 0 {
		return true
	} else {
		return false
	}
}

// Validate if latest move won the game.
// Return true in case of win, false otherwise.
//
// Function arguments:
// d - matrix representing current board
// turn - latest player (1 # player, -1 # computer)
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

// Choose next computer's move
// Return two ints, representing number of row and column choosen
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

// Ask user for difficulty level
// Return int representing difficulty level.
// 	# beginner      - option 1
//	# medium        - option 2
//	# unbeatable    - option 3
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

// Check if there is possibility to win/lose in one move.
// If true, it return true and int representing next field to choose.
// Otherwise, it returns false and -1.
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

// Find the best possible move in unbeatable level.
// Return int representing the best next move.
func (d *dataMembers) bestMove() int {
	fmt.Println("Number of missing fields:", d.missing)
	// slice with all matrix corners
	corners := [4]int{0, 2, 6, 8}
	// center field is the best field, if it is not occupy, choose it
	if d.board[1][1] == 0 {
		return 1*3 + 1
	}
	// choose random corner
	if d.missing == 8 {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(4)
		return corners[newPosition]
	}
	if d.missing == 7 {
		// find occupied by player field
		var playerField int
		for i := 0; i < 9; i++ {
			if d.board[i/3][i%3] == 1 {
				playerField = i
				break
			}
		}
		for _, val := range corners {
			// board without a chance to win
			if val == playerField {
				if playerField+3 <= 8 {
					return playerField + 3
				} else {
					return playerField - 3
				}
			}
		}
		// computer has a winning position
		// choose any corner neighbouring the filed occupied by player
		if playerField%3 != 2 {
			if playerField+1 != 4 {
				return playerField + 1
			} else {
				return playerField - 3
			}
		} else {
			return playerField - 3
		}
	}
	if d.missing == 6 {
		risk, move := d.preventLost(1)
		// defend yourself from losing
		if risk {
			return move
		} else {
			if d.board[1][1] == 1 {
				for _, val := range corners {
					// chose first free corner
					if d.board[val/3][val%3] == 0 {
						return val
					}
				}
			} else {
				// store two fields occupied by player
				var occupiedFields []int
				for i := 0; i < 9; i++ {
					if d.board[i/3][i%3] == 1 {
						occupiedFields = append(occupiedFields, i)
					}
				}
				// find nearest free field to two occupied fields by player
				minVal := 10
				index := 0
				for i := 0; i < 9; i++ {
					if i != 4 && d.board[i/3][i%3] == 0 {
						val1 := findDistance(i, occupiedFields[0])
						val2 := findDistance(i, occupiedFields[1])
						if val1+val2 < minVal {
							index = i
							minVal = val1 + val2
							if minVal == 2 {
								return index
							}
						}
					}
				}
				return index
			}
		}
	}
	if d.missing == 5 {
		win, move := d.preventLost(-1)
		// check possibility of winning
		if win {
			return move
		}
		risk, move2 := d.preventLost(1)
		// defend yourself from losing
		if risk {
			return move2
		}
		// store two fields occupied by player
		minVal := 10
		index := 0
		var occupiedFields []int
		for i := 0; i < 9; i++ {
			if d.board[i/3][i%3] == -1 {
				occupiedFields = append(occupiedFields, i)
			}
		}
		// find nearest free field to two occupied fields by player
		for i := 0; i < 9; i++ {
			if i != 4 && d.board[i/3][i%3] == 0 {
				val1 := findDistance(i, occupiedFields[0])
				val2 := findDistance(i, occupiedFields[1])
				if val1+val2 < minVal {
					index = i
					minVal = val1 + val2
					if minVal == 1 {
						return index
					}
				}
			}
		}
		return index
	}
	// this line will never be reached, function pre-exist before
	return 0
}

// Find closest free field to occupied fields.
// Return int representing the described field.
//
// Function arguments:
// x - 1st occupied field
// y - 2nd occupied field
func findDistance(x int, y int) int {
	ix := x / 3
	iy := y / 3
	jx := x % 3
	jy := y % 3
	if math.Abs(float64(ix)-float64(iy)) <= 1.5 && math.Abs(float64(jx)-float64(jy)) <= 1.5 {
		return 1
	} else {
		return 2
	}
}
