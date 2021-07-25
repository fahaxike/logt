package logt

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-delve/delve/pkg/dwarf/reader"
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
	panic("not implemented") // TODO: Implement
}

func getIreader(){
	runtime.
}


func getBaseStr()string{
    msg0:=time.Now().Format("2006-01-02 15:04:05:000")
	pc,file,line,ok:=runtime.Caller(2)
    var msg1
	if !ok{
		msg1="[null:null:null]"
	} else
	{
		msg1=fmt.Sprint("[%s:%s:%d]",filepath.Base(file),runtime.FuncForPC(pc).Name(),line)
	}
	return msg0+" "+msg1
}