package logger

import (
	"os"
	"path/filepath"
	"regexp"
	"fmt"
	"time"
	"runtime/debug"
	"runtime"
)

type LogLevel int

const(
	DebugLevel   LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

const (
	DebugLevelName   = "debug"
	InfoLevelName    = "info"
	WarningLevelName = "warning"
	ErrorLevelName   = "error"
)

var (
	filenameRegex = regexp.MustCompile(`[^\\/]+\.[^\\/]+`)

	logLevelById = map[LogLevel]string{
		DebugLevel  : DebugLevelName,
		InfoLevel   : InfoLevelName,
		WarningLevel: WarningLevelName,
		ErrorLevel  : ErrorLevelName,
	}

	logLevelByName = map[string]LogLevel{
		DebugLevelName  : DebugLevel,
		InfoLevelName   : InfoLevel,
		WarningLevelName: WarningLevel,
		ErrorLevelName  : ErrorLevel,
	}
	LogTmpl = "Hardcore | %v | %s: %s\n"
	logger *Logger
)

type LogMessage struct {
	Message string
	Level   LogLevel
	Args    []interface{}
}

func NewLogMessage(level LogLevel, message string, args ...interface{}) *LogMessage {
	logMessage := new(LogMessage)
	logMessage.Level = level
	logMessage.Message = message
	logMessage.Args = args
	return logMessage
}

type LogWriter interface {
	WriteString(string)
}

type StdoutWriter struct {}

func (this *StdoutWriter) WriteString(str string) {
	os.Stdout.WriteString(str)
}

type FileWriter struct {
	filename string
}

func (this *FileWriter) WriteString(str string) {
	f, err := os.OpenFile(this.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err == nil {
		_, err = f.WriteString(str)
		f.Close()
	}
}

type Logger struct {
	LogLevelName string
	Output       string
	level        LogLevel
	writer       LogWriter
	messages     chan *LogMessage
}

func Me() *Logger {
	if logger == nil {
		logger = new(Logger)
		logger.messages = make(chan *LogMessage)
		logger.level = DebugLevel
		for i := 0;i < runtime.NumCPU();i++ {
			go logger.write(i)
		}
		logger.writer = new(StdoutWriter)
	}
	return logger
}

func (this *Logger) Close() {
	close(this.messages)
}

func (this *Logger) write(id int) {
	for message := range this.messages {
		this.writeMessage(id, message)
	}
}

func (this *Logger) writeMessage(id int, message *LogMessage) {
	if this.writer != nil {
		this.writer.WriteString(
			fmt.Sprintf(
				LogTmpl,
				time.Now().Format("2006-01-02 15:04:05"),
				logLevelById[message.Level],
				fmt.Sprintf(message.Message, message.Args...),
			),
		)
		if message.Level == ErrorLevel {
			os.Exit(1)
		}
	}
}

func (this *Logger) initWriter() {
	this.writer = nil
	if filenameRegex.MatchString(this.Output) {
		dir := filepath.Dir(this.Output)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			Err("directory %s is not exists", dir)
		} else {
			this.writer = &FileWriter{this.Output}
		}
	} else if len(this.Output) == 0 || this.Output == "stdout" {
		this.writer = new(StdoutWriter)
	}
}

// посылает сервису логирования запись для логирования произвольного уровня
func log(message string, level LogLevel, args ...interface{}) {
	defer func(){recover()}()
	if Me().level <= level {
		if level > InfoLevel && Me().level == DebugLevel {
			message = fmt.Sprint(message, "\n", string(debug.Stack()))
		}
		Me().messages <- NewLogMessage(level, message, args...)
	}
}

func Err(message string, args ...interface{}) {
	log(message, ErrorLevel, args...)
}

func Warn(message string, args ...interface{}) {
	log(message, WarningLevel, args...)
}

func WarnWithErr(err error) {
	Warn("%v", err)
}

func Info(message string, args ...interface{}) {
	log(message, InfoLevel, args...)
}

func Debug(message string, args ...interface{}) {
	log(message, DebugLevel, args...)
}
