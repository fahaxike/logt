package logt

type logLevel uint8

const (
	Debug logLevel = iota
	Info
	Waring
	Error
)

type LogCon struct {
}
