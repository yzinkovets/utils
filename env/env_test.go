package env

import (
	"fmt"
	"os"
	"testing"
)

type testLogger struct {
	Expected string
	Got      string
}

func (l *testLogger) Valid() bool {
	return l.Expected == l.Got
}

func (l *testLogger) Fatalf(format string, v ...interface{}) {
	l.Got = fmt.Sprintf(format, v...)
}
func (l *testLogger) Reset() {
	l.Expected = ""
	l.Got = ""
}

func TestGetEnv(t *testing.T) {
	logger := &testLogger{
		Expected: "Please set TEST environment variable\n",
	}
	SetLogger(logger)
	Must("TEST")
	if !logger.Valid() {
		t.Errorf("Expected [%s], got [%s]", logger.Expected, logger.Got)
	}
	logger.Reset()

	os.Setenv("TEST", "test")
	if Get("TEST") != "test" {
		t.Errorf("Expected [test], got [%s]", Get("TEST"))
	}
}
