package benchmark

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/log"
)

var currentYear = time.Now().Year()

func benchmarkLogger_Debug(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("%s %s %s %s %s %s %s %s %s %d, test debug log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}

}

func benchmarkLogger_Info(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("%s %s %s %s %s %s %s %s %s %d, test info log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_Warn(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Warn("%s %s %s %s %s %s %s %s %s %d, test warn log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_Error(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error(errors.New("log err"), "%s %s %s %s %s %s %s %s %s %d,test error log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_ErrorIf(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorIf(errors.New("log err from ErrorIf"), "%s %s %s %s %s %s %s %s %s %d, test error log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_ErrorIf_NilError(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorIf(nil, "%s %s %s %s %s %s %s %s %s %d, this error log will not be printed", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_Log(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Log("%s %s %s %s %s %s %s %s %s %d, test log log", "Hello,", "this", "is", "a", "logger", "component", "based", "on", "golang", currentYear)
	}
}

func benchmarkLogger_Write(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DefaultLogger().Write([]byte("Hello, this is a logger component based on golang\n"))
	}
}
