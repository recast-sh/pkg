package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func NewStdoutFormatter() Formatter {
	return &StdoutFormatter{
		w: bufio.NewWriter(os.Stdout),
	}
}

type StdoutFormatter struct {
	w *bufio.Writer
}

const (
	stdoutTimestampFormat = "15:04:05"
)

func (s *StdoutFormatter) Format(l LogLevel, i int, entries ...interface{}) {
	// s.w.WriteByte(l.Char())
	// s.w.WriteByte(' ')
	now := time.Now()
	s.w.WriteString(now.Format(stdoutTimestampFormat))
	s.w.WriteByte(']')
	s.w.WriteByte(' ')
	writeEntries(s.w, entries...)
	s.Flush()
}

func writeEntries(w *bufio.Writer, entries ...interface{}) {
	str := fmt.Sprint(entries...)
	endsInNL := strings.HasSuffix(str, "\n")
	w.WriteString(str)
	if !endsInNL {
		w.WriteString("\n")
	}
}

func (s *StdoutFormatter) Flush() {
	s.w.Flush()
}
