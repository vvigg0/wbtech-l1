package main

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(msg string)
}

type FileLogger struct {
}

func (*FileLogger) Save(text string) {
	fmt.Printf("[TO FILE] %v\n", text)
}

type ConsoleLogger struct{}

func (*ConsoleLogger) Write(text string) {
	fmt.Printf("[TO CL] %v\n", text)
}

type ClLoggerAdapter struct {
	cL *ConsoleLogger
}

func NewClLoggerAdapter(clLog *ConsoleLogger) Logger {
	return &ClLoggerAdapter{clLog}
}
func (adapter *ClLoggerAdapter) Log(msg string) {
	if msg == "" {
		return
	}
	adapter.cL.Write(msg)
}

type FileLoggerAdapter struct {
	fL *FileLogger
}

func NewFileLoggerAdapter(fLog *FileLogger) Logger {
	return &FileLoggerAdapter{fLog}
}
func (adapter *FileLoggerAdapter) Log(msg string) {
	if msg == "" {
		return
	}
	adapter.fL.Save(msg)
}

type TimedLogger struct {
}

func (tLogger *TimedLogger) Record(t time.Time, payload []byte) {
	fmt.Printf("[WITH TIME] %v %v\n", t, string(payload))
}

type TimedLoggerAdapter struct {
	tL *TimedLogger
}

func NewTimedLogger(tLog *TimedLogger) Logger {
	return &TimedLoggerAdapter{tLog}
}
func (adapter *TimedLoggerAdapter) Log(msg string) {
	adapter.tL.Record(time.Now(), []byte(msg))
}
func main() {
	clLog := &ConsoleLogger{}
	fLog := &FileLogger{}
	tLog := &TimedLogger{}
	clLogAdapter := NewClLoggerAdapter(clLog)
	fLogAdapter := NewFileLoggerAdapter(fLog)
	tLogAdapter := NewTimedLogger(tLog)
	loggers := []Logger{clLogAdapter, fLogAdapter, tLogAdapter}
	for _, v := range loggers {
		v.Log("123")
	}
}
