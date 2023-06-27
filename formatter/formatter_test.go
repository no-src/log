package formatter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/no-src/log/content"
	"github.com/no-src/log/formatter"
	_ "github.com/no-src/log/formatter/json"
	_ "github.com/no-src/log/formatter/text"
	"github.com/no-src/log/level"
)

func TestInitDefaultFormatter_Concurrency(t *testing.T) {
	c := 10
	wg := sync.WaitGroup{}
	wg.Add(c * 2)
	for i := 0; i < c; i++ {
		go func() {
			formatter.InitDefaultFormatter(formatter.JsonFormatter)
			wg.Done()
		}()

		go func() {
			formatter.Default()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestNewJsonFormatter(t *testing.T) {
	f := formatter.NewJsonFormatter()
	data, err := f.Serialize(content.NewContent(level.DebugLevel, nil, false, time.RFC3339, "json formatter"))
	if err != nil {
		t.Errorf("test json formatter error => %v", err)
		return
	}
	expect := `{"level":"DEBUG","log":"json formatter"}` + "\n"
	actual := string(data)
	if expect != actual {
		t.Errorf("test json formatter failed, expect to get %s, but actual get %s", expect, actual)
		return
	}
}

func TestNewTextFormatter(t *testing.T) {
	testNewTextFormatter(t, formatter.NewTextFormatter())
}

func testNewTextFormatter(t *testing.T, f formatter.Formatter) {
	data, err := f.Serialize(content.NewContent(level.DebugLevel, nil, false, time.RFC3339, "text formatter"))
	if err != nil {
		t.Errorf("test text formatter error => %v", err)
		return
	}
	expect := `[DEBUG] text formatter` + "\n"
	actual := string(data)
	if expect != actual {
		t.Errorf("test text formatter failed, expect to get %s, but actual get %s", expect, actual)
		return
	}
}

func TestNew_Unsupported(t *testing.T) {
	formatter.InitDefaultFormatter(formatter.TextFormatter)
	testNewTextFormatter(t, formatter.New("unsupported"))
}
