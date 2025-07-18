package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	file *os.File
}

func formatMessage(message string, level Level) string {
	return fmt.Sprintf("%s :: %s :: %s", time.Now().String(), level, message)
}

func (logger *Logger) LogFileOnly(message string, level Level) error {
	if level == Fatal {
		logger.file.WriteString(formatMessage(message, level))
		os.Exit(1)
	} else {
		_, err := logger.file.WriteString(formatMessage(message, level))
		return err
	}
	return nil
}

func (logger *Logger) LogStdoutOnly(message string, level Level) {
	if level == Fatal {
		fmt.Println(formatMessage(message, level))
		os.Exit(1)
	} else {
		fmt.Println(formatMessage(message, level))
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

func CreateLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return &Logger{}, err
	}

	return &Logger{
		file: file,
	}, nil
}
