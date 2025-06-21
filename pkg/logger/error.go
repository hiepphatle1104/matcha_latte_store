package logger

func Error(msg string) {
	logger.Error().Msg(msg)
}
func Errorf(format string, args ...any) {
	logger.Error().Msgf(format, args...)
}
