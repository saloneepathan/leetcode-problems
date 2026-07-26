package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Spreadsheet struct {
	grid [][]int
	rows int
}

// Constructor initializes the spreadsheet
func Constructor(rows int) Spreadsheet {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, 26) // 26 columns A..Z
	}
	return Spreadsheet{grid: grid, rows: rows}
}

// helper: parse "A1" -> (rowIdx, colIdx)
func (this *Spreadsheet) parseCell(cell string) (int, int) {
	col := int(cell[0] - 'A')
	row, _ := strconv.Atoi(cell[1:])
	return row - 1, col
}

// SetCell sets a cell to a value
func (this *Spreadsheet) SetCell(cell string, value int) {
	r, c := this.parseCell(cell)
	this.grid[r][c] = value
}

// ResetCell resets a cell to 0
func (this *Spreadsheet) ResetCell(cell string) {
	r, c := this.parseCell(cell)
	this.grid[r][c] = 0
}

// helper: resolve operand (either int or cell reference)
func (this *Spreadsheet) resolve(op string) int {
	if unicode.IsDigit(rune(op[0])) { // it's a number
		val, _ := strconv.Atoi(op)
		return val
	}
	r, c := this.parseCell(op)
	return this.grid[r][c]
}

// GetValue evaluates a formula =X+Y
func (this *Spreadsheet) GetValue(formula string) int {
	formula = formula[1:] // strip leading '='
	splitIdx := -1
	for i, ch := range formula {
		if ch == '+' {
			splitIdx = i
			break
		}
	}
	left := formula[:splitIdx]
	right := formula[splitIdx+1:]
	return this.resolve(left) + this.resolve(right)
}

func main() {
	// Example usage (from the problem statement)
	spreadsheet := Constructor(3)

	fmt.Println(spreadsheet.GetValue("=5+7")) // 12

	spreadsheet.SetCell("A1", 10)
	fmt.Println(spreadsheet.GetValue("=A1+6")) // 16

	spreadsheet.SetCell("B2", 15)
	fmt.Println(spreadsheet.GetValue("=A1+B2")) // 25

	spreadsheet.ResetCell("A1")
	fmt.Println(spreadsheet.GetValue("=A1+B2")) // 15
}
