package benchmark

import (
	"errors"
	"testing"

	"github.com/no-src/log"
)

func benchmarkLogger_Debug(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Debug("%s %s, test debug log", "hello", "world")
	}

}

func benchmarkLogger_Info(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("%s %s, test info log", "hello", "world")
	}
}

func benchmarkLogger_Warn(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Warn("%s %s, test warn log", "hello", "world")
	}
}

func benchmarkLogger_Error(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Error(errors.New("log err"), "%s %s,test error log", "hello", "world")
	}
}

func benchmarkLogger_ErrorIf(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorIf(errors.New("log err from ErrorIf"), "%s %s, test error log", "hello", "world")
	}
}

func benchmarkLogger_ErrorIf_NilError(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.ErrorIf(nil, "%s %s, this error log will not be printed", "hello", "world")
	}
}

func benchmarkLogger_Log(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Log("%s %s, test log log", "hello", "world")
	}
}

func benchmarkLogger_Write(b *testing.B, initLogger func()) {
	initLogger()
	defer log.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.DefaultLogger().Write([]byte("hello logger"))
	}
}
