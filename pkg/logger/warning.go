package logger

func Warn(msg string) {
	logger.Warn().Msg(msg)
}

func Warnf(format string, args ...any) {
	logger.Warn().Msgf(format, args...)
}
