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

func main() {
	game := NewGame()

	for {
		game.PrintBoard()
		fmt.Printf("Joueur %d, entrez une colonne (0-6) : ", game.Player)
		var col int
		_, err := fmt.Scan(&col)
		if err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}
		if !game.Drop(col) {
			fmt.Println("Coup invalide. Réessayez.")
			continue
		}
		// potentiel connerie qui affiche une win
		game.SwitchPlayer()
	}

}
