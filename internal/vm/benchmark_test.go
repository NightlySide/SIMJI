package vm

import (
	"simji/internal/log"
	"strings"
	"testing"
)

func TestNewBenchmark(t *testing.T) {
	bm := StartBenchmark([]int{0x00000000}, 1)

	if len(bm.nbCycles) != 1 || len(bm.nbCyclesPerSecs) != 1 || bm.nbRuns != 1 || len(bm.totalTimes) != 1 {
		t.Error("Error in number of benchmark passes")
	}
}

func TestBMStats(t *testing.T) {
	bm := StartBenchmark([]int{0x00000000}, 1)
	bm.nbCyclesPerSecs = []float64{100.0, 0.0}

	if bm.average() != 50.0 {
		t.Error("Error while computing Benchmark average")
	}
	if bm.standardDeviation() != 70.71067811865476 {
		t.Error("Error while computing Benchmark STD")
	}
	if bm.median() != 50 {
		t.Error("Error while computing Benchmark median")
	}
	if bm.topOnePercent() != 0.0 {
		t.Error("Error while computing Benchmark top 1%")
	}
	if bm.bottomOnePercent() != 100.0 {
		t.Error("Error while computing Benchmark bottom 1%")
	}
}

func TestPrintResults(t *testing.T) {
	bm := StartBenchmark([]int{0x00000000}, 10)
	output := log.CaptureOutput(bm.PrintResults)
	lines := strings.Split(output, "\n")
	if len(lines) != 9 {
		t.Error("Wrong number of lines in the BM print results")
	}
}
