package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	file      *os.File
	formatter Formatter
}

func (logger *Logger) LogFileOnly(message string, level Level) error {
	if level == Fatal {
		logger.file.WriteString(logger.formatter.FormatForFile(message, level))
		os.Exit(1)
	} else {
		_, err := logger.file.WriteString(logger.formatter.FormatForFile(message, level))
		return err
	}
	return nil
}

func (logger *Logger) LogStdoutOnly(message string, level Level) {
	if level == Fatal {
		fmt.Println(logger.formatter.FormatForStdout(message, level))
		os.Exit(1)
	} else {
		fmt.Println(logger.formatter.FormatForStdout(message, level))
	}
}

func (logger *Logger) LogBoth(message string, level Level) {
	logger.LogStdoutOnly(message, level)
	logger.LogFileOnly(message, level)
}

func (logger *Logger) Close() {
	if logger.file != nil {
		logger.file.Close()
	}
}

func CreateDefaultLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{}, err
	}

	return &Logger{
		file:      file,
		formatter: &DefaultFormatter{},
	}, nil
}
