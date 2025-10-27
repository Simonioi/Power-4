package Power4

import (
	"fmt"
)


func CoupJoueur(plateau * [6][7]int, joueur int, colonne int){
	if colonne < 0 || colonne >= 7 {
		fmt.Println("Colonne invalide")
		return
	}
	if plateau[0][colonne] != 0 {
		fmt.Println("Colonne remplie")
		return
	}
	for ligne := 5; ligne >= 0; ligne = ligne - 1 {
        if plateau[ligne][colonne] == 0 {
            plateau[ligne][colonne] = joueur
            return
	    }
	
	}
}