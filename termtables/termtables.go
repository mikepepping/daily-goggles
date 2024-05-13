package termtables

import (
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	Columns  int
	Headings []string
	Rows     [][]string

	columnSizes []int
}

func (t *Table) Resize(columns int) {
	lenColSizes := len(t.columnSizes)

	if lenColSizes == 0 {
		t.columnSizes = make([]int, columns)
		t.Columns = columns
		return
	}

	if lenColSizes < columns {
		t.columnSizes = append(t.columnSizes, make([]int, columns-lenColSizes)...)
	} else if lenColSizes > columns {
		t.columnSizes = t.columnSizes[:columns]
	}

	for i, row := range t.Rows {
		t.Rows[i] = resizeRow(row, columns)
	}

	t.Columns = columns
}

func resizeRow(row []string, columns int) []string {
	rowLen := len(row)
	if rowLen < columns {
		return append(row, make([]string, columns-rowLen)...)
	} else if rowLen > columns {
		return row[:columns]
	}

	return row
}

func (t *Table) SetHeadings(headings []string) error {
	t.Headings = resizeRow(headings, t.Columns)

	// recalculate the columns sizes
	for i, cell := range headings {
		cellLen := len(cell)
		if t.columnSizes[i] < cellLen {
			t.columnSizes[i] = cellLen
		}
	}

	return nil
}

func (t *Table) AddRow(row []string) error {
	missingColumns := t.Columns - len(row)

	if missingColumns < 0 {
		return errors.New("row has too many columns")
	}

	// assign any empty cells with empty string
	row = resizeRow(row, t.Columns)

	// recalculate the columns sizes
	for i, cell := range row {
		cellLen := len(cell)
		if t.columnSizes[i] < cellLen {
			t.columnSizes[i] = cellLen
		}
	}

	t.Rows = append(t.Rows, row)

	return nil
}

func (t Table) Print() {
	tWidth := t.width()
	hDivider := strings.Join([]string{"+", strings.Repeat("-", tWidth), "+"}, "")

	// print heading
	fmt.Println(hDivider)
	paddedHeadings := []string{}
	for i, cell := range t.Headings {
		paddedHeadings = append(paddedHeadings, padRight(cell, t.columnSizes[i]))
	}
	fmt.Println("|", strings.Join(paddedHeadings, " | "), "|")
	fmt.Println(hDivider)

	// print rows
	for _, row := range t.Rows {
		paddedCells := []string{}
		for i, cell := range row {
			paddedCells = append(paddedCells, padRight(cell, t.columnSizes[i]))
		}
		fmt.Println("|", strings.Join(paddedCells, " | "), "|")
		fmt.Println(hDivider)
	}
}

func (t Table) width() int {
	width := 0
	for _, size := range t.columnSizes {
		width += size
	}

	width += (t.Columns * 2) // there is at least a single space character of padding for each cell
	width += (t.Columns - 1) // there are n+1 vertical dividers

	return width
}

func padRight(str string, maxLength int) string {
	padLen := maxLength - len(str)
	padding := strings.Repeat(" ", padLen)
	return str + padding
}
