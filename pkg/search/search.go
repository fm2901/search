package search

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Result struct {
	Phrase  string
	Ling    string
	LineNum int64
	ColNum  int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)

	for _, file := range files {
		go func(ctx context.Context, file string) {
			src, err := os.Open(file)
			if err != nil {
				log.Print(err)
				return
			}
			defer func() {
				if cerr := src.Close(); err != nil {
					err = cerr
				}
			}()
			reader := bufio.NewReader(src)
			for {
				line, err := reader.ReadString('\n')
				log.Print(line)
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Print(err)
					return
				}
			}
		}(ctx, file)
	}

	return ch
}
