package log

import (
	"fmt"
	"time"
	"net/http"
	"io"
	"os"
)

const (
	LOG_TIME_FORMAT = "2006-01-02 15:04:05"
)

type logWriter struct {
	std io.Writer
}

func (this *logWriter) write(content string) (int, error) {
	msg := fmt.Sprintf("[%s] %s", time.Now().Format(LOG_TIME_FORMAT), content)
	return fmt.Fprintln(this.std, msg)
}

func printLine(msg string, std io.Writer) {
	logger := &logWriter{std: std}
	logger.write(msg)
}

func FatalAndExit(v ...interface{}) {
	printLine(fmt.Sprintf("[FATAL AND EXIT] %s", fmt.Sprint(v...)), os.Stderr)
	os.Exit(1)
}

func Access(r *http.Request) {
	params := ""
	queryParams := r.URL.Query()
	if len(queryParams) > 0 {
		params = fmt.Sprint(queryParams)
	}

	printLine(fmt.Sprintf("[ACCESS] %s %s %s",
		r.Method, r.URL.Path, params), os.Stdout)
}

func Error(v ...interface{}) {
	printLine(fmt.Sprintf("[ERROR] %s", fmt.Sprint(v...)), os.Stderr)
}

func Info(v ...interface{}) {
	printLine(fmt.Sprintf("[INFO] %s", fmt.Sprint(v...)), os.Stdout)
}
