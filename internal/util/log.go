package util

import "log/slog"

type logValueWrapper struct {
	f func() any
}

func (lvw logValueWrapper) LogValue() slog.Value {
	return slog.AnyValue(lvw.f())
}

func LogValueFunc(f func() any) logValueWrapper {
	return logValueWrapper{f: f}
}
