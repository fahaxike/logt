package logt

import (
	"fmt"
	"runtime"

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
