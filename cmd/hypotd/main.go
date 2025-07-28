package main

import (
	"context"
	"log"

	"github.com/SuNNjek/hypotd/hyprpaper"
	"github.com/SuNNjek/hypotd/providers/bing"
)

func main() {
	path, err := bing.DownloadPotd()
	if err != nil {
		log.Fatalln(err)
	}

	if err := hyprpaper.SetWallpaper(context.Background(), path); err != nil {
		log.Fatalln(err)
	}
}
