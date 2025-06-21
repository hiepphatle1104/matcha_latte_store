package logger

import os "os"

func Fatal(msg string) {
	logger.Fatal().Msg(msg)
	os.Exit(1)
}
func Fatalf(format string, args ...any) {
	logger.Fatal().Msgf(format, args...)
	os.Exit(1)
}
