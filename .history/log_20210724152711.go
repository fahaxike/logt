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
	rootPath  string
	splitType string
	fileSize  int64
	fileTime  string
	logLevel  logLevel
}
