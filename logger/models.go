package logger

type Level string
type WriteMode int

const (
	StdoutOnly WriteMode = iota
	FileOnly
	Both
)
const (
	Debug Level = "Debug"
	Info  Level = "Info"
	Warn  Level = "Warn"
	Error Level = "Error"
)
