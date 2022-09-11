package logger

import (
	"log"
	"os"
)

type Log struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func NewLogger(path string) (*Log, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0660)
	if err != nil{
		file.Close()
		return nil, err
	}

	infoLog := log.New(file, "INFO: ", log.Ldate|log.Ltime)
	errLog := log.New(file, "ERR: ", log.Ldate|log.Ltime)

	logger := &Log{
		infoLog:  infoLog,
		errorLog: errLog,
	}

	return logger, nil
}

func (l *Log) Info(mes string){
	l.infoLog.Println(mes)
}

func (l *Log) Error(mes string) {
	l.errorLog.Println(mes)
}
