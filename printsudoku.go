package sudoku

import "github.com/01-edu/z01"


func PrintSudoku(table [9][9]int) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
			
            z01.PrintRune(rune(table[i][j])+'0')
            if j < 8 {
				
                z01.PrintRune(' ')
            }
        }
        z01.PrintRune('\n')
    }
}
