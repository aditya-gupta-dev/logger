package logger

type Level string
type WriteMode int

const (
	StdoutOnly WriteMode = iota
	FileOnly
	Both
)
const (
	Debug Level = "DEBUG"
	Info  Level = "INFO"
	Warn  Level = "WARN"
	Fatal Level = "FATAL"
)
