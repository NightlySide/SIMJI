package log

import (
	"strings"
	"testing"
)

func TestNewBar(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	if bar.current != 0 || bar.total != 100 || bar.graph != "█" {
		t.Error("Error while initializing the progress bar")
	}
}

func TestNewBarWithGraph(t *testing.T) {
	bar := new(Bar)
	bar.NewOptionWithGraph(0, 100, "x")

	if bar.graph != "x" {
		t.Error("Error while init with custom graph")
	}
}

func TestBarPercen(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	if bar.current != 0 || bar.total != 100 || bar.getPercent() != 0 {
		t.Error("Error while getting the bar percentage")
	}

	bar.NewOption(100, 100)
	if bar.getPercent() != 100 {
		t.Error("Error while getting the bar percentage")
	}
}

func TestFinishBar(t *testing.T) {
	bar := new(Bar)
	output := CaptureOutput(bar.Finish)
	if output != "\n" {
		t.Error("Finish function don't retrun empty new line")
	}
}

func TestPlayBar(t *testing.T) {
	bar := new(Bar)
	bar.NewOption(0, 100)
	output := CaptureOutput(func() { bar.Play(0) })
	if strings.TrimSpace(output) != "[                                                  ]  0%        0/100" {
		t.Error("Not getting expected bar output")
	}

	output = CaptureOutput(func() { bar.Play(100) })
	args := strings.Split(strings.TrimSpace(output), " ")
	if args[len(args)-1] != "100/100" {
		t.Error("Not getting expected bar output")
		t.Error(output)
	}
}
