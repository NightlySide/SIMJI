package log

import (
	"strings"
	"testing"
)

func TestSingleton(t *testing.T) {
	log1 := GetLogger()
	log2 := GetLogger()

	if log1 != log2 {
		t.Error("Singleton not working, multiple instances")
	}
}

func TestLogLevel(t *testing.T) {
	log := GetLogger()

	log.SetLevel(DEBUG)

	if log.level != DEBUG {
		t.Error("Cannot set logging level")
	}

	log.SetLevel(ERROR)
	output := CaptureOutput(func() { log.Info("should'nt be printed") })
	if output != "" {
		t.Error("sqfqfqf")
	}
}

func TestPrinting(t *testing.T) {
	log := GetLogger()
	GetLogger().SetLevel(DEBUG)
	phrase := "Hello world!"
	output := CaptureOutput(func() { log.Debug(phrase) })

	if output != phrase {
		t.Errorf("Log debug error - Expected: %s, got: %s\n", phrase, output)
	}
}

func TestTitle(t *testing.T) {
	log := GetLogger()
	output := CaptureOutput(func() { log.Title(DEBUG, "Hello world") })
	if strings.TrimSpace(output) != "===Hello world===" {
		t.Error("Error while formatting title")
	}
}

func TestLevelPrinting(t *testing.T) {
	log := GetLogger()
	output := CaptureOutput(func() { log.Info("test") })
	if strings.TrimSpace(output) != "[+] INFO: test" {
		t.Error("Error while formatting log INFO")
		t.Error(output)
	}

	output = CaptureOutput(func() { log.Warn("test") })
	if strings.TrimSpace(output) != "[/] WARN: test" {
		t.Error("Error while formatting log WARN")
		t.Error(output)
	}

	output = CaptureOutput(func() { log.Error("test") })
	if strings.TrimSpace(output) != "[-] ERROR: test" {
		t.Error("Error while formatting log ERROR")
		t.Error(output)
	}

	output = CaptureOutput(func() { log.Success("test") })
	if strings.TrimSpace(output) != "[+] test" {
		t.Error("Error while formatting log SUCESS")
		t.Error(output)
	}
}
