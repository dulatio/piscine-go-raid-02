package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	// declaring board
	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
	}
	// filling board with initial values
	for i, v := range os.Args {
		if i != 0 {
			if len(v) != 9 {
				fmt.Println("Error")
				return
			}
			for j, r := range v {
				if r != '.' {
					if r < '1' || r > '9' {
						fmt.Println("Error")
						return
					} else {
						if isSafe(board, i-1, j, int(r-'0')) {
							board[i-1][j] = int(r - '0')
						} else {
							fmt.Println("Error")
							return
						}
					}
				}
			}
		}
	}

	finalboard := make([][]int, 9)
	for i := range finalboard {
		finalboard[i] = make([]int, 9)
	}
	counter := 0
	solve(board, finalboard, &counter)
	if counter == 1 {
		printBoard(finalboard)
	} else {
		fmt.Println("Error")
	}
}

func solve(board [][]int, finalboard [][]int, counter *int) bool {
	row := -1
	col := -1
	isFull := true
	// in our case the length of thr board is 9
	n := len(board)
	// check if we have fields without numbers
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				row = i
				col = j
				isFull = false
				break
			}
		}
		if isFull == false {
			break
		}
	}
	if isFull {
		*counter++
		if finalboard[0][0] == 0 {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					finalboard[i][j] = board[i][j]
				}
			}
		}
		return false
	}
	// try values in empty fields recursively
	for i := 1; i <= n; i++ {
		if isSafe(board, row, col, i) {
			board[row][col] = i
			if solve(board, finalboard, counter) {
				return true
			} else {
				board[row][col] = 0
			}
		}
	}
	return false
}

func isSafe(board [][]int, row int, col int, num int) bool {
	// in our case n is 9
	n := len(board)
	// check the uniqueness of num horizontally
	for x := 0; x < n; x++ {
		if board[row][x] == num {
			return false
		}
	}
	// check the uniqueness of num vertically
	for y := 0; y < n; y++ {
		if board[y][col] == num {
			return false
		}
	}
	// apecific value for board of size 9
	sqrtn := 3
	boxRowStart := row - row%sqrtn
	boxColStart := col - col%sqrtn
	// check the uniqueness of num in the corresponding box
	for y := boxRowStart; y < boxRowStart+sqrtn; y++ {
		for x := boxColStart; x < boxColStart+sqrtn; x++ {
			if board[y][x] == num {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]int) {
	n := len(board)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			if x != 0 {
				fmt.Print(" ")
			}
			fmt.Print(board[y][x])
		}
		fmt.Println()
	}
}
