package logger

func Trace(msg string) {
	logger.Trace().Msg(msg)
}
func Tracef(format string, args ...any) {
	logger.Trace().Msgf(format, args...)
}
