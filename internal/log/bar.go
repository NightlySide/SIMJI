package log

import "fmt"

// Bar est une barre de chargement à afficher dans la console
type Bar struct {
	percent int
	current int
	total   int
	rate    string
	graph   string
}

// NewOption initialise une barre ainsi que son symbole
func (bar *Bar) NewOption(start int, total int) {
	bar.current = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph
	}
}

func (bar *Bar) getPercent() int {
	return int(float64(bar.current) / float64(bar.total) * 100)
}

// NewOptionWithGraph initialise la barre avec un symbole de chargement spécifique
func (bar *Bar) NewOptionWithGraph(start int, total int, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

// Play permet d'afficher la barre à un temps donné en fonction de la valeur donnée
func (bar *Bar) Play(current int) {
	bar.current = current
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.current, bar.total)
}

// Finish permet d'afficher une nouvelle ligne pour éviter les artefacts graphiques
func (bar *Bar) Finish() {
	fmt.Println()
}
