package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log := loggerSetup()
	startEnv(log)
	//InitAuth(log)
	httpServer := startServer(log)
	waitForShutdown(httpServer, log)
}

func loggerSetup() *log.Logger {

	fileLogger := &lumberjack.Logger{
		Filename:   "logFile.log",
		MaxSize:    500, // Max size in megabytes
		MaxBackups: 2,   // Max number of old log files to keep
		MaxAge:     20,  // Max age in days
		Compress:   true,
	}
	// Create a logger that writes to both the file and os.Stdout
	multiWriter := io.MultiWriter(fileLogger, os.Stdout)
	l := log.New(multiWriter, "NF-Go ", log.LstdFlags)
	defer fileLogger.Close()
	return l

}
