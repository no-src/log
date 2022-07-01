package benchmark

import (
	"testing"

	"github.com/no-src/log"
)

func initEmptyLogger() {
	log.InitDefaultLogger(log.NewEmptyLogger())
}

func BenchmarkEmptyLogger_Debug(b *testing.B) {
	benchmarkLogger_Debug(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_Info(b *testing.B) {
	benchmarkLogger_Info(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_Warn(b *testing.B) {
	benchmarkLogger_Warn(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_Error(b *testing.B) {
	benchmarkLogger_Error(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_ErrorIf(b *testing.B) {
	benchmarkLogger_ErrorIf(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_ErrorIf_NilError(b *testing.B) {
	benchmarkLogger_ErrorIf_NilError(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_Log(b *testing.B) {
	benchmarkLogger_Log(b, initEmptyLogger)
}

func BenchmarkEmptyLogger_Write(b *testing.B) {
	benchmarkLogger_Write(b, initEmptyLogger)
}
