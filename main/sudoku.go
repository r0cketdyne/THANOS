package main

import (
	"github.com/01-edu/z01"
	"sudoku"
	"os"
)

func main() {
	//baby chimera
	params := os.Args[1:]
	//если количество строк не равно 9, выводим Error
	if len(params) != 9 {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
		return
	}
	//yeet
	var table [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			
			if len(params[i]) != 9 {
				z01.PrintRune('E')
				z01.PrintRune('r')
				z01.PrintRune('r')
				z01.PrintRune('o')
				z01.PrintRune('r')
				z01.PrintRune('\n')
				return
			}
			
			if params[i][j] == '.' {
				table[i][j] = 0
			
			} else if params[i][j] >= '0' && params[i][j] <= '9' {
				table[i][j] = int(params[i][j] - '0')

			
			} else {
				z01.PrintRune('E')
				z01.PrintRune('r')
				z01.PrintRune('r')
				z01.PrintRune('o')
				z01.PrintRune('r')
				z01.PrintRune('\n')
				return
			}
		}
	}
	
	if !sudoku.FindSolution(table) {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
	}
}
