package Power4

import (
	"fmt"
)

const (
	Rows    = 6
	Columns = 7
)

type Game struct {
	Board  [Rows][Columns]int
	Player int
}

func NewGame() *Game {
	return &Game{Player: 1}
}

func (g *Game) PrintBoard() {
	fmt.Println("\n 0 1 2 3 4 5 6")
	for _, row := range g.Board {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(" .")
			case 1:
				fmt.Print(" X")
			case 2:
				fmt.Print(" O")
			}
		}
		fmt.Println()
	}
}

func (g *Game) Drop(col int) bool {
	if col < 0 || col >= Columns {
		return false
	}
	for row := Rows - 1; row >= 0; row-- {
		if g.Board[row][col] == 0 {
			g.Board[row][col] = g.Player
			return true
		}
	}
	return false
}

func (g *Game) SwitchPlayer() {
	g.Player = 3 - g.Player
}

func (g *Game) CheckWin() bool {
	return g.checkHorizontal() || g.checkVertical() || g.checkDiagonal()
}

func (g *Game) checkHorizontal() bool {
	for row := 0; row < Rows; row++ {
		for col := 0; col <= Columns-4; col++ {
			if g.Board[row][col] != 0 &&
				g.Board[row][col] == g.Board[row][col+1] &&
				g.Board[row][col] == g.Board[row][col+2] &&
				g.Board[row][col] == g.Board[row][col+3] {
				return true
			}
		}
	}
	return false
}

func (g *Game) checkVertical() bool {
	for row := 0; row <= Rows-4; row++ {
		for col := 0; col < Columns; col++ {
			if g.Board[row][col] != 0 &&
				g.Board[row][col] == g.Board[row+1][col] &&
				g.Board[row][col] == g.Board[row+2][col] &&
				g.Board[row][col] == g.Board[row+3][col] {
				return true
			}
		}
	}
	return false
}

func (g *Game) checkDiagonal() bool {
	for row := 0; row <= Rows-4; row++ {
		for col := 0; col <= Columns-4; col++ {
			if g.Board[row][col] != 0 &&
				g.Board[row][col] == g.Board[row+1][col+1] &&
				g.Board[row][col] == g.Board[row+2][col+2] &&
				g.Board[row][col] == g.Board[row+3][col+3] {
				return true
			}
		}
	}

	for row := 3; row < Rows; row++ {
		for col := 0; col <= Columns-4; col++ {
			if g.Board[row][col] != 0 &&
				g.Board[row][col] == g.Board[row-1][col+1] &&
				g.Board[row][col] == g.Board[row-2][col+2] &&
				g.Board[row][col] == g.Board[row-3][col+3] {
				return true
			}
		}
	}

	return false
}
