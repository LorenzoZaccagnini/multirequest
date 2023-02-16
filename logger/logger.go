package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() zerolog.Logger {
	//logger setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logFile, err := os.CreateTemp("./logs/", "log*.jsonl")
	if err != nil {
		// Can we log an error before we have our logger? :)
		log.Error().Err(err).Timestamp().Msg("there was an error creating a temporary file four our log")
	}
	fileLogger := zerolog.New(logFile).With().Logger()

	//print to console logs file path
	fmt.Println("log file used: " + logFile.Name())

	return fileLogger
}
