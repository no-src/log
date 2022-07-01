package benchmark

import (
	"testing"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
)

func initConsoleLogger() {
	log.InitDefaultLogger(log.NewConsoleLogger(level.DebugLevel))
}

func BenchmarkConsoleLogger_Debug(b *testing.B) {
	benchmarkLogger_Debug(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_Info(b *testing.B) {
	benchmarkLogger_Info(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_Warn(b *testing.B) {
	benchmarkLogger_Warn(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_Error(b *testing.B) {
	benchmarkLogger_Error(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_ErrorIf(b *testing.B) {
	benchmarkLogger_ErrorIf(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_ErrorIf_NilError(b *testing.B) {
	benchmarkLogger_ErrorIf_NilError(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_Log(b *testing.B) {
	benchmarkLogger_Log(b, initConsoleLogger)
}

func BenchmarkConsoleLogger_Write(b *testing.B) {
	benchmarkLogger_Write(b, initConsoleLogger)
}
