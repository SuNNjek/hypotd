package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/SuNNjek/hypotd/providers/bing"
)

func main() {
	path, err := bing.DownloadPotd()
	if err != nil {
		log.Fatalln(err)
	}

	unloadCmd := exec.Command("hyprctl", "hyprpaper", "unload", "all")
	if err := unloadCmd.Run(); err != nil {
		log.Fatalln(err)
	}

	preloadCmd := exec.Command("hyprctl", "hyprpaper", "preload", path)
	if err := preloadCmd.Run(); err != nil {
		log.Fatalln(err)
	}

	wallpaperCmd := exec.Command("hyprctl", "hyprpaper", "wallpaper", fmt.Sprintf(",%s", path))
	if err := wallpaperCmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
