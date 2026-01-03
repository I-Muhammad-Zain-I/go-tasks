package logger

import "fmt"

type Logger struct {
	Verbose bool
}

func New(verbose bool) *Logger {
	return &Logger{Verbose: verbose}
}

func (l *Logger) Info(format string, a ...any) {
	fmt.Printf("INFO: "+format+"\n", a...)
}

func (l *Logger) Debug(format string, a ...any) {
	if l.Verbose {
		fmt.Printf("DEBUG: "+format+"\n", a...)
	}
}

func (l *Logger) Error(format string, a ...any) {
	fmt.Printf("ERROR: "+format+"\n", a...)
}
