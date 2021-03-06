package main

import (
	"context"
	"log"

	"github.com/fm2901/search/pkg/search"
)

func main() {

	root := context.Background()
	ctx, _ := context.WithCancel(root)
	files := []string{
		"C:/homework/dz20/search/data/test1.txt",
		"C:/homework/dz20/search/data/test2.txt",
		"C:/homework/dz20/search/data/test3.txt",
	}
	result := search.All(ctx, "test", files)
	log.Print(result)
}
