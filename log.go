package logt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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
	RootPath  string   `json:"rootPath,omitempty"`
	SplitType string   `json:"splitType,omitempty"`
	FileSize  int64    `json:"fileSize,omitempty"`
	FileTime  string   `json:"fileTime,omitempty"`
	LogLevel  logLevel `json:"logLevel,omitempty"`
}
type Loger interface {
	LogDebug(format string, a ...interface{})
	LogInfo(format string, a ...interface{})
	LogWaring(format string, a ...interface{})
	LogError(format string, a ...interface{})
}

func init() {
	logjson, err := ioutil.ReadFile("./log.config")
	if err != nil {
		fmt.Printf("read config err:%v", err)
		Logt = LogCon{
			RootPath:  "./",
			SplitType: "time",
			FileSize:  1024,
			FileTime:  "day",
			LogLevel:  Info,
		}
	}
	err1 := json.Unmarshal(logjson, &Logt)
	if err1 != nil {
		fmt.Printf(" Unmarshal config err:%v", err1)
		Logt = LogCon{
			RootPath:  "./",
			SplitType: "time",
			FileSize:  1024,
			FileTime:  "day",
			LogLevel:  Info,
		}
	}

}
func (s *LogCon) LogDebug(format string, a ...interface{}) {
	msg := getMsg("Debug", format, a...)
	if enableWrite(Debug) {
		fullFile := getFile()
		writeFile(fullFile, msg)
	}
	fmt.Println(msg)
}

func (s *LogCon) LogInfo(format string, a ...interface{}) {
	msg := getMsg("Info", format, a...)
	if enableWrite(Info) {
		fullFile := getFile()
		writeFile(fullFile, msg)
	}
	fmt.Println(msg)
}

func (s *LogCon) LogWaring(format string, a ...interface{}) {
	msg := getMsg("Waring", format, a...)
	if enableWrite(Waring) {
		fullFile := getFile()
		writeFile(fullFile, msg)
	}
	fmt.Println(msg)
}

func (s *LogCon) LogError(format string, a ...interface{}) {

	msg := getMsg("Error", format, a...)
	if enableWrite(Error) {
		fullFile := getFile()
		writeFile(fullFile, msg)
	}
	fmt.Println(msg)
}

func writeFile(fileName string, msg string) {
	objFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		fmt.Println(err)
	} else {
		defer objFile.Close()
		objFile.WriteString(msg)
	}
}

func enableWrite(le logLevel) bool {
	if le >= Logt.LogLevel {
		return true
	} else {
		return false
	}
}
func getFile() string {

	var fileName string
	switch strings.ToLower(Logt.SplitType) {
	case "time":
		fileName = getFileNameByTime()
	case "size":

		fileName = getFileNameBySize()
	default:
		fileName = getFileNameByTime()
	}
	return filepath.Join(Logt.RootPath, fileName)

}
func getFileNameBySize() string {
	var fileName string = "log0.txt"
	fullPath := path.Join(Logt.RootPath, fileName)
	sta, err := os.Stat(fullPath)
	if err != nil {
		fmt.Printf("read logfile %s err:%v\n", fullPath, err)
	} else {
		if sta.Size() >= Logt.FileSize {
			backName := "logback_" + time.Now().Format("20060102150405") + ".txt"
			os.Rename(fullPath, path.Join(Logt.RootPath, backName))
		}
	}
	return fileName
}
func getFileNameByTime() string {
	var fileName string
	// var f1 string
	// f1 = strings.ToLower(Logt.FileTime)
	switch strings.ToLower(Logt.FileTime) {
	case "day":
		fileName = "log" + time.Now().Format("20060102") + ".txt"
	case "hour":
		fileName = "log" + time.Now().Format("2006010215") + ".txt"
	default:
		fileName = "log" + time.Now().Format("20060102") + ".txt"
	}
	return fileName
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
