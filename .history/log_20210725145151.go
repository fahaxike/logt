package logt

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type logLevel uint8

const (
	Debug logLevel = iota
	Info
	Waring
	Error
)
const (
	SplitByTime string = "time"
	SplitBySize string = "size"
)

var Logt LogCon

type LogCon struct {
	rootPath  string   `json:"rootPath,omitempty"`
	splitType string   `json:"splitType,omitempty"`
	fileSize  int64    `json:"fileSize,omitempty"`
	fileTime  string   `json:"fileTime,omitempty"`
	logLevel  logLevel `json:"logLevel,omitempty"`
}
type Loger interface {
	LogDebug(format string, a ...interface{})
	LogInfo(format string, a ...interface{})
	LogWaring(format string, a ...interface{})
	LogError(format string, a ...interface{})
}

func (s *LogCon) LogDebug(format string, a ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *LogCon) LogInfo(format string, a ...interface{}) {

	panic("not implemented") // TODO: Implement
}

func (s *LogCon) LogWaring(format string, a ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *LogCon) LogError(format string, a ...interface{}) {
	if enableWrite(Error) {

	}
}

func writeFile(fileName string, msg string) {
	objFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	defer objFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		objFile.WriteString(msg)
	}
}
func enableWrite(le logLevel) bool {
	if le >= Logt.logLevel {
		return true
	} else {
		return false
	}
}
func getMsg(le string, format string, a ...interface{}) string {
	msg0 := time.Now().Format("2006-01-02 15:04:05:000")
	pc, file, line, ok := runtime.Caller(2)
	var msg1 string
	if !ok {
		msg1 = "[null:null:null]"
	} else {
		msg1 = fmt.Sprintf("[%s:%s:%d]", filepath.Base(file), runtime.FuncForPC(pc).Name(), line)
	}
	msg3 := fmt.Sprintf(format, a...)
	return msg0 + " " + le + " " + msg1 + ":" + msg3
}
