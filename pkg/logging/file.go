package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "src/gin-pro/runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

/*
os.IsNotExist：能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
os.IsPermission：能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于I/O。如果出现错误，则为*PathError。*/