package log

type Logger interface {
	Enabled(level Level) bool
	Log(level Level, msg string, args ...any)
}
