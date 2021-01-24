package vm

import (
	"fmt"
	"math"
	"simji/internal/log"
	"sort"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Benchmark est une structure contenant le résultat
// d'exécution de la machine virtuelle
type Benchmark struct {
	nbRuns          int
	totalTimes      []time.Duration
	nbCycles        []int
	nbCyclesPerSecs []float64
}

// StartBenchmark permet de lancer le benchmark d'une suite
// d'instructions et retourne un objet capable de faire les statistiques
func StartBenchmark(program []int, nbRuns int) Benchmark {
	var bm Benchmark
	var machine VM
	var bar log.Bar

	bar.NewOption(0, nbRuns)
	for i := 0; i < nbRuns; i++ {
		bar.Play(i + 1)
		machine = NewVM(32, 1000)
		machine.LoadProg(program)
		machine.Run(false, false, false)

		bm.nbRuns++
		bm.totalTimes = append(bm.totalTimes, machine.totalTime)
		bm.nbCycles = append(bm.nbCycles, machine.cycles)
		bm.nbCyclesPerSecs = append(bm.nbCyclesPerSecs, float64(machine.cycles)/machine.totalTime.Seconds())
	}
	bar.Finish()

	// sorting results
	sort.Float64s(bm.nbCyclesPerSecs)

	return bm
}

func (bm *Benchmark) average() float64 {
	var sum float64
	for _, nbCyclesSec := range bm.nbCyclesPerSecs {
		sum += nbCyclesSec
	}

	return sum / float64(len(bm.nbCyclesPerSecs))
}

func (bm *Benchmark) median() float64 {
	middle := len(bm.nbCyclesPerSecs) / 2
	result := bm.nbCyclesPerSecs[middle]
	if len(bm.nbCyclesPerSecs)%2 == 0 {
		result = (result + bm.nbCyclesPerSecs[middle-1]) / 2
	}
	return result
}

func (bm *Benchmark) topOnePercent() float64 {
	onePercentIdx := int(len(bm.nbCyclesPerSecs) * 99 / 100)
	return bm.nbCyclesPerSecs[onePercentIdx]
}

func (bm *Benchmark) bottomOnePercent() float64 {
	onePercentIdx := int(len(bm.nbCyclesPerSecs) * 1 / 100)
	return bm.nbCyclesPerSecs[onePercentIdx]
}

func (bm *Benchmark) standardDeviation() float64 {
	mean := bm.average()

	var total float64
	for _, number := range bm.nbCyclesPerSecs {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(bm.nbCyclesPerSecs)-1)
	return math.Sqrt(variance)
}

// PrintResults permet d'afficher les résultat d'un benchmark
func (bm *Benchmark) PrintResults() {
	log.GetLogger().Title(log.INFO, "Benchmark results")

	pp := message.NewPrinter(language.French)

	pp.Printf(" Moyenne:\t%d op/sec\n", int(bm.average()))
	pp.Printf(" Médiane:\t%d op/sec\n", int(bm.median()))
	pp.Printf(" Ecart-type:\t%d op/sec\n", int(bm.standardDeviation()))
	pp.Printf(" Top 1%%:\t%d op/sec\n", int(bm.topOnePercent()))
	pp.Printf(" Bottom 1%%:\t%d op/sec\n", int(bm.bottomOnePercent()))

	fmt.Println("=======================")
}
