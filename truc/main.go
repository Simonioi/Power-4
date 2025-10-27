package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	p4 "power-4"
	"strconv"
	"sync"
)

var tmpl *template.Template
var game *p4.Game
var mu sync.Mutex

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		log.Println("template execute error:", err)
	}
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "menu.html", nil)
}

func playHandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	if game == nil {
		game = p4.NewGame()
	}
	mu.Unlock()

	if r.Method == http.MethodPost {
		colStr := r.FormValue("col")
		col, err := strconv.Atoi(colStr)
		if err == nil {
			mu.Lock()
			dropped := game.Drop(col)
			if dropped {

				if game.CheckWin() {
					winner := game.Player
					mu.Unlock()
					data := map[string]interface{}{"Winner": winner}
					renderTemplate(w, "victory.html", data)
					return
				}
				game.SwitchPlayer()
			}
			mu.Unlock()
		}
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}

	type Cell struct{ Couleur string }
	mu.Lock()
	grille := make([][]Cell, p4.Rows)
	for i := 0; i < p4.Rows; i++ {
		grille[i] = make([]Cell, p4.Columns)
		for j := 0; j < p4.Columns; j++ {
			switch game.Board[i][j] {
			case 1:
				grille[i][j].Couleur = "rouge"
			case 2:
				grille[i][j].Couleur = "jaune"
			default:
				grille[i][j].Couleur = ""
			}
		}
	}
	current := game.Player
	mu.Unlock()

	data := map[string]interface{}{
		"Grille": grille,
		"Player": current,
	}
	renderTemplate(w, "power.html", data)
}

func quitHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "goodbye.html", nil)
}

func newgameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	game = p4.NewGame()
	mu.Unlock()
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func main() {
	funcMap := template.FuncMap{
		"seq": func(start, end int) []int {
			if end < start {
				return []int{}
			}
			out := make([]int, 0, end-start+1)
			for i := start; i <= end; i++ {
				out = append(out, i)
			}
			return out
		},
	}

	var err error
	baseDir, _ := filepath.Abs(".")
	glob := filepath.Join(baseDir, "tmpl", "*.html")
	tmpl, err = template.New("tmpl").Funcs(funcMap).ParseGlob(glob)
	if err != nil {
		tmpl, err = template.New("tmpl").Funcs(funcMap).ParseGlob("tmpl/*.html")
		if err != nil {
			log.Fatal("Impossible de parser les templates:", err)
		}
	}

	staticDir := filepath.Join(baseDir, "static")
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", menuHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/newgame", newgameHandler)
	http.HandleFunc("/quit", quitHandler)

	log.Println("Serveur démarré sur http://localhost:8081 — appuyez sur CTRL+C pour arrêter")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
