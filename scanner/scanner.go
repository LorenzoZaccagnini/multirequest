package scanner

import (
	"bufio"
	"errors"
	"io"
)

func ReadFile(file io.Reader) ([]string, error) {
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if urls == nil {
		//return empty error
		return nil, errors.New("empty file")
	}
	return urls, scanner.Err()
}
