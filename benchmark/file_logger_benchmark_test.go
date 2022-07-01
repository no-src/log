package benchmark

import (
	"testing"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
)

func initFileLogger() {
	fl, err := log.NewFileLogger(level.DebugLevel, "./logs", "benchmark")
	if err != nil {
		panic("init file logger error")
	}
	log.InitDefaultLogger(fl)
}

func BenchmarkFileLogger_Debug(b *testing.B) {
	benchmarkLogger_Debug(b, initFileLogger)
}

func BenchmarkFileLogger_Info(b *testing.B) {
	benchmarkLogger_Info(b, initFileLogger)
}

func BenchmarkFileLogger_Warn(b *testing.B) {
	benchmarkLogger_Warn(b, initFileLogger)
}

func BenchmarkFileLogger_Error(b *testing.B) {
	benchmarkLogger_Error(b, initFileLogger)
}

func BenchmarkFileLogger_ErrorIf(b *testing.B) {
	benchmarkLogger_ErrorIf(b, initFileLogger)
}

func BenchmarkFileLogger_ErrorIf_NilError(b *testing.B) {
	benchmarkLogger_ErrorIf_NilError(b, initFileLogger)
}

func BenchmarkFileLogger_Log(b *testing.B) {
	benchmarkLogger_Log(b, initFileLogger)
}

func BenchmarkFileLogger_Write(b *testing.B) {
	benchmarkLogger_Write(b, initFileLogger)
}
