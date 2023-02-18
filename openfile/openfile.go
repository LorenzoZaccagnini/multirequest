package openfile

import (
	"multirequest/scanner"
	"os"

	"github.com/rs/zerolog"
)

func Init(filePath string, log zerolog.Logger) ([]string, error) {
	// open url file
	// if you're lazy just copy go run main.go -f urls.txt
	var urls []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Error().Err(err).Timestamp().Msg("invalid argument")
		return nil, err
	}

	defer file.Close()

	// read the file line by line
	urls, err = scanner.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Timestamp().Msg("invalid file structure")
		return nil, err
	}
	return urls, nil
}
