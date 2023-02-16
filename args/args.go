package args

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
)

func InitArgs(fileLogger zerolog.Logger) string {

	help := flag.Bool("h", false, "list of commands")
	var file_path string
	flag.StringVar(&file_path, "f", "./urls.txt", "file path")

	flag.Parse()

	//log file path
	fileLogger.Info().Str("file_path", file_path).Timestamp().Msg("file path")

	if *help {
		flag.PrintDefaults()
		fileLogger.Info().Timestamp().Msg("user requested help")
		os.Exit(0)
	}

	if file_path == "" {
		fileLogger.Error().Timestamp().Msg("file path is required")
		os.Exit(1)
	}

	return file_path
}
