package logger

func Panic(msg string) {
	logger.Panic().Msg(msg)
}
func Panicf(format string, args ...any) {
	logger.Panic().Msgf(format, args...)
}
