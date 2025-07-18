package main

import "github.com/aditya-gupta-dev/logger/logger"

func main() {
	logfile, err := logger.CreateDefaultLogger("service.log")
	if err != nil {
		panic(err)
	}

	logfile.LogFileOnly("Application Started..", logger.Info)
	logfile.LogFileOnly("Application Init..", logger.Debug)

	logfile.LogStdoutOnly("Comic", logger.Info)

	logfile.Close()

}
