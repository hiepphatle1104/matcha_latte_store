package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func InitLogger() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i any) string {
			return strings.ToUpper(fmt.Sprintf("| %s |", i))
		},
		FormatCaller: func(i any) string {
			caller := fmt.Sprintf("%v", i)
			parts := strings.Split(caller, "/")
			return fmt.Sprintf("%s >", parts[len(parts)-1])
		},
	}

	logger = zerolog.New(output).With().Caller().Timestamp().Logger()
	log.Logger = logger
}

func SetLevel(level zerolog.Level) {
	logger = logger.Level(level)
	log.Logger = logger
}
