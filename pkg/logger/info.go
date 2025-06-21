package logger

func Infof(format string, args ...any) {
	logger.Info().Msgf(format, args...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}
