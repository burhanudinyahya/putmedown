package main

import (
	"flag"
	"log"

	"github.com/burhanudinyahya/putmedown/internal/downloader"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Error: URL is required. Provide the video URL as the first argument.")
	}
	url := args[0]
	err := downloader.TiktokDownloader(url)
	if err != nil {
		log.Fatalf("Error downloading video: %v\n", err)
	}
}
