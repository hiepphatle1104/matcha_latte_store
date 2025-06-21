package logger

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Debugf(format string, args ...any) {
	logger.Debug().Msgf(format, args...)
}
