package scanner

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name            string
		args            args
		want            []string
		wantFileErr     bool
		urlsLenExpected int
		testLenError    bool
	}{
		{
			name: "test_ok",
			args: args{
				file: bytes.NewBufferString("https://www.google.com\nhttps://www.yahoo.com\nhttps://www.bing.com\nhttps://www.ask.com\nhttps://www.duckduckgo.com"),
			},
			wantFileErr:     false,
			urlsLenExpected: 5,
			testLenError:    false,
		},
		{
			name: "test_err",
			args: args{
				file: bytes.NewBufferString(""),
			},
			wantFileErr:     true,
			urlsLenExpected: 0,
			testLenError:    false,
		},
		{
			name: "test_err_url_len",
			args: args{
				file: bytes.NewBufferString("https://www.google.com\nhttps://www.yahoo.com\nhttps://www.bing.com\nhttps://www.ask.com\nhttps://www.duckduckgo.com"),
			},
			wantFileErr:     false,
			urlsLenExpected: 4,
			testLenError:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.file)
			//print err
			fmt.Println(err)
			if (err != nil) != tt.wantFileErr {
				t.Errorf("readFile() error = %v, wantFileErr %v", err, tt.wantFileErr)
				return
			}
			if len(got) != tt.urlsLenExpected && tt.wantFileErr == false && tt.testLenError == false {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
			if tt.testLenError == true {
				if len(got) == tt.urlsLenExpected {
					t.Errorf("readFile() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
