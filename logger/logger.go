package logger

import "os"

type Logger struct {
	file *os.File
}

func (logger *Logger) log(message string) {
	logger.file.WriteString(message)
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
