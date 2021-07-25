package logt

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
