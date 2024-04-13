package main

import (
	"fmt"
	"os"
	"strconv"
)

// const define two constants
const (
	size       = 9
	emptyValue = '.'
)

type Sudoku [size][size]int

func main() {
	args := os.Args[1:]
	if len(args) != size {
		fmt.Println("Error")
		return
	}
	// Create a Sudoku grid
	var sudoku Sudoku
	if !parseInput(args, &sudoku) {
		fmt.Println("Error")
		return
	}

	if !isValidSudoku(&sudoku) {
		fmt.Println("Error")
		return
	}

	solveSudoku(&sudoku)
	printSudoku(&sudoku)
}

// row len must be equal with size
func parseInput(args []string, sudoku *Sudoku) bool {
	for i, row := range args {
		if len(row) != size {
			return false
		}
		for j, char := range row {
			if char == emptyValue {
				sudoku[i][j] = 0
			} else {
				num, err := strconv.Atoi(string(char))
				if err != nil || num < 1 || num > 9 {
					// if not a valid digit between 1 and 9
					return false
				}
				sudoku[i][j] = num
			}
		}
	}
	return true
}

func isValidSudoku(sudoku *Sudoku) bool {
	// Check uniqueness
	for i := 0; i < size; i++ {
		if !isUnique(sudoku[i][:]) || !isUnique(sudoku.column(i)) {
			return false
		}
	}
	// Check 3x3 subgrids for uniqueness
	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			if !isUnique(sudoku.subgrid(i, j)) {
				return false
			}
		}
	}
	return true
}

// column returns a slice representing the specified column in the Sudoku grid.
func (s *Sudoku) column(colIdx int) []int {
	col := make([]int, size)
	for i := 0; i < size; i++ {
		col[i] = s[i][colIdx]
	}
	return col
}

// subgrid returns a slice representing the 3x3 subgrid starting at the specified row and column.
func (s *Sudoku) subgrid(rowIdx, colIdx int) []int {
	sub := make([]int, 0, size)
	for i := rowIdx; i < rowIdx+3; i++ {
		for j := colIdx; j < colIdx+3; j++ {
			sub = append(sub, s[i][j])
		}
	}
	return sub
}

func isUnique(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if num != 0 && seen[num] {
			return false
		}
		seen[num] = true
	}
	return true
}

func solveSudoku(sudoku *Sudoku) bool {
	var row, col int
	if !findEmptyLocation(sudoku, &row, &col) {
		return true
	}

	for num := 1; num <= size; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num

			if solveSudoku(sudoku) {
				return true
			}

			sudoku[row][col] = 0
		}
	}

	return false // No solution found
}

// finds an empty cell
func findEmptyLocation(sudoku *Sudoku, row, col *int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if sudoku[i][j] == 0 {
				*row, *col = i, j // Update row and column with empty cell coordinates
				return true
			}
		}
	}
	return false
}

// isSafe checks if placing a number in a cell is safe
func isSafe(sudoku *Sudoku, row, col, num int) bool {
	return !usedInRow(sudoku, row, num) &&
		!usedInColumn(sudoku, col, num) &&
		!usedInSubgrid(sudoku, row-row%3, col-col%3, num)
}

// usedInRow checks if a number is already used in the specified row.
func usedInRow(sudoku *Sudoku, row, num int) bool {
	for i := 0; i < size; i++ {
		if sudoku[row][i] == num {
			return true
		}
	}
	return false
}

// usedInColumn checks if a number is already used in the specified column.
func usedInColumn(sudoku *Sudoku, col, num int) bool {
	for i := 0; i < size; i++ {
		if sudoku[i][col] == num {
			return true
		}
	}
	return false
}

// usedInSubgrid checks if a number is already used in the specified 3x3 subgrid.
func usedInSubgrid(sudoku *Sudoku, startRow, startCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	return false
}

// printSudoku prints the Sudoku grid.
func printSudoku(sudoku *Sudoku) {
	for _, row := range sudoku {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
