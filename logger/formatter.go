package logger

import (
	"fmt"
	"time"
)

type Formatter interface {
	FormatForStdout(message string, level Level) string
	FormatForFile(message string, level Level) string
}

type DefaultFormatter struct{}

func (formatter *DefaultFormatter) FormatForStdout(message string, level Level) string {
	return fmt.Sprintf("%s :: %s", level, message)
}

func (formatter *DefaultFormatter) FormatForFile(message string, level Level) string {
	return fmt.Sprintf("%s :: %s :: %s", time.Now(), level, message)
}
