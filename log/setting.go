package log

import (
	"io"
	"log"
	"os"
)

func Output(file string, content error) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	writers := []io.Writer{
		logFile,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
	logger.Println(content)
}
