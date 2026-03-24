package log

type NopLogger struct{}

func (NopLogger) Enabled(Level) bool {
	return false
}

func (NopLogger) Log(Level, string, ...any) {}
