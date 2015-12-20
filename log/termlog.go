// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	stdout io.Writer = os.Stdout
    colors = map[Level]string {
        FINEST: "\033[32m",
        FINE: "\033[32m\033[2m",
        DEBUG: "\033[37m\033[2m",
        TRACE: "\033[36m",
        INFO: "\033[37m",
        WARNING: "\033[33m",
        ERROR: "\033[31m",
        CRITICAL: "\033[35m",
    }
)

// This is the standard writer that prints to standard output.
type ConsoleLogWriter struct {
	format string
	w      chan *LogRecord
}

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() *ConsoleLogWriter {
	consoleWriter := &ConsoleLogWriter{
		format: "[%T %D] [%L] %M\033[0m",
		w:      make(chan *LogRecord, LogBufferLength),
	}
	go consoleWriter.run(stdout)
	return consoleWriter
}
func (c *ConsoleLogWriter) SetFormat(format string) {
	c.format = format
}
func (c ConsoleLogWriter) run(out io.Writer) {
	for rec := range c.w {
		fmt.Fprint(out, FormatLogRecord(colors[rec.Level] + c.format, rec))
	}
}

// This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (c ConsoleLogWriter) LogWrite(rec *LogRecord) {
	c.w <- rec
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (c ConsoleLogWriter) Close() {
	close(c.w)
	time.Sleep(500 * time.Millisecond) // Try to give console I/O time to complete
}
