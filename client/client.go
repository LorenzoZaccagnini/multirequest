package client

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/zerolog"
)

func CheckRequest(url string, logger zerolog.Logger, wg *sync.WaitGroup) bool {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		logger.Error().Err(err).Str("url", url).Timestamp().Msg("http_error")
		return false
	}

	fmt.Println("response_code:", resp.StatusCode, "url:", url)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		logger.Info().Int("response_code", resp.StatusCode).Str("url", url).Timestamp().Msg("response_code")
		//print the same thing to the console
		return true
	} else {
		logger.Error().Err(errors.New("invalid response code")).Int("response_code:", resp.StatusCode).Str("url", url).Timestamp().Msg("response_code")
		return false
	}

}
