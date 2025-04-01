package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CustomLogger struct {
	log.Formatter
}

func (c CustomLogger) Format(entry *log.Entry) ([]byte, error) {
	entry.Time = entry.Time.UTC()
	return c.Formatter.Format(entry)
}

var Logger *log.Logger

func formatFileName(filePath string) string {
	files := strings.Split(filePath, "/")
	return files[len(files)-1]
}

func init() {
	Logger = log.New()
	Logger.SetReportCaller(true)
	Logger.SetLevel(log.InfoLevel)
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(CustomLogger{&log.JSONFormatter{CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
		return "", fmt.Sprintf("%s:%d", formatFileName(frame.File), frame.Line)
	}}})
}
